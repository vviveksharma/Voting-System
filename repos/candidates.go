package repos

import "voting-system/models"

func NewCandidateRequest() (Candidate, error) {
	return &CandidateImpl{}, nil
}

type Candidate interface {
	Create(value *models.DbCandidate) error
	Find(conditions *models.DbCandidate) (*models.DbCandidate, error)
	Update(value *models.DbCandidate) error
	GetResult() (*models.DbCandidate, error)
}
