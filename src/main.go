package main

import (
	"log"
	
	"github.com/gofiber/fiber/v2"

	"github.com/daan0101/discordapi/src/config"
	"github.com/daan0101/discordapi/src/routes"
)

func main() {
	config.Load()

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: false,
		StrictRouting: true,
		Immutable:     true,
		BodyLimit:     2 * 1024 * 1024, // 2 MB
	})


	routes.SetupAppRoutes(app)

	log.Fatal(app.Listen(":3000"))
}