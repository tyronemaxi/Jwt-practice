package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"jwt-practice/util"
)

func JwtAuth() gin.HandlerFunc{
	return func(c *gin.Context) {
		var claims *util.Claims
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "",
				"data": nil,
			})
			c.Abort()
			return
		}

		var err error
		claims, err = util.ParseJwtToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "",
				"data": nil,
			})

			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}

}
