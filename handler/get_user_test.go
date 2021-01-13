package handler

import (
	"context"
	glog "log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/pksingh/gin-curd-demo/log"
	"github.com/pksingh/gin-curd-demo/startup/appProps"
	"github.com/pksingh/gin-curd-demo/startup/db"
)

func TestGetUser(t *testing.T) {
	ctx := context.TODO()
	_ = os.Setenv("runEnv", "dev")
	_ = appProps.Load("./../resources")
	_ = log.Load(ctx)
	_ = db.Load(ctx)
	router := gin.New()
	gin.SetMode(gin.TestMode)
	router.GET("/user", GetUser)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user?user_id=1", nil)
	router.ServeHTTP(w, req)
	glog.Println("resp: ", w)
	assert.NotEmpty(t, w)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body)

	str := w.Body.String()
	assert.NotContains(t, str, "error")
	assert.NotContains(t, str, "record NOT Found")

	assert.Contains(t, str, "u_user_id")
	assert.Contains(t, str, "u_account_id")
	assert.Contains(t, str, "u_contact_id")
	assert.Contains(t, str, "u_loyalty_id")

	// assert.JSONEq(t,appInfo.String(), w.Body.String())
}

func TestGetUserInvalidRequest(t *testing.T) {
	ctx := context.TODO()
	_ = os.Setenv("runEnv", "dev")
	_ = appProps.Load("./../resources")
	_ = log.Load(ctx)
	_ = db.Load(ctx)
	router := gin.New()
	gin.SetMode(gin.TestMode)
	router.GET("/user", GetUser)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user", nil)
	router.ServeHTTP(w, req)
	glog.Println("resp: ", w)
	assert.NotEmpty(t, w)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotEmpty(t, w.Body)

	str := w.Body.String()
	assert.Contains(t, str, "error")
	assert.Contains(t, str, "all mandatory values NOT Passed")

	// assert.JSONEq(t,appInfo.String(), w.Body.String())
}

func TestGetUserNoRecordFound(t *testing.T) {
	ctx := context.TODO()
	_ = os.Setenv("runEnv", "dev")
	// _ = os.Setenv("DATABASE_NAME", "norec")
	_ = appProps.Load("./../resources")
	_ = log.Load(ctx)
	_ = db.Load(ctx)
	router := gin.New()
	gin.SetMode(gin.TestMode)
	router.GET("/user", GetUser)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user?user_id=12345", nil)
	router.ServeHTTP(w, req)
	glog.Println("resp: ", w)
	assert.NotEmpty(t, w)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body)

	str := w.Body.String()
	assert.Contains(t, str, "error")
	assert.Contains(t, str, "record NOT Found")
	// assert.JSONEq(t,appInfo.String(), w.Body.String())
}

func TestGetUserBadRequest(t *testing.T) {
	ctx := context.TODO()
	_ = os.Setenv("runEnv", "dev")
	_ = os.Setenv("DATABASE_NAME", "nil")
	_ = appProps.Load("./../resources")
	_ = log.Load(ctx)
	_ = db.Load(ctx)
	router := gin.New()
	gin.SetMode(gin.TestMode)
	router.GET("/user", GetUser)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user?user_id=12345", nil)
	router.ServeHTTP(w, req)
	glog.Println("resp: ", w)
	assert.NotEmpty(t, w)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotEmpty(t, w.Body)

	str := w.Body.String()
	assert.Contains(t, str, "error")
	assert.Contains(t, str, "ERROR")
	assert.Contains(t, str, "SQLSTATE")
	// assert.JSONEq(t,appInfo.String(), w.Body.String())
}
