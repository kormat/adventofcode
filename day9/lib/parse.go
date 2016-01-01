package day9

import (
	"log"
	"regexp"
	"sort"
	"strconv"
)

type Place struct {
	Name      string
	Distances map[string]int
}

type Places map[string]*Place

func (ps Places) ParseLines(lines []string) {
	rx := regexp.MustCompile(`^(\w+) to (\w+) = (\d+)$`)
	for i, line := range lines {
		result := rx.FindStringSubmatch(line)
		if result == nil {
			log.Fatal("Unable to parse line %d: %s", i, line)
		}
		names := result[1:3]
		sort.Strings(names)
		dist, err := strconv.Atoi(result[3])
		if err != nil {
			log.Fatal("Unable to parse distance on line %d: %s", i, result[3])
		}
		ps.add(names, dist)
	}
}

func (ps Places) add(names []string, dist int) (*Place, *Place) {
	p1 := ps.add1(names[0], names[1], dist)
	p2 := ps.add1(names[1], names[0], dist)
	return p1, p2
}

func (ps Places) add1(local, remote string, dist int) *Place {
	p, ok := ps[local]
	if !ok {
		p = &Place{Name: local}
		p.Distances = make(map[string]int)
		ps[local] = p
	}
	p.Distances[remote] = dist
	return p
}

func (ps Places) names() []string {
	var places []string
	for name, _ := range ps {
		places = append(places, name)
	}
	sort.Strings(places)
	return places
}
