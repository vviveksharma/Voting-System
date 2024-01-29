package services

import (
	"errors"
	"voting-system/models"
	"voting-system/repos"

	"github.com/google/uuid"
)

type IUserService interface {
	UserRegister(*models.UserResgiterRequestBody) (*models.UserRegisterResponseBody, error)
	UserLogin(*models.UserLoginRequestBody) (*models.UserLoginResponseBody, error)
}

func (us *UserService) setupUserInstance() error {
	var err error
	us.UserRepo, err = repos.NewUserRequest()
	if err != nil {
		return errors.New("error in the repo intilization")
	}
	return nil
}

type UserService struct {
	UserRepo repos.User
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) UserRegister(requestBody *models.UserResgiterRequestBody) (*models.UserRegisterResponseBody, error) {
	err := us.setupUserInstance()
	if err != nil {
		return nil, errors.New("error while intializig the user repos")
	}
	userDetails, err := us.UserRepo.Find(&models.DbUser{
		Email: requestBody.Email,
	})
	if err != nil {
		return nil, errors.New("error while finding the user")
	}
	if userDetails.Email == requestBody.Email {
		return nil, errors.New("user already exist please login")
	}
	voterId := uuid.New()
	err = us.UserRepo.Create(&models.DbUser{
		Username:  requestBody.UserName,
		SecretKey: requestBody.Password,
		Fname:     requestBody.FName,
		Lname:     requestBody.LName,
		Email:     requestBody.Email,
		VoterID:   voterId,
	})
	if err != nil {
		return nil, errors.New("unable to create the user from the repos layer")
	}
	return &models.UserRegisterResponseBody{
		Response: voterId.String(),
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
