package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var p2 = flag.Bool("2", false, "Calculate part 2 result")

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Fatalf("Expecting an input file (got %v instead)", args)
	}
	reader := openFile(args[0])
	data := decode(reader)
	count := walkUnknown(data)
	fmt.Printf("Sum of input: %d\n", count)
}

func openFile(path string) io.Reader {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	return file
}

func decode(reader io.Reader) interface{} {
	var data interface{}
	dec := json.NewDecoder(reader)
	if err := dec.Decode(&data); err != nil {
		log.Fatal("Error reading input:", err)
	}
	return data
}

func walkUnknown(data interface{}) int {
	count := 0
	switch val := data.(type) {
	case string:
		// Ignore.
	case float64:
		count += int(val)
	case []interface{}:
		count += walkArray(val)
	case map[string]interface{}:
		count += walkMap(val)
	default:
		fmt.Printf("Uh? %T:%v\n", data, data)
	}
	return count
}

func walkArray(data []interface{}) int {
	count := 0
	for _, val := range data {
		count += walkUnknown(val)
	}
	return count
}

func walkMap(data map[string]interface{}) int {
	count := 0
	for _, val := range data {
		switch val.(type) {
		case string:
			if *p2 && val.(string) == "red" {
				return 0
			}
		}
		count += walkUnknown(val)
	}
	return count
}
