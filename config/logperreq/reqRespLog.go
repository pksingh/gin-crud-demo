package logperreq

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// use custom context key to avoid hash collision with strings in context.
type customContextKey string

const (
	// correlationReqHeaderKey is the key which holds the correlation data (reqId) of
	// a single http.Request. If the request has a valid X-Request-Service-ID header, we derive it
	// from the header and use it for correlation or else we generate an uuid per request
	correlationReqHeaderKey customContextKey = "X-Request-Service-ID"

	// correlationLoggerKey is the key which holds the custom logger for the http.Request
	correlationLoggerKey customContextKey = "httpReqZapLogger"

	// interestedHttpReqKey is the key which differentiates a valid business use
	// case requests in http.Request and doesn't mark the management endpoints
	interestedHttpReqKey customContextKey = "isInterestedHttpReq"

	correlationIdNotFound = "req/resp correlation id not set."
)

var (
	// correlationReqHeaderKeyStr maintains string version of correlationReqHeaderKey
	// key as gin.Context can only set string type
	correlationReqHeaderKeyStr = customContextKeyStr(correlationReqHeaderKey)

	// CorrelationLoggerKeyStr maintains string version of correlationLoggerKey
	// key as gin.Context can only set string type
	CorrelationLoggerKeyStr = customContextKeyStr(correlationLoggerKey)

	// serviceNamePrefix is const prefix for every request ID, this contains
	// service name, branch name and k8s pod identifier
	serviceNamePrefix, _ = os.Hostname()

	// requestIncID is a counter for request ID
	requestIncID uint64
)

// InterestedEndpoints mark all the valid business use case requests in http.Request
func InterestedEndpoints() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		iCtx := context.WithValue(ctx.Request.Context(), interestedHttpReqKey, true)
		ctx.Request = ctx.Request.WithContext(iCtx)
		ctx.Next()
	}
}

// AttachRequestID will attach the reqID to the http.Request context,
// request header and add to http response header
// uniquePodId is a unique string generated on the startup of the app
// to avoid same reqIds on different pods for same git commit id
func AttachRequestID(uniquePodId string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if isInterestedHttpReq(ctx.Request.Context()) {
			reqID := ctx.GetHeader(correlationReqHeaderKeyStr)
			if len(strings.TrimSpace(reqID)) == 0 {
				reqID = makeNewReqId(uniquePodId)
			}
			ctx.Request.Header.Add(correlationReqHeaderKeyStr, reqID)
			ctx.Writer.Header().Add(correlationReqHeaderKeyStr, reqID)
			idCtx := context.WithValue(ctx.Request.Context(), correlationReqHeaderKey, reqID)
			ctx.Request = ctx.Request.WithContext(idCtx)
			ctx.Next()
		}
	}
}

// AttachRequestLogger will attach a custom zap logger (which contains txId) to http.Request
func AttachRequestLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if isInterestedHttpReq(ctx.Request.Context()) {
			reqLogger := logger.With(
				zap.String("reqId", GetRequestId(ctx.Request.Context())))
			ctx.Set(CorrelationLoggerKeyStr, reqLogger)
			logCtx := context.WithValue(ctx.Request.Context(), correlationLoggerKey, reqLogger)
			ctx.Request = ctx.Request.WithContext(logCtx)
			ctx.Next()
		}
	}
}

// helper to convert customContextKey type to string
func customContextKeyStr(key customContextKey) string {
	return string(key)
}

// GetRequestId retrieves correlationReqHeaderKey from context
func GetRequestId(ctx context.Context) string {
	corrReqId, ok := ctx.Value(correlationReqHeaderKey).(string)
	if !ok || len(corrReqId) == 0 {
		return correlationReqHeaderKeyStr + " " + correlationIdNotFound
	}
	return corrReqId
}

// isInterestedHttpReq checks if the http request is one of the valid business use case endpoints,
// we don't want to trace swagger endpoints, info endpoint, health endpoint, etc.
func isInterestedHttpReq(ctx context.Context) bool {
	v, found := ctx.Value(interestedHttpReqKey).(bool)
	return found && v
}

// makeNewReqId generates new requestId with 1 increment per request.
// uniquePodId is a random str generated during startup.
func makeNewReqId(uniquePodId string) (out string) {
	incId := atomic.AddUint64(&requestIncID, 1)
	unparsedReqID := fmt.Sprintf("%s-%s-%08d", serviceNamePrefix, uniquePodId, incId)
	reqIDContents := strings.Split(unparsedReqID, "-")
	if len(reqIDContents) > 4 {
		reqIDContents = reqIDContents[len(reqIDContents)-4:]
	}
	for _, s := range reqIDContents {
		out += s + "-"
	}
	// remove last "-"
	if last := len(out) - 1; last >= 0 && out[last] == '-' {
		out = out[:last]
	}
	return
}
