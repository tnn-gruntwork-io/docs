package nav

import (
	"strings"
	"github.com/gruntwork-io/docs/config"
)

// Given a Package Folder Name (e.g. "module-vpc") return the corresponding Package Name (e.g. "Network Topology")
func GetPackageFriendlyNameFromPackageFolderName(packages []config.GruntworkPackage, packageFolderName string) string {
	for _, gPackage := range packages {
		if gPackage.FolderName == packageFolderName {
			return gPackage.Name
		}
	}

	return packageFolderName
}

// Given a Package Folder Name (e.g. "module-vpc") return the corresponding Package Folder Name (e.g. "network-topology")
func GetPackageFolderNameFromPackageRepoName(packages []config.GruntworkPackage, packageRepoName string) string {
	for _, gPackage := range packages {
		if gPackage.GithubRepoName == packageRepoName {
			return gPackage.FolderName
		}
	}

	return packageRepoName
}

// Convert "network-topology" into "Network Topology"
func convertDashesToSpacesAndCapitalize(str string) string {
	str = strings.Replace(str, "-", " ", -1)
	str = strings.Title(str)

	return str
}

// Reorder the given set of folders to match the "TopLevelFolderOrdering"
func reorderFoldersToMatchTopLevelFolderOrdering(folders []*Folder, config config.Config) []*Folder {
	var newFolders []*Folder

	orderedFolderNames := config.TopLevelFolderOrdering

	for _, folderName := range orderedFolderNames {
		for _, folder := range folders {
			if folder.Name == folderName {
				newFolders = append(newFolders, folder)
			}
		}
	}

	return newFolders
}