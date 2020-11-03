package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/pksingh/gin-curd-demo/log"
	"github.com/pksingh/gin-curd-demo/startup/appProps"
	"github.com/pksingh/gin-curd-demo/startup/server"
)

func TestMain(m *testing.M) {
	// do stuff like setup before running all unit tests
	ctx := context.TODO()
	_ = os.Setenv("runEnv", "test")
	_ = appProps.Load("./resources")
	_ = log.Load(ctx)
	_ = server.Load(ctx)
	gin.SetMode(gin.TestMode)
	go main()

	c := m.Run()
	os.Exit(c)
}

func TestStartUpErr(t *testing.T) {

	someErr := fmt.Errorf("i want a startup err")

	// recover from panic if one occurred. Set err to nil otherwise.
	defer func() {
		err := recover()
		if e, ok := err.(error); ok {
			if !errors.Is(someErr, e) {
				t.Errorf("expecting %s error", someErr.Error())
			}
		}
	}()

	l, _ := zap.NewDevelopment()
	handleStartUpErr(l, someErr)
}
