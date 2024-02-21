package sort

import (
	"slices"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	want := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	slices.Sort(want)

	got := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	BubbleSort(got)

	if !slices.Equal(got, want) {
		t.Errorf("expected sorted to be %v, got %v", want, got)
	}
}
