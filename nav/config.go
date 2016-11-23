package nav

// Get a list of mappings from a folder name to a Package name. This is useful so that when we output filenames we
// use the "mapped" value instead of the original value
func getFolderNameToPackageNameMap() map[string]string {
	return map[string]string{
		"module-vpc": "network-topology",
	}
}

func getTopLevelFolderOrdering() []string {
	return []string{
		"introduction",
		"help",
		"packages",
	}
}