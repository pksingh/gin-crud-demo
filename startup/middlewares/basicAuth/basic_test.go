package basicAuth

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"

	"github.com/pksingh/gin-curd-demo/log"
	"github.com/pksingh/gin-curd-demo/startup/appProps"
)

func TestBasicAuthConfNotLoaded(t *testing.T) {
	t.Run("basic auth not loaded", func(t *testing.T) {
		assert.Nil(t, GetBasicAuthImpl())
	})
}

func TestBasicAuthApply(t *testing.T) {
	ctx := context.TODO()
	t.Setenv("runEnv", "prod")
	_ = appProps.Load("../../../resources")
	_ = log.Load(ctx)
	gin.SetMode(gin.TestMode)

	t.Run("basic auth env var not set", func(t *testing.T) {
		t.Setenv("CLIENT_BASIC_AUTH", ``)
		err := loadImpls(ctx)
		assert.ErrorContains(t, err, "invalid basic auth configuration")
	})

	t.Run("basic auth env var invalid", func(t *testing.T) {
		t.Setenv("CLIENT_BASIC_AUTH", `my-style`)
		err := loadImpls(ctx)
		assert.ErrorContains(t, err, "invalid basic auth configuration")
	})

	t.Run("basic auth env no username or password", func(t *testing.T) {
		t.Setenv("CLIENT_BASIC_AUTH", `[{"client-userid":"","client-pwd":"pass1","client-name":"name"}]`)
		err := loadImpls(ctx)
		assert.ErrorContains(t, err, "no basic auth users")
	})

	t.Run("happy path", func(t *testing.T) {
		t.Setenv("CLIENT_BASIC_AUTH", `[{"client-userid":"abcd","client-pwd":"efgh","client-name":"name"}]`)
		err := Load(ctx)
		assert.NoError(t, err)

		router := gin.Default()
		v1 := router.Group("/v1")
		assert.NotEmpty(t, GetBasicAuthImpl().Apply(ctx, v1))
	})
}
