package logger

import (
	"go.uber.org/zap"
)

/*
Level is an enum of common levels as we know them.
  - debug 0
  - info 1
  - warn 2
  - error 3

Selected level and all "above it" (ascending order) is used to log messages.
*/
type Level string

const (
	LevelDebug Level = "debug"
	LevelInfo  Level = "info"
	LevelWarn  Level = "warn"
	LevelError Level = "error"
)

type Logger interface {
	Debug(msg string, logContext ...Context) Logger
	Info(msg string, logContext ...Context) Logger
	Warn(msg string, logContext ...Context) Logger
	Error(msg string, logContext ...Context) Logger
}

type logger struct {
	Level     Level
	Format    Format
	zapLogger *zap.Logger
}

type Context = map[string]interface{}

func (l *logger) Debug(msg string, logContext ...Context) Logger {
	checkFields(l, logContext...)
	l.zapLogger.Debug(msg, getContextField(logContext...)...)
	return l
}

func (l *logger) Info(msg string, logContext ...Context) Logger {
	checkFields(l, logContext...)
	l.zapLogger.Info(msg, getContextField(logContext...)...)
	return l
}

func (l *logger) Warn(msg string, logContext ...Context) Logger {
	checkFields(l, logContext...)
	l.zapLogger.Warn(msg, getContextField(logContext...)...)
	return l
}

func (l *logger) Error(msg string, logContext ...Context) Logger {
	checkFields(l, logContext...)
	l.zapLogger.Error(msg, getContextField(logContext...)...)
	return l
}
