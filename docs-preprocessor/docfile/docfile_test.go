package docfile

import (
	//"testing"
	//
	//"github.com/stretchr/testify/assert"
)
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetRelativeFilePaths(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		path string
		body string
		expected string
	}{
		{	"packages/module-vpc/modules/vpc-app/README.md",
			`# VPC-App Terraform Module

			This Terraform Module launches a single VPC meant to house applications. By contrast, DevOps-related services such as
			Jenkins or InfluxDB should be in a "mgmt" VPC. (See the [vpc-mgmt](../vpc-mgmt) module.)`,
			`# VPC-App Terraform Module

			This Terraform Module launches a single VPC meant to house applications. By contrast, DevOps-related services such as
			Jenkins or InfluxDB should be in a "mgmt" VPC. (See the [vpc-mgmt](https://github.com/gruntwork-io/module-vpc/tree/master/modules/vpc-mgmt) module.)`,
		},
	}

	for _, testCase := range testCases {
		actual := sanitizeRelativeFilePaths(testCase.path, testCase.body)
		assert.Equal(t, testCase.expected, actual, testCase.body)
	}
}

func TestStripPackageInfoFromPath(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		rootPath string
		expected string
	}{
		{ "packages/module-vpc/modules/vpc-app/README.md", "modules/vpc-app/README.md" },
	}

	for _, testCase := range testCases {
		actual := stripPackageInfoFromPath(testCase.rootPath)
		assert.Equal(t, testCase.expected, actual, "rootPath = %s\n", testCase.rootPath)
	}
}

func TestNormalizeRelPaths(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		rootPath string
		relPath string
		expected string
	}{
		{ "packages/module-vpc/modules/vpc-app/README.md", "../vpc-mgmt", "packages/module-vpc/modules/vpc-mgmt" },
	}

	for _, testCase := range testCases {
		actual := getNormalizedRelPath(testCase.rootPath, testCase.relPath)
		assert.Equal(t, testCase.expected, actual, "rootPath = %s\nrelPath = %s\n", testCase.rootPath, testCase.relPath)
	}
}

//func TestCopyFile(t *testing.T) {
//	t.Parallel()
//
//	// Create a tempFile
//	file, err := ioutil.TempFile("", "docs-preprocessor")
//	if err != nil {
//		t.Fatal("Failed to create temp file.")
//	}
//
//	// Add random characters to distinguish the new file from the original
//	srcPath := file.Name()
//	dstPath := srcPath + "xyz"
//
//	copyFile(srcPath, dstPath)
//
//	assert.True(t, isFileExist(dstPath), "Expected %s to exist, but no file found at that path.", dstPath)
//}

// func TestProcessDocumentationFile(t *testing.T) {
// 	t.Parallel()

// 	testCases := []struct {
// 		inputFixturePath          string
// 		expectedOutputFixturePath string
// 	}{
// 		{"documentation.txt", "documentation-output.txt"},
// 		{"documentation-no-urls.md", "documentation-no-urls-output.md"},
// 		{"documentation-with-urls.md", "documentation-with-urls-output.md"},
// 		{"logo.png", "logo-output.png"},
// 	}

// 	for _, testCase := range testCases {
// 		actual, err := getContentsForDocumentationFile(testCase.inputFixturePath, &Opts{InputPath: GENERATOR_TESTS_FIXTURES_PATH})
// 		assert.Nil(t, err, "Error processing file %s: %v", testCase.inputFixturePath, err)

// 		expected := readProcessorTestsFixturesFile(t, testCase.expectedOutputFixturePath)
// 		assert.Equal(t, expected, actual, "inputFixturePath = %s, expectedOutputFixturePath = %s.\nExpected:\n%s\nActual:\n%s", testCase.inputFixturePath, testCase.expectedOutputFixturePath, string(expected), string(actual))
// 	}
// }

// func readProcessorTestsFixturesFile(t *testing.T, file string) []byte {
// 	bytes, err := ioutil.ReadFile(path.Join(GENERATOR_TESTS_FIXTURES_PATH, file))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	return bytes
// }
