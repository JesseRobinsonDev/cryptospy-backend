package main

import (
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"cryptospy-backend/api/users/routes"
)

func main() {

	//godotenv.Load(".env")

	gin.SetMode(gin.ReleaseMode)

	// Environment variables
	port := os.Getenv("PORT")
	origins := strings.Split(os.Getenv("ORIGINS"), ",")

	if port == "" { port = "8000" }

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	// Initialize component API routes
	routes.InitUserRoutes(router)

	// Starts gin server
	router.Run(":" + port)
}