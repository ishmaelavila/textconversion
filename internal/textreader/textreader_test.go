package textreader_test

import (
	"bufio"
	"errors"
	"ishmaelavila/amextocsv/internal/textreader"
	"os"
	"strings"
	"testing"
)

func TestTextReader(t *testing.T) {

	t.Run("returns an error if given an empty path", func(t *testing.T) {

		_, err := textreader.New("")

		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		if !errors.Is(err, textreader.ErrEmptyFilePath) {
			t.Fatalf("wrong error, got: %s want: %s", err, textreader.ErrEmptyFilePath)
		}
	})

	t.Run("reads a line", func(t *testing.T) {
		reader := textreader.Reader{}

		lineExpected := "May 6"

		linesToScan := strings.NewReader(`May 6
		Pending
		
		CHICK-FIL-A
		
		$27.48`)

		scanner := bufio.NewScanner(linesToScan)

		reader.Scanner = *scanner
		reader.OpenedFile = os.File{}

		text, err := reader.ReadLine()

		if err != nil {
			t.Fatalf("unexpected error reading line %s", err)
		}

		if text == nil {
			t.Fatalf("error, read text was nil")
		}

		if *text != lineExpected {
			t.Fatalf("error unexpected string, got %s, want %s", *text, lineExpected)
		}

	})
}
