package main

import (
	"testing"
)

func TestFoo(t *testing.T) {

	str := "3[a]2[bc]"

	result := encodedString(str)

	if len(result) == 0 {
		t.Fatal()
	}
}
