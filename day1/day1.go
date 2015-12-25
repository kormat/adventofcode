package main

import "os"
import "fmt"

func main() {
	args := os.Args[1]
	floor, to_minus := calcFloor(args)
	fmt.Printf("Santa ends up on floor %d, and entered the basement at instruction %d\n", floor, to_minus)
}

func calcFloor(dirs string) (int, int) {
	floor := 0
	to_minus := -1
	for i, c := range dirs {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		}
		if floor < 0 && to_minus == -1 {
			to_minus = i + 1
		}
		fmt.Printf("Idx: %d Char: %c Floor: %d To minus: %d\n", i, c, floor, to_minus)
	}
	return floor, to_minus
}
