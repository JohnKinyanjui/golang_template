package auth_service

import (
	"context"
	"errors"
	"golang_template/internal/db"
	"golang_template/pkg/middlewares"

	"github.com/google/uuid"
)

func ContinueWithGoogle(token string, params AuthParams) (string, error) {
	if token == "none" {
		return "", errors.New("token is not found")
	}

	dt, err := getUserInfo(token)
	if err != nil {
		return "", nil
	}

	email, ok := dt["email"].(string)
	if !ok || email == "" {
		return "", errors.New("email is not found")
	}

	googleUid, ok := dt["sub"].(string)
	if !ok || googleUid == "" {
		return "", errors.New("google uid is not found")
	}

	ctx := context.Background()
	if params.Option == SignIn {
		var userId uuid.UUID
		if err := db.PgConn.QueryRow(
			ctx,
			"select id from users where email = $1::text and google_uid = $2::text limit 1",
			email,
			googleUid,
		).Scan(&userId); err != nil {
			return "", err
		}

		return middlewares.GenerateJWT(userId.String())

	} else if params.Option == Update {
		var userId uuid.UUID
		if err := db.PgConn.QueryRow(
			ctx,
			"update users set email = $1::text, google_uid = $2::text where id = $3 returning id",
			email,
			googleUid,
			params.Id,
		).Scan(&userId); err != nil {
			return "", err
		}

		return middlewares.GenerateJWT(userId.String())
	}

	return "", errors.New("auth options is not available")
}
