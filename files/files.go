package files

import (
	"fmt"
	"path/filepath"
	"os"
)


func List(path string) ([]File, error) {

	files := make([]File, 0)

	isAbs := filepath.IsAbs(path)

	var err error
	if !isAbs {
		path, err = filepath.Abs(path)
	}
	
	if err == nil {
		err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err == nil {
				files = append(files, File{
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