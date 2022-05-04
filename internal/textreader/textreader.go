package textreader

import (
	"bufio"
	"errors"
	"log"
	"os"
)

type TextReader struct {
	Scanner    bufio.Scanner
	OpenedFile *os.File
}

func New(pathToTextFile string) (*TextReader, error) {

	reader := TextReader{}

	file, err := reader.OpenFile(pathToTextFile)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	reader.Scanner = *scanner

	return &reader, nil
}

func (t *TextReader) OpenFile(path string) (*os.File, error) {

	if path == "" {
		return nil, errors.New("file path must be longer than zero characters")
	}
	fileHandle, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	t.OpenedFile = fileHandle

	return fileHandle, nil
}

func (t *TextReader) ReadLine() (string, error) {

	if !t.Scanner.Scan() {
		err := t.Scanner.Err()
		if err == nil {
			log.Println("Scan completed and reached EOF")
			t.OpenedFile.Close()
		} else {
			log.Println("Other error")
			t.OpenedFile.Close()
			return "", err
		}
	}

	return t.Scanner.Text(), nil

}
