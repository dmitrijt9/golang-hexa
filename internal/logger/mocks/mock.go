package mocks

import (
	"context"
	"hexa-example-go/internal/logger"
)

type loggerMock struct{}

func New() logger.Logger {
	return &loggerMock{}
}

func (l *loggerMock) Debug(msg string, logContext ...logger.Context) logger.Logger {
	return l
}

func (l *loggerMock) Info(msg string, logContext ...logger.Context) logger.Logger {
	return l
}

func (l *loggerMock) Warn(msg string, logContext ...logger.Context) logger.Logger {
	return l
}

func (l *loggerMock) Error(msg string, logContext ...logger.Context) logger.Logger {
	return l
}

func (l *loggerMock) Printf(ctx context.Context, format string, v ...interface{}) {
}
