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
	app.Get("/artists/:id/concerts", handlers.GetArtistConcertsHandler)
	app.Post("/artists", handlers.CreateArtistHandler)
	app.Put("/artists/:id", handlers.UpdateArtistHandler)
	app.Patch("/artists/:id", handlers.PatchArtistHandler)
	app.Delete("/artists/:id", handlers.DeleteArtistHandler)

	app.Get("/concerts", handlers.GetConcertsHandler)
	app.Get("/concerts/:id", handlers.GetConcertByIDHandler)
	app.Post("/concerts", handlers.CreateConcertHandler)
	app.Patch("/concerts/:id", handlers.PatchConcertHandler)
	app.Delete("/concerts/:id", handlers.DeleteConcertHandler)

	log.Fatal(app.Listen(":3001"))
}
