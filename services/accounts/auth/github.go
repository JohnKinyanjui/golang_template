package auth_service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

const githubTokenURL = "https://github.com/login/oauth/access_token"
const githubAPIURL = "https://api.github.com"

type GitHubOAuthTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
}

type GithubUser struct {
	Id       int64  `json:"id"`
	Picture  string `json:"avatar_url"`
	UserName string `json:"name"`
	FullName string `json:"login"`
	Email    string `json:"email"`
}

func exchangeCodeForToken(code string) (*GitHubOAuthTokenResponse, error) {
	clientID := os.Getenv("GITHUB_CLIENT_ID")
	clientSecret := os.Getenv("GITHUB_CLIENT_SECRET")

	data := fmt.Sprintf("client_id=%s&client_secret=%s&code=%s", clientID, clientSecret, code)
	req, err := http.NewRequest("POST", githubTokenURL, strings.NewReader(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tokenResp GitHubOAuthTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&tokenResp)
	if err != nil {
		return nil, err
	}

	return &tokenResp, nil
}

func fetchGitHubUser(accessToken string) (*GithubUser, error) {
	// Reuse a single HTTP client
	client := &http.Client{}

	// Fetch user details
	req, err := http.NewRequest("GET", githubAPIURL+"/user", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get user info: %s", resp.Status)
	}

	var user GithubUser
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	// Fetch user emails
	email, err := getUserEmails(accessToken)
	if err != nil {
		return nil, err
	}
	user.Email = email

	return &user, nil
}

func getUserEmails(accessToken string) (string, error) {
	url := "https://api.github.com/user/emails"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get user emails: %s", resp.Status)
	}

	var emails []struct {
		Email    string `json:"email"`
		Primary  bool   `json:"primary"`
		Verified bool   `json:"verified"`
	}
	err = json.NewDecoder(resp.Body).Decode(&emails)
	if err != nil {
		return "", err
	}

	// Find the primary email or return the first one if no primary is found
	for _, email := range emails {
		if email.Primary {
			return email.Email, nil
		}
	}

	if len(emails) > 0 {
		return emails[0].Email, nil // Return first email if no primary
	}

	return "", nil // No emails found
}
