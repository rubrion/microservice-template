package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ExampleHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from microservice-template!",
	})
}
