package logger

import (
	"log"
	"os"
)

type Logger struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func New() *Logger {
	errorLog := log.New(os.Stderr, "[ERROR]:\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "[INFO]:\t", log.Ldate|log.Ltime|log.Lshortfile)

	return &Logger{
		errorLog: errorLog,
		infoLog:  infoLog,
	}
}

func (l *Logger) LogError(err error, message string) {
	l.errorLog.Printf("%s: %v", message, err)
}

func (l *Logger) LogInfo(message string) {
	l.infoLog.Println(message)
}

func (l *Logger) LogFatal(err error, message string) {
	l.errorLog.Fatalf("%s: %v", message, err)
}

func (l *Logger) LogPanic(err error, message string) {
	l.errorLog.Panicf("%s: %v", message, err)
}
