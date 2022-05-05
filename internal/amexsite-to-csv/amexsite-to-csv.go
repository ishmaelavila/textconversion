package amexsitetocsv

import (
	"errors"
	"ishmaelavila/amextocsv/internal/filereader"
)

var (
	linesPerTransaction      int   = 4
	IncompleteTransactionErr error = errors.New("missing a line, unable to complete transaction")
)

type AmexSiteConverter struct {
	fileReader *filereader.FileReader
}

func (a *AmexSiteConverter) ConvertToCsv() string {

	for {

	}

}

func readTransactions(reader filereader.FileReader) ([]string, error) {

	var transactionLines []string

	for i := 0; i < linesPerTransaction; i++ {
		line, err := reader.ReadLine()

		if err != nil {
			return nil, err
		}

		if line == nil {
			break
		}

		transactionLines = append(transactionLines, *line)
	}

	return transactionLines, nil

}
