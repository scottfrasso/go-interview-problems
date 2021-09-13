package main

import (
"testing"
)

func TestRacecarPath1(t *testing.T) {

	type raceTest struct {
		target int
		expected int
	}

	raceTests := []raceTest{
		{3, 2},
		{ 6, 3},
	}

	for _, tt := range raceTests {
		actual := racecar(tt.target)
		if actual != tt.expected {
			t.Errorf(`Expected %d but got %d`, tt.expected, actual)
		}
	}
}