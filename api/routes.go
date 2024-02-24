package routes

import (
	"voting-system/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, h *handlers.Handler) {
	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.SendString("Health Check is working fine")
	})

	app.Post("/users/register", h.UserRegister)
	app.Post("/users/login", h.UserLogin)
	app.Post("/users/validate-token", h.UserValidateToken)
	app.Post("/users/cast-vote", h.UserCastVote)
	app.Post("/admin/register", h.AdminRegisterEmployee)
	app.Post("/admin/login", h.AdminLoginEmployee)
	app.Post("/admin/add-candidate", h.AdminAddCandidate)
	app.Post("/admin/open-vote", h.AdminOpenVote)
	app.Get("/get-results", h.GetResult)
}
