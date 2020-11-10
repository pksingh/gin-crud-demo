package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/pksingh/gin-curd-demo/startup/appProps"
)

func TestGetInfo(t *testing.T) {
	_ = appProps.Load("./../resources")
	router := gin.New()
	gin.SetMode(gin.TestMode)
	router.GET("/info", GetInfo)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/info", nil)
	router.ServeHTTP(w, req)
	assert.NotEmpty(t, w)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body)
	assert.Contains(t, w.Body.String(), "appName")
	assert.ObjectsAreEqual(appInfo, w.Body)
	// val := w.Body.(AppInfo)
	// assert.Equal(t, "XYZ", w.Body.(AppInfo).AppName)
	// assert.EqualValues(t, appInfo{appInfo.AppName: "xyz"}, w.Body)

	//assert.ElementsMatch(t, w.Body, appInfo)
	// assert.JSONEq(t,appInfo.String(), w.Body.String())
}
