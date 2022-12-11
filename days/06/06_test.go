package main

import (
	"testing"
)

const input string = "3,4,3,1,2"

func TestPart1(t *testing.T) {
	expected := uint64(5934)
	if ans := Part1(input); ans != expected {
		t.Errorf("Expected %d, got: %d", expected, ans)
	}
}

func TestPart2(t *testing.T) {
	expected := uint64(26984457539)
	if ans := Part2(input); ans != expected {
		t.Errorf("Expected %d, got: %d", expected, ans)
	}
}
