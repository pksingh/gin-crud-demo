package db

import (
	"context"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pksingh/gin-curd-demo/log"
	"github.com/pksingh/gin-curd-demo/startup/appProps"
	"github.com/stretchr/testify/assert"
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
	_ = appProps.Load("../../resources")
	_ = log.Load(ctx)
	gin.SetMode(gin.TestMode)
	err := InitDB(ctx)
	assert.NoError(t, err)
}

func TestDBError(t *testing.T) {
	ctx := context.TODO()
	t.Setenv("runEnv", "dev")
	t.Setenv("DATABASE_NAME", "xyz")
	_ = appProps.Load("../../resources")
	_ = log.Load(ctx)
	gin.SetMode(gin.TestMode)
	err := InitDB(ctx)
	if assert.Error(t, err) {
		assert.ErrorContains(t, err, "failed to connect")
		assert.ErrorContains(t, err, "server error")
	}
}
