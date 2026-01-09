package auth_service

import (
	"errors"
	"golang_template/pkg/helpers"
)

type AuthResult struct {
	AccessToken string `json:"access_token"`
}

type GoogleUserDetails struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Locale        string `json:"locale"`
}

type EmailParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterEmailParams struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Password   string `json:"password"`
	Gender     string `json:"gender"`
	About      string `json:"about"`
	Location   string `json:"location"`
	LookingFor string `json:"looking_for"`
}

func (params EmailParams) validate() *helpers.ResultError {
	if params.Email == "" || params.Password == "" {
		return helpers.Error("email or password is empty", errors.New("email or password is empty"))
	}
	return nil
}

func (params RegisterEmailParams) validate() *helpers.ResultError {
	if params.Username == "" || params.Email == "" || params.Password == "" || params.Gender == "" {
		return helpers.Error("missing required fields", errors.New("missing required fields"))
	}
	if params.Gender != "male" && params.Gender != "female" {
		return helpers.Error("invalid gender", errors.New("invalid gender"))
	}
	return nil
}
