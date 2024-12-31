package routes

import (
	"github.com/gin-gonic/gin"
	"microservice-template/internal/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/example", controllers.ExampleHandler)
	}
}
