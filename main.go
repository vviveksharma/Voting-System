package main

import (
	"log"
	"time"
	routes "voting-system/api"
	"voting-system/api/handlers"
	"voting-system/api/services"
	"voting-system/db"
	"voting-system/models"
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
	// Swagger UI
	handler := handlers.NewHandler(Log).UserServiceInstance(services.NewUserService()).AdminSericeInstance(services.NewAdminService())
	routes.Routes(app, handler)
	go func() {
		// Simulate some work in the Goroutine
		time.Sleep(5 * time.Second)
		adminId, err := check()
		if err != nil {
			log.Print(err)
		}
		log.Print("Admin Id = ", adminId)
		// Signal that the Goroutine is done
		close(done)
	}()
	// Your API routes go here...
	log.Fatal(app.Listen(":8000"))
}

func check() (string, error) {
	adminrepos, err := repos.NewAdminRequest()
	if err != nil {
		log.Print("error in creation of repo")
		return "", err
	}
	count, err := adminrepos.FindAll()
	if err != nil {
		log.Print("error in findALL", err)
		return "", err
	}
	if len(count) != 0 {
		return "Please check your associated email already a Super Admin given", nil
	} else {
		response, err := adminrepos.Create(&models.DbAdmin{
			Role:         "SUPER-ADMIN",
			IsSuperAdmin: true,
		})
		if err != nil {
			log.Print("error in creating the admin creads for latest realease")
			return "", err
		}
		return response.Id.String(), nil
	}
}
