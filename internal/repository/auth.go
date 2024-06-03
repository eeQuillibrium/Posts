package repository

import (
	"context"
	"errors"

	"github.com/eeQuillibrium/posts/config"
	"github.com/eeQuillibrium/posts/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type authRepository struct {
	log *logger.Logger
	cfg *config.Config
	db  *sqlx.DB
}

func NewAuthRepository(
	log *logger.Logger,
	cfg *config.Config,
	db *sqlx.DB,
) Auth {
	return &authRepository{
		log: log,
		cfg: cfg,
		db:  db,
	}
}

func (ar *authRepository) CreateUser(
	ctx context.Context,
	login string,
	passhash []byte,
	name string,
) (int, error) {
	row := ar.db.QueryRowContext(ctx, "INSERT INTO Users (name, login, passhash)"+
		"VALUES ($1, $2, $3) RETURNING id", name, login, passhash)

	var userID int
	err := row.Scan(&userID)
	if err != nil {
		return 0, errors.New("authRepository.Register(): " + err.Error())
	}
	return userID, nil
}
