package logger

import (
	"log"
	"os"
)

type BuiltinLogger struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func NewBuiltinLogger() *BuiltinLogger {
	return &BuiltinLogger{
		InfoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
