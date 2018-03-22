package files

import (
	"os"
	"fmt"
)

type File struct{
	path string
}

func (f File) ToString() string {
	return f.path
}

func (f File) Open() (*os.File, error) {
	file, err := os.Open(f.path)
	return file, err
}

func (f File) Delete() error {
	err := os.Remove(f.path)
	if err != nil {
		fmt.Println("couldn't delete file:", f.path)
	} else {
		fmt.Println("file sucessfully deleted", f.path)
	}
	return err
}

