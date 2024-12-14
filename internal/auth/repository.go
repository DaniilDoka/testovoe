package auth

import (
	auth_models "testovoe/internal/auth/models"
)

type Repository interface {
	NewUser(*auth_models.SigninParams) error
	ConfirmEmail(mail string) error
}
