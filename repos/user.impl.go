package repos

import (
	"voting-system/db"
	"voting-system/models"
)


type UserImpl struct{
}

func(um *UserImpl) Create(value *models.DbUser) error {
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