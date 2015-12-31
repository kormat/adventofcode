package day7lib

import (
	"log"
	"regexp"
)

func (d *Day7) ParseLines(lines []string) {
	rx_gen := regexp.MustCompile(`^(.*) -> (.*)$`)
	rx_calc_assign := regexp.MustCompile(`^(\w+)$`)
	rx_calc_not := regexp.MustCompile(`^NOT (\w+)$`)
	rx_calc_gen := regexp.MustCompile(`^(\w+) (AND|OR|LSHIFT|RSHIFT) (\w+)$`)
	for i, line := range lines {
		result := rx_gen.FindStringSubmatch(line)
		if result == nil {
			log.Fatalf("Can't parse line %d: '%s'", i, line)
		}
		cmd := result[1]
		wire := d.wires.add(result[2])
		d.wg.Add(1)
		result = rx_calc_assign.FindStringSubmatch(cmd)
		if result != nil {
			go d.calcAssign(wire, result)
			continue
		}
		result = rx_calc_not.FindStringSubmatch(cmd)
		if result != nil {
			go d.calcNot(wire, result)
			continue
		}
		result = rx_calc_gen.FindStringSubmatch(cmd)
		if result != nil {
			go d.calcGeneral(wire, result)
			continue
		}
		log.Fatalf("Unable to parse command in line %d: '%s'", i, line)
	}
	d.wg.Wait()
}

func (d *Day7) Print() {
	d.wires.print()
}
