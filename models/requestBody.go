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
