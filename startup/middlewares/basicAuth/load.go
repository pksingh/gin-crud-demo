package basicAuth

import (
	"context"
	"encoding/json"
	"errors"
	"sync"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/pksingh/gin-curd-demo/log"
	"github.com/pksingh/gin-curd-demo/startup/appProps"
)

var (
	onceInit      = new(sync.Once)
	concreteImpls = make(map[string]interface{})
)

const (
	basicAuthInitKey = "basicAuthImpl"
)

// BasicAuthImpl Basic Auth Implementation type
type BasicAuthImpl struct {
	BasicAuthUsers gin.Accounts
	BasicAuthRealm string
}

// User is the structure for Basic auth user
type User struct {
	ClientUserId   string `json:"client-userid"`
	ClientPassword string `json:"client-pwd"`
	ClientName     string `json:"client-name"`
}

// Auth Basic Auth Implementation type
type Auth interface {
	Apply(ctx context.Context, grp *gin.RouterGroup) *gin.RouterGroup
}

// Load enables us load basic auth users
func Load(ctx context.Context) error {
	var appErr error

	onceInit.Do(func() {
		appErr = loadImpls(ctx)
	})
	return appErr
}

func loadImpls(ctx context.Context) error {
	logger := log.GetLogger(ctx)
	p := appProps.GetAll()
	if concreteImpls[basicAuthInitKey] == nil {
		clientBasicAuth := p.MustGetString("basic.auth.credentials")

		c := []byte(clientBasicAuth)
		var clientUserArr []*User

		// Check failure
		err := json.Unmarshal(c, &clientUserArr)
		if err != nil {
			return errors.New("invalid basic auth configuration")
		}

		basicAuthUsers := make(gin.Accounts)
		// create a map of users with key as userid and value as password
		for _, user := range clientUserArr {
			if len(user.ClientUserId) > 0 && len(user.ClientPassword) > 0 {
				basicAuthUsers[user.ClientUserId] = user.ClientPassword
			} else {
				logger.Error("bad configuration for basic auth. ignoring user",
					zap.String("clientName", user.ClientName))
			}
		}

		if len(basicAuthUsers) == 0 {
			return errors.New("no basic auth users configured")
		}

		concreteImpls[basicAuthInitKey] = BasicAuthImpl{
			BasicAuthUsers: basicAuthUsers,
		}
	}
	return nil
}

// GetBasicAuthImpl Getting basic auth implementation
func GetBasicAuthImpl() Auth {
	x, _ := concreteImpls[basicAuthInitKey].(Auth)
	return x
}
