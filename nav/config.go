package nav

import "strings"

// Get a list of mappings from a folder name to a Package name. This is useful so that when we output filenames we
// use the "mapped" value instead of the original value
func getFolderNameToPackageNameMap() map[string]string {
	return map[string]string{
		"module-vpc": "network-topology",
	}
}

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