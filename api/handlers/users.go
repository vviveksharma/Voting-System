package handlers

import (
	"errors"
	"log"
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

func (h *Handler) UserValidateToken(c *fiber.Ctx) error {
	var requestBody *models.UserValidateTokeRequestBody
	err := c.BodyParser(&requestBody)
	if err != nil {
		log.Print("Error while parsing the body: ", err.Error())
		return c.Status(fiber.StatusBadGateway).JSON(errors.New("error while parsing the request Body"))
	}
	if requestBody.Token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Failed(&fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: errors.New("token can't be empty").Error(),
		}))
	}
	response, err := h.UserService.UserValidateToken(requestBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Failed(&fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}))
	}
	return c.JSON(models.Success(response))
}

func (h *Handler) UserCastVote(c *fiber.Ctx) error {
	var requestBody *models.UserCastVoteRequestBody
	err := c.BodyParser(requestBody)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(errors.New("error while parsing the request Body"))
	}
	if requestBody.CandidateName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Failed(&fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: errors.New("candidateName cannot be empty").Error(),
		}))
	}
	response, err := h.UserService.UserCastVote(requestBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Failed(&fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}))
	}
	return c.JSON(models.Success(response))
}
