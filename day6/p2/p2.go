package p2

import (
	"fmt"
	"github.com/kormat/adventofcode/day6/common"
)

func ExecCommand(line, cmd int, coords day6.Coords, grid *day6.Grid) {
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
	fmt.Printf("%d. Brightness: %d\n", line, countBrightness(*grid))
}

func countBrightness(grid day6.Grid) int {
	total := 0
	for x := 0; x < day6.GRID_SIZE; x++ {
		for y := 0; y < day6.GRID_SIZE; y++ {
			total += grid[x][y]
		}
	}
	return total
}
