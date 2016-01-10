package p1

import (
	"strings"
)

const VOWELS = "aeiou"

var BANNED = [...]string{"ab", "cd", "pq", "xy"}

func IsStringNice(word string) bool {
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
