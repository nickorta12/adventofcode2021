package main

import (
	"testing"
)

const input string = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

func TestPart1(t *testing.T) {
	expected := 5
	if ans := Part1(input); ans != expected {
		t.Errorf("Expected %d, got: %d", expected, ans)
	}
}

func TestPart2(t *testing.T) {
	expected := 12
	if ans := Part2(input); ans != expected {
		t.Errorf("Expected %d, got: %d", expected, ans)
	}
}
