package nav

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFile_PopulateOutputPath(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		inputPath string
		expected  string
	}{
		{ inputPath: "global/_files/hello.txt", expected: "_files/hello.txt" },
		{ inputPath: "global/help/support.md", expected: "help/support.md" },
		{ inputPath: "packages/package-vpc/modules/vpc-app/README.md", expected: "packages/package-vpc/vpc-app/overview.md" },
	}

	for _, testCase := range testCases {
		file := NewFile(testCase.inputPath, "")

		err := file.PopulateOutputPath()
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, testCase.expected, file.OutputPath, "inputPath = %s\n", testCase.inputPath)
	}
}