package repos

import (
	"voting-system/models"

	"github.com/google/uuid"
)

func NewAdminRequest() (Admin, error) {
	return &AdminImpl{}, nil
}

type Admin interface {
	Create(*models.DbAdmin) (*models.DbAdmin, error)
	FindAll() ([]*models.DbAdmin, error)
	FindBy(conditions *models.DbAdmin) (*models.DbAdmin, error)
	IsAdmin(id uuid.UUID) (bool, error)
}
