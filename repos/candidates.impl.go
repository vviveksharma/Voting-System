package repos

import (
	"voting-system/db"
	"voting-system/models"
)

type CandidateImpl struct {
}

func (cd *CandidateImpl) Create(value *models.DbCandidate) error {
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

func (cd *CandidateImpl) Find(conditions *models.DbCandidate) (*models.DbCandidate, error) {
	dbConn, err := db.InitDB()
	if err != nil {
		return nil, err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return nil, transaction.Error
	}
	defer transaction.Rollback()
	var value models.DbCandidate
	state := transaction.Find(&value, conditions)
	if state.Error != nil {
		return nil, state.Error
	}
	return &value, nil
}

func (cd *CandidateImpl) Update(value *models.DbCandidate) error {
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
	transaction.Commit()
	return nil
}

func (cd *CandidateImpl) GetResult() (*models.DbCandidate, error) {
	dbConn, err := db.InitDB()
	if err != nil {
		return nil, err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return nil, transaction.Error
	}
	defer transaction.Rollback()
	var response models.DbCandidate
	transaction.Order("count desc").First(&response)
	return &response, nil
}
