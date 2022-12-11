package main

import (
	"testing"
)

const input string = `forward 5
down 5
forward 8
up 3
down 8
forward 2`

func TestPart1(t *testing.T) {
	expected := 150
	if ans := Part1(input); ans != expected {
		t.Errorf("Expected %d, got: %d", expected, ans)
	}
}

func TestPart2(t *testing.T) {
	expected := 900
	if ans := Part2(input); ans != expected {
		t.Errorf("Expected %d, got: %d", expected, ans)
	}
}
