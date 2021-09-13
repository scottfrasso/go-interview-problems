package main

import "testing"

func TestLongestMatrixPath(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3, 4, 5},
		{2, 3, 4, 4, 4},
		{2, 3, 4, 4, 4},
		{2, 3, 4, 4, 4},
		{2, 3, 4, 4, 4},
	}

	longestPath := longestMatrixPath(matrix)
	if longestPath != 5 {
		t.Fatalf(`Expected longest path to be 5 but got %d`, longestPath)
	}
}