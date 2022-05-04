package main

import (
	"fmt"
	"ishmaelavila/amextocsv/internal/textreader"
	"log"
)

func main() {

	reader, err := textreader.New("./test.txt")

	if err != nil {
		log.Fatalf("Error Opening File: %s", err)
		return
	}

	test, err := reader.ReadLine()

	if err != nil {
		log.Fatalf("Error Reading From File: %s", err)
	}

	fmt.Printf("output: %s", test)
}
