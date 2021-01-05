package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pksingh/gin-curd-demo/model"
)

func DeleteUser(c *gin.Context) {
	if err := model.DeleteSingleUser(c); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"error": "DELETE record SUCCESS"})
	}
}
