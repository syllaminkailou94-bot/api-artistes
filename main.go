package main

import (
	"log"

	"api-artistes/database"
	"api-artistes/handlers"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Get("/health", handlers.HealthHandler)

	app.Get("/artists", handlers.GetArtistsHandler)
	app.Get("/artists/:id", handlers.GetArtistByIDHandler)
	app.Post("/artists", handlers.CreateArtistHandler)
	app.Put("/artists/:id", handlers.UpdateArtistHandler)
	app.Delete("/artists/:id", handlers.DeleteArtistHandler)
	app.Use("/", static.New("./public"))
	app.Patch("/artists/:id", handlers.PatchArtistHandler)

	log.Fatal(app.Listen(":3001"))
}
