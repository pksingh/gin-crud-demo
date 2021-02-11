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

func TestDeleteUser(t *testing.T) {
	ctx := context.TODO()
	_ = os.Setenv("runEnv", "dev")
	_ = appProps.Load("./../resources")
	_ = log.Load(ctx)
	_ = db.Load(ctx)
	router := gin.New()
	gin.SetMode(gin.TestMode)
	router.DELETE("/user", DeleteUser)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/user?user_id=3", nil)
	router.ServeHTTP(w, req)
	glog.Println("resp: ", w)
	assert.NotEmpty(t, w)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body)

	str := w.Body.String()
	assert.NotContains(t, str, "error")
	assert.NotContains(t, str, "record NOT Found")

	assert.Contains(t, str, "message")
	assert.Contains(t, str, "DELETE record SUCCESS")
}

func TestDeleteUserInvalidRequest(t *testing.T) {
	ctx := context.TODO()
	_ = os.Setenv("runEnv", "dev")
	_ = appProps.Load("./../resources")
	_ = log.Load(ctx)
	_ = db.Load(ctx)
	router := gin.New()
	gin.SetMode(gin.TestMode)
	router.DELETE("/user", DeleteUser)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/user", nil)
	router.ServeHTTP(w, req)
	glog.Println("resp: ", w)
	assert.NotEmpty(t, w)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotEmpty(t, w.Body)

	str := w.Body.String()
	assert.Contains(t, str, "error")
	assert.Contains(t, str, "all mandatory values NOT Passed")
}

func TestDeleteUserBadRequest(t *testing.T) {
	ctx := context.TODO()
	_ = os.Setenv("runEnv", "dev")
	_ = os.Setenv("DATABASE_NAME", "nil")
	_ = appProps.Load("./../resources")
	_ = log.Load(ctx)
	_ = db.Load(ctx)
	router := gin.New()
	gin.SetMode(gin.TestMode)
	router.DELETE("/user", DeleteUser)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/user?user_id=12345", nil)
	router.ServeHTTP(w, req)
	glog.Println("resp: ", w)
	assert.NotEmpty(t, w)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotEmpty(t, w.Body)

	str := w.Body.String()
	assert.Contains(t, str, "error")
	assert.Contains(t, str, "executing")
	assert.Contains(t, str, "DELETE")
}
