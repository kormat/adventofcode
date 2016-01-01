package main

import (
	"fmt"
	"github.com/kormat/adventofcode/util"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines, err := util.ReadFileArg(os.Args[1:])
	if err {
		os.Exit(1)
	}
	paper := 0
	ribbon := 0

	for _, line := range lines {
		dims, err := parseDims(line)
		if err {
			os.Exit(1)
		}
		paper += calcPaper(dims)
		ribbon += calcRibbon(dims)
	}
	fmt.Printf("The elves need %d sqft of wrapping paper, and %d feet of ribbon.\n", paper, ribbon)
}

func parseDims(arg string) ([]int, bool) {
	var dims []int
	args := strings.Split(arg, "x")
	if len(args) != 3 {
		log.Printf("Dimensions not in expected NxNxN format: %s", arg)
		return dims, true
	}
	for _, c := range args {
		i, err := strconv.Atoi(c)
		if err != nil {
			log.Printf("Unable to parse '%s' as integer", c)
			return dims, true
		}
		dims = append(dims, i)
	}
	sort.Ints(dims)
	return dims, false
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
	feet := 2 * (dims[0] + dims[1])
	feet += dims[0] * dims[1] * dims[2]
	return feet
}
