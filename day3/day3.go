package main

import (
	"fmt"
	"github.com/kormat/adventofcode/util"
	"log"
	"os"
	"unicode/utf8"
)

type Location struct {
	X int
	Y int
}

func main() {
	lines, err := util.ReadFileArg()
	if err {
		os.Exit(1)
	}

	for i, dirs := range lines {
		fmt.Printf("%d. Directions: %d. Houses visited:\n", i, len(dirs))
		year1 := parseDirsY1(dirs)
		fmt.Printf("  Year 1: %d\n", year1)
		if len(dirs)%2 != 0 {
			continue
		}
		year2 := parseDirsY2(dirs)
		fmt.Printf("  Year 2: %d\n", year2)
	}
}

func parseDir(c rune, loc *Location) {
	switch c {
	case '^':
		loc.Y++
	case 'v':
		loc.Y--
	case '<':
		loc.X--
	case '>':
		loc.X++
	default:
		log.Fatal("Invalid direction '%c'", c)
	}
}

func parseDirsY1(dirs string) int {
	houses := make(map[Location]bool)
	loc := Location{0, 0}
	houses[loc] = true
	for _, c := range dirs {
		parseDir(c, &loc)
		houses[loc] = true
	}
	return len(houses)
}

func parseDirsY2(dirs string) int {
	houses := make(map[Location]bool)
	sloc := Location{0, 0}
	rloc := Location{0, 0}
	houses[sloc] = true
	for i := 0; i < len(dirs); {
		moveY2(dirs, &i, &sloc)
		houses[sloc] = true
		moveY2(dirs, &i, &rloc)
		houses[rloc] = true
	}
	return len(houses)
}

func moveY2(dirs string, i *int, loc *Location) {
	dir, width := utf8.DecodeRuneInString(dirs[*i:])
	parseDir(dir, loc)
	*i += width
}
