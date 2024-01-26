package main

import (
	"log"
	routes "voting-system/api"
	"voting-system/api/handlers"
	"voting-system/api/services"
	"voting-system/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(recover.New())

	db.InitDB()
	var log *log.Logger
	// Swagger UI
	handler := handlers.NewHandler(log).UserServiceInstance(services.NewUserService())
	routes.Routes(app, handler)
	// Your API routes go here...
	log.Fatal(app.Listen(":8000"))
}
