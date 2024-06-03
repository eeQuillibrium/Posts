package service

import (
	"context"
	"fmt"

	"github.com/eeQuillibrium/posts/config"
	"github.com/eeQuillibrium/posts/graph/model"
	"github.com/eeQuillibrium/posts/internal/repository"
	"github.com/eeQuillibrium/posts/pkg/logger"
)

type commentsService struct {
	log  *logger.Logger
	cfg  *config.Config
	repo repository.Comments
}

func NewCommentsService(
	log *logger.Logger,
	cfg *config.Config,
	repo repository.Comments,
) Comments {
	return &commentsService{
		log:  log,
		cfg:  cfg,
		repo: repo,
	}
}

func (cs *commentsService) CreateComment(
	ctx context.Context,
	comment *model.NewComment,
) (int, error) {
	commentID, err := cs.repo.CreateComment(ctx, comment)
	if err != nil {
		return 0, fmt.Errorf("commentsService.CreateComment():\n%w", err)
	}
	return commentID, nil
}

func (cs *commentsService) GetComments(
	ctx context.Context,
	postID int,
) ([]*model.Comment, error) {
	comments, err := cs.repo.GetComments(ctx, postID)
	if err != nil {
		return nil, fmt.Errorf("commentsService.GetComments():\n%w", err)
	}
	return comments, nil
}

func (cs *commentsService) GetByParentComment(
	ctx context.Context,
	parentID int,
) ([]*model.Comment, error) {
	comments, err := cs.repo.GetByParentComment(ctx, parentID)
	if err != nil {
		return nil, fmt.Errorf("commentsService.GetByComment():\n%w", err)
	}
	return comments, nil
}
