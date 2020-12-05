package db

import (
	"context"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pksingh/gin-curd-demo/log"
	"github.com/pksingh/gin-curd-demo/startup/appProps"
)

// Mock the config struct viper
type MockViper struct{}

func (viper *MockViper) GetString(key string) string {
	return "debug"
}

func (viper *MockViper) GetInt(key string) int {
	return 0
}

func (viper *MockViper) GetBool(key string) bool {
	return true
}

// func Open(driverName, dataSourceName string) (*DB, error)

func TestDB(t *testing.T) {
	ctx := context.TODO()
	t.Setenv("runEnv", "dev")
	_ = appProps.Load("../../../resources")
	_ = log.Load(ctx)
	gin.SetMode(gin.TestMode)
	loadLog(t)
	InitDB(ctx)
}

func TestDBError(t *testing.T) {
	ctx := context.TODO()
	t.Setenv("runEnv", "test")
	_ = appProps.Load("../../../resources")
	_ = log.Load(ctx)
	gin.SetMode(gin.TestMode)
	InitDB(ctx)
}

func loadLog(t *testing.T) {
	ctx := context.TODO()
	t.Setenv("runEnv", "dev")
	_ = appProps.Load("../../../resources")
	_ = log.Load(ctx)
}
