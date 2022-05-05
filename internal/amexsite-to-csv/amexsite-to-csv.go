package amexsitetocsv

import (
	"errors"
	"ishmaelavila/amextocsv/internal/filereader"
	"ishmaelavila/amextocsv/internal/textreader"
	"log"
	"time"
)

var (
	linesPerTransaction      int    = 3
	IncompleteTransactionErr error  = errors.New("missing a line, unable to complete transaction")
	lineOfDate               int    = 0
	lineOfDescription        int    = 1
	lineOfAmount             int    = 2
	amexDateLayout           string = "Jan 2"
)

var specialLines = map[string]string{
	"Pending":      "skip",
	"6% Cash Back": "skip",
	"Credit":       "skip",
	"":             "skip",
}

type Coverter struct {
	fileReader filereader.FileReader
}

func New(path string) (*Coverter, error) {

	textReader, err := textreader.New(path)

	if err != nil {
		return nil, err
	}

	return &Coverter{
		fileReader: textReader,
	}, nil
}

func (a *Coverter) ConvertToCsv() string {

	transactionsAsCSVString := ""
	allTransactions, err := a.readLines()

	if err != nil {
		log.Fatalf("error parsing transactions: %s", err)
	}

	for _, transactionLines := range allTransactions {

		currTransaction := a.convertLinesToTransaction(transactionLines)
		currTransaction += "\n"
		transactionsAsCSVString += currTransaction

	}

	return transactionsAsCSVString
}

func (a *Coverter) convertLinesToTransaction(lines []string) string {
	transaction := ""

	for index, arg := range lines {
		switch index {
		case lineOfDate:
			fallthrough
		case lineOfAmount:
			fallthrough
		case lineOfDescription:
			if index+1 == linesPerTransaction {
				transaction += arg
			} else {
				transaction += arg + ", "
			}

		}
	}
	return transaction
}

func isDate(line string) bool {
	_, err := time.Parse(amexDateLayout, line)

	if err != nil {
		return false
	}

	return true
}

func (a *Coverter) readLines() ([][]string, error) {

	var allTransactions [][]string
	var currentTransaction []string

	for {
		line, err := a.fileReader.ReadLine()

		if err != nil {
			return nil, err
		}

		if line == nil {
			allTransactions = append(allTransactions, currentTransaction)
			break
		}

		if action, ok := specialLines[*line]; ok {
			switch action {
			case "skip":
				continue
			}
		}

		if isDate(*line) && len(currentTransaction) > 0 {
			allTransactions = append(allTransactions, currentTransaction)
			currentTransaction = nil
			currentTransaction = append(currentTransaction, *line)
			continue
		}

		currentTransaction = append(currentTransaction, *line)
	}

	return allTransactions, nil

}
