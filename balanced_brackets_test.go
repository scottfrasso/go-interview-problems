package main

import "testing"

func TestBalancedBracketsFalse(t *testing.T) {
	brackets := []rune("{[(])}")

	if balancedBrackets(brackets) != false {
		t.Fatalf(`Expected %v to be unbalanced`, brackets)
	}
}

func TestBalancedBracketsTrue(t *testing.T) {
	brackets := []rune("{[]}")

	if balancedBrackets(brackets) != true {
		t.Fatalf(`Expected %v to be balanced`, brackets)
	}
}

func TestBalancedBracketsAlsoTrue(t *testing.T) {
	brackets := []rune("{{[[(())]]}}")

	if balancedBrackets(brackets) != true {
		t.Fatalf(`Expected %v to be balanced`, brackets)
	}
}