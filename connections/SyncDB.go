package connections

import (
	"fmt"
	"log"

	models "github.com/Hrishikesh-Panigrahi/GoCMS/models"
	"github.com/joho/godotenv"
)

func SyncDB() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DB.AutoMigrate(&models.User{}, &models.Post{})
	fmt.Println("Database Migrated")
}
