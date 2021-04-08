package utility

import (
	"io"
	"log"
	"os"
)

func getErrorFilePath() string {
	path, err := os.Getwd()
	if err != nil {
	}
	return path + "\\error.log"
}

func WriteToLog(a ...interface{}) {
	f, err := os.OpenFile(getErrorFilePath(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
	log.Println(a...)
}
