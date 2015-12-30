package main

import (
	"fmt"
	"github.com/kormat/adventofcode/util"
	"log"
	"os"
	"regexp"
	"strconv"
)

const GRID_SIZE = 1000
const (
	CMD_UNKNOWN = iota
	CMD_TURN_ON
	CMD_TURN_OFF
	CMD_TOGGLE
)

type Coords struct {
	x1, y1, x2, y2 int
}
type Grid [GRID_SIZE][GRID_SIZE]bool

func main() {
	lines, err := util.ReadFileArg()
	if err {
		os.Exit(1)
	}
	parseCommands(lines)
	//fmt.Printf("Summary: %d commands, %d lights on\n", len(lines), result)
}

func parseCommands(lines []string) int {
	var grid Grid
	for i, line := range lines {
		cmd, coords := parseCommand(line)
		execCommand(cmd, coords, &grid)
		fmt.Printf("%d. Lights on: %d\n", i, countLights(grid))
	}
	return 0
}

func parseCommand(line string) (int, Coords) {
	cmd, coords := parseCmdTurn(line)
	if cmd == CMD_UNKNOWN {
		cmd, coords = parseCmdToggle(line)
	}
	if cmd == CMD_UNKNOWN {
		log.Fatal("Unable to parse command: '%s'\n", line)
	}
	return cmd, coords
}

func parseCmdTurn(line string) (int, Coords) {
	rx := regexp.MustCompile(`turn (on|off) (\d+),(\d+) through (\d+),(\d+)`)
	cmd := CMD_UNKNOWN
	result := rx.FindStringSubmatch(line)
	if result == nil {
		return cmd, Coords{}
	}
	if result[1] == "on" {
		cmd = CMD_TURN_ON
	} else {
		cmd = CMD_TURN_OFF
	}
	return cmd, parseCoords(result[2:])
}

func parseCmdToggle(line string) (int, Coords) {
	rx := regexp.MustCompile(`toggle (\d+),(\d+) through (\d+),(\d+)`)
	cmd := CMD_UNKNOWN
	result := rx.FindStringSubmatch(line)
	if result == nil {
		return cmd, Coords{}
	}
	cmd = CMD_TOGGLE
	return cmd, parseCoords(result[1:])
}

func parseCoords(s []string) Coords {
	return Coords{
		parseInt(s[0]),
		parseInt(s[1]),
		parseInt(s[2]),
		parseInt(s[3]),
	}
}

func parseInt(c string) int {
	val, err := strconv.Atoi(c)
	if err != nil {
		log.Fatal("Unable to parse coordinate '%s'", c)
	}
	return val
}

func execCommand(cmd int, coords Coords, grid *Grid) {
	for x := coords.x1; x <= coords.x2; x++ {
		for y := coords.y1; y <= coords.y2; y++ {
			switch cmd {
			case CMD_TURN_ON:
				grid[x][y] = true
			case CMD_TURN_OFF:
				grid[x][y] = false
			case CMD_TOGGLE:
				grid[x][y] = !grid[x][y]
			}
		}
	}
}

func countLights(grid Grid) int {
	on := 0
	for x := 0; x < GRID_SIZE; x++ {
		for y := 0; y < GRID_SIZE; y++ {
			if grid[x][y] {
				on++
			}
		}
	}
	return on
}
