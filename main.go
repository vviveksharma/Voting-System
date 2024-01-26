package main

import (
	"log"
	"voting-system/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(recover.New())

	db.InitDB()
	// Swagger UI
	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.SendString("service is up and healthy")
	})

	// Your API routes go here...
	log.Fatal(app.Listen(":8000"))
}
