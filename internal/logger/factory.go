package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*
Format ensures that only one of the formats is used to print a log

There are two formats: "json" and "console"
*/
type Format string

const (
	// FormatJSON JSON formatted log with common fields + app and environment attrs. Usually used on production and processed by elastic / log stash / etc.
	FormatJSON Format = "json"
	// FormatConsole Classic colored formatting to console [time] [level] [message] [other fields]/n[stacktrace]. Usually used in development mode.
	FormatConsole Format = "console"
)

type Config struct {
	Level  Level
	Format Format
}

func New(conf Config) Logger {
	defaultLevel := toZapLevel(conf.Level)

	var core zapcore.Core
	if conf.Format == FormatConsole {
		core = newConsoleCore(defaultLevel)
	}

	if conf.Format == FormatJSON {
		core = newJSONCore(defaultLevel)
	}

	zapLogger := zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	return &logger{
		zapLogger: zapLogger,
		Level:     conf.Level,
		Format:    conf.Format,
	}
}

func newConsoleCore(level zapcore.Level) zapcore.Core {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	return zapcore.NewCore(zapcore.NewConsoleEncoder(config), os.Stdout, level)
}

func newJSONCore(level zapcore.Level) zapcore.Core {
	encConfig := zap.NewProductionEncoderConfig()
	encConfig.MessageKey = "message"
	encConfig.LevelKey = "level"
	encConfig.TimeKey = "timestamp"
	encConfig.CallerKey = "log_caller"
	encConfig.EncodeLevel = zapcore.LowercaseLevelEncoder

	enc := zapcore.NewJSONEncoder(encConfig)

	return zapcore.NewCore(enc, os.Stdout, level)
}
