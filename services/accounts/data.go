package accounts_service

import (
	"errors"
	"golang_template/pkg/helpers"
)

type MyAccount struct {
	Picture  string `json:"picture"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type CreateUserAgentParams struct {
	Picture     string `json:"picture"`
	Name        string `json:"name"`
	Description string `json:"description"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

func (params *CreateUserAgentParams) validate() *helpers.ResultError {
	if params.Name == "" || params.PhoneNumber == "" || params.Email == "" {
		return helpers.Error("name, phone number, or email is empty", errors.New("missing required fields"))
	}

	return nil
}
