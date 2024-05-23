package service

import (
	"context"

	"github.com/eeQuillibrium/posts/config"
	"github.com/eeQuillibrium/posts/graph/model"
	"github.com/eeQuillibrium/posts/internal/repository"
	"github.com/eeQuillibrium/posts/pkg/logger"
)

type Auth interface {
	Login(
		ctx context.Context,
		user *model.User,
	) (*model.User, error)
	Register(
		ctx context.Context,
		user *model.NewUser,
	) (*model.User, error)
}

type Topics interface {
	CreateTopic(
		ctx context.Context,
		topic *model.NewTopic,
	) (*model.Topic, error)
	GetTopics(
		ctx context.Context,
		getTopic *model.PaginationTopics,
	) ([]*model.Topic, error)
}

type Service struct {
	Topics
	Auth
}

func NewService(
	log *logger.Logger,
	cfg *config.Config,
	repo *repository.Repository,
) *Service {
	return &Service{
		Topics: NewTopicsService(log, cfg, repo.Topics),
		Auth:   NewAuthService(log, cfg, repo.Auth),
	}
}
