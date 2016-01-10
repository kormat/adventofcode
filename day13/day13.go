package main

import (
	"flag"
	"fmt"
	"github.com/kormat/adventofcode/day13/lib"
	"github.com/kormat/adventofcode/util"
	"os"
)

var p2 = flag.Bool("2", false, "Calculate part 2 result")

func main() {
	flag.Parse()
	lines, ok := util.ReadFileArg(flag.Args())
	if !ok {
		os.Exit(1)
	}
	day13.Parse(lines, *p2)
	net := day13.Seating()
	fmt.Printf("Net happiness: %d\n", net)
}
