package mocks

import "io/fs"

type DirEntryMock struct {
	FileName string
	Dir      bool
	FileMode fs.FileMode
	FileInfo fs.FileInfo
}

func (d DirEntryMock) Name() string {
	return d.FileName
}

func (d DirEntryMock) IsDir() bool {
	return d.Dir
}

func (d DirEntryMock) Type() fs.FileMode {
	return d.FileMode
}

func (d DirEntryMock) Info() (fs.FileInfo, error) {
	return d.FileInfo, nil
}
