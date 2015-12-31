package day7lib

import (
	"strconv"
	"sync"
)

type Day7 struct {
	wg    sync.WaitGroup
	wires WireMap
}

func NewDay7() *Day7 {
	d := &Day7{}
	d.wires.M = make(map[string]*Wire)
	return d
}

func (d *Day7) calcAssign(wire *Wire, matches []string) {
	defer d.wg.Done()
	wire.Value = d.readValue(matches[1])
	close(wire.Done)
}

func (d *Day7) calcNot(wire *Wire, matches []string) {
	defer d.wg.Done()
	wire.Value = ^d.readValue(matches[1])
	close(wire.Done)
}

func (d *Day7) calcGeneral(wire *Wire, matches []string) {
	defer d.wg.Done()
	a := d.readValue(matches[1])
	b := d.readValue(matches[3])
	op := matches[2]
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
	close(wire.Done)
}

// Utility functions:

func (d *Day7) readValue(s string) uint16 {
	val, err := strconv.Atoi(s)
	if err == nil {
		return uint16(val)
	}
	wire := d.wires.add(s)
	return wire.read()
}
