package textreader_test

import (
	"errors"
	"ishmaelavila/amextocsv/internal/textreader"
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
}
