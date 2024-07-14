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
	err = DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Role{}, &models.UserPostLink{})
	if err != nil {
		log.Fatal("Error migrating the database")
	}
	fmt.Println("Database Migrated")
	// seedData()
}

// func seedData() {
// 	roles := []models.UserPostImageLink{{UserID: 26, PostID: 1, ImageID:1}, {UserID: 27, PostID: 2, ImageID:2}}
// 	// user := []models.User{{Name: "blah", Email: "blah", Password: "blah", RoleID: 1}}
	
// 	DB.Save(&roles)
// 	// DB.Save(&user)
// }
