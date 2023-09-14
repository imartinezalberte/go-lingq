package logger

import (
	"context"

	"github.com/imartinezalberte/go-lingq/internal/utils"
	"go.uber.org/zap"
)

const ServiceLoggerKey string = "service"

type (
	ILogger interface {
		Debug(context.Context, string, ...zap.Field)
		Info(context.Context, string, ...zap.Field)
		Warn(context.Context, string, ...zap.Field)
		Error(context.Context, string, ...zap.Field)
		DPanic(context.Context, string, ...zap.Field)
		Panic(context.Context, string, ...zap.Field)
		Fatal(context.Context, string, ...zap.Field)
		WithFields(fields ...zap.Field) ILogger
		WithService(string) ILogger
		ToZap() *zap.Logger
		ToZapCtx(context.Context) *zap.Logger
	}

	logger struct {
		l                     *zap.Logger
		fields                []zap.Field
		propertiesFromContext []string
	}
)

func NewLogger(l *zap.Logger, defaultFields ...zap.Field) ILogger {
	iLogger := &logger{
		l:      l,
		fields: defaultFields,
	}

	return iLogger
}

func (l logger) WithFields(fields ...zap.Field) ILogger {
	l.l = l.l.With(fields...)
	return &l
}

func (l logger) WithService(service string) ILogger {
	l.l = l.l.With(zap.String(ServiceLoggerKey, service))
	return &l
}

func (logger logger) readCtx(ctx context.Context) []zap.Field {
	var result []zap.Field

	for _, key := range logger.propertiesFromContext {
		if v := ctx.Value(key); v != nil {
			result = append(result, zap.Any(key, v))
		}
	}

	return result
}

func (logger *logger) Debug(ctx context.Context, msg string, fields ...zap.Field) {
	logger.l.Debug(msg, logger.merge(append(fields, logger.readCtx(ctx)...)...)...)
}

func (logger *logger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	logger.l.Info(msg, logger.merge(append(fields, logger.readCtx(ctx)...)...)...)
}

func (logger *logger) Warn(ctx context.Context, msg string, fields ...zap.Field) {
	logger.l.Warn(msg, logger.merge(append(fields, logger.readCtx(ctx)...)...)...)
}

func (logger *logger) Error(ctx context.Context, msg string, fields ...zap.Field) {
	logger.l.Error(msg, logger.merge(append(fields, logger.readCtx(ctx)...)...)...)
}

func (logger *logger) DPanic(ctx context.Context, msg string, fields ...zap.Field) {
	logger.l.DPanic(msg, logger.merge(append(fields, logger.readCtx(ctx)...)...)...)
}

func (logger *logger) Panic(ctx context.Context, msg string, fields ...zap.Field) {
	logger.l.Panic(msg, logger.merge(append(fields, logger.readCtx(ctx)...)...)...)
}

func (logger *logger) Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	logger.l.Fatal(msg, logger.merge(append(fields, logger.readCtx(ctx)...)...)...)
}

func (logger *logger) ToZap() *zap.Logger {
	return logger.l.With(logger.fields...)
}

func (logger *logger) ToZapCtx(ctx context.Context) *zap.Logger {
	return logger.l.With(logger.merge(logger.readCtx(ctx)...)...)
}

// merge method returns all the fields from the request and overrides the ones that already
// exists in the child logger if necessary.
//
//	// Here we have already setted the zap logger
//	var zapLogger *zap.Logger
//
//	// We initialize the LoggerConfig and create a new logger from this configuration
//	lc := LoggerConfig{ DefaultFields: []LoggerField{ { Key: "integration", Val: "random_integration_here" } } }
//	l := lc.NewLogger(zapLogger, defaultFields...)
//
//	// This will display just the last integration key with the value "real_integration_here"
//	l.Info(context.Background(), zap.String("integration", "real_integration_here"))
func (logger logger) merge(newFields ...zap.Field) []zap.Field {
	return utils.RemoveDuplicatesKeepingLastSeen(append(logger.fields, newFields...), keyField)
}

func keyField(f zap.Field) string { return f.Key }
