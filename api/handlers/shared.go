package handlers

import (
	"voting-system/models"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetResult(c *fiber.Ctx) error {
	response, err := h.SharedService.SharedGetResult()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Failed(&fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}))
	}
	return c.JSON(models.Success(response))
}
