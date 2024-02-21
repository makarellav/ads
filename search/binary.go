package search

import "cmp"

func BinarySearch[T ~[]K, K cmp.Ordered](array T, target K) int {
	left := 0
	right := len(array) - 1

	for left <= right {
		mid := (left + right) / 2

		if array[mid] > target {
			right = mid - 1
		} else if array[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
