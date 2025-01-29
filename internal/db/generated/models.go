// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package query

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID             uuid.UUID        `json:"id"`
	Picture        pgtype.Text      `json:"picture"`
	FullName       string           `json:"full_name"`
	CreatedAt      pgtype.Timestamp `json:"created_at"`
	Email          pgtype.Text      `json:"email"`
	GoogleUid      pgtype.Text      `json:"google_uid"`
	GithubUid      pgtype.Text      `json:"github_uid"`
	PhoneNumber    pgtype.Text      `json:"phone_number"`
	Password       pgtype.Text      `json:"password"`
	GithubUsername pgtype.Text      `json:"github_username"`
	GithubToken    pgtype.Text      `json:"github_token"`
}
