package services

import (
	"errors"
	"voting-system/comman"
	"voting-system/models"
	"voting-system/repos"
)

type ISharedService interface {
	SharedGetResult() (*models.SharedGetResultResponseBody, error)
}

type SharedService struct {
	CandidateRepo repos.Candidate
}

func NewSharedService() *SharedService {
	return &SharedService{}
}

func (ss *SharedService) setupCandidateInstance() error {
	var err error
	ss.CandidateRepo, err = repos.NewCandidateRequest()
	if err != nil {
		return errors.New("error in the repo intilization")
	}
	return nil
}

func (ss *SharedService) SharedGetResult() (*models.SharedGetResultResponseBody, error) {
	err := ss.setupCandidateInstance()
	if err != nil {
		return nil, errors.New("error in setting up the dataBase instance: " + err.Error())
	}
	resp, err := ss.CandidateRepo.GetResult()
	if err != nil {
		return nil, errors.New("error while fetching response: " + err.Error())
	}
	name := comman.DeSerilizeData(resp.Name)
	return &models.SharedGetResultResponseBody{
		Name:  name,
		Count: resp.Count,
	}, nil
}
