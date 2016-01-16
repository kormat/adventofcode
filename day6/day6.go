package main

import (
	"flag"
	"github.com/kormat/adventofcode/day6/common"
	"github.com/kormat/adventofcode/day6/p1"
	"github.com/kormat/adventofcode/day6/p2"
	"github.com/kormat/adventofcode/util"
	"os"
)

var p2flag = flag.Bool("2", false, "Calculate part 2 result")

func main() {
	flag.Parse()
	lines, ok := util.ReadFileArg(flag.Args())
	if !ok {
		os.Exit(1)
	}
	parseCommands(lines)
}

func parseCommands(lines []string) {
	var grid day6.Grid
	for i, line := range lines {
		cmd, coords := day6.ParseCommand(line)
		if *p2flag {
			p2.ExecCommand(i, cmd, coords, &grid)
		} else {
			p1.ExecCommand(i, cmd, coords, &grid)
		}
	}
}
