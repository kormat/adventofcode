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

const maxL = 150

func main() {
	flag.Parse()
	args := flag.Args()
	lines, ok := util.ReadFileArg(args[:1])
	if !ok {
		os.Exit(1)
	}
	cntrs := parseContainers(lines)
	fmt.Printf("Containers: %v\n", cntrs)
	findCombinations(cntrs)
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

func findCombinations(cntrs []int) {
	used := []int{}
	avail := make([]int, len(cntrs))
	copy(avail, cntrs)
	fmt.Printf("used: %v. avail: %v\n", used, avail)
}

func findCombo(used, avail []int) {
	curr := sumCntrs(used)
	fmt.Printf("Curr: %d Used: %v Avail: %v\n", curr, used, avail)
	new_used := make([]int, len(used)+1)
	for len(avail) > 0 {
		next := avail[0]
		avail = avail[1:]
		new_used[len(used)] = next
		switch {
		case curr+next == maxL:
			fmt.Printf("Found match: %v\n", new_used)
		case curr+next > maxL:
			break
		default:
			findCombo(new_used, avail)
		}
	}
}

func sumCntrs(cntrs []int) int {
	var total int
	for _, c := range cntrs {
		total += c
	}
	return total
}
