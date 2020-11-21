package log

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/pksingh/gin-curd-demo/config/jlog"
	"github.com/pksingh/gin-curd-demo/config/logperreq"
	"github.com/pksingh/gin-curd-demo/startup/appProps"
)

func TestVerifyLogConfHappy(t *testing.T) {
	appProps.Load("./../resources")
	err := Load(context.TODO())
	assert.NoError(t, err, "err not expected")
}

func TestBadLogConf(t *testing.T) {
	mock := jlog.ZapAppConf{}
	lgr := mock.CustomizeLogger(&zap.Config{
		EncoderConfig: zapcore.EncoderConfig{
			EncodeTime: nil,
			TimeKey:    "fcfhgv",
		}})
	assert.NotNil(t, lgr, "bad zap config should default zap prod config")
}

func TestGetLoggerWithNilCtxt(t *testing.T) {
	l := GetLogger(nil)
	assert.NotEqual(t, zap.NewNop(), l)
}

func TestGetLoggerWithCustomTypeCtxt(t *testing.T) {
	newRootCtx := context.WithValue(context.TODO(),
		correlationLoggerKey, zap.NewNop())
	l := GetLogger(newRootCtx)
	assert.Equal(t, zap.NewNop(), l)
}

func TestGetLoggerWithStringCtxt(t *testing.T) {
	newRootCtx := context.WithValue(context.TODO(),
		logperreq.CorrelationLoggerKeyStr, zap.NewNop())
	l := GetLogger(newRootCtx)
	assert.Equal(t, zap.NewNop(), l)
}

func TestGetLoggerWithNotFoundInCtxt(t *testing.T) {
	ctx := context.TODO()
	err := Load(ctx)
	assert.NoError(t, err, "err not expected")
	assert.Equal(t, GetConfiguredLogger(), GetLogger(ctx))
}

func TestGetRequestIdNotFoundInCtxt(t *testing.T) {
	reqId := logperreq.GetRequestId(context.TODO())
	assert.NotEmpty(t, reqId,
		"request id shouldn't be empty even "+
			"if context doesnt have request id info")
}
