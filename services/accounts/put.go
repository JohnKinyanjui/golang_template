package accounts_service

import (
	"context"
	"errors"
	"golang_template/internal/db"
	query "golang_template/internal/db/generated"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

var ErrUsernameAlreadyExists = errors.New("username already exists")

type UpdateAccountParams struct {
	Username   *string `json:"username"`
	About      *string `json:"about"`
	Location   *string `json:"location"`
	LookingFor *string `json:"looking_for"`
}

func UpdateAccount(id string, params UpdateAccountParams) (interface{}, error) {
	ctx := context.Background()
	userID := uuid.MustParse(id)

	queries := query.New(db.PgConn)

	arg := query.UpdateMyAccountParams{
		ID:         userID,
		Username:   pgtype.Text{Valid: false},
		About:      pgtype.Text{Valid: false},
		Location:   pgtype.Text{Valid: false},
		LookingFor: pgtype.Text{Valid: false},
	}

	if params.Username != nil {
		arg.Username = pgtype.Text{String: *params.Username, Valid: true}
	}
	if params.About != nil {
		arg.About = pgtype.Text{String: *params.About, Valid: true}
	}
	if params.Location != nil {
		arg.Location = pgtype.Text{String: *params.Location, Valid: true}
	}
	if params.LookingFor != nil {
		arg.LookingFor = pgtype.Text{String: *params.LookingFor, Valid: true}
	}

	row, err := queries.UpdateMyAccount(ctx, arg)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return nil, ErrUsernameAlreadyExists
			}
		}
		return nil, err
	}

	res := MyAccount{
		Picture:    row.Picture,
		Username:   row.Username,
		Email:      row.Email,
		About:      row.About,
		Location:   row.Location,
		LookingFor: row.LookingFor,
	}

	return res, nil
}
