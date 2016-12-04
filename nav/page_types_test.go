package nav

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIsGlobalDocRegEx(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		path     string
		expected bool
	}{
		{"global/help/support.md", true},
		{"global/introduction/tools.md", true},
	}

	for _, testCase := range testCases {
		isMatch := checkRegex(testCase.path, IS_GLOBAL_DOC_REGEX)
		assert.Equal(t, testCase.expected, isMatch, "path = %s", testCase.path)
	}
}

func TestIsModuleDocRegEx(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		path     string
		expected bool
	}{
		{"packages/module-vpc/modules/vpc-app/_docs/example.md", true},
		{"packages/some_otherPackageName/modules/modname/_docs/README.md", true},
		{"packages5/some_otherPackageName/modules/modname/README.md", false},
		{"packages/module-vpc/modules/vpc-app/_docs", false},
		{"packages/module-vpc/modules/vpc-app/docs", false},
		{"packages/module-vpc/modules/vpc-app/docs/example.md", false},
		{"packages/module-vpc/modules/vpc-app/example.md", false},
		{"packages5/module-vpc/modules/vpc-app/example.md", false},
	}

	for _, testCase := range testCases {
		isMatch := checkRegex(testCase.path, IS_MODULE_DOC_REGEX)
		assert.Equal(t, testCase.expected, isMatch, "path = %s", testCase.path)
	}
}

func TestIsModuleOverviewDocRegEx(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		path     string
		expected bool
	}{
		{"packages/module-vpc/modules/vpc-mgmt-network-acls/README.md", true},
		{"packages/module-vpc/modules/vpc-peering/README.md", true},
	}

	for _, testCase := range testCases {
		isMatch := checkRegex(testCase.path, IS_MODULE_OVERVIEW_DOC_REGEX)
		assert.Equal(t, testCase.expected, isMatch, "path = %s", testCase.path)
	}
}

func TestIsModuleExampleDocRegEx(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		path     string
		expected bool
	}{
		{"packages/module-vpc/examples/vpc-app/Example.md", true},
		{"packages/module-vpc/examples/vpc-app/README.md", false},
		{"packages/module-vpc/examples/vpc-app/_docs/Example.md", true},
		{"packages/module-vpc/examples/vpc-app/docs/Example.md", true},
		{"packages/module-vpc/examples/vpc-app/docs/Example.txt", false},
		{"packages/module-vpc/examples/vpc-app/Example.txt", false},
		{"packages/module-vpc/examples/Example.txt", false},
	}

	for _, testCase := range testCases {
		isMatch := checkRegex(testCase.path, IS_MODULE_EXAMPLE_DOC_REGEX)
		assert.Equal(t, testCase.expected, isMatch, "path = %s", testCase.path)
	}
}

func TestIsModuleExampleOverviewDocRegEx(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		path     string
		expected bool
	}{
		{"packages/module-vpc/examples/vpc-app/README.md", true},
		{"packages/package-vpc/examples/vpc-app/README.md", true},
		{"packages/something_else/examples/some_module_name/README.md", true},
		{"packages/something_else/examples/some_module_name/overview.md", false},
		{"packages/package-name/some_module_name/README.md", false},
	}

	for _, testCase := range testCases {
		isMatch := checkRegex(testCase.path, IS_MODULE_EXAMPLE_OVERVIEW_DOC_REGEX)
		assert.Equal(t, testCase.expected, isMatch, "path = %s", testCase.path)
	}
}

func TestIsPackageDocRegEx(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		path     string
		expected bool
	}{
		{"packages/module-vpc/modules/_docs/README.md", true},
		{"packages/module-vpc/modules/_docs/doc-name.md", true},
		{"packages/module-vpc/modules/_docs/subfolder/README.md", true},
		{"packages/module-vpc/_docs/README.md", false},
		{"packages/module-vpc/_docs/subfolder/README.md", false},
		{"packages/module-vpc/docs/README.md", false},
		{"packages/module-vpc/README.md", false},
	}

	for _, testCase := range testCases {
		isMatch := checkRegex(testCase.path, IS_PACKAGE_DOC_REGEX)
		assert.Equal(t, testCase.expected, isMatch, "path = %s", testCase.path)
	}
}

func TestIsPackageOverviewDocRegEx(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		path     string
		expected bool
	}{
		{"packages/module-vpc/README.md", true},
		{"packages/module-vpc/examples/README.md", false},
		{"packages/module-vpc/examples/vpc-app/README.md", false},
		{"packages/module-vpc/overview.md", false},
		{"packages/package-vpc/README.md", true},
		{"packages/package-_.vpc/overview.md", false},
	}

	for _, testCase := range testCases {
		isMatch := checkRegex(testCase.path, IS_PACKAGE_OVERVIEW_DOC_REGEX)
		assert.Equal(t, testCase.expected, isMatch, "path = %s", testCase.path)
	}
}

