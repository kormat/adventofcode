package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode/utf8"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		dirs := scanner.Text()
		year1 := parseDirsY1(dirs)
		year2 := parseDirsY2(dirs)
		fmt.Printf("From %d directions, %d houses were visited in Year 1, %d in Year 2.\n", len(dirs), year1, year2)
	}
}

func parseDir(c rune, x, y *int) {
	switch c {
	case '^':
		*y++
	case 'v':
		*y--
	case '<':
		*x--
	case '>':
		*x++
	}
}

func parseDirsY1(dirs string) int {
	houses := make(map[[2]int]bool)
	x, y := 0, 0
	houses[[2]int{x, y}] = true
	for _, c := range dirs {
		parseDir(c, &x, &y)
		houses[[2]int{x, y}] = true
	}
	return len(houses)
}

func parseDirsY2(dirs string) int {
	houses := make(map[[2]int]bool)
	sx, sy := 0, 0
	rx, ry := 0, 0
	houses[[2]int{sx, sy}] = true
	for i := 0; i < len(dirs); {
		dir, width := utf8.DecodeRuneInString(dirs[i:])
		parseDir(dir, &sx, &sy)
		houses[[2]int{sx, sy}] = true
		i += width
		dir, width = utf8.DecodeRuneInString(dirs[i:])
		parseDir(dir, &rx, &ry)
		houses[[2]int{rx, ry}] = true
		i += width
	}
	return len(houses)
}
