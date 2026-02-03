package main

import (
	"example.com/REST-API/db"
	"example.com/REST-API/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//loading the .env file.
	err := godotenv.Load()

	if err != nil {
		panic("Error!! failed to load .env")
	}

	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}


