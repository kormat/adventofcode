package main

import (
	"flag"
	"fmt"
	"github.com/kormat/adventofcode/util"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Sue map[string]int

var sues map[int]*Sue
var p2 = flag.Bool("2", false, "Calculate part 2 result")
var attrs = [...]string{
	"children", "cats", "samoyeds", "pomeranians", "akitas",
	"vizslas", "goldfish", "trees", "cars", "perfumes",
}

func main() {
	flag.Parse()
	args := flag.Args()
	sue_lines, ok := util.ReadFileArg(args[:1])
	if !ok {
		os.Exit(1)
	}
	msg_lines, ok := util.ReadFileArg(args[1:2])
	if !ok {
		os.Exit(1)
	}
	sues = make(map[int]*Sue)
	parseSues(sue_lines)
	msg := parseMsg(msg_lines)
	matchSue(msg)
}

func parseSues(lines []string) {
	for i, line := range lines {
		id_data := strings.SplitN(line, ": ", 2)
		id := parseId(i, id_data[0])
		dataMap := parseData(i, strings.Split(id_data[1], ", "))
		sues[id] = newSue(dataMap)
	}
}

func parseId(line int, id_s string) int {
	rx := regexp.MustCompile(`^Sue (\d+)$`)
	result := rx.FindStringSubmatch(id_s)
	if result == nil {
		log.Fatalf("Unable to parse id on line %d: %s", line, id_s)
	}
	return parseInt(line, result[1], "ID")
}

func parseData(line int, datas []string) map[string]int {
	m := make(map[string]int)
	rx := regexp.MustCompile(`^(\w+): (\d+)$`)
	for _, data := range datas {
		result := rx.FindStringSubmatch(data)
		if result == nil {
			log.Fatalf("Unable to parse id on line %d: %s", line, data)
		}
		name := result[1]
		m[name] = parseInt(line, result[2], name)
	}
	return m
}

func parseInt(line int, input, desc string) int {
	ret, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("Unable to parse %s '%s' on line %d: %s", desc, input, line, err)
	}
	return ret
}

func newSue(data map[string]int) *Sue {
	s := make(Sue)
	for k, v := range data {
		s[k] = v
	}
	return &s
}

func (s Sue) String() string {
	var tmp []string
	for _, k := range attrs {
		val, ok := s[k]
		if ok {
			tmp = append(tmp, fmt.Sprintf("  %s: %d", k, val))
		}
	}
	return strings.Join(tmp, "\n")
}

func parseMsg(lines []string) *Sue {
	rx := regexp.MustCompile(`^(\w+): (\d+)$`)
	m := make(map[string]int)
	for i, line := range lines {
		result := rx.FindStringSubmatch(line)
		if result == nil {
			log.Fatalf("Unable to parse attribute on line %d: %s", i, line)
		}
		name := result[1]
		val := parseInt(i, result[2], "value")
		m[name] = val
	}
	return newSue(m)
}

func matchSue(msg *Sue) {
	for id := 1; id <= 500; id++ {
		sue := sues[id]
		matched := true
		for _, k := range attrs {
			msg_val := (*msg)[k]
			sue_val, ok := (*sue)[k]
			if !ok {
				continue
			}
			if *p2 {
				if !p2Compare(k, msg_val, sue_val) {
					matched = false
					break
				}
			} else if sue_val != msg_val {
				matched = false
				break
			}
		}
		if matched {
			fmt.Printf("Found her!\nSue %d:\n%s\n", id, sue)
			break
		}
	}
}

func p2Compare(key string, msg_val, sue_val int) bool {
	switch key {
	case "cats", "trees":
		return sue_val > msg_val
	case "pomeranians", "goldfish":
		return sue_val < msg_val
	default:
		return sue_val == msg_val
	}
}
