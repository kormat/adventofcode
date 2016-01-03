package day13

import (
	"log"
	"regexp"
	"strconv"
)

type person struct {
	Name string
	Map  map[string]int
}

type peopleMap map[string]*person

var people peopleMap

func Parse(lines []string, part2 bool) {
	people = make(peopleMap)
	if part2 {
		addPerson("Me", part2)
	}
	rx := regexp.MustCompile(`^(\w+) would (\w+) (\d+) happiness units by sitting next to (\w+).$`)
	for i, line := range lines {
		result := rx.FindStringSubmatch(line)
		if result == nil {
			log.Fatalf("Unable to parse line %d: %s", i, line)
		}
		persA := addPerson(result[1], part2)
		change := parseUnits(result[2], result[3])
		persA.Map[result[4]] = change
	}
}

func addPerson(name string, part2 bool) *person {
	p, ok := people[name]
	if !ok {
		p = &person{Name: name}
		p.Map = make(map[string]int)
		if part2 && name != "Me" {
			p.Map["Me"] = 0
			people["Me"].Map[name] = 0
		}
		people[name] = p
	}
	return p
}

func parseUnits(dir string, count string) int {
	val, err := strconv.Atoi(count)
	if err != nil {
		log.Fatal("Unable to parse happiness units: %s", count)
	}
	switch dir {
	case "gain":
		// Do nothing
	case "lose":
		val *= -1
	default:
		log.Fatal("Invalid happiness direction: %s", dir)
	}

	return val
}
