package logger

import "go.uber.org/zap"

type Logger struct {
	zap.SugaredLogger
}

func NewLogger() *Logger {
	return &Logger{*zap.NewExample().Sugar()}
}
