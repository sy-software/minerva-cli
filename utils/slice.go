package utils

// Contains checks if a slice of strings contains a given value
func Contains(slice []string, val string) bool {
	return IndexOf(slice, val) != -1
}

// IndexOf returns the index of a value in a string slice. If the value is not present returns -1
func IndexOf(slice []string, val string) int {
	for index, item := range slice {
		if item == val {
			return index
		}
	}

	return -1
}
