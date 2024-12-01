package common

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func ReadInput(day int) string {
	isTest := flag.Bool("test", false, "Use test data")
	flag.Parse()

	directory := fmt.Sprintf("./day%02d/", day)
	filename := directory + "input.txt"
	if *isTest {
		filename = directory + "test_input.txt"
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
