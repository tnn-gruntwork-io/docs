package docfile

import (
	"fmt"
	"regexp"
	"strings"
	"path/filepath"
)

//const RELATIVE_FILE_PATHS_REGEX = `([A-Za-z0-9_/.-]+/)([A-Za-z0-9_.-]*)`
const FILE_PATHS_REGEX = `(?:http:/|https:/)?(/[A-Za-z0-9_/.-]+)|([A-Za-z0-9_/.-]+/[A-Za-z0-9_.-]*)`
const NON_PACKAGE_PORTION_OF_URL_REGEX = `^(packages/[\w -]+)/(.*)$`
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

// Given a doc file at the given filePath with the given body, convert all paths in the body (e.g. "/foo" or "../bar")
// to fully qualified URLs.
func convertPathsToUrls(filePath string, body string) string {
	newBody := body

	linkPaths := getAllLinkPaths(body)

	for _, linkPath := range linkPaths {
		packageFilePath, nonPackageFilePath := extractPackageInfoFromFilePath(filePath)
		normalizedLinkPath := getNormalizedLinkPath(packageFilePath, nonPackageFilePath, linkPath)
		absUrl := PACKAGE_ABSOLUTE_URL + "/" + normalizedLinkPath
		newBody = strings.Replace(newBody, linkPath, absUrl, -1)
	}

	return newBody
}

// Given a body of text find all instances of link paths (e.g. /foo or ../bar)
func getAllLinkPaths(body string) []string {
	var relPaths []string

	regex := regexp.MustCompile(FILE_PATHS_REGEX)
	submatches := regex.FindAllStringSubmatch(body, -1)

	if len(submatches) == 0 {
		return relPaths
	}

	for _, submatch := range submatches {
		relPath := submatch[0]

		// Cowardly use string search because Golang regular expressions don't support positive lookahead.
		if ! strings.Contains(relPath, "http://") && ! strings.Contains(relPath, "https://") {
			relPaths = append(relPaths, relPath)
		}
	}

	return relPaths
}

// Given a filePath like packages/modules/foo/bar, return the package portion (packages/modules) and non-package portion (foo/bar)
func extractPackageInfoFromFilePath(filePath string) (string, string) {
	var packagePath string
	var nonPackagePath string

	regex := regexp.MustCompile(NON_PACKAGE_PORTION_OF_URL_REGEX)
	submatches := regex.FindAllStringSubmatch(filePath, -1)

	packagePath = submatches[0][1]
	nonPackagePath = submatches[0][2]
	return packagePath, nonPackagePath
}

// Given a filePath (the package portion such as /packages/package-vpc, and the non-package portion such as /foo/bar)
// and a linkPath, get a normalized version of the linkPath.
// Example: linkPath = /x/y ==> x/y
// Example: packageFilePath = packages/package-vpc, nonPackageFilePath = /foo/bar, linkPath = ../p/q ==> packages/package-vpc/foo/p/q
func getNormalizedLinkPath(packageFilePath, nonPackageFilePath, linkPath string) string {
	if strings.HasPrefix(linkPath, "/") {
		return strings.Replace(linkPath, "/", "", 1)
	}

	if strings.HasPrefix(linkPath, "../") {
		return filepath.Join(nonPackageFilePath, "../", linkPath)
	}

	return ""
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
