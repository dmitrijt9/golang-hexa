package logger

import (
	"fmt"
	"runtime"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func toZapLevel(level Level) zapcore.Level {
	switch level {
	case LevelDebug:
		return zapcore.DebugLevel
	case LevelInfo:
		return zapcore.InfoLevel
	case LevelWarn:
		return zapcore.WarnLevel
	case LevelError:
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func getContextField(logContext ...Context) []zapcore.Field {
	f := []zapcore.Field{}

	if len(logContext) > 0 {
		f = append(f, zap.Any("context", logContext[0]))
	}

	return f
}

func checkFields(l Logger, logContext ...Context) {
	_, file, line, _ := runtime.Caller(2)
	caller := fmt.Sprintf("%s:%d", file, line)
	if len(logContext) > 1 {
		l.Warn("logger: Added just first map to context object. Cannot process more than one map with context fields. Use just second argument to pass a map with context fields.", Context{"caller": caller, "number of maps given": len(logContext)})
	}
}
