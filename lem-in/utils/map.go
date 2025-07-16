package utils

// this is used so i don't pass by refrence
func CopyMap(original map[string]bool) map[string]bool {
	newMap := make(map[string]bool)
	for key, value := range original {
		newMap[key] = value
	}
	return newMap
}

func CopySliceSlice(paths [][]string) [][]string {
	use := make([][]string, len(paths))

	for i := range paths {
		use[i] = make([]string, len(paths[i]))
		copy(use[i], paths[i])
	}

	return use
}
