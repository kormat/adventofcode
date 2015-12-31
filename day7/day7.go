package main

import (
	"fmt"
	"github.com/kormat/adventofcode/util"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"sync"
)

type Wire struct {
	Name  string
	Done  chan bool
	Value uint16
}

type WireMap map[string]*Wire

type CmdFunc struct {
	rx *regexp.Regexp
	f  func(*Wire, []string)
}

var RxCmdMap = [...]CmdFunc{
	CmdFunc{regexp.MustCompile(`^(\w+)$`), calcAssign},
	CmdFunc{regexp.MustCompile(`^NOT (\w+)$`), calcNot},
	CmdFunc{regexp.MustCompile(`^(\w+) (AND|OR|LSHIFT|RSHIFT) (\w+)$`), calcGeneral},
}

var wg sync.WaitGroup
var wires = make(WireMap)
var wires_m = sync.Mutex{}
var wire_names []string

func main() {
	lines, err := util.ReadFileArg()
	if err {
		os.Exit(1)
	}
	parseLines(lines)
	wg.Wait()
	fmt.Printf("All calculations finished:\n")
	sort.Strings(wire_names)
	wires_m.Lock()
	for _, name := range wire_names {
		fmt.Printf("  %s: %d\n", name, wires[name].Value)
	}
	wires_m.Unlock()
}

func parseLines(lines []string) {
	rx := regexp.MustCompile(`^(.*) -> (.*)$`)
	for i, line := range lines {
		success := false
		result := rx.FindStringSubmatch(line)
		if result == nil {
			log.Fatalf("Can't parse line %d: '%s'", i, line)
		}
		cmd := result[1]
		wire := addWire(result[2])
		for _, cmdf := range RxCmdMap {
			result = cmdf.rx.FindStringSubmatch(cmd)
			if result == nil {
				continue
			}
			wg.Add(1)
			go wrapCalc(cmdf.f, wire, result)
			success = true
			break
		}
		if !success {
			log.Fatalf("Unable to parse command in line %d: '%s'", i, line)
		}
	}
}

func wrapCalc(f func(*Wire, []string), wire *Wire, results []string) {
	defer wg.Done()
	f(wire, results[1:])
	close(wire.Done)
}

// Calculation functions:

func calcAssign(wire *Wire, matches []string) {
	wire.Value = readValue(matches[0])
}

func calcNot(wire *Wire, matches []string) {
	wire.Value = ^readValue(matches[0])
}

func calcGeneral(wire *Wire, matches []string) {
	a := readValue(matches[0])
	b := readValue(matches[2])
	op := matches[1]
	switch op {
	case "AND":
		wire.Value = a & b
	case "OR":
		wire.Value = a | b
	case "LSHIFT":
		wire.Value = a << b
	case "RSHIFT":
		wire.Value = a >> b
	}
}

// Utility functions:

func addWire(name string) *Wire {
	wires_m.Lock()
	elem, ok := wires[name]
	if !ok {
		elem = &Wire{Name: name, Done: make(chan bool)}
		wires[name] = elem
		wire_names = append(wire_names, name)
	}
	wires_m.Unlock()
	return elem
}

func readValue(s string) uint16 {
	val, err := strconv.Atoi(s)
	if err == nil {
		return uint16(val)
	}
	wire := addWire(s)
	return readWireValue(wire)
}

func readWireValue(wire *Wire) uint16 {
	<-wire.Done
	return wire.Value
}
