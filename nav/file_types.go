package nav

import (
	"regexp"
	"fmt"
	"github.com/gruntwork-io/docs/errors"
)

// FileType RegEx's
const IS_GLOBAL_IMAGE_FILE_REGEX = `^global/([\w -/]*_images)/([\w -]+\.(jpg|jpeg|gif|png|svg))$`
const IS_GLOBAL_IMAGE_FILE_REGEX_NUM_CAPTURE_GROUPS = 3

const IS_MODULE_IMAGE_FILE_REGEX = `^packages/([\w -]+)/modules/_images/([\w -]+\.(jpg|jpeg|gif|png|svg))$`
const IS_MODULE_IMAGE_FILE_REGEX_NUM_CAPTURE_GROUPS = 3

func GetOutputPathOfGlobalImageFile(inputPath string) string {
	var outputPath string

	regex := regexp.MustCompile(IS_GLOBAL_IMAGE_FILE_REGEX)
	submatches := regex.FindAllStringSubmatch(inputPath, -1)

	if len(submatches) == 0 || len(submatches[0]) != IS_GLOBAL_IMAGE_FILE_REGEX_NUM_CAPTURE_GROUPS + 1 {
		return outputPath, errors.WithStackTrace(&WrongNumberOfCaptureGroupsFoundInPathRegEx{ inputPath: inputPath, regExName: "IS_GLOBAL_IMAGE_FILE_REGEX", regEx: IS_GLOBAL_IMAGE_FILE_REGEX })
	}

	// If we were parsing inputPath = global/help/_images/sample.jpg...
	subfolderName := submatches[0][1] // = help/_images
	imageName := submatches[0][2] // = sample.jpg

	return fmt.Sprintf("%s/%s", subfolderName, imageName), nil
}