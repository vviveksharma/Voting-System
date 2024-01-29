package handlers

import (
	"log"
	"voting-system/api/services"
)

type Handler struct {
	Logger      *log.Logger
	UserService services.IUserService
}

func NewHandler(logger *log.Logger) *Handler {
	return &Handler{Logger: logger}
}

func (h *Handler) UserServiceInstance(us services.IUserService) *Handler {
	h.UserService = us
	return h
}

