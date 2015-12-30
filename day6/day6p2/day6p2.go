package main

import (
	"fmt"
	"github.com/kormat/adventofcode/day6"
	"github.com/kormat/adventofcode/util"
	"os"
)

type Grid [day6.GRID_SIZE][day6.GRID_SIZE]int

func main() {
	lines, err := util.ReadFileArg()
	if err {
		os.Exit(1)
	}
	parseCommands(lines)
}

func parseCommands(lines []string) int {
	var grid Grid
	for i, line := range lines {
		cmd, coords := day6.ParseCommand(line)
		execCommand(cmd, coords, &grid)
		fmt.Printf("%d. Total brightness: %d\n", i, countBrightness(grid))
	}
	return 0
}

func execCommand(cmd int, coords day6.Coords, grid *Grid) {
	for x := coords.X1; x <= coords.X2; x++ {
		for y := coords.Y1; y <= coords.Y2; y++ {
			switch cmd {
			case day6.CMD_TURN_ON:
				grid[x][y] += 1
			case day6.CMD_TURN_OFF:
				grid[x][y] -= 1
			case day6.CMD_TOGGLE:
				grid[x][y] += 2
			}
			if grid[x][y] < 0 {
				grid[x][y] = 0
			}
		}
	}
}

func countBrightness(grid Grid) int {
	total := 0
	for x := 0; x < day6.GRID_SIZE; x++ {
		for y := 0; y < day6.GRID_SIZE; y++ {
			total += grid[x][y]
		}
	}
	return total
}
