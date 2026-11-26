package database

import (
	"log"

	"api-artistes/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	DB, err = gorm.Open(sqlite.Open("artists.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Impossible d'ouvrir la base :", err)
	}

	if err := DB.AutoMigrate(&models.Artist{}); err != nil {
		log.Fatal("Erreur de migration :", err)
	}

	var count int64
	DB.Model(&models.Artist{}).Count(&count)
	if count == 0 {
		DB.Create(&[]models.Artist{
			{Name: "Daft Punk", Genre: "Électro"},
			{Name: "Stromae", Genre: "Pop"},
			{Name: "Nekfeu", Genre: "Rap"},
		})
	}
}
