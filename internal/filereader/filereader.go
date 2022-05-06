package filereader

type FileReader interface {
	ReadLine() (*string, error)
}
