// Package main provides ...
package main

import (
	l "github.com/gtldhawalgandhi/go-training/3.Intermediate/logger"
)

func helo() {
	l.D("Helo debug")
	l.F("My Fatal example")
}

// -trimpath will cut short file name everywhere in our code when displaying
// go build -trimpath -o app && ./app
// OR
// go run -trimpath .
func main() {
	l.SetFileLogger("myLog.txt", l.TRACE)
	defer l.CleanUp()
	l.I("Entering main file")

	l.T("My Trace data")
	l.E("My error")
	helo()
}
