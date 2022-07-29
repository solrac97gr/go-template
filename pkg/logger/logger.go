package logger

import "log"

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Info(i interface{}) {
	log.Println(i)
}

func (l *Logger) Error(i interface{}) {
	log.Fatalln(i)
}
