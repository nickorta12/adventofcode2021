package aoc

import (
	"testing"
)

func TestTranspose(t *testing.T) {
	input := []string{"abcde", "12345", "qwert"}
	expected := []string{"a1q", "b2w", "c3e", "d4r", "e5t"}
	transposed := Transpose(input)

	for i, line := range expected {
		if line != transposed[i] {
			t.Errorf("Expected %s, got %s", line, transposed[i])
		}
	}
}
