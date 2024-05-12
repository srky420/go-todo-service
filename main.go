package main

import (
	"log"
	"os"
	"todo-service/models"
	"todo-service/routes"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// Create new Gin router instance
	router := gin.Default()

	// Load .env
	envErr := godotenv.Load(".env")
	if envErr != nil {
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

	// Create tables
	tablesErr := models.CreateTables()
	if tablesErr != nil {
		log.Fatal(tablesErr)
	}

	// Create admin
	adminErr := models.CreateAdmin()
	if adminErr != nil {
		log.Fatal(adminErr)
	}

	// Attach routes
	routes.AuthRoutes(router)
	routes.TodoRoutes(router)
	routes.AdminRoutes(router)

	// Run the server
	router.Run("localhost:8080")
}
