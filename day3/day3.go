package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		count := parseDirs(dirs)
		fmt.Printf("%d houses visited over %d deliveries\n", count, len(dirs)+1)
	}
}

func parseDirs(dirs string) int {
	houses := make(map[[2]int]int)
	x, y := 0, 0
	houses[[2]int{x, y}] += 1
	for _, c := range dirs {
		switch c {
		case '^':
			y++
		case 'v':
			y--
		case '<':
			x--
		case '>':
			x++
		}
		houses[[2]int{x, y}] += 1
	}
	return len(houses)
}
