package main

import (
	"fmt"
	"github.com/kormat/adventofcode/util"
	"os"
)

func main() {
	lines, ok := util.ReadFileArg(os.Args[1:])
	if !ok {
		os.Exit(1)
	}
	nice := 0
	naughty := 0
	for i, word := range lines {
		result := isStringNice(word)
		if result {
			nice++
		} else {
			naughty++
		}
		fmt.Printf("%d. %s: %v\n", i, word, result)
	}
	fmt.Printf("Summary: %d words, %d nice, %d naughty\n", len(lines), nice, naughty)
}

func isStringNice(chars string) bool {
	if hasValidPairs(chars) && hasValidRepeat(chars) {
		return true
	}
	return false
}

func hasValidPairs(chars string) bool {
	pairs := make(map[string][]int)
	for i := 0; i < len(chars)-1; i++ {
		pair := chars[i : i+2]
		pairs[pair] = append(pairs[pair], i)
	}
	for _, idxs := range pairs {
		if len(idxs) < 2 {
			continue
		}
		first := idxs[0]
		for _, second := range idxs[1:] {
			if second > first+1 {
				return true
			}
		}
	}
	return false
}

func hasValidRepeat(chars string) bool {
	for i := 0; i < len(chars)-2; i++ {
		if chars[i] == chars[i+2] {
			return true
		}
	}
	return false
}
