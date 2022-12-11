package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/nickorta12/adventofcode2021/aoc"
)

const main_template string = `package main

import (
	"fmt"

	"github.com/nickorta12/adventofcode2021/aoc"
)

func Part1(input string) int {
	return 0
}

func Part2(input string) int {
	return 0
}

func main() {
	input := aoc.ReadInputFile("%02d")

	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}
`

const test_template string = `package main

import (
	"testing"
)

const input string = ""

func TestPart1(t *testing.T) {
	expected := 0
	if ans := Part1(input); ans != expected {
		t.Errorf("Expected %d, got: %d", expected, ans)
	}
}

func TestPart2(t *testing.T) {
	expected := 0
	if ans := Part2(input); ans != expected {
		t.Errorf("Expected %d, got: %d", expected, ans)
	}
}
`

func pathExists(name string) bool {
	if _, err := os.Stat(name); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		aoc.Fatal("Error for path %s: %s", name, err)
	}

	return false
}

func main() {
	if len(os.Args) != 2 {
		aoc.Fatal("./adventofcode2021 DAY")
	}
	day, err := strconv.Atoi(os.Args[1])
	dayStr := fmt.Sprintf("%02d", day)
	if err != nil {
		aoc.Fatal("Error: %s", err)
	}

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		aoc.Fatal("No caller info")
	}
	dir := filepath.Dir(filename)
	binDir := filepath.Join(dir, "days", dayStr)

	inputPath := filepath.Join(dir, "inputs", dayStr+".txt")
	binPath := filepath.Join(binDir, dayStr+".go")
	testPath := filepath.Join(binDir, dayStr+"_test.go")

	if pathExists(binPath) || pathExists(testPath) {
		aoc.Fatal("Code for day %d already exists", day)
	}

	if pathExists(inputPath) {
		aoc.Fatal("Input for day %d already exists", day)
	}

	if !pathExists(binDir) {
		if err = os.Mkdir(binDir, 0755); err != nil {
			aoc.Fatal("Making bin dir %s: %s", binDir, err)
		}
	}

	cmd := exec.Command("aoc", "-y", "2021", "-d", fmt.Sprint(), "-i", inputPath, "download")

	if err = cmd.Run(); err != nil {
		aoc.Fatal("Running aoc command for input files: %s", err)
	}

	if err = os.WriteFile(binPath, []byte(fmt.Sprintf(main_template, day)), 0644); err != nil {
		aoc.Fatal("Making main file %s: %s", binPath, err)
	}
	if err = os.WriteFile(testPath, []byte(test_template), 0644); err != nil {
		aoc.Fatal("Making test file %s: %s", testPath, err)
	}

}
