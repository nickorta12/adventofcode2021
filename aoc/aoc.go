package aoc

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
)

func ReadInputFile(day string) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		Fatal("No caller info")
	}
	dir := path.Dir(path.Dir(filename))
	input := path.Join(dir, "inputs", day+".txt")
	contents, err := os.ReadFile(input)

	if err != nil {
		Fatal("File not found: " + input)
	}

	return strings.TrimSpace(string(contents[:]))
}

func Fatal(format string, s ...any) {
	fmt.Printf(format+"\n", s...)
	os.Exit(1)
}

func Lines(input string) []string {
	return strings.Split(input, "\n")
}

func Transpose(slice []string) []string {
	xl := len(slice[0])
	yl := len(slice)
	intermediate := make([][]rune, xl)

	for i := 0; i < xl; i++ {
		intermediate[i] = make([]rune, yl)
	}

	for i, line := range slice {
		for j, char := range line {
			intermediate[j][i] = char
		}
	}

	result := make([]string, xl)
	for i, chars := range intermediate {
		result[i] = string(chars)
	}
	return result
}
