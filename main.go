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
	go func() {
		time.Sleep(5 * time.Second)
		_, err := db.InitDB()
		if err != nil {
			log.Println("error in starting the DataBase: ", err)
		}
		adminId, err := check()
		if err != nil {
			log.Print(err)
		}
		log.Print("Admin Id = ", adminId)
		close(done)
	}()
	var Log *log.Logger
	handler := handlers.NewHandler(Log).UserServiceInstance(services.NewUserService()).AdminSericeInstance(services.NewAdminService()).SharedServiceInstance(services.NewSharedService())
	routes.Routes(app, handler)

	log.Println(app.Listen(":8000"))
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
