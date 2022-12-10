package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nickorta12/adventofcode2021/aoc"
)

func linesToInts(input string) []int {
	depths := make([]int, 0)

	for _, line := range strings.Split(input, "\n") {
		depth, err := strconv.Atoi(line)
		if err != nil {
			panic("Not an int: " + line)
		}

		depths = append(depths, depth)
	}

	return depths
}

func Part1(input string) int {
	var prev, total int = 0, 0
	for _, current := range linesToInts(input) {
		if prev != 0 && current > prev {
			total++
		}

		prev = current
	}

	return total
}

func Part2(input string) int {
	var prev, total int = 0, 0

	depths := linesToInts(input)
	for i := 0; i < len(depths)-2; i++ {
		current := depths[i] + depths[i+1] + depths[i+2]
		if prev != 0 && current > prev {
			total++
		}

		prev = current
	}

	return total

}

func main() {
	input := aoc.ReadInputFile("01")

	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}
