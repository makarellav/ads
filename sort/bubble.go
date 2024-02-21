package sort

import "cmp"

func BubbleSort[T cmp.Ordered](array []T) {
	var swapped bool

	for i := 0; i < len(array); i++ {
		swapped = false

		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				swapped = true
			}
		}

		if !swapped {
			return
		}
	}
}
