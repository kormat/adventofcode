package main

import (
	"fmt"
	"github.com/kormat/adventofcode/day6"
	"github.com/kormat/adventofcode/util"
	"os"
)

type Grid [day6.GRID_SIZE][day6.GRID_SIZE]bool

func main() {
	lines, ok := util.ReadFileArg(os.Args[1:])
	if !ok {
		os.Exit(1)
	}
	parseCommands(lines)
}

func parseCommands(lines []string) int {
	var grid Grid
	for i, line := range lines {
		cmd, coords := day6.ParseCommand(line)
		execCommand(cmd, coords, &grid)
		fmt.Printf("%d. Lights on: %d\n", i, countLights(grid))
	}
	return 0
}

func execCommand(cmd int, coords day6.Coords, grid *Grid) {
	for x := coords.X1; x <= coords.X2; x++ {
		for y := coords.Y1; y <= coords.Y2; y++ {
			switch cmd {
			case day6.CMD_TURN_ON:
				grid[x][y] = true
			case day6.CMD_TURN_OFF:
				grid[x][y] = false
			case day6.CMD_TOGGLE:
				grid[x][y] = !grid[x][y]
			}
		}
	}
}

func countLights(grid Grid) int {
	on := 0
	for x := 0; x < day6.GRID_SIZE; x++ {
		for y := 0; y < day6.GRID_SIZE; y++ {
			if grid[x][y] {
				on++
			}
		}
	}
	return on
}
