package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func DefaultLogger(level zapcore.Level) ILogger {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(level)
	l, err := config.Build(zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	if err != nil {
		panic(err)
	}

	return NewLogger(l)
}
