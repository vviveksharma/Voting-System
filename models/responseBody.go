package models

type UserRegisterResponseBody struct {
	Response string `json:"response"`
}

type UserLoginResponseBody struct {
	Response string `json:"response"`
}

type UserValidateTokenResponseBody struct {
	VoterId string
}

type UserCastVoteResponseBody struct {
	Response string
}
