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
		_path := fmt.Sprintf("%s, size:%d", path, info.Size())
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
