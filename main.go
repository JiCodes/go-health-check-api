package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB 

func connectDB() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
    log.Fatal("Error loading .env file")
  }

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func HealthCheckHandler(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Pragma", "no-cache")
	c.Header("X-Content-Type-Options", "nosniff")

	if err := db.Exec("SELECT 1").Error; err == nil {
		c.Status(http.StatusOK) 
	} else {
		c.Status(http.StatusServiceUnavailable)
	}
}

func CheckMethodMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/healthz" && c.Request.Method != http.MethodGet {
			c.Status(http.StatusMethodNotAllowed)
			return
		}
		c.Next()
	}
}

func CheckPayloadMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.ContentLength > 0 {
			c.Status(http.StatusBadRequest)
			return
		}
		c.Next()
	}
}

func main() {

	var err error
	db, err = connectDB()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
	}

	router := gin.Default();

	router.Use(CheckMethodMiddleware())
	router.Use(CheckPayloadMiddleware())

	router.GET("/healthz", HealthCheckHandler)
	router.Run(":8080")
	fmt.Println("Server listening on :8080")
}