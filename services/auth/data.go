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

func (params EmailParams) validate() *helpers.ResultError {
	if params.Email == "" || params.Password == "" {
		return helpers.Error("email or password is empty", errors.New("email or password is empty"))
	}
	return nil
}
