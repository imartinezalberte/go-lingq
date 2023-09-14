package logger

import "go.uber.org/zap"

func DefaultLogger() ILogger {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	l, err := config.Build(zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	if err != nil {
		panic(err)
	}

	return NewLogger(l)
}
