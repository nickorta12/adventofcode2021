package main

import (
	"testing"
)

const input string = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

func TestPart1(t *testing.T) {
	expected := 198
	if ans := Part1(input); ans != expected {
		t.Errorf("Expected %d, got: %d", expected, ans)
	}
}

func TestPart2(t *testing.T) {
	expected := 230
	if ans := Part2(input); ans != expected {
		t.Errorf("Expected %d, got: %d", expected, ans)
	}
}
