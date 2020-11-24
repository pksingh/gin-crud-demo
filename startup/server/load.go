package server

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"sync"

	"github.com/gin-gonic/gin"

	"github.com/pksingh/gin-curd-demo/config/logperreq"
	"github.com/pksingh/gin-curd-demo/log"
	"github.com/pksingh/gin-curd-demo/startup/appProps"
	"github.com/pksingh/gin-curd-demo/startup/middlewares/basicAuth"
)

var (
	onceInit      = new(sync.Once)
	concreteImpls = make(map[string]interface{})
)

// RouterConfKey
const (
	RouterConfKey = "ServerStartupConfiguration"
)

// ServeConf holds all the configuration data needed for starting HTTP Gin server.
type ServeConf struct {
	RunTimeProfilingEnabled bool
	InfoEndpointEnabled     bool

	InterestedEndpoint  gin.HandlerFunc
	AttachRequestID     gin.HandlerFunc
	AttachRequestLogger gin.HandlerFunc

	Auth basicAuth.Auth
}

// Load enables us enable/disable specific endpoints from app configurations
func Load(ctx context.Context) error {
	var appErr error
	onceInit.Do(func() {
		appErr = loadImpls(ctx)
	})
	return appErr
}

func loadImpls(_ context.Context) error {
	logger := log.GetConfiguredLogger()
	if concreteImpls[RouterConfKey] == nil {
		randStr, _ := randSeq(6)
		p := appProps.GetAll()
		concreteImpls[RouterConfKey] = ServeConf{
			RunTimeProfilingEnabled: p.GetBool("runtime.profiling.enabled", false),
			InfoEndpointEnabled:     p.GetBool("info.endpoint.enabled", false),

			InterestedEndpoint:  logperreq.InterestedEndpoints(),
			AttachRequestLogger: logperreq.AttachRequestLogger(logger),
			AttachRequestID:     logperreq.AttachRequestID(randStr),

			Auth: basicAuth.GetBasicAuthImpl(),
		}
	}
	return nil
}

// GetStartupServerConf Getting startup configuration
func GetStartupServerConf() ServeConf {
	v, _ := concreteImpls[RouterConfKey].(ServeConf)
	return v
}

// randSeq returns a URL-safe, base64 encoded
// securely generated random string.
func randSeq(s int) (string, error) {
	b, err := generateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b)[:s], err
}

// generateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func generateRandomBytes(n int) ([]byte, error) {
	if n < 1 {
		return nil, errors.New("invalid length")
	}
	b := make([]byte, n)
	_, err := rand.Read(b)
	return b, err
}
