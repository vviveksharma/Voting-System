package models

type UserResgiterRequestBody struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	FName    string `json:"first_name"`
	LName    string `json:"last_name"`
	Password string `json:"password"`
}

type UserLoginRequestBody struct {
	VoterId string `json:"voter_id"`
}

type UserValidateTokeRequestBody struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type UserCastVoteRequestBody struct {
	Email         string `json:"email"`
	CandidateName string `json:"candidate_name"`
}

type AdminEmployeeRegisterRequesteBody struct {
	AdminId string `json:"admin_Id"`
	Email   string `json:"email"`
	Role    string `json:"role"`
}

type AdminEmployeeLoginRequesteBody struct {
	EmplId string `json:"empl_id"`
}

type AdminOpenVoteRequestBody struct {
	Id   string `json:"id"`
	Role string `json:"role"`
	Time int64  `json:"time"`
}
