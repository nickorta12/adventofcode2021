package main

import (
	"fmt"
	"strconv"

	"github.com/nickorta12/adventofcode2021/aoc"
)

func zerosOnes(lines []string) ([]int, []int) {
	lineLen := len(lines[0])

	zeros := make([]int, lineLen)
	ones := make([]int, lineLen)

	for y, col := range aoc.Transpose(lines) {
		for _, char := range col {
			if char == '0' {
				zeros[y] += 1
			} else if char == '1' {
				ones[y] += 1
			} else {
				aoc.Fatal("Invalid char: %c", char)
			}
		}
	}

	return zeros, ones

}

func Part1(input string) int {
	lines := aoc.Lines(input)
	zeros, ones := zerosOnes(lines)
	lineLen := len(zeros)

	gamma := make([]rune, lineLen)
	epsilon := make([]rune, lineLen)

	for i := 0; i < lineLen; i++ {
		if ones[i] > zeros[i] {
			gamma[i] = '1'
			epsilon[i] = '0'
		} else {
			gamma[i] = '0'
			epsilon[i] = '1'
		}
	}

	gammaVal, err := strconv.ParseInt(string(gamma), 2, 64)
	if err != nil {
		aoc.Fatal("%s", err)
	}
	epsilonVal, err := strconv.ParseInt(string(epsilon), 2, 64)
	if err != nil {
		aoc.Fatal("%s", err)
	}

	return int(gammaVal * epsilonVal)
}

// func calcGravity(gravChars []rune, lines []string, transposed []string) string {
// 	validIndexes := make([]bool, len(lines))
// 	for i := 0; i < len(lines); i++ {
// 		validIndexes[i] = true
// 	}
// 	numValid := len(lines)

// 	for i, expectedChar := range gravChars {
// 		for j, char := range transposed[i] {
// 			if !validIndexes[j] {
// 				continue
// 			}
// 			validIndexes[j] = char == expectedChar
// 		}
// 	}
// 	return ""
// }

func Part2(input string) int {
	lines := aoc.Lines(input)
	// transposed := aoc.Transpose(lines)
	zeros, ones := zerosOnes(lines)
	lineLen := len(zeros)

	gravChars := make([]rune, lineLen)
	co2Chars := make([]rune, lineLen)

	for i := 0; i < lineLen; i++ {
		numOnes := ones[i]
		numZeros := zeros[i]

		var gravity, co2 rune

		if numOnes >= numZeros {
			gravity = '1'
			co2 = '0'
		} else {
			gravity = '0'
			co2 = '1'
		}

		gravChars[i] = gravity
		co2Chars[i] = co2
	}

	// gravity := calcGravity(gravChars)

	return 0
}

func main() {
	input := aoc.ReadInputFile("03")

	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}
