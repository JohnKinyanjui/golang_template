package auth_service

import (
	"context"
	"errors"
	"golang_template/internal/db"
	query "golang_template/internal/db/generated"
	"golang_template/pkg/helpers"
	"golang_template/pkg/middlewares"
	"log"

	"github.com/jackc/pgx/v5/pgconn"
)

func RegisterWithEmail(params RegisterEmailParams) (*AuthResult, *helpers.ResultError) {
	if err := params.validate(); err != nil {
		return nil, err
	}

	ctx := context.Background()

	queries := query.New(db.PgConn)
	id, err := queries.CreateUserWithEmail(ctx, query.CreateUserWithEmailParams{
		FullName: params.FullName,
		Email:    params.Email,
		Phone:    params.Phone,
		Password: params.Password,
	})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return nil, helpers.Error("user already exists", errors.New("duplicate"))
			}
		}
		return nil, helpers.Error("unable to register", err)
	}

	token, err := middlewares.GenerateJWT(id.String())
	if err != nil {
		log.Println("unable to generate session:", err)
		return nil, helpers.Error("unable to generate session", err)
	}

	return &AuthResult{AccessToken: token}, nil
}
