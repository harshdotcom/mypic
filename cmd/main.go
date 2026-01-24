package main

import (
	"fmt"
	"log"
	"mypic/config"
	"mypic/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	config.ConnectDB()
	r := gin.Default()

	// r.Use(middlewares.CORSMiddleware())

	r.Static("/uploads", "./uploads")
	routes.RegisterRoutes(r)
	fmt.Println("Server starting on http://localhost:8080")
	r.Run(":8080")
}
