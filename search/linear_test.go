package search

import (
	"testing"
)

func TestLinear(t *testing.T) {
	array := []int{1, 2, 128, 299, 3, 4, 5, 356}

	tests := []struct {
		target  int
		wantIdx int
	}{
		{
			1,
			0,
		},
		{
			10,
			-1,
		},
		{
			299,
			3,
		},
		{
			0,
			-1,
		},
		{
			356,
			7,
		},
	}

	for i, tt := range tests {
		gotIdx := LinearSearch(array, tt.target)

		if gotIdx != tt.wantIdx {
			t.Errorf("tests[%d]: expected to find %d by index %d, got index %d", i, tt.target, tt.wantIdx, gotIdx)
		}
	}
}
