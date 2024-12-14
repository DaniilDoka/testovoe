package auth_repository

import (
	"testovoe/internal/auth"
	auth_models "testovoe/internal/auth/models"
	"testovoe/pkg/pg"
)

type postgres struct {
	db *pg.Pg
}

func NewRepository(pgConn *pg.Pg) auth.Repository {
	return &postgres{
		db: pgConn,
	}
}

func (p *postgres) ConfirmEmail(mail string) error {
	err := p.db.Exec(`update "user" set status = 1 where email=$1`, mail)
	return err
}

func (p *postgres) NewUser(params *auth_models.SigninParams) error {
	err := p.db.Exec(`insert into "user"(email, nickname, password) values($1,$2,$3)`, params.Mail, params.Nickname, params.Password)

	return err
}
