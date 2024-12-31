package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"microservice-template/internal/routes"
	"microservice-template/pkg/logger"
)

func main() {
	logger.Init()
	router := gin.Default()
	routes.RegisterRoutes(router)

	port := "8080"
	log.Printf("Starting server on port %s...", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
