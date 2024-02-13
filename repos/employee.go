package repos

import "voting-system/models"

func NewEmployeeRequest() (Employee, error) {
	return &EmployeeImpl{}, nil
}

type Employee interface {
	Create(*models.DbEmployee) (*models.DbEmployee, error)
	FindBy(*models.DbEmployee) (*models.DbEmployee, error)
}