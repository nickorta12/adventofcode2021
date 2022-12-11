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
