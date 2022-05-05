package main

import (
	"fmt"
	amexsitetocsv "ishmaelavila/amextocsv/internal/amexsite-to-csv"
	"log"
	"os"
)

func main() {

	args := os.Args
	argsExpected := 2

	if len(args) > argsExpected {
		log.Printf("Invalid number of arguments, expected %d, got %d, ignoring extra arguments", argsExpected, len(args))
	}

	if len(args) < argsExpected {
		log.Fatalf("Invalid number of arguments, please provide a path")
	}

	path := args[1]

	converter, err := amexsitetocsv.New(path)

	if err != nil {
		log.Fatalf("could not initalize amex site to csv converter: %s", err)
	}

	csvString := converter.ConvertToCsv()

	fmt.Println(csvString)

	f, err := os.Create("./output.csv")

	if err != nil {
		log.Fatalf("error opening output file: %s", err)
	}

	_, err = f.WriteString(csvString)

	if err != nil {
		log.Fatalf("error writing to output file %s", err)
	}

	os.Exit(0)

}
