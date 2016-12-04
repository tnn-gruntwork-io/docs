package nav

import (
	"regexp"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIsGlobalImageFileRegEx(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		path     string
		expected bool
	}{
		{"global/help/_images/sample.jpg", true},
		{"global/help/_images/sample.png", true},
		{"global/help/_images/sample.gif", true},
		{"global/help/sample.gif", false},
		{"packages/module-vpc/modules/_docs/images/sample.jpg", false},
		{"packages/module-vpc/modules/_docs/images/sample.md", false},
		{"global/sample.png", false},
		{"global/help/images/sample.doc", false},
		{"global/_images/sample.jpg", true},
	}

	for _, testCase := range testCases {
		isMatch := checkRegex(testCase.path, IS_GLOBAL_IMAGE_FILE_REGEX)
		assert.Equal(t, testCase.expected, isMatch, "path = %s", testCase.path)
	}
}

func TestIsGlobalNonImageFileRegEx(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		path     string
		expected bool
	}{
		{"global/_files/sample.json", true},
		{"global/introduction/_files/something.js", true},
		{"global/something.md", false},
		{"global/introduction/something.md", false},
		{"global/introduction/something.js", false},
	}

	for _, testCase := range testCases {
		isMatch := checkRegex(testCase.path, IS_GLOBAL_NONIMAGE_FILE_REGEX)
		assert.Equal(t, testCase.expected, isMatch, "path = %s", testCase.path)
	}
}

func TestIsModuleImageFileRegEx(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		path     string
		expected bool
	}{
		{"packages/package-vpc/modules/_images/sample.jpg", true},
	}

	for _, testCase := range testCases {
		isMatch := checkRegex(testCase.path, IS_MODULE_IMAGE_FILE_REGEX)
		assert.Equal(t, testCase.expected, isMatch, "path = %s", testCase.path)
	}
}

func TestIsModuleNonImageFileRegEx(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		path     string
		expected bool
	}{
		{"packages/package-vpc/modules/_files/sample.pdf", true},
		{"packages/package-vpc/sample.pdf", false},
	}

	for _, testCase := range testCases {
		isMatch := checkRegex(testCase.path, IS_MODULE_NONIMAGE_FILE_REGEX)
		assert.Equal(t, testCase.expected, isMatch, "path = %s", testCase.path)
	}
}

// Check whether the given path matches the given RegEx. We panic if there's an error (versus returning a bool and an
// error) to keep the if-else statements throughout the code simpler.
func checkRegex(relPath string, regexStr string) bool {
	regex := regexp.MustCompile(regexStr)
	return regex.MatchString(relPath)
}