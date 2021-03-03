package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
)

type level int

const (
	TRACE level = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

var logLevel level

// SetLogLevel ..
func SetLogLevel(l level) {
	logLevel = l
}

var logger = log.New(os.Stdout, "Log: ", log.Ldate|log.Ltime|log.Lshortfile)

// SetLogger will set a custom writer for ourapp to log thing
func SetLogger(wc io.WriteCloser) {
	logger = log.New(wc, "Log: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// T ....
func T(v ...interface{}) {
	if logLevel <= TRACE { // 2 <= 0
		logger.SetPrefix("TRACE: ")
		logger.Println(v...)
	}
}

// D ....
func D(v ...interface{}) {
	if logLevel <= DEBUG {
		logger.SetPrefix("DEBUG: ")
		logger.Println(v...)
	}
}

// I ....
func I(v ...interface{}) {
	if logLevel <= INFO {
		logger.SetPrefix("INFO: ")
		logger.Println(v...)
	}
}

// W ....
func W(v ...interface{}) {
	if logLevel <= WARN {
		logger.SetPrefix("WARN: ")
		logger.Println(v...)
	}
}

// E ....
func E(v ...interface{}) {
	if logLevel <= ERROR {
		logger.SetPrefix("ERROR: ")
		logger.Println(v...)
		printStackTrace(10)
	}
}

// F ....
func F(v ...interface{}) {
	if logLevel <= FATAL {
		logger.SetPrefix("FATAL: ")
		printStackTrace(10)
		logger.Fatalln(v...)
	}
}

func printStackTrace(maxStackLength int) {
	stackBuf := make([]uintptr, maxStackLength)
	length := runtime.Callers(3, stackBuf[:])
	stack := stackBuf[:length]

	trace := ""
	frames := runtime.CallersFrames(stack)
	// debug.PrintStack()
	for {
		frame, more := frames.Next()
		if !strings.Contains(frame.File, "runtime/") {
			trace = trace + fmt.Sprintf("\n\tFile: %s, Line: %d. Function: %s", frame.File, frame.Line, frame.Function)
		}

		if !more {
			break
		}
	}
	logger.Println(trace)
}
