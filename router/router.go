package router

import (
	"github.com/gin-gonic/gin"
	"jwt-practice/api"
)

func InitRouter() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	// user auth
	auth := router.Group("/user")
	{
		auth.POST("/auth", api.UserAuth)
	}
	// v1
	apiV1 := router.Group("/api/v1")
	apiV1.Use(api.JwtAuth())
	{
		apiV1.GET("/home", api.GetHome)
	}



	router.Run(":8080")
}