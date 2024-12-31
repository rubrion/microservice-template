package routes

import (
	"microservice-template/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/example", controllers.ExampleHandler)
	}
}
