package files

import (
	"fmt"
	"os"
	"path/filepath"
)

func List(path string) ([]IFile, error) {

	files := make([]IFile, 0)

	isAbs := filepath.IsAbs(path)

	var err error
	if !isAbs {
		path, err = filepath.Abs(path)
	}

	if err == nil {
		err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err == nil {
				files = append(files, IFile{
					path,
				})
			} else {
				fmt.Println("error while walking path", err)
			}
			return nil
		})
	}

	return files, err
}
