package files

import (
	"fmt"
	"path/filepath"
	"os"
	"log"
)


func executionPath() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	return dir, err
}


func List(path string) ([]File, error) {

	files := make([]File, 0)

	// isAbs := filepath.IsAbs(path)

	dir, err := executionPath()
	if err == nil {
		fmt.Println(dir)

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