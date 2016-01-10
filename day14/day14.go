package main

import (
	"flag"
	"github.com/kormat/adventofcode/day14/lib"
	"github.com/kormat/adventofcode/util"
	"log"
	"os"
	"strconv"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		log.Fatalf("Expected two parameters, an input file, and a duration in seconds. Got: %s",
			flag.Args())
	}
	lines, ok := util.ReadFileArg(args[0:1])
	if !ok {
		os.Exit(1)
	}
	duration, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatalf("Unable to parse duration '%s': %s", args[1], err)
	}
	deers := day14.Parse(lines)
	day14.Race(deers, duration)
}
