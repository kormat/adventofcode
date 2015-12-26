package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	lines := readFile()
	for i, line := range lines {
		floor, to_minus, err := calcFloor(line)
		if err {
			os.Exit(1)
		}
		fmt.Printf("%d. Santa ends up on floor %d, and entered the basement at instruction %d\n", i, floor, to_minus)
	}
}

func readFile() []string {
	var lines []string
	if len(os.Args) != 2 {
		log.Fatal("Expecting a single input file as argument")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func calcFloor(dirs string) (int, int, bool) {
	floor := 0
	to_minus := -1
	for i, c := range dirs {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		default:
			log.Printf("Unrecognised instruction '%c' at index %d", c, i)
			return 0, 0, true
		}
		if to_minus == -1 && floor < 0 {
			to_minus = i + 1
		}
	}
	return floor, to_minus, false
}
