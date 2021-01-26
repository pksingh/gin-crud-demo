package handler

import (
	"context"
	glog "log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/pksingh/gin-curd-demo/log"
	"github.com/pksingh/gin-curd-demo/startup/appProps"
	"github.com/pksingh/gin-curd-demo/startup/db"
)

func TestInsertUserBadRequest(t *testing.T) {
	ctx := context.TODO()
	_ = os.Setenv("runEnv", "dev")
	_ = os.Setenv("DATABASE_NAME", "nil")
	_ = appProps.Load("./../resources")
	_ = log.Load(ctx)
	_ = db.Load(ctx)
	router := gin.New()
	gin.SetMode(gin.TestMode)
	router.POST("/user", InsertUser)
	w := httptest.NewRecorder()
	body := "{\"u_user_id\":5,\"u_account_id\":21,\"u_contact_id\":22,\"u_loyalty_id\":23,\"u_is_active_id\":true,\"u_reference_id\":21,\"u_user_type\":\"u\",\"u_account_type\":\"a\",\"u_loyalty_type\":\"l\",\"u_member_type\":\"m\",\"u_brand_type\":\"b\",\"u_create_rcd_at\":\"2022-09-13T15:11:48+05:30\",\"u_create_rcd_by_who\":\"w\",\"u_create_rcd_by_app\":\"a\",\"u_update_rcd_at\":\"2022-09-13T15:12:12+05:30\",\"u_update_rcd_by_who\":\"w\",\"u_update_rcd_by_app\":\"a\",\"u_data_source\":\"s\"}"
	// body := "{\"u_user_id\":1,\"u_account_id\":21,\"u_contact_id\":22,\"u_loyalty_id\":23,\"u_is_active_id\":true,\"u_reference_id\":21}"
	req, _ := http.NewRequest(http.MethodPost, "/user?user_id=1", strings.NewReader(body))
	router.ServeHTTP(w, req)
	glog.Println("resp: ", w)
	assert.NotEmpty(t, w)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotEmpty(t, w.Body)

	str := w.Body.String()
	assert.NotContains(t, str, "record NOT Found")

	assert.NotContains(t, str, "message")
	assert.NotContains(t, str, "SUCCESS")

	assert.Contains(t, str, "error")
	assert.Contains(t, str, "executing")
	assert.Contains(t, str, "INSERT")
	// assert.JSONEq(t,appInfo.String(), w.Body.String())
}
