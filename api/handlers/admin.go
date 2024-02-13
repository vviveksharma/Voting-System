package handlers

import (
	"errors"
	"voting-system/models"

	"github.com/gofiber/fiber/v2"
)

// Api used By Super Admin to add employees in the DB
func (h *Handler) AdminRegisterEmployee(c *fiber.Ctx) error {
	var requestBody models.AdminEmployeeRegisterRequesteBody
	err := c.BodyParser(&requestBody)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(errors.New("error while parsing the request Body"))
	}
	if requestBody.AdminId == "" || requestBody.Email == "" || requestBody.Role == "" {
		return c.Status(fiber.ErrBadRequest.Code).JSON(models.Failed(&fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: errors.New("email, AdminId and role for that Employee can't be empty").Error(),
		}))
	}
	response, err := h.AdminService.AdminRegisterEmployee(&requestBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Failed(&fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}))
	}
	return c.JSON(models.Success(response))
}

func (h *Handler) AdminLoginEmployee(c *fiber.Ctx) error {
	var requestBody models.AdminEmployeeLoginRequesteBody
	err := c.BodyParser(&requestBody)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(errors.New("error while parsing the request Body"))
	}
	if requestBody.EmplId == "" {
		return c.Status(fiber.ErrBadRequest.Code).JSON(models.Failed(&fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: errors.New("employee id for that Employee can't be empty").Error(),
		}))
	}
	response, err := h.AdminService.AdminLoginEmployee(&requestBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Failed(&fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}))
	}
	return c.JSON(models.Success(response))
}
