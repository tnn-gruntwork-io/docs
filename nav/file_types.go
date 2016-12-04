package nav

import (
	"regexp"
	"fmt"
	"github.com/gruntwork-io/docs/errors"
	"github.com/gruntwork-io/docs/config"
)

// Get a map of all RegExes that match an inputPath that represents a File (versus representing a Page)
func getFileRegExes() map[string]getOutputPathFuncType {
	fileRegExes := map[string]getOutputPathFuncType{
		IS_GLOBAL_IMAGE_FILE_REGEX: GetOutputPathOfGlobalImageFile,
		IS_GLOBAL_NONIMAGE_FILE_REGEX: GetOutputPathOfGlobalNonimageFile,
		IS_MODULE_IMAGE_FILE_REGEX: GetOutputPathOfModuleImageFile,
		IS_MODULE_NONIMAGE_FILE_REGEX: GetOutputPathOfModuleNonimageFile,
	}

	return fileRegExes
}

const IS_GLOBAL_IMAGE_FILE_REGEX = `^global/([\w -/]*_images)/([\w -]+\.(jpg|jpeg|gif|png|svg))$`
const IS_GLOBAL_IMAGE_FILE_REGEX_NUM_CAPTURE_GROUPS = 3

func GetOutputPathOfGlobalImageFile(inputPath string, config config.Config) (string, error) {
	var outputPath string

	regex := regexp.MustCompile(IS_GLOBAL_IMAGE_FILE_REGEX)
	submatches := regex.FindAllStringSubmatch(inputPath, -1)

	if len(submatches) == 0 || len(submatches[0]) != IS_GLOBAL_IMAGE_FILE_REGEX_NUM_CAPTURE_GROUPS + 1 {
		return outputPath, errors.WithStackTrace(&WrongNumberOfCaptureGroupsReturnedFromFileRegEx{inputPath: inputPath, regExName: "IS_GLOBAL_IMAGE_FILE_REGEX", regEx: IS_GLOBAL_IMAGE_FILE_REGEX })
	}

	// If we were parsing inputPath = global/help/_images/sample.jpg...
	subfolderName := submatches[0][1] // = help/_images
	imageName := submatches[0][2] // = sample.jpg

	outputPath = fmt.Sprintf("%s/%s", subfolderName, imageName)

	return outputPath, nil
}

const 	IS_GLOBAL_NONIMAGE_FILE_REGEX = `^global/([\w -/]*_files)/([\w -]+\.(css|js|txt|pdf|doc|docx|xls|xlsx|rtf|csv|json|xml|yml|yaml|key|ppt|pptx))$`
const IS_GLOBAL_NONIMAGE_FILE_REGEX_NUM_CAPTURE_GROUPS = 3

func GetOutputPathOfGlobalNonimageFile(inputPath string, config config.Config) (string, error) {
	var outputPath string

	regex := regexp.MustCompile(IS_GLOBAL_NONIMAGE_FILE_REGEX)
	submatches := regex.FindAllStringSubmatch(inputPath, -1)

	if len(submatches) == 0 || len(submatches[0]) != IS_GLOBAL_NONIMAGE_FILE_REGEX_NUM_CAPTURE_GROUPS + 1 {
		return outputPath, errors.WithStackTrace(&WrongNumberOfCaptureGroupsReturnedFromFileRegEx{inputPath: inputPath, regExName: "IS_GLOBAL_NONIMAGE_FILE_REGEX", regEx: IS_GLOBAL_NONIMAGE_FILE_REGEX })
	}

	// If we were parsing inputPath = global/help/_files/bueno.txt...
	subfolderName := submatches[0][1] // = help/_files
	fileName := submatches[0][2] // = bueno.txt

	outputPath = fmt.Sprintf("%s/%s", subfolderName, fileName)

	return outputPath, nil
}

const IS_MODULE_IMAGE_FILE_REGEX = `^packages/([\w -]+)/modules/_images/([\w -]+\.(jpg|jpeg|gif|png|svg))$`
const IS_MODULE_IMAGE_FILE_REGEX_NUM_CAPTURE_GROUPS = 3

func GetOutputPathOfModuleImageFile(inputPath string, config config.Config) (string, error) {
	var outputPath string

	regex := regexp.MustCompile(IS_MODULE_IMAGE_FILE_REGEX)
	submatches := regex.FindAllStringSubmatch(inputPath, -1)

	if len(submatches) == 0 || len(submatches[0]) != IS_MODULE_IMAGE_FILE_REGEX_NUM_CAPTURE_GROUPS + 1 {
		return outputPath, errors.WithStackTrace(&WrongNumberOfCaptureGroupsReturnedFromFileRegEx{inputPath: inputPath, regExName: "IS_MODULE_IMAGE_FILE_REGEX", regEx: IS_MODULE_IMAGE_FILE_REGEX })
	}

	// If we were parsing inputPath = packages/package-vpc/modules/_images/sample.jpg...
	packageName := submatches[0][1] // = package-vpc
	imageName := submatches[0][2] // = sample.jpg

	outputPath = fmt.Sprintf("packages/%s/_images/%s", packageName, imageName)

	return outputPath, nil
}

const IS_MODULE_NONIMAGE_FILE_REGEX = `^packages/([\w -]+)/modules/_files/([\w -]+\.(css|js|txt|pdf|doc|docx|xls|xlsx|rtf|csv|json|xml|yml|yaml|key|ppt|pptx))$`
const IS_MODULE_NONIMAGE_FILE_REGEX_NUM_CAPTURE_GROUPS = 3

func GetOutputPathOfModuleNonimageFile(inputPath string, config config.Config) (string, error) {
	var outputPath string

	regex := regexp.MustCompile(IS_MODULE_NONIMAGE_FILE_REGEX)
	submatches := regex.FindAllStringSubmatch(inputPath, -1)

	if len(submatches) == 0 || len(submatches[0]) != IS_MODULE_NONIMAGE_FILE_REGEX_NUM_CAPTURE_GROUPS + 1 {
		return outputPath, errors.WithStackTrace(&WrongNumberOfCaptureGroupsReturnedFromFileRegEx{inputPath: inputPath, regExName: "IS_MODULE_NONIMAGE_FILE_REGEX", regEx: IS_MODULE_NONIMAGE_FILE_REGEX })
	}

	// If we were parsing inputPath = packages/package-vpc/modules/_files/namaste.txt...
	packageName := submatches[0][1] // = package-vpc
	fileName := submatches[0][2] // = namaste.txt

	outputPath = fmt.Sprintf("packages/%s/_files/%s", packageName, fileName)

	return outputPath, nil
}

// Custom errors

type WrongNumberOfCaptureGroupsReturnedFromFileRegEx struct {
	inputPath string
	regExName string
	regEx     string
}
func (err WrongNumberOfCaptureGroupsReturnedFromFileRegEx) Error() string {
	return fmt.Sprintf("The wrong number of capture groups was found. This may be because the path did not match the RegEx.\ninputPath = %s\nregExName = %s\nregEx = %s\n", err.inputPath, err.regExName, err.regEx)
}