package models

type UserRegisterResponseBody struct {
	Response string `json:"response"`
}

type UserLoginResponseBody struct {
	Response string `json:"response"`
}

type UserValidateTokenResponseBody struct {
	Response string `json:"response"`
}

type UserCastVoteResponseBody struct {
	Response string
}
