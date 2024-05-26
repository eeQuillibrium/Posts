package service

import (
	"context"
	"errors"

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

func (s *commentsService) CreateComment(
	ctx context.Context,
	comment *model.NewComment,
) (int, error) {
	commentID, err := s.repo.CreateComment(ctx, comment)
	if err != nil {
		return 0, errors.New("commentsService.CreateComment():\n" + err.Error())
	}
	return commentID, nil
}

func (s *commentsService) GetComments(
	ctx context.Context,
	postID int,
) ([]*model.Comment, error) {
	comments, err := s.repo.GetComments(ctx, postID)
	if err != nil {
		return nil, errors.New("commentsService.GetComments():\n"+ err.Error())
	}
	return comments, nil
}

func (s *commentsService) GetByComment(
	ctx context.Context,
	commentID int,
) ([]*model.Comment, error) {
	comments, err := s.repo.GetByComment(ctx, commentID)
	if err != nil {
		return nil, errors.New("commentsService.GetByComment():\n"+ err.Error())
	}
	return comments, nil
}
