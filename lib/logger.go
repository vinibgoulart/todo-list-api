package lib

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	*log.Logger
}

func (l *Logger) Info(v ...interface{}) {
	l.SetPrefix("INFO: ")
	l.Println(v...)
}

func (l *Logger) Warning(v ...interface{}) {
	l.SetPrefix("WARNING: ")
	l.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.SetPrefix("ERROR: ")
	l.Println(v...)
}

func NewLogger(target string) *Logger {
	logger := log.New(os.Stdout, fmt.Sprintf("clerk.%s: ", target), log.LstdFlags)
	return &Logger{logger}
}
