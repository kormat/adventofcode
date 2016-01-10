package day14

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

type Deer struct {
	Name       string
	Speed      int
	FlyPeriod  int
	RestPeriod int
}

func Parse(lines []string) []Deer {
	var deers []Deer

	rx := regexp.MustCompile(
		`^(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds\.$`)
	for i, line := range lines {
		result := rx.FindStringSubmatch(line)
		if result == nil {
			log.Fatalf("Unable to parse line %d: %s", i, line)
		}
		speed := parseInt(i, result[2], "speed")
		fly := parseInt(i, result[3], "flight period")
		rest := parseInt(i, result[4], "rest period")
		deers = append(deers, Deer{result[1], speed, fly, rest})
	}
	return deers
}

func parseInt(line int, input, desc string) int {
	ret, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("Unable to parse %s '%s' on line %d: %s", desc, input, line, err)
	}
	return ret
}

func (d *Deer) String() string {
	return fmt.Sprintf("Name: %s. Speed: %d km/s. Fly period: %ds. Rest period: %ds.",
		d.Name, d.Speed, d.FlyPeriod, d.RestPeriod)
}

func (d *Deer) travel(t int) int {
	full_period := d.FlyPeriod + d.RestPeriod
	remainder := t % full_period
	if remainder < d.FlyPeriod {
		return d.Speed
	}
	return 0
}
