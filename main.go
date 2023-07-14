package main

import (
	fileManager "RenameFiles/FileManager"
	"log"
	"os"
	"path/filepath"
)

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fileUtils := fileManager.FileUtils{FileReader: &fileManager.FileReaderImpl{}}

	allFilesCurrentDirectory := fileUtils.GetAllFilenamesCurrentDirectory()

	fileOrganiser := fileManager.FileOrganiser{}
	odds, even := fileOrganiser.SplitIntoOddEvenLists(allFilesCurrentDirectory)
	commandMap := fileOrganiser.CreateRenameCommandMap(odds, even)

	currentDir, err := os.Getwd()

	HandleError(err)

	for oldName, newName := range commandMap {
		oldPath := filepath.Join(currentDir, oldName)
		newPath := filepath.Join(currentDir, newName)

		HandleError(os.Rename(oldPath, newPath))

	}
}
