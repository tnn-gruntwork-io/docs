package config

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetSliceFromJson(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		jsonString string
		expectedSlice []GruntworkPackage
	}{
		{
			`{
			  "packages": [
			    {
			      "name": "Network Topology",
			      "alias": "module-vpc",
			      "github_url": "https://github.com/gruntwork-io/module-vpc",
			      "git_ref": "0846eaef79c7853d5cab6ae9c47f8a43cf25c70a"
			    },
			    {
			      "name": "Monitoring & Alerting",
			      "alias": "module-aws-monitoring-public",
			      "github_url": "https://github.com/gruntwork-io/module-aws-monitoring-public",
			      "git_ref": "c15a7224466a73b6b87f5ffba8f0253afe0b81a1"
			    }
			  ]
			}`,
			[]GruntworkPackage{
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

			},
		},
	}

	for _, testCase := range testCases {
		actualSlice, err := GetPackagesFromJson(testCase.jsonString)
		if ! assert.Nil(t, err, "Unexpected error") {
			assert.Equal(t, testCase.expectedSlice, actualSlice, "jsonString = %s\n", testCase.jsonString)
		}
	}
}
