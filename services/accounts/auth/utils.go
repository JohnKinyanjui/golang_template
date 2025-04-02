package auth_service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type AuthOptions string

const (
	SignIn AuthOptions = "signin"
	Update AuthOptions = "update"
)

type AuthParams struct {
	Id     uuid.UUID
	Option AuthOptions
}

func getUserInfo(accessToken string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://www.googleapis.com/oauth2/v3/userinfo?access_token=%s", accessToken)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch user info")
	}

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	return userInfo, nil
}
