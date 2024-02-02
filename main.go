package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Pragma", "no-cache")
	c.Header("X-Content-Type-Options", "nosniff")

	c.Status(http.StatusOK)
}

func main() {

	router := gin.Default();
	router.GET("/healthz", HealthCheckHandler)
	fmt.Println("Server listening on :8080")
	router.Run(":8080")
}