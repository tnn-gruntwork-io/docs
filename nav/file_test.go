package nav

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/gruntwork-io/docs/config"
)

func TestFile_PopulateOutputPath(t *testing.T) {
	t.Parallel()

	config := &config.Config{
		TopLevelFolderOrdering: nil,
		Packages: []config.GruntworkPackage{
			{
				Name: "Network Topology",
				FolderName: "network-topology",
				GithubRepoOwner: "gruntwork-io",
				GithubRepoName: "module-vpc",
				GitRef: "0846eaef79c7853d5cab6ae9c47f8a43cf25c70a",
			},
			{
				Name: "Monitoring & Alerting",
				FolderName: "monitoring-and-alerting",
				GithubRepoOwner: "gruntwork-io",
				GithubRepoName: "module-aws-monitoring",
				GitRef: "c15a7224466a73b6b87f5ffba8f0253afe0b81a1",
			},

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

		err := file.PopulateOutputPath(*config)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, testCase.expected, file.OutputPath, "inputPath = %s\n", testCase.inputPath)
	}
}