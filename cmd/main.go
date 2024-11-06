package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"cinema-project-go/internal/database"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitDB()

	r := gin.Default()

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
