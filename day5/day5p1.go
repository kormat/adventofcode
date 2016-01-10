package main

import (
	"fmt"
	"github.com/kormat/adventofcode/util"
	"os"
	"strings"
)

const VOWELS = "aeiou"

var BANNED = [...]string{"ab", "cd", "pq", "xy"}

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

func isStringNice(word string) bool {
	vowels := 0
	dups := false
	for i := 0; i < len(word); i++ {
		checkVowel(word[i], &vowels)
		if i == len(word)-1 {
			break
		}
		pair := word[i : i+2]
		checkDups(pair, &dups)
		if checkBanned(pair) {
			return false
		}
	}
	if vowels >= 3 && dups {
		return true
	}
	return false
}

func checkVowel(c byte, vowels *int) {
	if strings.Contains(VOWELS, string(c)) {
		*vowels += 1
	}
}

func checkDups(chars string, dups *bool) {
	if chars[0] == chars[1] {
		*dups = true
	}
}

func checkBanned(chars string) bool {
	for _, b := range BANNED {
		if b == chars {
			return true
		}
	}
	return false
}
