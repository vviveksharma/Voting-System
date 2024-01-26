package repos

import (
	"voting-system/db"
	"voting-system/models"
)

type UserImpl struct {
}

func (um *UserImpl) Create(value *models.DbUser) error {
	dbConn, err := db.InitDB()
	if err != nil {
		return err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return transaction.Error
	}
	defer transaction.Rollback()
	state := transaction.Create(value)
	if state.Error != nil {
		return state.Error
	}
	transaction.Commit()
	return nil
}

func (um *UserImpl) Find(conditions *models.DbUser) (*models.DbUser, error) {
	dbConn, err := db.InitDB()
	if err != nil {
		return nil, err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return nil, transaction.Error
	}
	defer transaction.Rollback()
	var value models.DbUser
	state := transaction.Find(&value, conditions)
	if state.Error != nil {
		return nil, state.Error
	}
	return &value, nil
}

func (um *UserImpl) Update(value models.DbUser) error {
	dbConn, err := db.InitDB()
	if err != nil {
		return err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return transaction.Error
	}
	defer transaction.Rollback()
	state := transaction.Save(&value)
	if state.Error != nil {
		return state.Error
	}
	return nil
}
