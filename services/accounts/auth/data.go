package auth_service

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

type GoogleUserDetails struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Locale        string `json:"locale"`
}
