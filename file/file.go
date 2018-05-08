package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type File struct {
	path string
}

func New(path string) *File {
	file := File{path}
	return &file
}

func (f *File) ToString() string {
	return f.path
}

func (f *File) Open() (*os.File, error) {
	if f.Exists() {
		file, err := os.Open(f.path)
		return file, err
	}
	return nil, errors.New("file doesn't exist")
}

func (f *File) Delete() error {
	err := os.Remove(f.path)
	if err != nil {
		fmt.Println("couldn't delete file:", f.path)
	} else {
		fmt.Println("file sucessfully deleted", f.path)
	}
	return err
}

func (f *File) Create() (*os.File, error) {
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

func (f *File) Exists() bool {
	_, err := os.Stat(f.path)
	if err != nil {
		fmt.Println("File.Exists()", err)
	}
	return err == nil
}

func (f *File) Name() string {
	return filepath.Base(f.path)
}

func (f *File) IsDir() (bool, error) {
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

func (f *File) Type() string {
	// TODO: error if is dir
	return filepath.Ext(f.path)
}

func (f *File) Rename() (*File, error) {
	// TODO
	return f, nil
}

func (f *File) Write(content string) error {
	// TODO
	return nil
}

func (f *File) Append(content string) error {
	// TODO
	return nil
}

func (f *File) Clear() error {
	// TODO
	return nil
}

func (f *File) NumberOfLines() int {
	// TODO
	return -1
}

func (f *File) ReadLine(number int) (string, error) {
	// TODO
	return "", nil
}

func (f *File) ReadLines(from int, to int) (string, error) {
	// TODO
	return "", nil
}
