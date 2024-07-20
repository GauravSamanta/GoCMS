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
	err = DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Role{}, &models.UserPostLink{}, &models.Location{})
	if err != nil {
		log.Fatal("Error migrating the database")
	}
	fmt.Println("Database Migrated")
	// seedData()
}

// func seedData() {
// 	// roles := []models.UserPostImageLink{{UserID: 26, PostID: 1, ImageID:1}, {UserID: 27, PostID: 2, ImageID:2}}
// 	user := []models.Location{{CityName: "Bhubaneswar", StateName: "Odisha", CountryName: "India", ZipCode: "751001"},
// 		{CityName: "Cuttack", StateName: "Odisha", CountryName: "India", ZipCode: "753001"},
// 		{CityName: "Puri", StateName: "Odisha", CountryName: "India", ZipCode: "752001"},
// 		{CityName: "Berhampur", StateName: "Odisha", CountryName: "India", ZipCode: "760001"},
// 		{CityName: "Sambalpur", StateName: "Odisha", CountryName: "India", ZipCode: "768001"},
// 		{CityName: "Mumbai", StateName: "Maharashtra", CountryName: "India", ZipCode: "400001"},
// 		{CityName: "Pune", StateName: "Maharashtra", CountryName: "India", ZipCode: "411001"},
// 		{CityName: "Thane", StateName: "Maharashtra", CountryName: "India", ZipCode: "400601"}}

// 	DB.Save(&user)
// 	// DB.Save(&user)
// }
