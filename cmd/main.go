package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	const songsDir = "songs"

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {

	})

	app.Listen(":8080")
}
