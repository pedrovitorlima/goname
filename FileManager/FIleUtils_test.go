package FileManager

import (
	"RenameFiles/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/fs"
	"os"
	"testing"
)

type FileReaderMock struct {
	mock.Mock
}

func (fr *FileReaderMock) ReadFiles() ([]fs.DirEntry, error) {
	args := fr.Called()
	return args.Get(0).([]os.DirEntry), args.Error(1)
}

func TestReadAllFileNamesWithPrefix(t *testing.T) {
	fileReaderMock := FileReaderMock{}

	fileEntries := []os.DirEntry{
		mocks.DirEntryMock{FileName: "odd-1.jpeg"},
		mocks.DirEntryMock{FileName: "odd-2.jpeg"},
		mocks.DirEntryMock{FileName: "even" +
			"-1.jpeg"},
		mocks.DirEntryMock{FileName: "filename.jpeg"},
	}

	fileReaderMock.On("ReadFiles").Return(fileEntries, nil)

	fileUtils := FileUtils{&fileReaderMock}

	allFiles := fileUtils.GetAllFilenamesCurrentDirectory()

	expectedFileNames := []string{"odd-1.jpeg", "odd-2.jpeg", "even-1.jpeg"}
	assert.ElementsMatch(t, expectedFileNames, allFiles)
}
