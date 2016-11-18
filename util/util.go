package util

// Return true if `slice` contains the given `strToFind`; missing generics about now...
func StrSliceContains(slice []string, strToFind string) bool {
	for _, str := range slice {
		if str == strToFind {
			return true
		}
	}
	return false
}

// Given [x y z] return the [y z], the tail.
func GetStrSliceTail(slice []string) []string {
	return slice[1:]
}
