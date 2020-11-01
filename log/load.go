package log

import (
	"context"
	"sync"

	"go.uber.org/zap"

	"github.com/pksingh/gin-curd-demo/config/jlog"
	"github.com/pksingh/gin-curd-demo/startup/appProps"
)

var (
	// onceInit guarantees initialization of custom zap Logger only once
	onceInit      = new(sync.Once)
	concreteImpls = make(map[string]interface{})
)

const (
	configuredZapLogKey = "configuredZapLogger"
)

const (
	ctxtLogNotFound   = "logger not found in context."
	logConfDefZapProd = "defaulting to zap-prod configuration."
	ctxNil            = "input context is nil."
)

// Load implementations
func Load(ctx context.Context) error {
	var appErr error
	onceInit.Do(func() {
		appErr = loadImpls(ctx)
	})
	return appErr
}

// loadImpls calls function in od-go/v2/jlog.BaseZapConf to initialize zapBaseLogger
//
//	zapConf is modified depending on environment and properties by calling jlog.CustomizeLogger
//
// Multiple times modification of the zapConf and multi initialization of zapBaseLogger
// is avoided by the protection provided by onceInit
func loadImpls(_ context.Context) error {
	if concreteImpls[configuredZapLogKey] == nil {
		props := appProps.GetAll()
		zapAppConf := jlog.ZapAppConf{
			ShowCallerInLogs:       props.GetBool("log.show.caller", false),
			ShowStackTraceInLogs:   props.GetBool("log.show.stacktrace", false),
			UseUnstructuredLogging: props.GetBool("log.use.unstructured.logging", false),
			AppLogLevel:            props.GetString("log.level", "info"),
		}

		// get base default zap configuration
		baseZapConf := jlog.BaseZapConf()

		concreteImpls[configuredZapLogKey] = zapAppConf.CustomizeLogger(baseZapConf)
	}
	return nil
}

// GetConfiguredLogger getting zap logger
func GetConfiguredLogger() *zap.Logger {
	v := concreteImpls[configuredZapLogKey]
	return v.(*zap.Logger)
}
