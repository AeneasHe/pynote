package fs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var RootPath string = "~/"

type StopError struct {
}

func (err *StopError) Error() string {
	return "stop"
}
func ReadFile(filename string) string {
	//filename := "/Users/aeneas/Github/Cofepy/youdao/0新文档/ERP.md"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Open file Failed", err)
		return ""
	}
	defer func() {
		file.Close()
	}()
	var b []byte = make([]byte, 4096)
	n, err := file.Read(b)
	if err != nil {
		fmt.Println("Open file Failed", err)
	}
	data := string(b[:n])
	return data
}

func ShowPath(path string, flag string) []string {
	var names []string

	if path == "." {
		path = RootPath
	}

	println("====>path", path)
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if len(names) > 10 {
			return &StopError{}
		}
		if flag == "folder" {
			if info.IsDir() {

				path, err := filepath.Abs(path)
				names = append(names, path)
				if err != nil {
					return err
				}
			}
		}
		if flag == "file" {
			if !info.IsDir() {
				if strings.HasSuffix(path, ".md") {
					path, err := filepath.Abs(path)
					names = append(names, path)
					if err != nil {
						return err
					}
				}
			}
		}
		if flag == "all" {
			path, err := filepath.Abs(path)
			names = append(names, path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return names
}
