package nav

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/gruntwork-io/docs/gruntwork_package"
)

func TestFile_PopulateOutputPath(t *testing.T) {
	t.Parallel()

	packages := []gruntwork_package.GruntworkPackage{
		{
			Name: "Network Topology",
			Alias: "module-vpc",
			GithubUrl: "https://github.com/gruntwork-io/module-vpc",
			GitRef: "0846eaef79c7853d5cab6ae9c47f8a43cf25c70a",
		},
		{
			Name: "Network Topology",
			Alias: "module-vpc",
			GithubUrl: "https://github.com/gruntwork-io/module-vpc",
			GitRef: "0846eaef79c7853d5cab6ae9c47f8a43cf25c70a",
		},

	}

	testCases := []struct {
		inputPath string
		expected  string
	}{
		{ inputPath: "global/_files/hello.txt", expected: "_files/hello.txt" },
		{ inputPath: "global/help/support.md", expected: "help/support.md" },
		{ inputPath: "packages/package-vpc/modules/vpc-app/README.md", expected: "packages/package-vpc/modules/vpc-app/overview.md" },
	}

	for _, testCase := range testCases {
		file := NewFile(testCase.inputPath, "")

		err := file.PopulateOutputPath(packages)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, testCase.expected, file.OutputPath, "inputPath = %s\n", testCase.inputPath)
	}
}