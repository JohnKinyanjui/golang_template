package auth_service

import (
	"context"
	"golang_template/internal/db"
	query "golang_template/internal/db/generated"
	"golang_template/pkg/helpers"
	"golang_template/pkg/middlewares"
	"log"
)

func ContinueWithEmail(params EmailParams) (*AuthResult, *helpers.ResultError) {
	if err := params.validate(); err != nil {
		return nil, err
	}

	ctx := context.Background()
	id, err := db.Query.AuthenticateByEmail(ctx, query.AuthenticateByEmailParams{
		Email:    params.Email,
		Password: params.Password,
	})
	if err != nil {
		return nil, helpers.Error("unable to authenticate by email", err)
	}

	token, err := middlewares.GenerateJWT(id.String())
	if err != nil {
		log.Println("unable to generate session:", err)
		return nil, helpers.Error("unable to generate session", err)
	}

	return &AuthResult{AccessToken: token}, nil
}
