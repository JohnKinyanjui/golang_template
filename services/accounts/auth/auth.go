package auth_service

import (
	"context"
	"errors"
	"fmt"
	"golang_template/internal/db"
	"golang_template/pkg/middlewares"
	"log"
)

func ContinueWithGithub(code string) (map[string]string, error) {
	log.Println("is code found: ", code)

	if code == "none" {
		return nil, errors.New("code is not found")
	}
	var id string

	t, err := exchangeCodeForToken(code)
	if err != nil {
		return nil, err
	}

	pr, err := fetchGitHubUser(t.AccessToken)
	if err != nil {
		return nil, err
	}

	githubUid := fmt.Sprintf("%d", pr.Id)

	ctx := context.Background()
	err = db.PgConn.QueryRow(ctx, checkIfGithubAccountExists, githubUid).Scan(&id)
	if err != nil {
		err = db.PgConn.QueryRow(context.Background(), signUpGithubQuery,
			pr.Picture, pr.FullName, pr.Email, githubUid, t.RefreshToken).
			Scan(&id)

		if err != nil {
			return nil, err
		}

		token, err := middlewares.GenerateJWT(id)
		if err != nil {
			return nil, err
		}

		return map[string]string{
			"access_token": token,
		}, nil
	}

	token, err := middlewares.GenerateJWT(id)
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"access_token": token,
	}, nil
}
