package main

import (
	"log"
	"time"
	routes "voting-system/api"
	"voting-system/api/handlers"
	"voting-system/api/services"
	"voting-system/db"
	"voting-system/repos"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(recover.New())
	done := make(chan struct{})
	db.InitDB()
	var Log *log.Logger
	var test bool
	// Swagger UI
	handler := handlers.NewHandler(Log).UserServiceInstance(services.NewUserService())
	routes.Routes(app, handler)
	go func() {
		// Simulate some work in the Goroutine
		time.Sleep(5 * time.Second)
		test = check()
		log.Print(test)
		// Signal that the Goroutine is done
		close(done)
	}()
	// Your API routes go here...
	log.Fatal(app.Listen(":8000"))
}

func check() bool {
	response := false
	adminrepos, err := repos.NewAdminRequest()
	if err != nil {
		log.Print("error in creation of repo")
		return false
	}
	count, err := adminrepos.FindAll()
	if err != nil {
		log.Print("error in findALL", err)
		return false
	}
	if len(count) != 0 {
		response = true
	}
	return response
}
