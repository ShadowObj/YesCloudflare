package log

import (
	"fmt"
	"log"
)

type Logger struct {
	logger  *log.Logger
	logChan chan<- string
}

func Default(ch chan<- string) *Logger {
	return &Logger{
		logger:  log.Default(),
		logChan: ch,
	}
}

func (l *Logger) Fatalf(format string, v ...any) {
	f := fmt.Sprintf(format, v...)
	l.logChan <- f
	l.logger.Fatal(f)
}

func (l *Logger) Printf(format string, v ...any) {
	f := fmt.Sprintf(format, v...)
	l.logChan <- f
	l.logger.Print(f)
}
