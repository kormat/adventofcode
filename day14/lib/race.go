package day14

import (
	"fmt"
)

func Race(deers []Deer, duration int) {
	dists := make(map[string]int)
	scores := make(map[string]int)
	for t := 0; t < duration; t++ {
		for _, d := range deers {
			dists[d.Name] += d.travel(t)
		}
		awardPoints(dists, scores)
	}
	for _, d := range deers {
		fmt.Printf("%s travelled %vkm, and got %d points\n", d.Name, dists[d.Name], scores[d.Name])
	}
	fmt.Println("-----------")
	furthest, max_dist := findMax(dists)
	fmt.Printf("Furthest travelled: %s, with %vkm\n", furthest, max_dist)
	highest, max_score := findMax(scores)
	fmt.Printf("Highest score: %s, with %v points\n", highest, max_score)
}

func awardPoints(dists, scores map[string]int) {
	var leaders []string
	max_dist := 0
	for name, dist := range dists {
		if dist > max_dist {
			max_dist = dist
			leaders = []string{name}
		} else if dist == max_dist {
			leaders = append(leaders, name)
		}
	}
	for _, l := range leaders {
		scores[l]++
	}
}

func findMax(data map[string]int) (string, int) {
	var best string
	max := 0
	for name, val := range data {
		if val > max {
			best = name
			max = val
		}
	}
	return best, max
}
