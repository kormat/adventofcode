package day13

import (
	"fmt"
)

var result int

func Seating() int {
	var names []string
	for key, _ := range people {
		names = append(names, key)
	}
	fmt.Printf("People: %s\n", names)
	calcSeating([]string{}, names)

	return result
}

func calcSeating(done, todo []string) {
	if len(todo) == 0 {
		addResult(done)
		return
	}
	for i, name := range todo {
		new_done := append(copyList(done), name)
		new_todo := delListItem(todo, i)
		calcSeating(new_done, new_todo)
	}

}

func delListItem(list []string, idx int) []string {
	after := copyList(list[:idx])
	return append(after, list[idx+1:]...)
}

func copyList(list []string) []string {
	return append([]string(nil), list...)
}

func addResult(names []string) {
	var net int
	for i := 0; i < len(names)-1; i++ {
		net += lookupMap(names[i], names[i+1])
	}
	// Handle the wrap-around
	net += lookupMap(names[0], names[len(names)-1])
	if net > result {
		fmt.Printf("Best yet names: %s\n", names)
		result = net
	}
}

func lookupMap(name1, name2 string) int {
	return people[name1].Map[name2] + people[name2].Map[name1]
}
