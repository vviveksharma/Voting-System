package handlers

import (
	"errors"
	"voting-system/models"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) UserRegister(c *fiber.Ctx) error {
	var requestBody *models.UserResgiterRequestBody
	err := c.BodyParser(&requestBody)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(errors.New("error while parsing the request Body"))
	}
	if requestBody.Email == "" || requestBody.Password == "" || requestBody.UserName == "" || requestBody.FName == "" || requestBody.LName == "" {
		return c.Status(fiber.ErrBadRequest.Code).JSON(models.Failed(&fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: errors.New("email, password, Username, lastname and firstName cannot be empty").Error(),
		}))
	}
	response, err := h.UserService.UserRegister(requestBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Failed(&fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}))
	}
	return c.JSON(models.Success(response))
}

func (h *Handler) UserLogin(c *fiber.Ctx) error {
	var requestBody *models.UserLoginRequestBody
	err := c.BodyParser(&requestBody)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(errors.New("error while parsing the request Body"))
	}
	if requestBody.VoterId == "" {
		return c.Status(fiber.ErrBadRequest.Code).JSON(models.Failed(&fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: errors.New("voter_Id cannot be empty").Error(),
		}))
	}
	response, err := h.UserService.UserLogin(requestBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Failed(&fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}))
	}
	return c.JSON(models.Success(response))
}
