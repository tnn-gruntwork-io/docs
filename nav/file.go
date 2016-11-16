package nav

import (
	"regexp"
	"fmt"
	"github.com/gruntwork-io/terragrunt/errors"
	"github.com/gruntwork-io/docs/logger"
	"path/filepath"
	"github.com/gruntwork-io/docs/file"
)

// A File represents a non-markdown generic file on the file system. Examples includes images, txt files, PDFs, etc.
type File struct {
	InputPath      string  // the original path of the file (relative to root input folder)
	OutputPath     string  // the path where this page will exist when finally output

	inputPathRegEx string  // the RegEx used to interpret the type of file this is based on its inputPath
	isFile         bool    // true if this file matches a "file" RegEx
	isPage         bool    // true if this file matches a "page" RegEx
}

// The type signature for a function that takes an inputPath and returns an outputPath
type getOutputPathFuncType func(string) (string, error)

// Populate the OutputPath property by looking up the appropriate RegEx.
// Store the results of our search in a private property (isFile or isPage) to avoid duplicating the RegEx check in other functions.
func (f *File) PopulateOutputPath() error {
	var err error

	for regexStr, getOutputPathFunc := range getFileRegExes() {
		regex := regexp.MustCompile(regexStr)
		if regex.MatchString(f.InputPath) {
			f.inputPathRegEx = regexStr
			f.isFile = true
			f.OutputPath, err = getOutputPathFunc(f.InputPath)
			if err != nil {
				return errors.WithStackTrace(err)
			}
		}
	}

	for regexStr, getOutputPathFunc := range getPageRegExes() {
		regex := regexp.MustCompile(regexStr)
		if regex.MatchString(f.InputPath) {
			f.inputPathRegEx = regexStr
			f.isPage = true
			f.OutputPath, err = getOutputPathFunc(f.InputPath)
			if err != nil {
				return errors.WithStackTrace(err)
			}
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
func (f *File) GetAsPage() *Page {
	return NewPage(f)
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
func NewFile(inputPath string) *File {
	return &File{
		InputPath: inputPath,
	}
}

// Custom errors

type FileInputPathDidNotMatchAnyRegEx string
func (inputPath FileInputPathDidNotMatchAnyRegEx) Error() string {
	return fmt.Sprintf("The path %s did not match any RegEx.\n", inputPath)
}