package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pksingh/gin-curd-demo/model"
)

func GetUser(c *gin.Context) {
	if res, err := model.GetSingleUser(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, res)
	}
}
