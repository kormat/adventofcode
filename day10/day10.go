package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

var cycles = flag.Int("n", 5, "number of times to run the cycle")

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Fatalf("Expecting a single seed as cmdline argument (got %v instead)", args)
	}
	var value []int
	for _, c := range args[0] {
		val, err := strconv.Atoi(string(c))
		if err != nil {
			log.Fatalf("Unable to parse char '%c' as a number: %v", c, err)
		}
		value = append(value, val)
	}
	fmt.Printf("Starting value: %v\n", value)
	for i := 0; i < *cycles; i++ {
		value = lookAndSay(value)
		fmt.Printf("%d. Len: %d\n", i, len(value))
	}
}

func lookAndSay(val []int) []int {
	var out []int
	for i := 0; i < len(val); i++ {
		curr := val[i]
		count := 1
		for j := i + 1; j < len(val); j++ {
			if val[j] != curr {
				break
			}
			count++
			i++
		}
		out = append(out, count)
		out = append(out, curr)
	}
	return out
}
