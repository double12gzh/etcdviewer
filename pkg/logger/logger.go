// Package logger
/* Copyright © 2022 JeffreyGuan <double12gzh@gmail.com> */

package logger

import (
	"fmt"

	"github.com/fatih/color"
)

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Info(msg string, args ...interface{}) {
	_, _ = color.New(color.FgHiCyan).Println(fmt.Sprintf(msg, args...))
}

func (l *Logger) Error(err error) {
	_, _ = color.New(color.FgHiRed).Println(fmt.Sprintf("%#+v", err))
}
