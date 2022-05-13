package main

import (
	"flag"
	"fmt"
	amexsitetocsv "ishmaelavila/amextocsv/internal/amexsite-to-csv"
	"log"
	"os"
	"strings"
)

func main() {

	args := os.Args
	argsExpected := 2

	if len(args) > argsExpected {
		log.Printf("invalid number of arguments, expected %d, got %d, ignoring extra arguments", argsExpected-1, len(args))
	}

	if len(args) < argsExpected {
		log.Fatalf("invalid number of arguments. Usage amextocsv")
	}

	path := args[1]

	reverseOutput := flag.Bool("r", true, "reverse output of CSV")

	converter, err := amexsitetocsv.New(path)

	if err != nil {
		log.Fatalf("could not initalize amex site to csv converter: %s", err)
	}

	csvString := converter.ConvertToCsv()

	if *reverseOutput {
		csvString = reverse(csvString)
	}

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

func reverse(s string) string {
	lines := strings.Split(s, "\n")

	reverseOrderSlice := []string{}

	for i := range lines {
		line := strings.TrimSpace(lines[len(lines)-1-i])
		reverseOrderSlice = append(reverseOrderSlice, line+"\n")
	}

	reversedString := strings.Join(reverseOrderSlice, "")
	return reversedString
}
