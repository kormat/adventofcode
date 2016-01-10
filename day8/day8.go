package main

import (
	"fmt"
	"github.com/kormat/adventofcode/util"
	"os"
	"strings"
)

func main() {
	lines, ok := util.ReadFileArg(os.Args[1:])
	if !ok {
		os.Exit(1)
	}
	var total_lit, total_mem, total_reenc int
	for _, line := range lines {
		total_lit += len(line)
		total_mem += countEscaped(line)
		total_reenc += countReencoded(line)
	}
	fmt.Printf("Lines: %d Literals: %d\n", len(lines), total_lit)
	fmt.Printf("Memory: %d Literal overhead: %d\n", total_mem, total_lit-total_mem)
	fmt.Printf("Reencoded: %d Overhead: %d\n", total_reenc, total_reenc-total_lit)
}

func countEscaped(line string) int {
	var chars int
	for i := 1; i < len(line)-1; {
		s := line[i:]
		switch {
		case strings.HasPrefix(s, `\\`), strings.HasPrefix(s, `\"`):
			chars++
			i += 2
		case strings.HasPrefix(s, `\x`):
			chars++
			i += 4
		default:
			chars++
			i += 1
		}
	}
	return chars
}

func countReencoded(line string) int {
	chars := 3 // Opening "\"
	for i := 1; i < len(line)-1; i++ {
		switch line[i] {
		case '"', '\\': // Become \" and \\ respectively
			chars += 2
		default:
			chars += 1
		}
	}
	chars += 3 // Closing \""
	return chars
}
