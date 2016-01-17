package main

import (
	"flag"
	"fmt"
	"github.com/kormat/adventofcode/util"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	flag.Parse()
	args := flag.Args()
	lines, ok := util.ReadFileArg(args[:1])
	if !ok {
		os.Exit(1)
	}
	cntrs := parseContainers(lines)
	fmt.Printf("Containers: %v\n", cntrs)
}

func parseContainers(lines []string) []int {
	var cntrs []int
	for i, line := range lines {
		cntrs = append(cntrs, parseInt(i, line, "container"))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(cntrs)))
	return cntrs
}

func parseInt(line int, input, desc string) int {
	ret, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("Unable to parse %s '%s' on line %d: %s", desc, input, line, err)
	}
	return ret
}
