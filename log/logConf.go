package log

import (
	"context"

	"go.uber.org/zap"

	"github.com/pksingh/gin-curd-demo/config/logperreq"
)

// use custom context key to avoid hash collision with strings in context.
type customContextKey string

const (
	// correlationLoggerKey is the key which holds the custom logger for the http.Request
	correlationLoggerKey customContextKey = "httpReqZapLogger"
)

// GetLogger Logger returns a zap logger which is tied to that http.Request. Using this
// logger across the application gives tracing info of a request.
// ctx input interchangeably works with both inputs -  gin.Context or context.Context.
// It returns a default logger if no context information is found
func GetLogger(ctx context.Context) *zap.Logger {
	if ctx != nil {
		logger, ok := ctx.Value(correlationLoggerKey).(*zap.Logger)
		if ok && logger != nil {
			return logger
		}
		logger, ok = ctx.Value(logperreq.CorrelationLoggerKeyStr).(*zap.Logger)
		if ok && logger != nil {
			return logger
		}
		zapBaseLogger := GetConfiguredLogger()
		zapBaseLogger.Warn(ctxtLogNotFound + " " + logConfDefZapProd)
		return zapBaseLogger
	}
	zapBaseLogger := GetConfiguredLogger()
	zapBaseLogger.Error(ctxNil + " " + logConfDefZapProd)
	return zapBaseLogger
}
