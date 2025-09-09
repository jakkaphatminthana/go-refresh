package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitializeLogger() {
	var errLoggerBuild error
	config := zap.NewProductionConfig()

	config.EncoderConfig.TimeKey = "receiveTimestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = ""
	config.EncoderConfig.LevelKey = "severity"

	Logger, errLoggerBuild = config.Build()
	if errLoggerBuild != nil {
		panic(errLoggerBuild)
	}
}
