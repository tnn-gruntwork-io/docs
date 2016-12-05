package github

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGithubCommit_Download(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		path       string
		newDirName string
		expected   string
	}{
		{"_gruntwork-io-module-vpc-0846eaef79c7853d5cab6ae9c47f8a43cf25c70a/modules/vpc-mgmt-network-acls/overview.html", "module-vpc", "module-vpc/modules/vpc-mgmt-network-acls/overview.html"},
	}

	for _, testCase := range testCases {
		actual := replaceTopMostDirectory(testCase.path, testCase.newDirName)
		assert.Equal(t, testCase.expected, actual, "path = %s", testCase.path)
	}
}
