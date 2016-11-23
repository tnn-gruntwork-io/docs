package nav

import (
	"strings"
	"regexp"
	"fmt"
	"github.com/gruntwork-io/docs/errors"
)

// Get a map of all RegExes that match an inputPath that represents a Page (versus representing a File)
func getPageRegExes() map[string]getOutputPathFuncType {
	pageRegExes := map[string]getOutputPathFuncType{
		IS_GLOBAL_DOC_REGEX: GetOutputPathOfGlobalDoc,
		IS_MODULE_DOC_REGEX: GetOutputPathOfModuleDoc,
		IS_MODULE_OVERVIEW_DOC_REGEX: GetOutputPathOfModuleOverviewDoc,
		IS_MODULE_EXAMPLE_DOC_REGEX: GetOutputPathOfModuleExampleDoc,
		IS_MODULE_EXAMPLE_OVERVIEW_DOC_REGEX: GetOutputPathOfModuleExampleOverviewDoc,
		IS_PACKAGE_DOC_REGEX: GetOutputPathOfPackageDoc,
		IS_PACKAGE_OVERVIEW_DOC_REGEX: GetOutputPathOfPackageOverviewDoc,
	}

	return pageRegExes
}

const IS_GLOBAL_DOC_REGEX = `^global/(.*\.md)$`

func GetOutputPathOfGlobalDoc(inputPath string) (string, error) {
	return strings.Replace(inputPath, "global/", "", -1), nil
}

const IS_MODULE_DOC_REGEX = `^packages/([\w -]+)/modules/([\w -]+)/_docs/([\w -]+\.md)$`
const IS_MODULE_DOC_REGEX_NUM_CAPTURE_GROUPS = 3

func GetOutputPathOfModuleDoc(inputPath string) (string, error) {
	var outputPath string

	regex := regexp.MustCompile(IS_MODULE_DOC_REGEX)
	submatches := regex.FindAllStringSubmatch(inputPath, -1)

	if len(submatches) == 0 || len(submatches[0]) != IS_MODULE_DOC_REGEX_NUM_CAPTURE_GROUPS + 1 {
		return outputPath, errors.WithStackTrace(&WrongNumberOfCaptureGroupsReturnedFromPageRegEx{ inputPath: inputPath, regExName: "IS_MODULE_DOC_REGEX", regEx: IS_MODULE_DOC_REGEX })
	}

	// If we were parsing inputPath = packages/module-vpc/modules/vpc-app/module-doc.md...
	packageName := submatches[0][1] // = module-vpc
	moduleName := submatches[0][2] 	// = vpc-app
	docName := submatches[0][3] 	// = module-doc.md

	packageName = getOutputPackageFolderNameFromInputPackageFolderName(packageName)

	outputPath = fmt.Sprintf("packages/%s/%s/%s", packageName, moduleName, docName)

	return outputPath, nil
}

const IS_MODULE_OVERVIEW_DOC_REGEX = `^packages/([\w -]+)/modules/([\w -]+)/README.md$`
const IS_MODULE_OVERVIEW_DOC_REGEX_NUM_CAPTURE_GROUPS = 2

func GetOutputPathOfModuleOverviewDoc(inputPath string) (string, error) {
	var outputPath string

	regex := regexp.MustCompile(IS_MODULE_OVERVIEW_DOC_REGEX)
	submatches := regex.FindAllStringSubmatch(inputPath, -1)

	if len(submatches) == 0 || len(submatches[0]) != IS_MODULE_OVERVIEW_DOC_REGEX_NUM_CAPTURE_GROUPS + 1 {
		return outputPath, errors.WithStackTrace(&WrongNumberOfCaptureGroupsReturnedFromPageRegEx{ inputPath: inputPath, regExName: "IS_MODULE_OVERVIEW_DOC_REGEX", regEx: IS_MODULE_OVERVIEW_DOC_REGEX })
	}

	// If we were parsing inputPath = packages/module-vpc/modules/vpc-app/README.md...
	packageName := submatches[0][1] // = module-vpc
	moduleName := submatches[0][2] 	// = vpc-app

	packageName = getOutputPackageFolderNameFromInputPackageFolderName(packageName)

	outputPath = fmt.Sprintf("packages/%s/modules/%s/overview.md", packageName, moduleName)

	return outputPath, nil
}

const IS_MODULE_EXAMPLE_DOC_REGEX = `^packages/([\s\w- ]+)/examples/([\s\w -]+)/(.*[^README].md)$`
const IS_MODULE_EXAMPLE_DOC_REGEX_NUM_CAPTURE_GROUPS = 3

func GetOutputPathOfModuleExampleDoc(inputPath string) (string, error) {
	var outputPath string

	regex := regexp.MustCompile(IS_MODULE_EXAMPLE_DOC_REGEX)
	submatches := regex.FindAllStringSubmatch(inputPath, -1)

	if len(submatches) == 0 || len(submatches[0]) != IS_MODULE_EXAMPLE_DOC_REGEX_NUM_CAPTURE_GROUPS + 1 {
		return outputPath, errors.WithStackTrace(&WrongNumberOfCaptureGroupsReturnedFromPageRegEx{ inputPath: inputPath, regExName: "IS_MODULE_EXAMPLE_DOC_REGEX", regEx: IS_MODULE_EXAMPLE_DOC_REGEX })
	}

	// If we were parsing inputPath = packages/module-vpc/examples/vpc-app/example-doc.md...
	packageName := submatches[0][1] // = module-vpc
	moduleName := submatches[0][2] 	// = vpc-app
	docName := submatches[0][3] 	// = example-doc.md

	packageName = getOutputPackageFolderNameFromInputPackageFolderName(packageName)

	outputPath = fmt.Sprintf("packages/%s/modules/%s/%s", packageName, moduleName, docName)

	return outputPath, nil
}

