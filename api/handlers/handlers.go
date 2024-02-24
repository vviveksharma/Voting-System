package handlers

import (
	"log"
	"voting-system/api/services"
)

type Handler struct {
	Logger        *log.Logger
	UserService   services.IUserService
	AdminService  services.IAdminService
	SharedService services.ISharedService
}

func NewHandler(logger *log.Logger) *Handler {
	return &Handler{Logger: logger}
}

func (h *Handler) UserServiceInstance(us services.IUserService) *Handler {
	h.UserService = us
	return h
}

func (h *Handler) AdminSericeInstance(ad services.IAdminService) *Handler {
	h.AdminService = ad
	return h
}

func (h *Handler) SharedServiceInstance(ss services.ISharedService) *Handler {
	h.SharedService = ss
	return h
}
