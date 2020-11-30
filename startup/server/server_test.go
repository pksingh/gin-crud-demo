package server

import (
	"context"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/pksingh/gin-curd-demo/startup/middlewares/basicAuth"
)

func TestCreateServer(_ *testing.T) {
	gin.SetMode(gin.TestMode)
	mockBa := basicAuth.BasicAuthImpl{
		BasicAuthUsers: gin.Accounts{"a": "b"},
	}
	mock := ServeConf{
		Auth:                    mockBa,
		RunTimeProfilingEnabled: true,
		InfoEndpointEnabled:     true,
	}
	mock.CreateServer(context.TODO())
}
