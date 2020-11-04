package logperreq

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func TestMiddlewaresInjections(t *testing.T) {
	svr := gin.New()
	svr.Use(InterestedEndpoints())
	svr.Use(AttachRequestID("test"))
	svr.Use(AttachRequestLogger(zap.NewNop()))
	resp := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/", nil)
	svr.ServeHTTP(resp, req)
	expectedXReqId := resp.Header().Get(correlationReqHeaderKeyStr)
	if expectedXReqId == "" {
		t.Errorf("expectedXReqId shouldn't be empty")
	}
}

// verify the traceId generated in kube environments
func TestNewReqId(t *testing.T) {
	serviceNamePrefix = "dev-mp-sales-tax-service-abcdef-cccc"
	requestIncID = 23
	out := makeNewReqId("vfnsvjnsl")
	expected := "abcdef-cccc-vfnsvjnsl-00000024"
	if out != expected {
		t.Errorf("new reqId incorrectly generated")
	}
}

func TestGetRequestIdNoCtx(t *testing.T) {
	got := GetRequestId(context.TODO())
	expected := "X-Request-Service-ID req/resp correlation id not set."
	if got != expected {
		t.Errorf("GetRequestId() without context should return %s, but recieved %s", expected, got)
	}
}
