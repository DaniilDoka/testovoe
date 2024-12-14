package auth_usercase

import (
	"fmt"
	"math/rand"
	"regexp"
	"sync"
	"testovoe/config"
	"testovoe/internal/auth"
	auth_models "testovoe/internal/auth/models"
	"testovoe/pkg/logger"

	"golang.org/x/crypto/bcrypt"
)

type usecase struct {
	repo      auth.Repository
	mu        sync.Mutex
	mailCodes map[string]int
	cfg       *config.Config
	logger    *logger.Logger
}

func NewUsecase(repo auth.Repository, cfg *config.Config, logger *logger.Logger) auth.Usecase {
	return &usecase{
		repo:      repo,
		mailCodes: make(map[string]int),
		cfg:       cfg,
		logger:    logger,
	}
}

func (u *usecase) ConfirmEmail(mail string, code int) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	if v, ok := u.mailCodes[mail]; v != code || ok == false {
		return fmt.Errorf("Invalid code")
	}

	err := u.repo.ConfirmEmail(mail)

	return err
}

func (u *usecase) sendMailConfirm(mail string) error {
	return nil
}

func (u *usecase) Signin(params *auth_models.SigninParams) (*auth_models.SigninResponse, error) {
	if regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(params.Mail) == false {
		return nil, fmt.Errorf("Invalid email")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	params.Password = string(hashedPassword)
	if err := u.repo.NewUser(params); err != nil {
		return nil, err
	}

	u.mu.Lock()
	code := rand.Intn(900000) + 100000
	u.logger.Infof("Email %s code %d", params.Mail, code)
	u.mailCodes[params.Mail] = code
	u.mu.Unlock()

	if err := u.sendMailConfirm(params.Mail); err != nil {
		return nil, err
	}

	return &auth_models.SigninResponse{
		Status: "Need email confirm",
	}, nil
}
