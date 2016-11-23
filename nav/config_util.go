package nav

import "strings"

// Given a Package Folder Name (e.g. "module-vpc") return the corresponding Package Name (e.g. "Network Topology")
func getOutputPackageFolderNameFromInputPackageFolderName(packageFolderName string) string {
	folderNameToPackageNameMap := getFolderNameToPackageNameMap()

	if packageName, ok := folderNameToPackageNameMap[packageFolderName]; ok {
		return packageName
	} else {
		return packageFolderName
	}
}

// Convert "network-topology" into "Network Topology"
func convertDashesToSpacesAndCapitalize(str string) string {
	str = strings.Replace(str, "-", " ", -1)
	str = strings.Title(str)
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