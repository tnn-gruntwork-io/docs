package nav

import (
	"regexp"
	"fmt"
	"github.com/gruntwork-io/terragrunt/errors"
	"github.com/gruntwork-io/docs/logger"
	"path/filepath"
	"github.com/gruntwork-io/docs/file"
	"github.com/gruntwork-io/docs/gruntwork_package"
)

// A File represents a generic file on the file system. Examples include markdown files, images, txt files, PDFs, etc.
type File struct {
	FullInputPath  string // the original input path of the file, relative to the OS directory where this program was launched
	InputPath      string // the original input path of the file, relative to the user-specified input folder
	OutputPath     string // the path where this page will exist when finally output

	inputPathRegEx string // the RegEx used to interpret the type of file this is based on its inputPath
	isFile         bool   // true if this file matches a "file" RegEx
	isPage         bool   // true if this file matches a "page" RegEx
}

// The type signature for a function that takes an inputPath and returns an outputPath
type getOutputPathFuncType func(string, []gruntwork_package.GruntworkPackage) (string, error)

// Populate the OutputPath property by looking up the appropriate RegEx.
// Store the results of our search in a private property (isFile or isPage) to avoid duplicating the RegEx check in other functions.
func (f *File) PopulateOutputPath(gruntworkPackages []gruntwork_package.GruntworkPackage) error {
	var err error

	for regexStr, getOutputPathFunc := range getFileRegExes() {
		regex := regexp.MustCompile(regexStr)
		if regex.MatchString(f.InputPath) {
			f.inputPathRegEx = regexStr
			f.isFile = true
			f.OutputPath, err = getOutputPathFunc(f.InputPath, gruntworkPackages)
			if err != nil {
				return errors.WithStackTrace(err)
			}

			return nil
		}
	}

	for regexStr, getOutputPathFunc := range getPageRegExes() {
		regex := regexp.MustCompile(regexStr)
		if regex.MatchString(f.InputPath) {
			f.inputPathRegEx = regexStr
			f.isPage = true
			f.OutputPath, err = getOutputPathFunc(f.InputPath, gruntworkPackages)
			if err != nil {
				return errors.WithStackTrace(err)
			}

			return nil
		}
	}

	return FileInputPathDidNotMatchAnyRegEx(f.InputPath)
}

// Return true if this file matches a "file" RegEx
func (f *File) IsFile() bool {
	if f.isFile {
		return true
	}

	for regexStr, _ := range getFileRegExes() {
		regex := regexp.MustCompile(regexStr)
		if regex.MatchString(f.InputPath) {
			return true
		}
	}

	return false
}

// Return true if this file matches a "page" RegEx
func (f *File) IsPage() bool {
	if f.isPage {
		return true
	}

	for regexStr, _ := range getPageRegExes() {
		regex := regexp.MustCompile(regexStr)
		if regex.MatchString(f.InputPath) {
			return true
		}
	}

	return false
}

// Get a struct of type Page based on this File
func (f *File) GetAsPage(rootNavFolder *Folder) *Page {
	return NewPage(f, rootNavFolder)
}

// Write the file to rootOutputPath/file.outputPath
func (f *File) WriteToOutputPath(rootInputPath, rootOutputPath string) error {
	absInputPath := filepath.Join(rootInputPath, f.InputPath)
	absOutputPath := filepath.Join(rootOutputPath, f.OutputPath)

	logger.Logger.Printf("Writing FILE %s to %s\n", absInputPath, absOutputPath)

	err := file.CopyFile(absInputPath, absOutputPath)
	if err != nil {
		return errors.WithStackTrace(err)
	}

	return nil
}

// Return a new instance of a File
func NewFile(inputPath, fullInputPath string) *File {
	return &File{
		FullInputPath: fullInputPath,
		InputPath: inputPath,
	}
}

// Custom errors

type FileInputPathDidNotMatchAnyRegEx string
func (inputPath FileInputPathDidNotMatchAnyRegEx) Error() string {
	return fmt.Sprintf("The path %s did not match any RegEx.\n", inputPath)
}