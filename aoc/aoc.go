package aoc

import (
	"os"
	"path"
	"runtime"
	"strings"
)

func ReadInputFile(day string) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller info")
	}
	dir := path.Dir(path.Dir(filename))
	input := path.Join(dir, "inputs", day+".txt")
	contents, err := os.ReadFile(input)

	if err != nil {
		panic("File not found: " + input)
	}

	return strings.TrimSpace(string(contents[:]))
}
