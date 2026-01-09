package accounts_service

type MyAccount struct {
	Picture    string `json:"picture"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	About      string `json:"about"`
	Location   string `json:"location"`
	LookingFor string `json:"looking_for"`
}
