package main

import (
	"fmt"
	"github.com/kormat/adventofcode/day7/lib"
	"github.com/kormat/adventofcode/util"
	"os"
)

func main() {
	lines, err := util.ReadFileArg(os.Args[1:])
	if err {
		os.Exit(1)
	}
	d := day7lib.NewDay7()
	d.ParseLines(lines)
	fmt.Printf("All calculations finished:\n")
	d.Print()
}
