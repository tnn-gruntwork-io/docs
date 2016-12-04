package config

import (
	"encoding/json"
	"fmt"
	"github.com/gruntwork-io/docs/file"
)

type Config struct {
	Packages               []GruntworkPackage `json:"packages"`
	TopLevelFolderOrdering []string          `json:"top_level_folder_ordering"`
}

type GruntworkPackage struct {
	Name            string `json:"name"`
	Alias           string `json:"alias"`
	GithubRepoOwner string `json:"github_repo_owner"`
	GithubRepoName  string `json:"github_repo_name"`
	GitRef          string `json:"git_ref"`
}

func GetConfigFromJsonFile(path string) (Config, error) {
	var config Config

	jsonString, err := file.ReadFile(path)
	if err != nil {
		return config, fmt.Errorf("Error while reading Config File: %s", err)
	}

	err = json.Unmarshal([]byte(jsonString), &config)
	if err != nil {
		return config, fmt.Errorf("Error while unmarshalling JSON for Config File: %s", err)
	}

	return config, nil
}

func (p *GruntworkPackage) GetGithubUrl() string {
	return fmt.Sprintf("https://github.com/%s/%s", p.GithubRepoOwner, p.GithubRepoName)
}

