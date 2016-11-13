package docfile

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestConvertPathsToUrls(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		path string
		body string
		expected string
	}{
		{	"packages/module-vpc/modules/vpc-app/README.md",
			`Check out the [examples folder](/examples).`,
			`Check out the [examples folder](https://github.com/gruntwork-io/module-vpc/tree/master/examples).`,
		},
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
		actual := convertPathsToUrls(testCase.path, testCase.body)
		assert.Equal(t, testCase.expected, actual, testCase.body)
	}
}

func TestGetAllRelativePaths(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		body     string
		expected []string
	}{
		{
			`# VPC-App Terraform Module
			This Terraform Module launches a single VPC meant to house applications. By contrast, DevOps-related services such as
			Jenkins or InfluxDB should be in a "mgmt" VPC. (See the [vpc-mgmt](../vpc-mgmt) module.)`,
			[]string{"../vpc-mgmt" },
		},
		{
			`## How do you use this module?
			Check out the [examples folder](/examples)`,
			[]string{"/examples" },
		},
		{
			`A [VPC](https://aws.amazon.com/vpc/) or Virtual Private Cloud is a logically isolated section of your AWS cloud. Each
			VPC defines a virtual network within which you run your AWS resources, as well as rules for what can go in and out of
			that network. This includes subnets, route tables that tell those subnets how to route inbound and outbound traffic,
			security groups, firewalls for the subnet (known as "Network ACLs"), and any other network components such as VPN connections.`,
			[]string(nil),
		},
		{
			`A [VPC](https://aws.amazon.com/vpc/) or Virtual Private Cloud is a logically isolated section of your AWS cloud. Each
			VPC defines a virtual network within which you run your [AWS resources](../resources), as well as rules for what can go in and out of
			that network. This includes subnets, route tables that tell those subnets how to route inbound and outbound traffic,
			security groups, firewalls for the subnet (known as "Network ACLs"), and any other network components such as VPN connections.`,
			[]string{"../resources"},
		},
		{
			`A [VPC](https://aws.amazon.com/vpc/) or [Virtual Private Cloud](../../vpc/) is something special.`,
			[]string{"../../vpc/"},
		},
		{
			`A [VPC](https://aws.amazon.com/vpc/) or [Virtual Private Cloud](../../vpc/) is [something](/something) special.`,
			[]string{"../../vpc/", "/something"},
		},
	}

	for _, testCase := range testCases {
		actual := getAllLinkPaths(testCase.body)
		assert.Equal(t, testCase.expected, actual, "Test case used the following body:\n%s", testCase.body)
	}
}

func TestExtractPackageInfoFromPath(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		rootPath               string
		expectedPackagePath    string
		expectedNonPackagePath string
	}{
		{"packages/module-vpc/modules/vpc-app/README.md", "packages/module-vpc", "modules/vpc-app/README.md" },
	}

	for _, testCase := range testCases {
		actualPackagePath, actualNonPackagePath := extractPackageInfoFromFilePath(testCase.rootPath)
		assert.Equal(t, testCase.expectedPackagePath, actualPackagePath, "rootPath = %s\n", testCase.rootPath)
		assert.Equal(t, testCase.expectedNonPackagePath, actualNonPackagePath, "rootPath = %s\n", testCase.rootPath)
	}
}

func TestGetNormalizedLinkPaths(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		packageFilePath    string
		nonPackageFilePath string
		linkPath           string
		expected           string
	}{
		{"packages/module-vpc/", "modules/vpc-app/README.md", "../vpc-mgmt", "modules/vpc-mgmt" },
		{"packages/package-vpc", "", "/examples", "examples" },
		{"packages/package-vpc", "/foo/bar", "../p/q", "/p/q" },
	}

	for _, testCase := range testCases {
		actual := getNormalizedLinkPath(testCase.packageFilePath, testCase.nonPackageFilePath, testCase.linkPath)
		assert.Equal(t, testCase.expected, actual, "packagesPath = %s\nnonPackagesPath = %s\nrelPath = %s\n", testCase.packageFilePath, testCase.nonPackageFilePath, testCase.linkPath)
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
