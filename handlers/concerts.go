package handlers

import (
	"errors"
	"log"
	"strconv"

	"api-artistes/database"
	"api-artistes/models"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func GetConcertsHandler(c fiber.Ctx) error {
	var concerts []models.Concert

	if err := database.DB.Find(&concerts).Error; err != nil {
		log.Println("Erreur GORM :", err)
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	return c.JSON(concerts)
}

func CreateConcertHandler(c fiber.Ctx) error {
	var newConcert models.Concert

	if err := c.Bind().Body(&newConcert); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	if newConcert.ArtistID == 0 || newConcert.City == "" || newConcert.Date == "" {
		return c.Status(400).JSON(fiber.Map{"error": "artist_id, city and date are required"})
	}

	var artist models.Artist
	if err := database.DB.First(&artist, newConcert.ArtistID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{"error": "Artist not found"})
		}
		log.Println("Erreur GORM :", err)
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	newConcert.ID = 0

	if err := database.DB.Create(&newConcert).Error; err != nil {
		log.Println("Erreur GORM :", err)
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	return c.Status(201).JSON(newConcert)
}
func GetArtistConcertsHandler(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid artist ID"})
	}

	var artist models.Artist
	if err := database.DB.First(&artist, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{"error": "Artist not found"})
		}
		log.Println("Erreur GORM :", err)
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	var concerts []models.Concert
	if err := database.DB.Where("artist_id = ?", id).Find(&concerts).Error; err != nil {
		log.Println("Erreur GORM :", err)
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	return c.JSON(concerts)
}
func GetConcertByIDHandler(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid concert ID"})
	}

	var concert models.Concert
	if err := database.DB.First(&concert, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{"error": "Concert not found"})
		}
		log.Println("Erreur GORM :", err)
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	return c.JSON(concert)
}

func DeleteConcertHandler(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid concert ID"})
	}

	result := database.DB.Delete(&models.Concert{}, id)
	if result.Error != nil {
		log.Println("Erreur GORM :", result.Error)
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Concert not found"})
	}

	return c.SendStatus(204)
}
func PatchConcertHandler(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid concert ID"})
	}

	var patch models.ConcertPatch
	if err := c.Bind().Body(&patch); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	if patch.ArtistID == nil && patch.City == nil && patch.Venue == nil && patch.Date == nil {
		return c.Status(400).JSON(fiber.Map{"error": "No fields to update"})
	}

	if patch.City != nil && *patch.City == "" {
		return c.Status(400).JSON(fiber.Map{"error": "City cannot be empty"})
	}
	if patch.Date != nil && *patch.Date == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Date cannot be empty"})
	}

	var concert models.Concert
	if err := database.DB.First(&concert, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{"error": "Concert not found"})
		}
		log.Println("Erreur GORM :", err)
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	if patch.ArtistID != nil {
		var artist models.Artist
		if err := database.DB.First(&artist, *patch.ArtistID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.Status(404).JSON(fiber.Map{"error": "Artist not found"})
			}
			log.Println("Erreur GORM :", err)
			return c.Status(500).JSON(fiber.Map{"error": "Database error"})
		}
		concert.ArtistID = *patch.ArtistID
	}
	if patch.City != nil {
		concert.City = *patch.City
	}
	if patch.Venue != nil {
		concert.Venue = *patch.Venue
	}
	if patch.Date != nil {
		concert.Date = *patch.Date
	}

	if err := database.DB.Save(&concert).Error; err != nil {
		log.Println("Erreur GORM :", err)
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	return c.JSON(concert)
}
