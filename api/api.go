package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHome(c *gin.Context) {
	username := c.MustGet("username")

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  fmt.Sprintf("welcome %s to the home page", username),
	})
}
