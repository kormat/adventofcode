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
	process(1, lines, p1.IsStringNice)
	process(2, lines, p2.IsStringNice)
}

func process(part int, lines []string, f func(word string) bool) {
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
	fmt.Printf("Part %v: %d words, %d nice, %d naughty\n", part, len(lines), nice, naughty)
}
