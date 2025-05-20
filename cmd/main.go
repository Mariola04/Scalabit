package main

import (
	
	"log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/Mariola04/Scalabit/internal/handlers"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env")
	}
}

func main() {

	// Initializes HTTP
	router := gin.Default()

	// Defines API Routes
	router.POST("/repos", handlers.CreateRepo)
	router.DELETE("/repos/:owner/:repo", handlers.DeleteRepo)
	router.GET("/repos", handlers.ListRepos)
	router.GET("/repos/:owner/:repo/pulls", handlers.ListPullRequests)

	// Initialize Port (8080)
	log.Println("Server initialized on http://localhost:8080")
	router.Run(":8080")
}
