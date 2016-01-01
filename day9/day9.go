package main

import (
	"flag"
	"fmt"
	"github.com/kormat/adventofcode/day9/lib"
	"github.com/kormat/adventofcode/util"
	"os"
)

var shortest = flag.Bool("short", true, "find shortest path")

func main() {
	flag.Parse()
	lines, err := util.ReadFileArg(flag.Args())
	if err {
		os.Exit(1)
	}
	places := make(day9.Places)
	places.ParseLines(lines)
	route, dist := places.FindBestRoute(*shortest)
	fmt.Printf("Best route is %d, via %s\n", dist, route)
}
