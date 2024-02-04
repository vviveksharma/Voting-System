package services

import (
	"errors"
	"log"
	"voting-system/comman"
	"voting-system/models"
	"voting-system/repos"

	"github.com/google/uuid"
)

type IUserService interface {
	UserRegister(*models.UserResgiterRequestBody) (*models.UserRegisterResponseBody, error)
	UserLogin(*models.UserLoginRequestBody) (*models.UserLoginResponseBody, error)
	UserValidateToken(*models.UserValidateTokeRequestBody) (*models.UserValidateTokenResponseBody, error)
	UserCastVote(requestBody *models.UserCastVoteRequestBody) (*models.UserCastVoteResponseBody, error)
}

func (us *UserService) setupUserInstance() error {
	var err error
	us.UserRepo, err = repos.NewUserRequest()
	if err != nil {
		return errors.New("error in the repo intilization")
	}
	return nil
}

func (us *UserService) setupCandidateInstance() error {
	var err error
	us.CandidateRepo, err = repos.NewCandidateRequest()
	if err != nil {
		return errors.New("error in the repo intilization")
	}
	return nil
}

type UserService struct {
	UserRepo      repos.User
	CandidateRepo repos.Candidate
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) UserRegister(requestBody *models.UserResgiterRequestBody) (*models.UserRegisterResponseBody, error) {
	err := us.setupUserInstance()
	if err != nil {
		return nil, errors.New("error while intializig the user repos")
	}
	token := uuid.New()
	hashedPassword, err := comman.GenerateHash(requestBody.Password)
	if err != nil {
		return nil, err
	}
	err = us.UserRepo.Create(&models.DbUser{
		Username:  requestBody.UserName,
		SecretKey: hashedPassword,
		Fname:     requestBody.FName,
		Lname:     requestBody.LName,
		Email:     requestBody.Email,
		Token:     token,
	})
	if err != nil {
		return nil, errors.New("unable to create the user from the repos layer")
	}
	return &models.UserRegisterResponseBody{
		Response: token.String(),
	}, nil
}

func (us *UserService) UserLogin(requestBody *models.UserLoginRequestBody) (*models.UserLoginResponseBody, error) {
	err := us.setupUserInstance()
	if err != nil {
		return nil, errors.New("error while intializig the user repos")
	}
	parsedUUID, err := uuid.Parse(requestBody.VoterId)
	if err != nil {
		return nil, errors.New("error while parsing the requestBody uuid to UUID format")
	}

	userDetails, err := us.UserRepo.Find(&models.DbUser{
		VoterID: parsedUUID,
	})

	if err != nil {
		return nil, errors.New("error while finding the user")
	}

	if !userDetails.IsValidate {
		return nil, errors.New("please validate the token sent on your email")
	}
	return &models.UserLoginResponseBody{
		Response: "User Logged in Successfully",
	}, nil
}

func (us *UserService) UserValidateToken(requestBody *models.UserValidateTokeRequestBody) (*models.UserValidateTokenResponseBody, error) {
	err := us.setupUserInstance()
	if err != nil {
		return nil, errors.New("error while setting up the repo instance")
	}
	userDetails, err := us.UserRepo.Find(&models.DbUser{
		Email: requestBody.Email,
	})
	if err != nil {
		return nil, errors.New("unable to find the user")
	}
	if userDetails.Token.String() != requestBody.Token {
		return nil, errors.New("invalid token entered please check email")
	}
	voterId := uuid.New()
	userDetails.IsValidate = true
	userDetails.VoterID = voterId
	err = us.UserRepo.Update(userDetails)
	if err != nil {
		log.Print("the error = ", err)
		return nil, errors.New("unable to update the entry in the DatBase")
	}
	return &models.UserValidateTokenResponseBody{
		Response: voterId.String(),
	}, nil
}

func (us *UserService) UserCastVote(requestBody *models.UserCastVoteRequestBody) (*models.UserCastVoteResponseBody, error) {
	err := us.setupUserInstance()
	if err != nil {
		return nil, errors.New("error while intilizing the user repo instance")
	}
	err = us.setupCandidateInstance()
	if err != nil {
		return nil, errors.New("error while intilizing the candidate repo instance")
	}
	userDetails, err := us.UserRepo.Find(&models.DbUser{
		Email: requestBody.Email,
	})
	if err != nil {
		return nil, errors.New("error while finding the user with this Email")
	}
	if !userDetails.IsValidate {
		return nil, errors.New("please validate the user ")
	}
	if !userDetails.IsLoggedIn {
		return nil, errors.New("please loggin To vote")
	}
	candidateCheck, err := us.CandidateRepo.Find(&models.DbCandidate{
		Name: requestBody.CandidateName,
	})
	if err != nil {
		return nil, errors.New("")
	}
	candidateCheck.Count += 1
	err = us.CandidateRepo.Update(candidateCheck)
	if err != nil {
		return nil, errors.New("error while updating the cadidate record")
	}
	return nil, nil
}
