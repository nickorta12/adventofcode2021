package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nickorta12/adventofcode2021/aoc"
)

type Sub struct {
	x, y, aim int
}

func (s *Sub) parseLine(line string) {
	parts := strings.Split(line, " ")
	dir := parts[0]
	num, err := strconv.Atoi(parts[1])
	if err != nil {
		aoc.Fatal("Invalid line: %s\n%s", line, err)
	}

	if dir == "forward" {
		s.x += num
	} else if dir == "down" {
		s.y += num
	} else if dir == "up" {
		s.y -= num
	} else {
		aoc.Fatal("Invalid direction: %s", dir)
	}
}

func (s *Sub) parseLineAim(line string) {
	parts := strings.Split(line, " ")
	dir := parts[0]
	num, err := strconv.Atoi(parts[1])
	if err != nil {
		aoc.Fatal("Invalid line: %s\n%s", line, err)
	}

	if dir == "forward" {
		s.x += num
		s.y += num * s.aim
	} else if dir == "down" {
		s.aim += num
	} else if dir == "up" {
		s.aim -= num
	} else {
		aoc.Fatal("Invalid direction: %s", dir)
	}
}

func (s *Sub) final() int {
	return s.x * s.y
}

func Part1(input string) int {
	sub := Sub{}

	for _, line := range strings.Split(input, "\n") {
		sub.parseLine(line)
	}

	return sub.final()
}

func Part2(input string) int {
	sub := Sub{}

	for _, line := range strings.Split(input, "\n") {
		sub.parseLineAim(line)
	}

	return sub.final()
}

func main() {
	input := aoc.ReadInputFile("02")

	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}
