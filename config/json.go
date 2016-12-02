package config

import (
	"encoding/json"
)

type JsonFile struct {
	Packages               []GruntworkPackage `json:"packages"`
	TopLevelFolderOrdering []string
}

type GruntworkPackage struct {
	Name      string `json:"name"`
	Alias     string `json:"alias"`
	GithubUrl string `json:"github_url"`
	GitRef    string `json:"git_ref"`
}

func GetPackagesFromJson(jsonString string) ([]GruntworkPackage, error) {
	var gPackages []GruntworkPackage
	var jsonFile JsonFile

	err := json.Unmarshal([]byte(jsonString), &jsonFile)
	if err != nil {
		return gPackages, err
	}

	gPackages = jsonFile.Packages

	return gPackages, nil
}

func GetTopLevelFolderOrderingFromJson(jsonString string) ([]string, error) {
	var topLevelFolders []string
	var jsonFile JsonFile

	err := json.Unmarshal([]byte(jsonString), &jsonFile)
	if err != nil {
		return topLevelFolders, err
	}

	return jsonFile.TopLevelFolderOrdering, nil
}


