package main

import (
	"testing"
)

const input string = `199
200
208
210
200
207
240
269
260
263`

func TestPart1(t *testing.T) {
	if ans := Part1(input); ans != 7 {
		t.Errorf("Expected 7, got: %d", ans)
	}
}

func TestPart2(t *testing.T) {
	if ans := Part2(input); ans != 5 {
		t.Errorf("Expected 5, got: %d", ans)
	}
}
