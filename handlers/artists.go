package handlers

import (
	"errors"
	"strconv"

	"api-artistes/database"
	"api-artistes/models"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func GetArtistsHandler(c fiber.Ctx) error {
	var artists []models.Artist

	if err := database.DB.Find(&artists).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	return c.JSON(artists)
}

func GetArtistByIDHandler(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid artist ID"})
	}

	var artist models.Artist
	if err := database.DB.First(&artist, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{"error": "Artist not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	return c.JSON(artist)
}

func CreateArtistHandler(c fiber.Ctx) error {
	var newArtist models.Artist

	if err := c.Bind().Body(&newArtist); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	if newArtist.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Name is required"})
	}

	newArtist.ID = 0

	if err := database.DB.Create(&newArtist).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	return c.Status(201).JSON(newArtist)
}

func UpdateArtistHandler(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid artist ID"})
	}

	var updated models.Artist
	if err := c.Bind().Body(&updated); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	if updated.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Name is required"})
	}

	var artist models.Artist
	if err := database.DB.First(&artist, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{"error": "Artist not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	artist.Name = updated.Name
	artist.Genre = updated.Genre

	if err := database.DB.Save(&artist).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	return c.JSON(artist)
}

func DeleteArtistHandler(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid artist ID"})
	}

	result := database.DB.Delete(&models.Artist{}, id)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Artist not found"})
	}

	return c.SendStatus(204)
}
func PatchArtistHandler(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid artist ID"})
	}

	var patch models.ArtistPatch
	if err := c.Bind().Body(&patch); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	if patch.Name == nil && patch.Genre == nil {
		return c.Status(400).JSON(fiber.Map{"error": "No fields to update"})
	}

	if patch.Name != nil && *patch.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Name cannot be empty"})
	}

	var artist models.Artist
	if err := database.DB.First(&artist, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{"error": "Artist not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	if patch.Name != nil {
		artist.Name = *patch.Name
	}
	if patch.Genre != nil {
		artist.Genre = *patch.Genre
	}

	if err := database.DB.Save(&artist).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	return c.JSON(artist)
}
