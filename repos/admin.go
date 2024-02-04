package repos

import "voting-system/models"

func NewAdminRequest() (Admin, error) {
	return &AdminImpl{}, nil
}

type Admin interface {
	Create(*models.DbAdmin) (*models.DbAdmin, error)
	FindAll() ([]*models.DbAdmin, error)
}
