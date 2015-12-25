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
	total := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		total += calcPaper(scanner.Text())
	}
	fmt.Printf("The elves need %d sqft of wrapping paper.\n", total)
}

func calcPaper(arg string) int {
	var dims []int
	sqft := 0
	for _, c := range strings.Split(arg, "x") {
		i, _ := strconv.Atoi(c)
		dims = append(dims, i)
	}
	sort.Ints(dims)
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
