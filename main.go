package main

import (
	"log"
	"os"
	"todo-service/models"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// Create new Gin router instance
	router := gin.Default()

	// Load .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	// Create database config
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Addr:                 os.Getenv("DBADDR"),
		DBName:               os.Getenv("DBNAME"),
		AllowNativePasswords: true,
	}

	// Connect to DB and initialize DB variable
	models.InitDB(cfg)

	// Run the server
	router.Run("localhost:8080")

}
