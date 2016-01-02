package main

import (
	"fmt"
	"github.com/kormat/adventofcode/day11/lib"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Expecting a command ('check' or 'gen') and an input string (got %v instead)", os.Args[1:])
	}
	cmd := os.Args[1]
	input := os.Args[2]
	switch cmd {
	case "check":
		if day11.Validate(input) {
			fmt.Printf("%s: Valid\n", input)
		} else {
			fmt.Printf("%s: Invalid\n", input)
		}
	case "gen":
		c, passwd := day11.Generate(input)
		fmt.Printf("After %d iterations, new password: %s\n", c, passwd)
	}
}
