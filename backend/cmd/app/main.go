package main

import (
	"log"
	"os"

	"github.com/akekapong78/workflow/internal/item"
	"github.com/akekapong78/workflow/internal/middleware"
	"github.com/akekapong78/workflow/internal/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// Load env
	err := godotenv.Load("./../.env")
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

	// Router
	r := gin.Default()

	// CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:8000", // Frontend url
		"http://127.0.0.1:8000",
	}

	r.Use(cors.New(config))

	// New Controller
	userController := user.NewController(db, os.Getenv("JWT_SECRET"), os.Getenv("BACKEND_DOMAIN"))
	itemController := item.NewController(db)

	// Routes Group
	v1 := r.Group("/api/v1")

	// User Group
	userGroup := v1.Group("/users")
	// Without middleware
	userGroup.POST("/register", userController.Register)
	userGroup.POST("/login", userController.Login)
	// With middleware
	userGroup.Use(middleware.Guard(os.Getenv("JWT_SECRET")))
	userGroup.GET("/", userController.GetUsers)
	userGroup.GET("/:id", userController.GetUser)

	// Item Group
	itemGroup := v1.Group("/items")
	itemGroup.Use(middleware.Guard(os.Getenv("JWT_SECRET")))
	itemGroup.POST("/", itemController.CreateItem)
	itemGroup.GET("/:id", itemController.GetItem)
	itemGroup.GET("/", itemController.GetItems)
	itemGroup.PUT("/:id", itemController.UpdateItem)
	itemGroup.PATCH("/:id", middleware.CheckAdminRole, itemController.UpdateItemStatus)

	// Get the port from the environment variable
	port := os.Getenv("PORT_API")
	if port == "" {
		port = "8080" // Default port
	}

	// Start server
	if err := r.Run(":" + port); err != nil {
		log.Panic(err)
	}
}
