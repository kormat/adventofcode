package main

import (
	"fmt"
	"github.com/kormat/adventofcode/day5/p1"
	"github.com/kormat/adventofcode/day5/p2"
	"github.com/kormat/adventofcode/util"
	"os"
)

func main() {
	lines, ok := util.ReadFileArg(os.Args[1:])
	if !ok {
		os.Exit(1)
	}
	nice, naughty := process(lines, p1.IsStringNice)
	fmt.Printf("Part 1: %d words, %d nice, %d naughty\n", len(lines), nice, naughty)
	nice, naughty = process(lines, p2.IsStringNice)
	fmt.Printf("Part 2: %d words, %d nice, %d naughty\n", len(lines), nice, naughty)
}

func process(lines []string, f func(word string) bool) (int, int) {
	nice := 0
	naughty := 0
	for _, word := range lines {
		result := f(word)
		if result {
			nice++
		} else {
			naughty++
		}
	}
	return nice, naughty
}
