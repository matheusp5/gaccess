package logger

import "github.com/fatih/color"

type Logger struct{}

func (logger *Logger) Info(msg string) {
	color.Green("[INFO] " + msg)
}

func (logger *Logger) Error(msg string) {
	color.Red("[ERROR] " + msg)
}

func NewLogger() Logger {
	logger := Logger{}
	return logger
}
