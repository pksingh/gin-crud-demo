package basicAuth

import (
	"context"

	"github.com/gin-gonic/gin"
)

// Apply - BasicAuth Implementation
func (b BasicAuthImpl) Apply(_ context.Context, grp *gin.RouterGroup) *gin.RouterGroup {
	return grp.Group("/", gin.BasicAuthForRealm(b.BasicAuthUsers, b.BasicAuthRealm))
}
