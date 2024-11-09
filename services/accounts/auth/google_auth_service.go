package auth_service

import (
	"errors"

	"github.com/surrealdb/surrealdb.go"
)

func ContinueWithGoogle(token string) (string, error) {
	if token == "none" {
		return "", errors.New("token is not found")
	}

	// data, err := getUserInfo(token)
	// if err != nil {
	// 	return "", err
	// }

	// exists := checkIfUseExists(data["email"].(string))
	// if exists {
	// 	token, err := db.Conn.SignIn(&surrealdb.Auth{
	// 		Namespace: "root",
	// 		Database:  "root",
	// 		Scope:     "root",
	// 	})
	// 	if err != nil {
	// 		return "", err
	// 	}

	// 	return token, nil
	// }

	// token, err = db.Conn.SignUp(&surrealdb.Auth{
	// 	Namespace: "root",
	// 	Database:  "root",
	// 	Scope:     "root",
	// })
	// if err != nil {
	// 	return "", err
	// }

	return token, nil
}

func ValidateToken(conn *surrealdb.DB) (interface{}, error) {

	return nil, nil
}
