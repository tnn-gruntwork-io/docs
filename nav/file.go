package nav

import "regexp"

type File struct {
	InputPath      string  // the original path of the file (relative to root input folder)
	OutputPath     string  // the path where this page will exist when finally output

	inputPathRegEx string  // the RegEx used to interpret the type of file this is based on its inputPath
	isFile         bool    // true if this file matches a "file" RegEx
	isPage         bool    // true if this file matches a "page" RegEx
}

type getOutputPathFuncType func(string) string

// Look up all RegExes, attempt to get a match, generate the corresponding outputPath, and populate it
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

func getFileRegExes() map[string]getOutputPathFuncType {
	fileRegExes := map[string]getOutputPathFuncType{
		IS_GLOBAL_IMAGE_FILE_REGEX: GetOutputPathOfGlobalImageFile,
	}

	return fileRegExes
}