package server

import (
	"context"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/pksingh/gin-curd-demo/log"
	"github.com/pksingh/gin-curd-demo/startup/appProps"
)

func TestCreateServerInit(t *testing.T) {

	ctx := context.TODO()
	t.Setenv("runEnv", "dev")
	_ = appProps.Load("../../resources")
	_ = log.Load(ctx)
	gin.SetMode(gin.TestMode)

	assert.NoError(t, Load(ctx))
	assert.NotEmpty(t, GetStartupServerConf())
}

func Test_randSeq(t *testing.T) {
	desiredLen := 8
	if got, err := randSeq(desiredLen); err != nil || len(got) != desiredLen {
		t.Errorf("randSeq() = got str of len %d, want str of len %d", len(got), desiredLen)
	}
}

func Test_generateRandomBytes(t *testing.T) {
	_, err := generateRandomBytes(-1)
	assert.Error(t, err)
}
