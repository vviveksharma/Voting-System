package repos

import (
	"voting-system/db"
	"voting-system/models"

	"github.com/google/uuid"
)

type AdminImpl struct {
}

func (ad *AdminImpl) Create(value *models.DbAdmin) (*models.DbAdmin, error) {
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

func (ad *AdminImpl) FindAll() ([]*models.DbAdmin, error) {
	dbConn, err := db.InitDB()
	if err != nil {
		return nil, err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return nil, transaction.Error
	}
	var response []*models.DbAdmin
	defer transaction.Rollback()
	result := transaction.Find(&response)
	if result.Error != nil {
		return nil, result.Error
	}
	return response, nil
}

func (ad *AdminImpl) FindBy(conditions *models.DbAdmin) (*models.DbAdmin, error) {
	dbConn, err := db.InitDB()
	if err != nil {
		return nil, err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return nil, transaction.Error
	}
	var response *models.DbAdmin
	defer transaction.Rollback()
	result := transaction.Find(&response, conditions)
	if result.Error != nil {
		return nil, result.Error
	}
	return response, nil
}

func (ad *AdminImpl) IsAdmin(id uuid.UUID) (bool, error) {
	dbConn, err := db.InitDB()
	if err != nil {
		return false, err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return false, transaction.Error
	}
	var response *models.DbAdmin
	defer transaction.Rollback()
	result := transaction.Find(&response, models.DbAdmin{
		Id: id,
	})
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
