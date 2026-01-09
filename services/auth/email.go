package auth_service

import (
	"context"
	"errors"
	"golang_template/internal/db"
	"golang_template/pkg/helpers"
	"golang_template/pkg/middlewares"
	"log"

	"github.com/google/uuid"
)

func ContinueWithEmail(params EmailParams) (*AuthResult, *helpers.ResultError) {
	if err := params.validate(); err != nil {
		return nil, err
	}

	ctx := context.Background()
	var id uuid.UUID
	if err := db.PgConn.QueryRow(
		ctx,
		"select id from users where email = $1::text and password_hash = crypt($2::text, password_hash) limit 1",
		params.Email,
		params.Password,
	).Scan(&id); err != nil {
		return nil, helpers.Error("unable to authenticate by email", errors.New("invalid credentials"))
	}

	_, _ = db.PgConn.Exec(
		ctx,
		"update users set last_login_at = (current_timestamp at time zone 'utc') where id = $1",
		id,
	)

	token, err := middlewares.GenerateJWT(id.String())
	if err != nil {
		log.Println("unable to generate session:", err)
		return nil, helpers.Error("unable to generate session", err)
	}

	return &AuthResult{AccessToken: token}, nil
}
