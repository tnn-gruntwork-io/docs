package github

import (
	"fmt"
	"net/http"
	"github.com/gruntwork-io/docs/errors"
	"io/ioutil"
	"bytes"
	"path/filepath"
	"archive/zip"
	"os"
	"strings"
)

// Represents a specific commit of a GitHub Repo
type GithubCommit struct {
	RepoOwner string // e.g. "gruntwork-io"
	RepoName  string // e.g. "module-vpc"
	GitRef    string // e.g. "c185238djaf8df8asdfja"
}

// Download the given GitHub Commit
func (c *GithubCommit) Download(dstPath string, githubToken string) error {
	zipFilePath, err := downloadGithubZipFile(c, githubToken)
	if err != nil {
		return errors.WithStackTrace(err)
	}
	defer deleteAllFiles(zipFilePath)

	if err = extractZipFile(zipFilePath, dstPath, c.RepoName); err != nil {
		return errors.WithStackTrace(err)
	}

	return nil
}

// Extract the contents of the zip file at zipFilePath into localPath
func extractZipFile(zipFilePath, localPath, repoName string) error {
	// Open the zip file for reading.
	reader, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return err
	}
	defer reader.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range reader.File {
		// Replace the top-level directory name in the zip file with the repo name
		// By default, GitHub unzips to a top-level directory name like gruntwork-io-module-aws-monitoring-c15a7224466a73b6b87f5ffba8f0253afe0b81a1
		// so we need to replace this with just "module-aws-monitoring"
		relFileOutputPath := replaceTopMostDirectory(f.Name, repoName)

		if f.FileInfo().IsDir() {
			// Create a directory
			os.MkdirAll(filepath.Join(localPath, relFileOutputPath), 0777)
		} else {
			// Read the file into a byte array
			readCloser, err := f.Open()
			if err != nil {
				return fmt.Errorf("Failed to open file %s: %s", f.Name, err)
			}

			byteArray, err := ioutil.ReadAll(readCloser)
			if err != nil {
				return fmt.Errorf("Failed to read file %s: %s", f.Name, err)
			}

			// Write the file
			err = ioutil.WriteFile(filepath.Join(localPath, relFileOutputPath), byteArray, 0644)
			if err != nil {
				return fmt.Errorf("Failed to write file: %s", err)
			}
		}
	}

	return nil
}

// Download the zip file at the given URL to a temporary local directory.
// Returns the absolute path to the downloaded zip file.
// IMPORTANT: You must call "defer os.RemoveAll(dir)" in the calling function when done with the downloaded zip file!
func downloadGithubZipFile(commit *GithubCommit, githubToken string) (string, error) {
	var zipFilePath string

	// Create a temp directory. Note that ioutil.TempDir has a peculiar interface. We need not specify any meaningful
	// values to achieve our goal of getting a temporary directory.
	tempDir, err := ioutil.TempDir("", "")
	if err != nil {
		return zipFilePath, errors.WithStackTrace(err)
	}

	// Download the zip file, possibly using the GitHub oAuth Token
	req, err := makeGitHubZipFileRequest(commit, githubToken)
	if err != nil {
		return zipFilePath, errors.WithStackTrace(err)
	}

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return zipFilePath, errors.WithStackTrace(err)
	}
	if resp.StatusCode != 200 {
		return zipFilePath, fmt.Errorf("Failed to download file at the url %s. Received HTTP Response %d.", req.URL.String(), resp.StatusCode)
	}
	if resp.Header.Get("Content-Type") != "application/zip" {
		return zipFilePath, fmt.Errorf("Failed to download file at the url %s. Expected HTTP Response's \"Content-Type\" header to be \"application/zip\", but was \"%s\"", req.URL.String(), resp.Header.Get("Content-Type"))
	}

	// Copy the contents of the downloaded file to our empty file
	respBodyBuffer := new(bytes.Buffer)
	respBodyBuffer.ReadFrom(resp.Body)
	err = ioutil.WriteFile(filepath.Join(tempDir, "repo.zip"), respBodyBuffer.Bytes(), 0644)
	if err != nil {
		return zipFilePath, errors.WithStackTrace(err)
	}

	zipFilePath = filepath.Join(tempDir, "repo.zip")

	return zipFilePath, nil
}

// Return an HTTP request that will fetch the given GitHub repo's zip file for the given tag, possibly with the gitHubOAuthToken
// in the header.
func makeGitHubZipFileRequest(commit *GithubCommit, githubToken string) (*http.Request, error) {
	var request *http.Request

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/zipball/%s", commit.RepoOwner, commit.RepoName, commit.GitRef)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return request, errors.WithStackTrace(err)
	}

	if githubToken != "" {
		request.Header.Set("Authorization", fmt.Sprintf("token %s", githubToken))
	}

	return request, nil
}

func replaceTopMostDirectory(path, newDirName string) string {
	var newPath string

	pathParts := strings.Split(path, "/")
	newPath = strings.Join(pathParts[1:], "/")
	newPath = filepath.Join(newDirName, newPath)

	return newPath
}

func deleteAllFiles(path string) error {
	return os.RemoveAll(path)
}
