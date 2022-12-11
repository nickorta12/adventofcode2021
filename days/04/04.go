package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nickorta12/adventofcode2021/aoc"
)

type Moves []int

func parseMoves(line string) Moves {
	chars := strings.Split(line, ",")
	ret := make([]int, len(chars))

	for i, char := range chars {
		num, err := strconv.Atoi(char)
		if err != nil {
			aoc.Fatal("Not a number: %s", char)
		}
		ret[i] = num
	}

	return ret
}

type Card struct {
	numbers  [25]int
	selected [25]bool
}

func (c *Card) callMove(move int) {
	for i, num := range c.numbers {
		if move == num {
			c.selected[i] = true
		}
	}
}

func (c *Card) checkIfWon() bool {
	rowCount := 0
	var colCount [5]int

	for i, s := range c.selected {
		x := i % 5
		if x == 0 {
			rowCount = 0
		}
		if s {
			rowCount++
			colCount[x]++
		}

		if rowCount == 5 || colCount[x] == 5 {
			return true
		}
	}

	return false
}

// func (c *Card) print() {
// 	bold := color.New(color.FgRed).SprintFunc()
// 	var char string
// 	for i, num := range c.numbers {
// 		char = fmt.Sprintf("%2d", num)
// 		if c.selected[i] {
// 			char = bold(char)
// 		}

// 		if (i > 0) && ((i+1)%5 == 0) {
// 			fmt.Println(char)
// 		} else {
// 			fmt.Printf("%s ", char)
// 		}

// 	}

// 	fmt.Println()

// }

func (c *Card) sumUnmarked() int {
	total := 0
	for i, s := range c.selected {
		if !s {
			total += c.numbers[i]
		}
	}

	return total
}

func parseCard(lines []string) Card {
	var ret [25]int
	i := 0
	for _, line := range lines {
		for _, s := range strings.Split(line, " ") {
			if len(s) == 0 {
				continue
			}

			num, err := strconv.Atoi(s)
			if err != nil {
				aoc.Fatal("Not a number: %s, line: %s", s, line)
			}

			ret[i] = num
			i++
		}
	}

	return Card{numbers: ret}
}

func parseCards(lines []string) []Card {
	cards := make([]Card, 0)
	var cardLines [5]string
	cardLineIndex := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			cards = append(cards, parseCard(cardLines[:]))
			cardLineIndex = 0
			continue
		}
		cardLines[cardLineIndex] = line
		cardLineIndex++
	}

	cards = append(cards, parseCard(cardLines[:]))

	return cards
}

func Part1(input string) int {
	lines := aoc.Lines(input)
	moves := parseMoves(lines[0])
	cards := parseCards(lines[2:])

	for _, move := range moves {
		for i := 0; i < len(cards); i++ {
			card := &cards[i]
			card.callMove(move)
			if card.checkIfWon() {
				return move * card.sumUnmarked()
			}
		}
	}

	return 0
}

func Part2(input string) int {
	lines := aoc.Lines(input)
	moves := parseMoves(lines[0])
	cards := parseCards(lines[2:])

	cardsHaveWon := make([]bool, len(cards))
	cardsRemaining := len(cards)

	for _, move := range moves {
		for i := 0; i < len(cards); i++ {
			if cardsHaveWon[i] {
				continue
			}
			card := &cards[i]
			card.callMove(move)
			if card.checkIfWon() {
				cardsHaveWon[i] = true
				cardsRemaining--
			}

			if cardsRemaining == 0 {
				return move * card.sumUnmarked()
			}
		}
	}
	return 0
}

func main() {
	input := aoc.ReadInputFile("04")

	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}
