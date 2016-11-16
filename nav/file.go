package nav

import "regexp"

// A File represents a non-markdown generic file on the file system. Examples includes images, txt files, PDFs, etc.
type File struct {
	InputPath      string  // the original path of the file (relative to root input folder)
	OutputPath     string  // the path where this page will exist when finally output

	inputPathRegEx string  // the RegEx used to interpret the type of file this is based on its inputPath
	isFile         bool    // true if this file matches a "file" RegEx
	isPage         bool    // true if this file matches a "page" RegEx
}

// The type signature for a function that takes an inputPath and returns an outputPath
type getOutputPathFuncType func(string) string

// Populate the OutputPath property by looking up the appropriate RegEx.
// Store the results of our search in a private property (isFile or isPage) to avoid duplicating the RegEx check in other functions.
func (f *File) PopulateOutputPath() error {
	for regexStr, getOutputPathFunc := range getFileRegExes() {
		regex := regexp.MustCompile(regexStr)
		if regex.MatchString(f.InputPath) {
			f.inputPathRegEx = regexStr
			f.isFile = true
			f.OutputPath = getOutputPathFunc(f.InputPath)
		}
	}

	// TODO: Return custom error that no file type was found
	return nil
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

// TODO
// Return true if this file matches a "page" RegEx
func (f *File) IsPage() bool {
	return false
}

// TODO
func (f *File) GetAsPage() *Page {
	// - For each pageRegEx, if one matches, return true
	regEx := "bla"
	return NewPage(f, regEx)
}

// TODO
func (f *File) WriteToOutputPath() error {
	return nil
}

func NewFile(inputPath string) *File {
	return &File{
		InputPath: inputPath,
	}
}