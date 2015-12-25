package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	paper := 0
	ribbon := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		dims := parseDims(scanner.Text())
		paper += calcPaper(dims)
		ribbon += calcRibbon(dims)
	}
	fmt.Printf("The elves need %d sqft of wrapping paper, and %d feet of ribbon.\n", paper, ribbon)
}

func parseDims(arg string) []int {
	var dims []int
	for _, c := range strings.Split(arg, "x") {
		i, _ := strconv.Atoi(c)
		dims = append(dims, i)
	}
	sort.Ints(dims)
	return dims
}

func calcPaper(dims []int) int {
	sqft := 0
	for i, val1 := range dims {
		if i == len(dims) {
			break
		}
		for _, val2 := range dims[i+1:] {
			sqft += 2 * val1 * val2
		}
	}
	sqft += dims[0] * dims[1]
	return sqft
}

func calcRibbon(dims []int) int {
	feet := 0
	feet += 2 * (dims[0] + dims[1])
	feet += dims[0] * dims[1] * dims[2]
	return feet
}
