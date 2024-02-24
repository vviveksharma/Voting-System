package services

import (
	"errors"
	"fmt"
	"log"
	"voting-system/comman"
	"voting-system/models"
	"voting-system/repos"

	"github.com/google/uuid"
)

const (
	REGISTER_SUBJECT = "Your employee Id for logging in the Portal"
	BODY_SUBJECT     = "Your employee id "
)

type IAdminService interface {
	AdminRegisterEmployee(*models.AdminEmployeeRegisterRequesteBody) (*models.AdminEmployeeRegisterResponseBody, error)
	AdminLoginEmployee(*models.AdminEmployeeLoginRequesteBody) (*models.AdminEmployeeLoginResponseBody, error)
	AdminOpenVote(*models.AdminOpenVoteRequestBody) (*models.AdminOpenVoteResponseBody, error)
	AdminAddCandidate(requestBody *models.AdminAddCandidateRequestBody) (*models.AdminAddCandidateResponseBody, error)
}

func (ads *AdminService) setupAdminInstance() error {
	var err error
	ads.AdminRepo, err = repos.NewAdminRequest()
	if err != nil {
		return errors.New("error in the repo intilization")
	}
	return nil
}

func (ads *AdminService) setupEmployeeInstance() error {
	var err error
	ads.EmployeeRepo, err = repos.NewEmployeeRequest()
	if err != nil {
		return errors.New("error in the repo intilization")
	}
	return nil
}

func (ads *AdminService) setupCandidateInstance() error {
	var err error
	ads.CandidateRepo, err = repos.NewCandidateRequest()
	if err != nil {
		return errors.New("error in the repo intilization")
	}
	return nil
}

type AdminService struct {
	AdminRepo     repos.Admin
	EmployeeRepo  repos.Employee
	CandidateRepo repos.Candidate
}

func NewAdminService() *AdminService {
	return &AdminService{}
}

// Api used to add the user in the Employee can only be done by the Admin
func (ads *AdminService) AdminRegisterEmployee(requestBody *models.AdminEmployeeRegisterRequesteBody) (*models.AdminEmployeeRegisterResponseBody, error) {
	err := ads.setupAdminInstance()
	if err != nil {
		return nil, errors.New("error while connecting to admin repo: " + err.Error())
	}
	err = ads.setupEmployeeInstance()
	if err != nil {
		return nil, errors.New("error while connecting to employee repo: " + err.Error())
	}
	parsedUUID, err := uuid.Parse(requestBody.AdminId)
	if err != nil {
		fmt.Println("Error parsing UUID:", err)
		return nil, errors.New("error while parsing the string to uuid from the requestBody" + err.Error())
	}
	_, err = ads.AdminRepo.FindBy(&models.DbAdmin{
		Id: parsedUUID,
	})
	if err != nil && err.Error() == "Record not Found" {
		return nil, errors.New("please check the adminId entered")
	} else if err != nil {
		return nil, errors.New("error while finding the adminId in Db: " + err.Error())
	}

	check, err := ads.EmployeeRepo.FindBy(&models.DbEmployee{
		Email: requestBody.Email,
		Role:  requestBody.Role,
	})
	if err != nil {
		return nil, errors.New("error while finding the user if it already exists:" + err.Error())
	}
	if check.Email == requestBody.Email {
		return nil, errors.New("user already exist with the role and email")
	}
	response, err := ads.EmployeeRepo.Create(&models.DbEmployee{
		Role:  requestBody.Role,
		Email: requestBody.Email,
	})
	if err != nil {
		return nil, errors.New("error while adding the new employee in the DataBase: " + err.Error())
	}
	resp, err := comman.SendEmail(requestBody.Email, BODY_SUBJECT+response.Id.String(), REGISTER_SUBJECT)
	if err != nil {
		return nil, err
	}
	log.Print("the response = ", resp)
	return &models.AdminEmployeeRegisterResponseBody{
		Response: "Check email for this",
	}, nil
}

func (ads *AdminService) AdminLoginEmployee(requestBody *models.AdminEmployeeLoginRequesteBody) (*models.AdminEmployeeLoginResponseBody, error) {
	err := ads.setupEmployeeInstance()
	if err != nil {
		return nil, errors.New("error while connecting to employee repo: " + err.Error())
	}
	emplId := uuid.MustParse(requestBody.EmplId)
	_, err = ads.EmployeeRepo.FindBy(&models.DbEmployee{
		Id: emplId,
	})
	if err != nil && err.Error() == "Record Not Found" {
		return nil, errors.New("unable to find this ID ")
	} else if err != nil {
		return nil, errors.New("unable to find the user for login")
	}
	return &models.AdminEmployeeLoginResponseBody{
		Response: "Login Successfull",
	}, nil
}

func (ads *AdminService) AdminOpenVote(requestBody *models.AdminOpenVoteRequestBody) (*models.AdminOpenVoteResponseBody, error) {
	id := uuid.MustParse(requestBody.Id)
	err := ads.setupAdminInstance()
	if err != nil {
		return nil, errors.New("error while connecting to admin repo: " + err.Error())
	}
	err = ads.setupEmployeeInstance()
	if err != nil {
		return nil, errors.New("error while connecting to employee repo: " + err.Error())
	}
	if requestBody.Role == "SUPER-ADMIN" {
		_, err := ads.AdminRepo.IsAdmin(id)
		if err != nil {
			return nil, errors.New("error in finding the admin: " + err.Error())
		}
		return &models.AdminOpenVoteResponseBody{
			Response: "Opened Voting",
		}, nil
	}
	resp, err := ads.EmployeeRepo.FindBy(&models.DbEmployee{
		Id: id,
	})
	if err != nil {
		return nil, errors.New("error while finding the user in the employee DB")
	}
	if resp.Role != "CASTER" {
		return nil, errors.New("can't start the voting inaccessible feature")
	}
	return nil, nil
}

func (ads *AdminService) AdminAddCandidate(requestBody *models.AdminAddCandidateRequestBody) (*models.AdminAddCandidateResponseBody, error) {
	err := ads.setupCandidateInstance()
	if err != nil {
		return nil, errors.New("error while setting up the repo instance for the" + err.Error())
	}
	searlizedData := comman.Searlize(requestBody.CandidateName)
	err = ads.CandidateRepo.Create(&models.DbCandidate{
		Name:  searlizedData,
		Count: 0,
		Image: "",
	})
	if err != nil {
		return nil, errors.New("error in creating the candidate entry " + err.Error())
	}
	return &models.AdminAddCandidateResponseBody{
		Response: "Candidate Created Successfully",
	}, nil
}
