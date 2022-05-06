package router

import (
	"github.com/gin-gonic/gin"
	"jwt-practice/api"
	"jwt-practice/middleware"
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
	apiV1.Use(middleware.JwtAuth())
	{
		apiV1.GET("/home", api.GetHome)
	}



	router.Run(":8080")
}