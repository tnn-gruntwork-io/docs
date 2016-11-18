package nav

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPage_PopulateTitle(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		outputPath string
		expected   string
	}{
		{outputPath: "packages/package-vpc/modules/vpc-app/overview.md", expected: "Overview" },
	}

	for _, testCase := range testCases {
		page := NewPageWithOutputPath(testCase.outputPath)
		actual := page.getTitle()

		assert.Equal(t, testCase.expected, actual, "outputPath = %s\n", testCase.outputPath)
	}
}

func TestPage_getPathRelativeToPackageRoot(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		inputPath string
		expected  string
	}{
		{inputPath: "packages/package-vpc/modules/vpc-app/README.md", expected: "/modules/vpc-app/README.md" },
		{inputPath: "packages/package-vpc/examples/vpc-app/README.md", expected: "/examples/vpc-app/README.md" },
		{inputPath: "packages/package-vpc/README.md", expected: "/README.md" },
	}

	for _, testCase := range testCases {
		actual, err := getPathRelativeToPackageRoot(testCase.inputPath)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, testCase.expected, actual, "inputPath = %s\n", testCase.inputPath)
	}
}

func TestPage_getPackageName(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		inputPath   string
		expected    string
		expectedErr error
	}{
		{inputPath: "packages/package-vpc/modules/vpc-app/README.md", expected: "package-vpc", expectedErr: nil },
		{inputPath: "packages/package-vpc/README.md", expected: "package-vpc", expectedErr: nil },
		{inputPath: "package/package-vpc/README.md", expected: "/README.md", expectedErr: &WrongNumberOfCaptureGroupsReturnedFromPageRegEx{} },
	}

	for _, testCase := range testCases {
		actual, err := getPackageName(testCase.inputPath)

		if testCase.expectedErr == nil {
			assert.Nil(t, err, "inputPath = %s\n", testCase.inputPath)
			assert.Equal(t, testCase.expected, actual, "inputPath = %s\n", testCase.inputPath)
		}

		if testCase.expectedErr != nil {
			assert.NotNil(t, err, "inputPath = %s\n", testCase.inputPath)
		}
	}
}

func TestPage_isPackageInputPath(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		inputPath string
		expected  bool
	}{
		{inputPath: "packages/package-vpc/modules/vpc-app/README.md", expected: true },
		{inputPath: "packages/package-vpc/README.md", expected: true },
		{inputPath: "package/package-vpc/README.md", expected: false },
		{inputPath: "global/introduction", expected: false },
		{inputPath: "/packages/introduction", expected: false },
	}

	for _, testCase := range testCases {
		actual := isPackageInputPath(testCase.inputPath)

		assert.Equal(t, testCase.expected, actual, "inputPath = %s\n", testCase.inputPath)
	}
}

func TestPage_ConvertPackageLinkToUrl(t *testing.T) {
	t.Parallel()

	const PACKAGE_GITHUB_REPO_URL_PREFIX_PACKAGE_VPC = "https://github.com/gruntwork-io/package-vpc/tree/master"
	const PACKAGE_GITHUB_REPO_URL_PREFIX_PACKAGE_ASG = "https://github.com/gruntwork-io/package-asg/tree/master"

	testCases := []struct {
		inputPath   string
		linkPath    string
		expected    string
		expectedErr error
	}{
		{inputPath: "packages/package-vpc/modules/vpc-app/README.md", linkPath: "../vpc-mgmt", expected: PACKAGE_GITHUB_REPO_URL_PREFIX_PACKAGE_VPC + "/modules/vpc-mgmt" },
		{inputPath: "packages/package-vpc/modules/vpc-app/README.md", linkPath: "../../vpc-mgmt", expected: PACKAGE_GITHUB_REPO_URL_PREFIX_PACKAGE_VPC + "/vpc-mgmt" },
		{inputPath: "packages/package-vpc/modules/vpc-app/README.md", linkPath: "/examples", expected: PACKAGE_GITHUB_REPO_URL_PREFIX_PACKAGE_VPC + "/examples" },
		{inputPath: "packages/package-asg/modules/asg-cluster/_docs/concepts.md", linkPath: "/modules/asg-alerts", expected: PACKAGE_GITHUB_REPO_URL_PREFIX_PACKAGE_ASG + "/modules/asg-alerts" },
		{inputPath: "packages/package-asg/modules/asg-cluster/_docs/subfolder1/subfolder2/concepts.md", linkPath: "../modules/asg-alerts", expected: PACKAGE_GITHUB_REPO_URL_PREFIX_PACKAGE_ASG + "/modules/asg-cluster/_docs/subfolder1/modules/asg-alerts" },
		{inputPath: "global/introduction/overview.md", linkPath: "/examples", expected: "/examples" },
		{inputPath: "packages/package-vpc/modules/vpc-app/README.md", linkPath: "./", expected: PACKAGE_GITHUB_REPO_URL_PREFIX_PACKAGE_VPC + "/modules/vpc-app/README.md" },
	}

	for _, testCase := range testCases {
		actual, err := convertPackageLinkToUrl(testCase.inputPath, testCase.linkPath)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, testCase.expected, actual, "inputPath = %s\n\tlinkPath = %s\n", testCase.inputPath, testCase.linkPath)
	}
}

func TestPage_getAllLinkPaths(t *testing.T) {
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

func TestConvertMarkdownLinksToUrls(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		inputPath string
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
		actual, err := convertMarkdownLinksToUrls(testCase.inputPath, testCase.body)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, testCase.expected, actual, testCase.body)
	}
}

func NewPageWithOutputPath(outputPath string) *Page {
	page := NewFile("", "")
	page.OutputPath = outputPath
	return NewPage(page)
}

func NewPageWithInputPath(inputPath string) *Page {
	page := NewFile("", "")
	page.InputPath = inputPath
	return NewPage(page)
}