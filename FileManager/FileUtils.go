package FileManager

import (
	"log"
	"strings"
)

type FileUtils struct {
	FileReader FileReader
}

func (u *FileUtils) GetAllFilenamesCurrentDirectory() []string {
	files, err := u.FileReader.ReadFiles()

	if err != nil {
		log.Fatal(err)
	}

	var listOfFiles []string
	for _, file := range files {
		listOfFiles = append(listOfFiles, file.Name())
	}

	return filterWithPrefixOfInterest(listOfFiles)
}

func filterWithPrefixOfInterest(files []string) []string {
	var filtered = make([]string, 0)

	for _, fileName := range files {
		if strings.HasPrefix(fileName, "even-") ||
			strings.HasPrefix(fileName, "odd-") {
			filtered = append(filtered, fileName)
		}
	}

	return filtered
}
