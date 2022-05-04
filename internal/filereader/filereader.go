package filereader

import "io"

type FileReader interface {
	OpenFile() *io.ReadCloser
	ReadLine() *string
	CloseFile() bool
}
