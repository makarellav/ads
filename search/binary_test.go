package search

import (
	"slices"
	"testing"
)

func TestBinary(t *testing.T) {
	array := []int{1, 2, 3, 4, 500, 1000, 1500}

	tests := []struct {
		target  int
		wantIdx int
	}{
		{
			1,
			0,
		},
		{
			500,
			4,
		},
		{
			69420,
			-1,
		},
		{
			1500,
			6,
		},
	}

	for i, tt := range tests {
		wantIdx, ok := slices.BinarySearch(array, tt.target)
		gotIdx := BinarySearch(array, tt.target)

		if !ok && gotIdx != -1 {
			t.Errorf("tests[%d]: expected not to find %d, found it by index %d", i, tt.target, gotIdx)
		}

		if ok && wantIdx != gotIdx {
			t.Errorf("tests[%d]: expected to find %d by index %d, got index %d", i, tt.target, tt.wantIdx, gotIdx)
		}
	}
}
