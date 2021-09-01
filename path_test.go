package main

import (
	"testing"
)

func TestFastestPath3(t *testing.T) {

	path := make([]int,0)
	fastestPath := followPath(3, path)

	expectedPath := []int{1,2}

	if len(fastestPath) != len(expectedPath) {
		t.Fatalf(`Expected %v but got %v`, expectedPath, fastestPath)
	}

	for i, v := range expectedPath {
		if v != fastestPath[i] {
			t.Fatalf(`Expected %v but got %v`, expectedPath, fastestPath)
		}
	}
}
