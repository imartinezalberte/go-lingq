package rest

import (
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	log "github.com/imartinezalberte/go-lingq/internal/logger"
)

const (
	defaultTimeout          = 10 * time.Second
	defaultRetryCount       = 3
	defaultRetryWaitTime    = 1 * time.Second
	defaultRetryMaxWaitTime = 2 * time.Second
	defaultLogLevel         = zapcore.InfoLevel
)

// NewClient func returns an already to use resty.Client.
func NewClient(cl *http.Client, logger log.ILogger, host string) (*resty.Client, error) {
	c := resty.NewWithClient(cl).
		SetBaseURL(host).
		SetTimeout(defaultTimeout).
		SetRetryCount(defaultRetryCount).
		SetRetryWaitTime(defaultRetryWaitTime).
		SetRetryMaxWaitTime(defaultRetryMaxWaitTime).
		SetDebug(true)

	c.OnBeforeRequest(OnBeforeRequest(logger)).
		OnAfterResponse(OnAfterResponse(logger)).
		AddRetryCondition(func(r *resty.Response, err error) bool {
			return err != nil || r.StatusCode() >= 429
		})

	return c, nil
}

func DefaultClient(host string) (*resty.Client, error) {
	return NewClient(http.DefaultClient, log.DefaultLogger(), host)
}

func OnBeforeRequest(logger log.ILogger) func(*resty.Client, *resty.Request) error {
	return func(_ *resty.Client, r *resty.Request) error {
		logger.Info(r.Context(), "",
			zap.String("type", "request"),
			zap.String("method", r.Method),
			zap.String("path", r.URL),
			zap.Any("headers", r.Header),
			zap.Any("pathParams", r.PathParams),
			zap.Any("query", r.QueryParam),
			zap.Any("body", r.Body),
		)

		return nil
	}
}

func OnAfterResponse(logger log.ILogger) func(*resty.Client, *resty.Response) error {
	return func(_ *resty.Client, r *resty.Response) error {
		logger.Info(r.Request.Context(), "",
			zap.String("type", "response"),
			zap.Int("statusCode", r.StatusCode()),
			zap.Duration("time", r.Time()),
			zap.String("method", r.Request.Method),
			zap.String("path", r.Request.RawRequest.URL.Path),
			zap.Any("headers", r.Header()),
			zap.Any("query", r.Request.QueryParam),
			zap.Any("body", r.Body()))

		return nil
	}
}
