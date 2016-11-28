package nav

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/gruntwork-io/docs/file"
	"path/filepath"
)

const TEST_FIXTURE_INPUT_FILES_PATH = "test-fixtures/inputs"
const TEST_FIXTURE_OUTPUT_FILES_PATH = "test-fixtures/outputs"

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
		{inputPath: "packages/module-vpc/modules/vpc-mgmt-network-acls/README.md", expected: true },
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

	bodyFixture1Path := filepath.Join(TEST_FIXTURE_INPUT_FILES_PATH, "vpc-app-example.md")
	expectedBodyFixture1Path := filepath.Join(TEST_FIXTURE_OUTPUT_FILES_PATH, "vpc-app-example.md")

	bodyFixture1 := readFile(t, bodyFixture1Path)
	expectedBodyFixture1 := readFile(t, expectedBodyFixture1Path)

	testCases := []struct {
		inputPath string
		body      string
		expected  string
	}{
		{"packages/module-vpc/modules/vpc-app/README.md",
			`Check out the [examples folder](/examples).`,
			`Check out the [examples folder](https://github.com/gruntwork-io/module-vpc/tree/master/examples).`,
		},
		{"packages/module-vpc/modules/vpc-app/README.md",
			`# VPC-App Terraform Module
			This Terraform Module launches a single VPC meant to house applications. By contrast, DevOps-related services such as
			Jenkins or InfluxDB should be in a "mgmt" VPC. (See the [vpc-mgmt](../vpc-mgmt) module.)`,
			`# VPC-App Terraform Module
			This Terraform Module launches a single VPC meant to house applications. By contrast, DevOps-related services such as
			Jenkins or InfluxDB should be in a "mgmt" VPC. (See the [vpc-mgmt](https://github.com/gruntwork-io/module-vpc/tree/master/modules/vpc-mgmt) module.)`,
		},
		{
			"packages/module-vpc/examples/vpc-app/README.md",
			`# App VPC Example

			This shows an example of how to use the [vpc-app](/modules/vpc-app) module to launch a VPC that can be used as a
			production or staging environment for your apps.

			For more information on the core concepts behind the VPC, see the [vpc-app](/modules/vpc-app) module. For common
			usage-patterns with vpc-app, such as running multiple VPCs, launching EC2 instances in a VPC, setting up Network ACLs,
			and more, see the [vpc-network-acls](../vpc-network-acls) and [vpc-peering](../vpc-peering) examples.`,
			`# App VPC Example

			This shows an example of how to use the [vpc-app](https://github.com/gruntwork-io/module-vpc/tree/master/modules/vpc-app) module to launch a VPC that can be used as a
			production or staging environment for your apps.

			For more information on the core concepts behind the VPC, see the [vpc-app](https://github.com/gruntwork-io/module-vpc/tree/master/modules/vpc-app) module. For common
			usage-patterns with vpc-app, such as running multiple VPCs, launching EC2 instances in a VPC, setting up Network ACLs,
			and more, see the [vpc-network-acls](https://github.com/gruntwork-io/module-vpc/tree/master/examples/vpc-network-acls) and [vpc-peering](https://github.com/gruntwork-io/module-vpc/tree/master/examples/vpc-peering) examples.`,
		},
		{
			"packages/module-vpc/examples/vpc-app/README.md",
			bodyFixture1,
			expectedBodyFixture1,
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

func TestPage_ReplaceMdFileExtensionWithHtmlFileExtension(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		path     string
		expected string
	}{
		{path: "packages/package-vpc/vpc-app/overview.md", expected: "packages/package-vpc/vpc-app/overview.html" },
		{path: "packages/package-vpc/vpc-app/overview.md.md", expected: "packages/package-vpc/vpc-app/overview.md.html" },
		{path: "packages/package-vpc/vpc-app/.md.overview.md.md", expected: "packages/package-vpc/vpc-app/.md.overview.md.html" },
		{path: "packages/package-vpc.md/vpc-app/overview.md", expected: "packages/package-vpc.md/vpc-app/overview.html" },
	}

	for _, testCase := range testCases {
		actual, err := replaceMdFileExtensionWithHtmlFileExtension(testCase.path)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, testCase.expected, actual, "path = %s\n", testCase.path)
	}
}

func TestPage_GetRelPathToFolder(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		pageOutputPath   string
		folderOutputPath string
		expected         string
	}{
		{pageOutputPath: "packages/module-vpc/network-acl-outbound/overview.html", folderOutputPath: "packages/module-vpc/network-acl-outbound", expected: ".." },
		{pageOutputPath: "packages/module-vpc/vpc-peering/examples.html", folderOutputPath: "introduction/overview.html", expected: "../../../introduction/overview.html" },
	}

	for _, testCase := range testCases {
		page := NewPageWithOutputPath(testCase.pageOutputPath)
		folder := NewFolderWithOutputPath(testCase.folderOutputPath)

		actual := page.GetRelPathToFolder(folder)

		assert.Equal(t, testCase.expected, actual, "pageOutputPath = %s\nfolderOutputPath = %s\n", testCase.pageOutputPath, testCase.folderOutputPath)
	}
}

