package logger

import (
	"hexa-example-go/internal/config"
	"log"
	"net"
	"os"

	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger(loggerConfig config.LoggerConfig) *zap.Logger {
	defaultLevel := zapcore.InfoLevel
	if loggerConfig.Level != "" {
		parsedLevel, err := zapcore.ParseLevel(loggerConfig.Level)
		if err == nil {
			defaultLevel = parsedLevel
		}
	}

	var cores []zapcore.Core
	if loggerConfig.ConsoleEnabled {
		consoleCore := initConsoleCore(defaultLevel)
		cores = append(cores, consoleCore)
	}

	if loggerConfig.FilebeatEnabled {
		filebeatCore := initFilebeatCore(defaultLevel, loggerConfig)
		cores = append(cores, filebeatCore)
	}

	core := zapcore.NewTee(cores...)
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

func initConsoleCore(level zapcore.Level) zapcore.Core {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	return zapcore.NewCore(zapcore.NewConsoleEncoder(config), os.Stdout, level)
}

func initFilebeatCore(level zapcore.Level, loggerConfig config.LoggerConfig) zapcore.Core {
	c, err := net.Dial("tcp", loggerConfig.FilebeatUrl)

	if err != nil {
		log.Panic(err)
	}

	logSync := &TcpLogSyncer{}
	logSync.SetTcpConnection(c)

	parsingMetadata := map[string]interface{}{
		"index":    loggerConfig.FilebeatIndex,
		"appname":  loggerConfig.FileBeatAppName,
		"keepDays": 7,
	}
	customFields := []zapcore.Field{zap.String("app", "aspira-waitress"), zap.String("environment", "development"), zap.Any("parsing_metadata", parsingMetadata)}

	encoderConfig := ecszap.NewDefaultEncoderConfig()
	return ecszap.NewCore(encoderConfig, logSync, level).With(customFields)
}
