package server

import (
	"context"
	"net/http/pprof"

	"github.com/gin-gonic/gin"

	"github.com/pksingh/gin-curd-demo/handler"
)

// CreateServer makes the server that main will run.
// Includes all routing rules to business endpoints.
func (s ServeConf) CreateServer(c context.Context) *gin.Engine {

	// router := gin.New()
	router := gin.Default() // Enabled Req/Resp logging on console
	s.addDefaultEndpoints(router)
	s.addProfiling(router)

	// versioning to support http api request interfaces future extensibility.
	versionRtr := router.Group("/v1")

	// create a group for all endpoints which contains business logic and
	// apply logging, tracing, and metrics middlewares
	// TODO - change "/myapp"
	myAppRtr := versionRtr.Group("/myapp")

	s.addBusinessEndpointsMiddlewares(myAppRtr)

	return router
}

func (s ServeConf) addProfiling(router *gin.Engine) {
	if s.RunTimeProfilingEnabled {
		debugRouter := router.Group("/debug/pprof")
		debugRouter.GET("/", gin.WrapF(pprof.Index))
		debugRouter.GET("/cmdline", gin.WrapF(pprof.Cmdline))
		debugRouter.GET("/profile", gin.WrapF(pprof.Profile))
		debugRouter.GET("/symbol", gin.WrapF(pprof.Symbol))
		debugRouter.GET("/goroutine", gin.WrapH(pprof.Handler("goroutine")))
		debugRouter.GET("/heap", gin.WrapH(pprof.Handler("heap")))
		debugRouter.GET("/threadcreate", gin.WrapH(pprof.Handler("threadcreate")))
		debugRouter.GET("/block", gin.WrapH(pprof.Handler("block")))
	}
}

func (s ServeConf) addDefaultEndpoints(router *gin.Engine) {
	_ = router.SetTrustedProxies(nil)
	if s.InfoEndpointEnabled {
		router.GET("/info", handler.GetInfo)
	}
	router.GET("/health", handler.GetHealth)
}

func (s ServeConf) addBusinessEndpointsMiddlewares(router *gin.RouterGroup) {
	router.Use(s.InterestedEndpoint, s.AttachRequestID, s.AttachRequestLogger)
}
