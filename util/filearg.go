package util

import (
	"bufio"
	"log"
	"os"
)

func ReadFileArg() ([]string, bool) {
	var lines []string
	if len(os.Args) != 2 {
		log.Printf("Expecting a single input file as argument (got %v instead)", os.Args[1:])
		return lines, true
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Printf("Error opening file: %s", err)
		return lines, true
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, false
}