const IS_MODULE_EXAMPLE_OVERVIEW_DOC_REGEX = `^packages/([\s\w -]+)/examples/([\s\w -]+)/README.md$`
const IS_MODULE_EXAMPLE_OVERVIEW_DOC_REGEX_NUM_CAPTURE_GROUPS = 2

func GetOutputPathOfModuleExampleOverviewDoc(inputPath string) (string, error) {
	var outputPath string

	regex := regexp.MustCompile(IS_MODULE_EXAMPLE_OVERVIEW_DOC_REGEX)
	submatches := regex.FindAllStringSubmatch(inputPath, -1)

	if len(submatches) == 0 || len(submatches[0]) != IS_MODULE_EXAMPLE_OVERVIEW_DOC_REGEX_NUM_CAPTURE_GROUPS + 1 {
		return outputPath, errors.WithStackTrace(&WrongNumberOfCaptureGroupsReturnedFromPageRegEx{ inputPath: inputPath, regExName: "IS_MODULE_EXAMPLE_OVERVIEW_DOC_REGEX", regEx: IS_MODULE_EXAMPLE_OVERVIEW_DOC_REGEX })
	}

	// If we were parsing inputPath = packages/module-vpc/examples/vpc-app/README.md...
	packageName := submatches[0][1] // = module-vpc
	moduleName := submatches[0][2] 	// = vpc-app

	packageName = getOutputPackageFolderNameFromInputPackageFolderName(packageName)

	outputPath = fmt.Sprintf("packages/%s/modules/%s/examples.md", packageName, moduleName)

	return outputPath, nil
}

const IS_PACKAGE_DOC_REGEX = `^packages/([\w -]+)/modules/_docs/([\w -/]+\.md)$`
const IS_PACKAGE_DOC_REGEX_NUM_CAPTURE_GROUPS = 2

func GetOutputPathOfPackageDoc(inputPath string) (string, error) {
	var outputPath string

	regex := regexp.MustCompile(IS_PACKAGE_DOC_REGEX)
	submatches := regex.FindAllStringSubmatch(inputPath, -1)

	if len(submatches) == 0 || len(submatches[0]) != IS_PACKAGE_DOC_REGEX_NUM_CAPTURE_GROUPS + 1 {
		return outputPath, errors.WithStackTrace(&WrongNumberOfCaptureGroupsReturnedFromPageRegEx{ inputPath: inputPath, regExName: "IS_PACKAGE_DOC_REGEX", regEx: IS_PACKAGE_DOC_REGEX })
	}

	// If we were parsing inputPath = packages/package-vpc/modules/_docs/doc-name.md...
	packageName := submatches[0][1] // = package-vpc
	docName := submatches[0][2] 	// = doc-name.md

	packageName = getOutputPackageFolderNameFromInputPackageFolderName(packageName)

	outputPath = fmt.Sprintf("packages/%s/%s", packageName, docName)

	return outputPath, nil
}

const IS_PACKAGE_OVERVIEW_DOC_REGEX = `^packages/([\s\w -]+)/README.md$`
const IS_PACKAGE_OVERVIEW_DOC_REGEX_NUM_CAPTURE_GROUPS = 1

func GetOutputPathOfPackageOverviewDoc(inputPath string) (string, error) {
	var outputPath string

	regex := regexp.MustCompile(IS_PACKAGE_OVERVIEW_DOC_REGEX)
	submatches := regex.FindAllStringSubmatch(inputPath, -1)

	if len(submatches) == 0 || len(submatches[0]) != IS_PACKAGE_OVERVIEW_DOC_REGEX_NUM_CAPTURE_GROUPS + 1 {
		return outputPath, errors.WithStackTrace(&WrongNumberOfCaptureGroupsReturnedFromPageRegEx{ inputPath: inputPath, regExName: "IS_PACKAGE_OVERVIEW_DOC_REGEX", regEx: IS_PACKAGE_OVERVIEW_DOC_REGEX })
	}

	// If we were parsing inputPath = packages/package-vpc/README.md...
	packageName := submatches[0][1] // = package-vpc

	packageName = getOutputPackageFolderNameFromInputPackageFolderName(packageName)

	outputPath = fmt.Sprintf("packages/%s/overview.md", packageName)

	return outputPath, nil
}

// Custom errors

type WrongNumberOfCaptureGroupsReturnedFromPageRegEx struct {
	inputPath string
	regExName string
	regEx     string
}
func (err WrongNumberOfCaptureGroupsReturnedFromPageRegEx) Error() string {
	return fmt.Sprintf("The wrong number of capture groups was found. This may be because the path did not match the RegEx.\ninputPath = %s\nregExName = %s\nregEx = %s\n", err.inputPath, err.regExName, err.regEx)
}