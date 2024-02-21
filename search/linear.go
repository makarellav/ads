package search

func LinearSearch[T comparable](array []T, target T) int {
	for i, v := range array {
		if v == target {
			return i
		}
	}

	return -1
}
