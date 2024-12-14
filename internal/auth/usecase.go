package auth

import auth_models "testovoe/internal/auth/models"

type Usecase interface {
	Signin(*auth_models.SigninParams) (*auth_models.SigninResponse, error)
	ConfirmEmail(mail string, code int) error
}
