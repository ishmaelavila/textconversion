package textreader

import (
	"bufio"
	"errors"
	"log"
	"os"
)

var (
	ErrEmptyFilePath error = errors.New("file path must not have a length of less than 1")
)

type Reader struct {
	Scanner    bufio.Scanner
	OpenedFile os.File
}

func New(pathToTextFile string) (*Reader, error) {

	reader := Reader{}

	file, err := openFile(pathToTextFile)

	if err != nil {
		return nil, err
	}

	reader.OpenedFile = *file

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	reader.Scanner = *scanner

	return &reader, nil
}

func openFile(path string) (*os.File, error) {

	if path == "" {
		return nil, ErrEmptyFilePath
	}

	fileHandle, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	return fileHandle, nil
}

func (t *Reader) ReadLine() (*string, error) {

	if !t.Scanner.Scan() {
		err := t.Scanner.Err()
		if err == nil {
			log.Println("Scan completed and reached EOF")
			t.OpenedFile.Close()
			return nil, nil
		} else {
			log.Println("Other error")
			t.OpenedFile.Close()
			return nil, err
		}
	}

	text := t.Scanner.Text()
	return &text, nil

}
