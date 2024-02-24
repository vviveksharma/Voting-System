package repos

import (
	"voting-system/db"
	"voting-system/models"
)

type EmployeeImpl struct {
}

func (empl *EmployeeImpl) Create(value *models.DbEmployee) (*models.DbEmployee, error) {
	dbConn, err := db.InitDB()
	if err != nil {
		return nil, err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return nil, transaction.Error
	}
	defer transaction.Rollback()
	state := transaction.Create(&value)
	if state.Error != nil {
		return nil, state.Error
	}
	transaction.Commit()
	return value, nil
}

func (empl *EmployeeImpl) FindBy(conditions *models.DbEmployee) (*models.DbEmployee, error) {
	dbConn, err := db.InitDB()
	if err != nil {
		return nil, err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return nil, transaction.Error
	}
	defer transaction.Rollback()
	var value *models.DbEmployee
	state := transaction.Find(&value, conditions)
	if state.Error != nil {
		return nil, state.Error
	}
	transaction.Commit()
	return value, nil
}
