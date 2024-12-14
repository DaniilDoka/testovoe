package auth_models

type SigninParams struct {
	Mail     string `json:"email"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type SigninResponse struct {
	Status string `json:"status"`
}
