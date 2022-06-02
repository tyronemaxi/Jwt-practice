package api

import (
	"github.com/gin-gonic/gin"
	"website/util"
	"net/http"
)

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func UserAuth(c *gin.Context) {
	var token string
	var err error

	testUser := UserInfo{
		"tyrone",
		"123456",
	}

	username, password := c.PostForm("username"), c.PostForm("password")

	// check the user correct
	if username == testUser.Username && password == testUser.Password {
		token, err = util.GenerateJwtToken(username, password)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"code": http.StatusForbidden,
				"msg":  TokenGenerateFail,
				"data": gin.H{
					"token": nil,
				},
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": gin.H{
			"token": token,
		},
	})
}
