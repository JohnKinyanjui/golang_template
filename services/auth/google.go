package auth_service

import (
	"context"
	"errors"
	"golang_template/internal/db"
	query "golang_template/internal/db/generated"
	"golang_template/pkg/middlewares"
)

func ContinueWithGoogle(token string, params AuthParams) (string, error) {
	if token == "none" {
		return "", errors.New("token is not found")
	}

	dt, err := getUserInfo(token)
	if err != nil {
		return "", nil
	}

	ctx := context.Background()
	if params.Option == SignIn {
		userId, err := db.Query.SignInWithGoogle(ctx, query.SignInWithGoogleParams{
			Email:     dt["email"].(string),
			GoogleUid: dt["google_uid"].(string),
		})
		if err != nil {
			return "", err
		}

		return middlewares.GenerateJWT(userId.String())

	} else if params.Option == Update {
		userId, err := db.Query.UpdateGoogleDetails(ctx, query.UpdateGoogleDetailsParams{
			ID:        params.Id,
			Email:     dt["email"].(string),
			GoogleUid: dt["google_uid"].(string),
		})
		if err != nil {
			return "", err
		}

		return middlewares.GenerateJWT(userId.String())
	}

	return "", errors.New("auth options is not available")
}
