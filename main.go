package main

import (
	"log"

	"api-artistes/database"
	"api-artistes/handlers"

	"github.com/gofiber/fiber/v3"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Get("/health", handlers.HealthHandler)

	app.Get("/artists", handlers.GetArtistsHandler)
	app.Get("/artists/:id", handlers.GetArtistByIDHandler)
	app.Post("/artists", handlers.CreateArtistHandler)
	app.Put("/artists/:id", handlers.UpdateArtistHandler)
	app.Patch("/artists/:id", handlers.PatchArtistHandler)
	app.Delete("/artists/:id", handlers.DeleteArtistHandler)

	app.Get("/concerts", handlers.GetConcertsHandler)
	app.Post("/concerts", handlers.CreateConcertHandler)
	app.Get("/artists/:id/concerts", handlers.GetArtistConcertsHandler)
	app.Get("/concerts/:id", handlers.GetConcertByIDHandler)
	app.Delete("/concerts/:id", handlers.DeleteConcertHandler)
	app.Patch("/concerts/:id", handlers.PatchConcertHandler)

	log.Fatal(app.Listen(":3001"))
}
