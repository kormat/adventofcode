package day6

import (
	"log"
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

type Grid [GRID_SIZE][GRID_SIZE]int

type Coords struct {
	X1, Y1, X2, Y2 int
}

func ParseCommand(line string) (int, Coords) {
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
