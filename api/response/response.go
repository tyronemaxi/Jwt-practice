package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, jsonStr interface{}) {
	code := http.StatusOK
	msg := "success"

	if jsonStr == nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code": code,
				"msg": msg,
			},
		)
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code": code,
				"msg": msg,
				"data": jsonStr,
			},
		)
	}
}

func Fail(c *gin.Context, jsonStr interface{}) {
	code := http.StatusInternalServerError
	msg := "Fail"

	c.JSON(
		http.StatusOK,
		gin.H{
			"code": code,
			"msg": msg,
			"data": jsonStr,
		},
	)
}



