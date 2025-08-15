package main

import "testing"

func TestFactorial(t *testing.T) {
	input := 5
	expected := 120

	got := factorial(input)

	if got != expected {
		t.Errorf("got %v, want %v", got, expected)
	}
}
