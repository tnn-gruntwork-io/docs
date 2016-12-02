package gruntwork_package

import (
	"encoding/json"
)

type GruntworkPackage struct {
	Name      string `json:"name"`
	Alias     string `json:"alias"`
	GithubUrl string `json:"github_url"`
	GitRef    string `json:"git_ref"`
}

type JsonFile struct {
	Packages []GruntworkPackage `json:"packages"`
}

func GetSliceFromJson(jsonString string) ([]GruntworkPackage, error) {
	var gPackages []GruntworkPackage
	var jsonFile JsonFile

	err := json.Unmarshal([]byte(jsonString), &jsonFile)
	if err != nil {
		return gPackages, err
	}

	gPackages = jsonFile.Packages

	return gPackages, nil
}


