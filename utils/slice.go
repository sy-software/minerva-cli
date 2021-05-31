package utils

func Contains(slice []string, val string) int {
	for index, item := range slice {
		if item == val {
			return index
		}
	}

	return -1
}
