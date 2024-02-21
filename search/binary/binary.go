package binary

import "cmp"

// {1, 2, 3, 4, 500, 1000, 1500}
//              -4          -6

func Search[T ~[]K, K cmp.Ordered](array T, target K) int {
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
