package day9

import (
	"fmt"
)

type Best struct {
	Val   int
	Route []string
}

var best Best
var shortest bool

func (ps Places) FindBestRoute(short bool) ([]string, int) {
	names := ps.names()
	shortest = short
	fmt.Printf("Names: %s Shortest? %v\n", names, shortest)
	ps.FindRoute([]string{}, names, 0)
	return best.Route, best.Val
}

func (ps Places) FindRoute(done []string, todo []string, curr_dist int) {
	if len(todo) == 0 {
		bestUpdate(done, curr_dist)
		return
	}
	for i, name := range todo {
		new_done := append(copyList(done), name)
		new_todo := delListItem(todo, i)
		new_dist := curr_dist + ps.nextDist(done, name)
		ps.FindRoute(new_done, new_todo, new_dist)
	}
}

func (ps Places) nextDist(done []string, next string) int {
	if len(done) == 0 {
		return 0
	}
	curr := done[len(done)-1]
	return ps[curr].Distances[next]
}

func delListItem(list []string, idx int) []string {
	after := copyList(list[:idx])
	return append(after, list[idx+1:]...)
}

func copyList(list []string) []string {
	return append([]string(nil), list...)
}

func bestUpdate(route []string, dist int) {
	if (best.Val == 0) || (shortest && dist < best.Val) || (!shortest && dist > best.Val) {
		best.Route = route
		best.Val = dist
	}
}
