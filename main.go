package main

import (
	"database/sql"
	"log"
	"os"
	"todo-service/models"
	"todo-service/routes"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
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

	// Insert Admin into the table if not exists
	_, err := models.GetUserByEmail(os.Getenv("ADMIN_EMAIL"))
	if err != nil {
		if err == sql.ErrNoRows {

			// Generate password hash for admin
			hashedPass, hashErr := bcrypt.GenerateFromPassword([]byte(os.Getenv("ADMIN_PASS")), bcrypt.DefaultCost)
			if hashErr != nil {
				log.Fatal(hashErr)
			}

			// Insert into admin into DB
			_, inserErr := models.DB.Exec(`
				INSERT INTO user
					(username, email, password, is_admin)
					VALUES
					(?, ?, ?, ?);
			`, os.Getenv("ADMIN_USER"), os.Getenv("ADMIN_EMAIL"), hashedPass, true)

			if inserErr != nil {
				log.Fatal(inserErr)
			}
		} else {
			log.Fatal(err)
		}
	}

	routes.AuthRoutes(router)
	routes.TodoRoutes(router)
	routes.AdminRoutes(router)

	// Run the server
	router.Run("localhost:8080")

}
