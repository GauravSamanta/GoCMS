package connections

import (
	"fmt"
	"log"

	"github.com/Hrishikesh-Panigrahi/GoCMS/models"
	"github.com/joho/godotenv"
)

func SyncDB() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	err = DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Role{}, &models.UserPostLink{}, &models.Location{}, &models.Comments{}, &models.LinkUserPostComment{})
	if err != nil {
		log.Fatal("Error migrating the database")
	}
	fmt.Println("Database Migrated")
	// seedData()
}

func seedData() {
	// // roles := []models.UserPostImageLink{{UserID: 26, PostID: 1, ImageID:1}, {UserID: 27, PostID: 2, ImageID:2}}
	// user := []models.Comments{{Comment: "This is a comment from user 1"},
	// 	{Comment: "This is a comment from user 2"},
	// 	{Comment: "This is a comment from user 3"},
	// }
	// DB.Save(&user)
	// DB.Save(&user)
}
