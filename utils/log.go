package utils

import (
	"fmt"
	"log"
	"os"
)

func InitLog() {

	logFile := "/Users/aeneas/Github/Cofepy/pynote/pynote/tmp/log.log"
	src, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println("err", err)
	}

	log.SetOutput(src)
	log.Println("\n\n=================当前目录:", GetStartPath())
}
