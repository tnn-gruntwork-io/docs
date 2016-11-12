package docfile

import (
	"fmt"
	"regexp"
	"strings"
	"path/filepath"
)

const RELATIVE_FILE_PATHS_REGEX = `([A-Za-z0-9_/.-]+/)([A-Za-z0-9_.-]*)`
const NON_PACKAGE_PORTION_OF_URL_REGEX = `^packages/[\w -]+/(.*)$`
const PACKAGE_ABSOLUTE_URL = "https://github.com/gruntwork-io/module-vpc/tree/master"

// The DocFile interface represents a Gruntwork documentation file
type DocFile interface {
	// IsMatch returns true if the DocFile's properties make it a valid instance of the given DocFile type.
	IsMatch() bool

	// Copy writes a document file to the appropriate location relative to the outputPathRoot on the local filesystem.
	Copy(outputPathRoot string) error
}

func CreateAllDocFileTypes(absPath string, relPath string) []DocFile {
	docTypes := []DocFile{
		NewGlobalDoc(absPath, relPath),
		NewGlobalImageDoc(absPath, relPath),
		NewModuleDoc(absPath, relPath),
		NewModuleExampleDoc(absPath, relPath),
		NewModuleExampleOverviewDoc(absPath, relPath),
		NewModuleImageDoc(absPath, relPath),
		NewModuleOverviewDoc(absPath, relPath),
		NewPackageDoc(absPath, relPath),
		NewPackageOverviewDoc(absPath, relPath),
	}

	return docTypes
}

// Check whether the given path matches the given RegEx. We panic if there's an error (versus returning a bool and an
// error) to keep the if-else statements throughout the code simpler.
func checkRegex(relPath string, regexStr string) bool {
	regex := regexp.MustCompile(regexStr)
	return regex.MatchString(relPath)
}

func sanitizeRelativeFilePaths(rootPath string, body string) string {
	sanitizedBody := body

	regex := regexp.MustCompile(RELATIVE_FILE_PATHS_REGEX)
	submatches := regex.FindAllStringSubmatch(body, -1)

	if len(submatches) > 0  {
		for _, submatch := range submatches {
			// If the URI found is ../vpc-mgmt, then...
			relPath := submatch[0] // = ../vpc-mgmt

			rootPath := stripPackageInfoFromPath(rootPath)
			normalizedRelPath := getNormalizedRelPath(rootPath, relPath)
			absUrl := PACKAGE_ABSOLUTE_URL + "/" + normalizedRelPath
			sanitizedBody = strings.Replace(sanitizedBody, relPath, absUrl, -1)
		}
	}

	return sanitizedBody
}

// Given packages/foo/bar, return foo/bar
func stripPackageInfoFromPath(rootPath string) string {
	var updatedPath string

	regex := regexp.MustCompile(NON_PACKAGE_PORTION_OF_URL_REGEX)
	submatches := regex.FindAllStringSubmatch(rootPath, -1)

	updatedPath = submatches[0][1]
	return updatedPath
}

// Convert all instances of "../" in a path into the corresponding absolute path.
// Example: Given /a/b/c and ../d, yield /a/b/d
func getNormalizedRelPath(rootPath, relPath string) string {
	return filepath.Join(rootPath, "../" + relPath)
}

// Custom errors

type WrongNumberOfCaptureGroupsFoundInPathRegEx struct {
	docTypeName string
	path string
	regEx string
}
func (err WrongNumberOfCaptureGroupsFoundInPathRegEx) Error() string {
	return fmt.Sprintf("The wrong number of capture groups was found. This may be because the path did not match the RegEx.\ndocType = %s\npath = %s\nRegEx = %s", err.docTypeName	, err.path, err.regEx)
}

type NoCaptureGroupsFoundInNonPackagePortionOfURLRegEx struct {
	path string
	regEx string
}
func (err NoCaptureGroupsFoundInNonPackagePortionOfURLRegEx) Error() string {
	return fmt.Sprintf("No capture groups were found. This is most likely because the RegEx encountered a path for which it had not previously been tested.\nregEx = %s\npath = %s", err.regEx, err.path)
}
