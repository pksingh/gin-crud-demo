package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pksingh/gin-curd-demo/model"
)

// ListUsers returns List of Users.
func ListUsers(c *gin.Context) {
	if res, err := model.GetAllUsers(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, res)
	}
}
