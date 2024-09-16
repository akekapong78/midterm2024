package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akekapong78/workflow/internal/item"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// load env
	// Absolute path or relative path to the .env file
	// err := godotenv.Load("/path/to/your/.env")
	// if err != nil {
	// 		log.Fatalf("Error loading .env file")
	// }

	// // Access environment variables
	// dbUser := os.Getenv("POSTGRES_USER")
	// dbPass := os.Getenv("POSTGRES_PASSWORD")

	err := godotenv.Load("../../../.env")
  if err != nil {
    log.Fatal("Error loading .env file")
  }
  
	// Connect database
	db, err := gorm.Open(
		postgres.Open(
			os.Getenv("DATABASE_URL"),
		),
	)
	if err != nil {
		log.Panic(err)
	}
	
	// Controller
	controller := item.NewController(db)

	// Router
	r := gin.Default()

	config := cors.DefaultConfig()

	// frontend URL
	config.AllowOrigins = []string{
		"http://localhost:8000",
		"http://127.0.0.1:8000",
	}

	r.Use(cors.New(config))
	
	// Register router
	r.POST("/items", controller.CreateItem)
	
	// Start server
	if err := r.Run(); err != nil {
		log.Panic(err)
	}
}