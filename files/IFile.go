package files

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
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
		return file, err
	}
	return nil, nil
}

func (f IFile) Exists() bool {
	_, err := os.Stat(f.path)
	if err != nil {
		fmt.Println("IFile.Exists()", err)
	}
	return !os.IsExist(err)
}

func (f IFile) Name() string {
	return filepath.Base(f.path)
}

func (f IFile) IsDir() (bool, error) {
	file, openError := f.Open()
	defer file.Close()
	if openError != nil {
		return false, openError
	}
	fi, err := file.Stat()
	switch {
	case err != nil:
		return false, err
	case fi.IsDir():
		return true, nil
	default:
		return false, nil
	}
	return false, nil
}

/**
 * ROADMAP
 */

func (f IFile) Type() string {
	// TODO: error if is dir
	return filepath.Ext(f.path)
}

func (f IFile) Rename() (IFile, error) {
	// TODO
	return f, nil
}

func (f IFile) Write(content string) error {
	// TODO
	return nil
}

func (f IFile) Append(content string) error {
	// TODO
	return nil
}

func (f IFile) Clear() error {
	// TODO
	return nil
}

func (f IFile) NumberOfLines() int {
	// TODO
	return -1
}

func (f IFile) ReadLine(number int) (string, error) {
	// TODO
	return "", nil
}

func (f IFile) ReadLines(from int, to int) (string, error) {
	// TODO
	return "", nil
}
