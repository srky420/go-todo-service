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

	// Insert Admin into the table if not exists
	// Generate password hash
	// hashedPass, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("ADMIN_PASS")), bcrypt.DefaultCost)
	// _, err = models.DB.Exec(`
	// 	INSERT INTO user
	// 		(username, email, password, is_admin)
	// 		VALUES
	// 		(?, ?, ?, ?);
	// `, os.Getenv("ADMIN_USER"), os.Getenv("ADMIN_EMAIL"), hashedPass, true)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	routes.AuthRoutes(router)
	routes.TodoRoutes(router)

	// Run the server
	router.Run("localhost:8080")

}
