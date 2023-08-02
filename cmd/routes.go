package main

import (
	"github.com/Junkes887/go-db/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/fact", handlers.ListFacts)

	app.Post("/fact", handlers.CreateFact)
}
