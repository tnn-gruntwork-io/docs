package nav

import (
	"strings"
	"github.com/gruntwork-io/docs/config"
)

// Given a Package Folder Name (e.g. "module-vpc") return the corresponding Package Name (e.g. "Network Topology")
func getOutputPackageFolderNameFromInputPackageFolderName(packages []config.GruntworkPackage, packageFolderName string) string {
	for _, gPackage := range packages {
		if gPackage.Alias == packageFolderName {
			return convertSpacesToDashesAndLowerCase(gPackage.Name)
		}
	}

	return packageFolderName
}

// Convert "network-topology" into "Network Topology"
func convertDashesToSpacesAndCapitalize(str string) string {
	str = strings.Replace(str, "-", " ", -1)
	str = strings.Title(str)
	return str
}

// Convert "Network Topology" into "network-topology"
func convertSpacesToDashesAndLowerCase(str string) string {
	str = strings.Replace(str, " ", "-", -1)
	str = strings.ToLower(str)
	return str
}

// Reorder the given set of folders to match the "TopLevelFolderOrdering"
func reorderFoldersToMatchTopLevelFolderOrdering(folders []*Folder) []*Folder {
	var newFolders []*Folder

	orderedFolderNames := getTopLevelFolderOrdering()

	for _, folderName := range orderedFolderNames {
		for _, folder := range folders {
			if folder.Name == folderName {
				newFolders = append(newFolders, folder)
			}
		}
	}

	return newFolders
}