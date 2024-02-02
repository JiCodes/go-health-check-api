package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func main() {

	router := gin.Default();
	router.GET("/healthz", HealthCheckHandler)
	fmt.Println("Server listening on :8080")
	router.Run(":8080")
}