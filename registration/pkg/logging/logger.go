package logging

import (
    "log"
    "os"
)

type Logger struct {
    logger *log.Logger
}

func NewLogger() *Logger {
    return &Logger{
        logger: log.New(os.Stdout, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile),
    }
}

func (l *Logger) Error(v ...interface{}) {
    l.logger.SetPrefix("ERROR: ")
    l.logger.Println(v...)
}

func (l *Logger) Info(v ...interface{}) {
    l.logger.SetPrefix("INFO: ")
    l.logger.Println(v...)
}

func (l *Logger) Warn(v ...interface{}) {
    l.logger.SetPrefix("WARN: ")
    l.logger.Println(v...)
}
