package FileManager

import (
	"log"
	"os"
)

type FileReaderImpl struct {
}

type FileReader interface {
	ReadFiles() ([]os.DirEntry, error)
}

func (fr *FileReaderImpl) ReadFiles() ([]os.DirEntry, error) {
	currentPath, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	return os.ReadDir(currentPath)
}
