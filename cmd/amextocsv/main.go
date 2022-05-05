package main

import (
	"fmt"
	"ishmaelavila/amextocsv/internal/textreader"
	"log"
	"os"
)

func main() {

	args := os.Args
	argsExpected := 2
	log.Printf("args: %s", args)

	if len(args) > argsExpected || len(args) <= 1 {
		log.Printf("Invalid amount of arguments, expected %d, got %d, ignoring extra arguments", argsExpected, len(args))
	}

	path := args[1]

	reader, err := textreader.New(path)

	if err != nil {
		log.Fatalf("Error Opening File: %s", err)
		return
	}

	for {
		line, err := reader.ReadLine()

		if err != nil {
			log.Fatalf("something went wrnog while reading file: %s", err)
		}

		if line == nil {
			break
		}

		fmt.Println(*line)
	}

}
