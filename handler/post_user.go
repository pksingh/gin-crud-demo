package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pksingh/gin-curd-demo/model"
)

func InsertUser(c *gin.Context) {
	if err := model.InsertSingleUser(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "INSERT one record SUCCESS"})
	}
}
