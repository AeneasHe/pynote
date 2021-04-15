package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetStartPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("目录:", dir)
	return strings.Replace(dir, "\\", "/", -1)
}

func GetExePath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	log.Println("可执行文件", exPath)
	return exPath
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