func TestPage_GetAllGruntworkGithubUrls(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		body     string
		expected []string
	}{
		{
			body: `
			<div id="docs">

			    <div id="sidebar">
				<div class="logo">
				    <a class="site" href="/" title="Gruntwork"></a>
				</div>
				<div class="nav-border"></div>

				<nav>
				    <ul><li class='folder top_level_folder'><a href='#'>Introduction</a><ul class='hidden'><li class='page'><a href='../../introduction/overview.html'>Overview</a></li><li class='page'><a href='../../introduction/tools.html'>Tools</a></li></ul></li><li class='folder top_level_folder'><a href='#'>Help</a><ul class='hidden'><li class='page'><a href='../../help/support.html'>Support</a></li></ul></li><li class='folder top_level_folder'><a href='#'>Packages</a><ul><li class='folder package_folder'><a href='#'>Network Topology</a><ul class='hidden'><li class='page'><a class='active' href='#'>Overview</a></li><li class='page'><a href='core-concepts.html'>Core Concepts</a></li></ul><ul class='hidden'><li class='folder'><a href='#'>Modules</a><ul><li class='folder module_folder'><a href='#'>network-acl-inbound</a><ul class='hidden'><li class='page'><a href='modules/network-acl-inbound/overview.html'>Overview</a></li></ul></li><li class='folder module_folder'><a href='#'>network-acl-outbound</a><ul class='hidden'><li class='page'><a href='modules/network-acl-outbound/overview.html'>Overview</a></li></ul></li><li class='folder module_folder'><a href='#'>vpc-app</a><ul class='hidden'><li class='page'><a href='modules/vpc-app/overview.html'>Overview</a></li><li class='page'><a href='modules/vpc-app/examples.html'>Examples</a></li></ul></li><li class='folder module_folder'><a href='#'>vpc-app-network-acls</a><ul class='hidden'><li class='page'><a href='modules/vpc-app-network-acls/overview.html'>Overview</a></li></ul></li><li class='folder module_folder'><a href='#'>vpc-mgmt</a><ul class='hidden'><li class='page'><a href='modules/vpc-mgmt/overview.html'>Overview</a></li><li class='page'><a href='modules/vpc-mgmt/examples.html'>Examples</a></li></ul></li><li class='folder module_folder'><a href='#'>vpc-mgmt-network-acls</a><ul class='hidden'><li class='page'><a href='modules/vpc-mgmt-network-acls/overview.html'>Overview</a></li></ul></li><li class='folder module_folder'><a href='#'>vpc-network-acls</a><ul class='hidden'><li class='page'><a href='modules/vpc-network-acls/examples.html'>Examples</a></li></ul></li><li class='folder module_folder'><a href='#'>vpc-peering</a><ul class='hidden'><li class='page'><a href='modules/vpc-peering/overview.html'>Overview</a></li><li class='page'><a href='modules/vpc-peering/examples.html'>Examples</a></li></ul></li><li class='folder module_folder'><a href='#'>vpc-peering-external</a><ul class='hidden'><li class='page'><a href='modules/vpc-peering-external/overview.html'>Overview</a></li><li class='page'><a href='modules/vpc-peering-external/examples.html'>Examples</a></li></ul></li></ul></li></ul></li></ul></li></ul>
				</nav>
			    </div>

			    <div id="content-wrapper" tabindex="0">
				<header id="main-header">
				    <nav>
					<ul class="external">
					    <li><a href="http://gruntwork.io">Gruntwork.io</a></li>

					    <li class="button">
						<a href="http://www.gruntwork.io/#get-in-touch">Contact Us</a>
					    </li>
					</ul>
				    </nav>
				    <p>https://github.com/gruntwork-io/module-vpc/tree/master/README.md</p>
				</header>

				<article id="content" class="markdown-body entry-content" >
				    <h1><a name="vpc-modules" class="anchor" href="#vpc-modules" rel="nofollow" aria-hidden="true"><span class="octicon octicon-link"></span></a>VPC Modules</h1>

				<p>This repo contains modules for creating best-practices Virtual Private Clouds (VPCs) on AWS.</p>
				<h4><a name="main-modules" class="anchor" href="#main-modules" rel="nofollow" aria-hidden="true"><span class="octicon octicon-link"></span></a>
				Main Modules</h4>

				<p>The two main modules are:</p>

				<ul>
				<li><a href="https://github.com/gruntwork-io/module-vpc/tree/master/modules/vpc-app" rel="nofollow">vpc-app</a>: Launch a VPC meant to house applications. The VPC includes 3 &#34;tiers&#34; of subnets
				(public, private app, private persistence), routing rules, security groups, network ACLs, and NAT gateways.</li>
				<li><a href="https://github.com/gruntwork-io/module-vpc/tree/master/modules/vpc-mgmt" rel="nofollow">vpc-mgmt</a>: Launch a VPC meant to house DevOps and other management services. The VPC includes
				2 &#34;tiers&#34; of subnets (public, private), routing rules, security groups, network ACLs, and NAT gateways.</li>
				</ul>
				<h4><a name="supporting-modules" class="anchor" href="#supporting-modules" rel="nofollow" aria-hidden="true"><span class="octicon octicon-link"></span></a>
				Supporting Modules</h4>
			</div>
			`,
			expected: []string{
				"https://github.com/gruntwork-io/module-vpc/tree/master/README.md",
				"https://github.com/gruntwork-io/module-vpc/tree/master/modules/vpc-app",
				"https://github.com/gruntwork-io/module-vpc/tree/master/modules/vpc-mgmt",
			},
		},
	}

	for index, testCase := range testCases {
		actual := getGruntworkGithubUrls(testCase.body)

		assert.Equal(t, testCase.expected, actual, "See body #%d (zero-indexed) in the test source code.", index)
	}
}

func NewPageWithOutputPath(outputPath string) *Page {
	page := NewFile("", "")
	page.OutputPath = outputPath
	return NewPage(page, nil)
}

func NewFolderWithOutputPath(outputPath string) *Folder {
	folder := NewFolder("", "")
	folder.OutputPath = outputPath
	return folder
}

func readFile(t *testing.T, srcPath string) string {
	body, err := file.ReadFile(srcPath)
	if err != nil {
		t.Fatalf("Error while reading file at %s: %s\n", srcPath, err.Error())
	}

	return body
}