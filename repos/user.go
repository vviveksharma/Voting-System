package repos

import (
	"voting-system/models"
)

func NewUserRequest() (User, error) {
	return &UserImpl{}, nil
}

type User interface {
	Create(*models.DbUser) error
}
