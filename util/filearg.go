package util

import (
	"bufio"
	"log"
	"os"
)

func ReadFileArg(args []string) ([]string, bool) {
	var lines []string
	if len(args) != 1 {
		log.Printf("Expecting a single input file as argument (got %v instead)", args)
		return lines, false
	}
	file, err := os.Open(args[0])
	if err != nil {
		log.Printf("Error opening file: %s", err)
		return lines, false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, true
}
