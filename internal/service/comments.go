package service

import (
	"context"

	"github.com/eeQuillibrium/posts/config"
	"github.com/eeQuillibrium/posts/graph/model"
	"github.com/eeQuillibrium/posts/internal/repository"
	"github.com/eeQuillibrium/posts/pkg/logger"
)

type comments struct {
	log  *logger.Logger
	cfg  *config.Config
	repo repository.Comments
}

func NewCommentsService(
	log *logger.Logger,
	cfg *config.Config,
	repo repository.Comments,
) Comments {
	return &comments{
		log:  log,
		cfg:  cfg,
		repo: repo,
	}
}

func (s *comments) CreateComment(
	ctx context.Context,
	comment *model.NewComment,
) (int, error) {
	return s.repo.CreateComment(ctx, comment)
}

func (s *comments) GetComments(
	ctx context.Context,
	postID int,
) ([]*model.Comment, error) {
	return s.repo.GetComments(ctx, postID)
}

func (s *comments) GetByComment(
	ctx context.Context,
	commentID int,
) ([]*model.Comment, error) {
	return s.repo.GetByComment(ctx, commentID)
}