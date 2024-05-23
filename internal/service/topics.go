package service

import (
	"context"

	"github.com/eeQuillibrium/posts/config"
	"github.com/eeQuillibrium/posts/graph/model"
	"github.com/eeQuillibrium/posts/internal/repository"
	"github.com/eeQuillibrium/posts/pkg/logger"
)

type topics struct {
	log  *logger.Logger
	cfg  *config.Config
	repo repository.Topics
}

func NewTopicsService(
	log *logger.Logger,
	cfg *config.Config,
	repo repository.Topics,
) Topics {
	return &topics{
		log:  log,
		cfg:  cfg,
		repo: repo,
	}
}

func (s *topics) CreateTopic(
	ctx context.Context,
	topic *model.NewTopic,
) (*model.Topic, error) {
	topicID, err := s.repo.CreateTopic(ctx, topic)
	if err != nil {
		return nil, err
	}
	s.log.Infof("created topic id: %d", topicID)
	return &model.Topic{
		ID:       topicID,
		UserID:   topic.UserID,
		Text:     topic.Text,
		Header:   topic.Header,
		IssuedAt: topic.IssuedAt,
	}, nil
}

func (s *topics) GetTopics(
	ctx context.Context,
	getTopic *model.PaginationTopics,
) ([]*model.Topic, error) {
	topics, err := s.repo.GetTopics(ctx, getTopic.Offset, getTopic.Limit)
	s.log.Info(topics)
	return topics, err
}
