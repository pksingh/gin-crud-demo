package main

import (
	"context"

	"go.uber.org/zap"

	"github.com/pksingh/gin-curd-demo/log"
	"github.com/pksingh/gin-curd-demo/startup/appProps"
	"github.com/pksingh/gin-curd-demo/startup/db"
	"github.com/pksingh/gin-curd-demo/startup/middlewares/basicAuth"
	"github.com/pksingh/gin-curd-demo/startup/server"
)

func main() {

	// Create a background context for a long-running application
	ctx := context.Background()

	// Create a default zap startupLogger for logging app startup operations.
	startupLogger, _ := zap.NewProduction()
	startupLogger = startupLogger.With(zap.String("action", "app startup"))

	// Initially, always load all app properties
	handleStartUpErr(startupLogger, appProps.Load("./resources"))

	// load log config
	handleStartUpErr(startupLogger, log.Load(ctx))

	// load db with all properties into appProps
	handleStartUpErr(startupLogger, db.Load(ctx))

	// load basic auth creds
	handleStartUpErr(startupLogger, basicAuth.Load(ctx))

	// load server with all properties into server.ServeConf
	handleStartUpErr(startupLogger, server.Load(ctx))

	// create router from server.ServeConf
	router := server.GetStartupServerConf().CreateServer(ctx)

	handleStartUpErr(startupLogger, router.Run())
	startupLogger.Info("service up and listening",
		zap.String("port", appProps.GetAll().MustGetString("app.port")))
}

// handleStartUpErr makes sure that app fails to start in case of
// invalid app configurations
func handleStartUpErr(logger *zap.Logger, err error) {
	if err != nil {
		logger.Panic("failed to start application",
			zap.NamedError("appCrashed", err))
	}
}
