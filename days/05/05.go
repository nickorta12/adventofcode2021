package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nickorta12/adventofcode2021/aoc"
)

type coord struct {
	x, y int
}

func parseCoord(input string) (coord, error) {
	parts := strings.Split(input, ",")
	if len(parts) != 2 {
		return coord{}, fmt.Errorf("invalid coordinate string: %s", input)
	}

	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return coord{}, err
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return coord{}, err
	}

	return coord{x, y}, nil
}

func (c coord) equals(o coord) bool {
	return c.x == o.x && c.y == o.y
}

type line struct {
	start, end coord
}

func parseLine(input string) (line, error) {
	parts := strings.Split(input, " -> ")
	if len(parts) != 2 {
		return line{}, fmt.Errorf("invalid line: %s", input)
	}

	start, err := parseCoord(parts[0])
	if err != nil {
		return line{}, err
	}
	end, err := parseCoord(parts[1])
	if err != nil {
		return line{}, err
	}

	return line{start, end}, nil
}

func parseLines(input string) ([]line, error) {
	var err error
	rawLines := aoc.Lines(input)
	lines := make([]line, len(rawLines))
	for i, rawLine := range rawLines {
		lines[i], err = parseLine(rawLine)
		if err != nil {
			return nil, err
		}
	}

	return lines, nil
}

func (l *line) isStraight() bool {
	return (l.start.x == l.end.x) || (l.start.y == l.end.y)
}

func (l *line) coords() []coord {
	ret := make([]coord, 0)
	current := l.start

	for {
		ret = append(ret, current)
		if current.equals(l.end) {
			return ret
		}

		if current.x < l.end.x {
			current.x++
		} else if current.x > l.end.x {
			current.x--
		}

		if current.y < l.end.y {
			current.y++
		} else if current.y > l.end.y {
			current.y--
		}
	}

}

type grid struct {
	points map[coord]int
}

func newGrid(lines []line) grid {
	g := grid{}
	if g.points == nil {
		g.points = make(map[coord]int)
	}
	for _, line := range lines {
		for _, coord := range line.coords() {
			g.points[coord] += 1
		}
	}

	return g
}

func (g *grid) print() {
	xMax, yMax := 0, 0
	for coord := range g.points {
		if coord.x > xMax {
			xMax = coord.x
		}
		if coord.y > yMax {
			yMax = coord.y
		}
	}

	for y := 0; y < yMax+2; y++ {
		for x := 0; x < xMax+2; x++ {
			val, ok := g.points[coord{x, y}]
			if !ok {
				fmt.Printf("%s", ".")
			} else {
				fmt.Printf("%d", val)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g *grid) getOverOne() int {
	total := 0
	for _, v := range g.points {
		if v > 1 {
			total++
		}
	}

	return total
}

func Part1(input string) int {
	lines, err := parseLines(input)
	if err != nil {
		aoc.Fatal("%s", err)
	}
	straightLines := make([]line, 0)
	for _, line := range lines {
		if line.isStraight() {
			straightLines = append(straightLines, line)
		}
	}
	g := newGrid(straightLines)
	// g.print()

	return g.getOverOne()
}

func Part2(input string) int {
	lines, err := parseLines(input)
	if err != nil {
		aoc.Fatal("%s", err)
	}
	g := newGrid(lines)
	g.print()

	return g.getOverOne()
}

func main() {
	input := aoc.ReadInputFile("05")

	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}
