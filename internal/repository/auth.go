package repository

import (
	"context"

	"github.com/eeQuillibrium/posts/config"
	"github.com/eeQuillibrium/posts/graph/model"
	"github.com/eeQuillibrium/posts/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type auth struct {
	log *logger.Logger
	cfg *config.Config
	db  *sqlx.DB
}

func NewAuthRepository(
	log *logger.Logger,
	cfg *config.Config,
	db *sqlx.DB,
) *auth {
	return &auth{
		log: log,
		cfg: cfg,
		db:  db,
	}
}

func (r *auth) Register(
	ctx context.Context,
	login string,
	passhash []byte,
	name string,
) (*model.User, error) {
	row := r.db.QueryRowContext(ctx, "INSERT INTO Users (name, login, passhash)"+
		"VALUES ('', $1, $2) RETURNING id", name, login, passhash)

	var userID int
	err := row.Scan(&userID)

	return &model.User{
		ID: userID,
		Login: login,
		Name: name,
	}, err
}
func (r *auth) Login(
	ctx context.Context,
	login string,
	passhash string,
) (*model.User, error) {
	user := model.User{}

	err := r.db.GetContext(ctx, &user, "SELECT * FROM Users WHERE login=$1", login)

	return &user, err
}
