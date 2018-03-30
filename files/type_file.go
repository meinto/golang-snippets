package files

import (
	"errors"
	"fmt"
	"os"
)

type IFile struct {
	path string
}

func (f IFile) ToString() string {
	return f.path
}

func (f IFile) Open() (*os.File, error) {
	if f.Exists() {
		file, err := os.Open(f.path)
		return file, err
	}
	return nil, errors.New("file doesn't exist")
}

func (f IFile) Delete() error {
	err := os.Remove(f.path)
	if err != nil {
		fmt.Println("couldn't delete file:", f.path)
	} else {
		fmt.Println("file sucessfully deleted", f.path)
	}
	return err
}

func (f IFile) Create() (*os.File, error) {
	if !f.Exists() {
		var file, err = os.Create(f.path)
		if err != nil {
			fmt.Println("error while creating file", f.path)
			fmt.Println("error message", err.Error())
		}
		defer file.Close()
		return file, err
	}
	return nil, nil
}

func (f IFile) Exists() bool {
	_, err := os.Stat(f.path)
	return os.IsExist(err)
}
