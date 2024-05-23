package repository

import (
	"context"

	"github.com/eeQuillibrium/posts/config"
	"github.com/eeQuillibrium/posts/graph/model"
	"github.com/eeQuillibrium/posts/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type Topics interface {
	CreateTopic(
		ctx context.Context,
		topic *model.NewTopic,
	) (int, error)
	GetTopics(
		ctx context.Context,
		offset int,
		limit int,
	) ([]*model.Topic, error)
}

type Auth interface {
	Login(
		ctx context.Context,
		login string,
		passhash string,
	) (*model.User, error)
	Register(
		ctx context.Context,
		login string,
		passhash []byte,
		name string,
	) (*model.User, error)
}

type Repository struct {
	Topics
	Auth
}

func NewRepository(
	log *logger.Logger,
	cfg *config.Config,
	db *sqlx.DB,
) *Repository {
	return &Repository{
		Topics: NewTopicsRepository(log, cfg, db),
		Auth:   NewAuthRepository(log, cfg, db),
	}
}
