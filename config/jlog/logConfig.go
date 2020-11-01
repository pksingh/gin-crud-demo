package jlog

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapAppConf provides a way to customize a zap logger.
type ZapAppConf struct {
	ShowCallerInLogs       bool
	ShowStackTraceInLogs   bool
	UseUnstructuredLogging bool
	AppLogLevel            string
}

// set of all log levels accepted by zap
const (
	Debug  = "debug"
	Info   = "info"
	Warn   = "warn"
	Error  = "error"
	DPanic = "dpanic"
	Panic  = "panic"
	Fatal  = "fatal"
)

const (
	logConfDefZapProd = "defaulting to zap-prod configuration."
	logConfBad        = "failed to initialize zap logger with given configuration."
)

// CustomizeLogger - add logic here if more tuning is required to control log behaviour per env.
// in case of invalid zap configuration from user, we default to zapProd
// default configuration in VerifyLogConf
func (z ZapAppConf) CustomizeLogger(zapConf *zap.Config) *zap.Logger {
	zapConf.Level = zap.NewAtomicLevelAt(fetchLogLevel(z.AppLogLevel))
	zapConf.DisableCaller = !z.ShowCallerInLogs
	zapConf.DisableStacktrace = !z.ShowStackTraceInLogs
	if z.UseUnstructuredLogging || len(os.Getenv("configEnvironment")) == 0 {
		zapConf.Encoding = "console"
	}
	zapLogger, zapConfigErr := zapConf.Build()
	if zapConfigErr != nil {
		*zapConf = zap.NewProductionConfig()
		zapLogger, _ = zapConf.Build()
		zapLogger.Error(logConfBad+logConfDefZapProd, zap.Error(zapConfigErr))
	}
	return zapLogger
}

func fetchLogLevel(appLogLevel string) zapcore.Level {
	switch strings.ToLower(appLogLevel) {
	case Debug:
		return zapcore.DebugLevel
	case Info:
		return zapcore.InfoLevel
	case Warn:
		return zapcore.WarnLevel
	case Error:
		return zapcore.ErrorLevel
	case DPanic:
		return zapcore.DPanicLevel
	case Panic:
		return zapcore.PanicLevel
	case Fatal:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

// BaseZapConf gives base default zap configuration
func BaseZapConf() *zap.Config {
	return &zap.Config{
		Level:             zap.NewAtomicLevelAt(zapcore.InfoLevel),
		DisableCaller:     true,
		DisableStacktrace: true,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:       "lvl",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeTime:     zapcore.RFC3339TimeEncoder,
			MessageKey:     "msg",
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
}
