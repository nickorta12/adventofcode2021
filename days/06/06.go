package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nickorta12/adventofcode2021/aoc"
)

type aquarium struct {
	fishes map[uint8]uint64
}

func newAquarium(input string) aquarium {
	initialTimers := strings.Split(input, ",")
	fishes := make(map[uint8]uint64, 8)
	for _, timer := range initialTimers {
		timerVal, err := strconv.Atoi(timer)
		if err != nil {
			aoc.Fatal("Not a valid int: %s", timerVal)
		}
		fishes[uint8(timerVal)]++
	}

	return aquarium{fishes}
}

func (a *aquarium) tick() {
	newFishes := make(map[uint8]uint64, 8)
	for k, v := range a.fishes {
		if k != 0 {
			newFishes[k-1] = v
		}
	}

	newFishes[6] += a.fishes[0]
	newFishes[8] += a.fishes[0]

	a.fishes = newFishes
}

func (a *aquarium) num() uint64 {
	total := uint64(0)
	for _, v := range a.fishes {
		total += v
	}

	return total
}

func Part1(input string) uint64 {
	a := newAquarium(input)
	for i := 0; i < 80; i++ {
		a.tick()
	}
	return a.num()
}

func Part2(input string) uint64 {
	a := newAquarium(input)
	for i := 0; i < 256; i++ {
		a.tick()
	}
	return a.num()
}

func main() {
	input := aoc.ReadInputFile("06")

	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}
