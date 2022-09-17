package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var files []string

	root := "/root/go_project/src/github.com/crazydeng"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		size := fmt.Sprintf("%d BYTES", info.Size())
		mb := float64(info.Size()) / (1024.0 * 1024)
		if mb > 1.0 {
			// lager than 1mb
			size = fmt.Sprintf("%0.2f MB", mb)
		}
		_path := fmt.Sprintf("%s, size:%s", path, size)
		files = append(files, _path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
}
