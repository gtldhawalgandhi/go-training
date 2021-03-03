// Package main provides ...
package main

import (
	"fmt"
	"io"
	"os"

	l "github.com/gtldhawalgandhi/go-training/3.Intermediate/logger"
)

func main() {
	var logFile io.WriteCloser
	// WSL // file permissions
	// Windows: perm has no effect
	// 1 1 0 >> R W X >> 6 6 0
	logFile, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		fmt.Println(err)
	}
	defer logFile.Close()
	l.SetLogger(logFile)
	l.SetLogLevel(l.INFO)
	l.I("Entering main file")

	l.T("My logger ")
	l.E("My error  ")
	l.F("My Fatal example  ")
}
