package service

import (
	"context"

	"github.com/eeQuillibrium/posts/config"
	"github.com/eeQuillibrium/posts/graph/model"
	"github.com/eeQuillibrium/posts/internal/repository"
	"github.com/eeQuillibrium/posts/pkg/logger"
)

type posts struct {
	log  *logger.Logger
	cfg  *config.Config
	repo repository.Posts
}

func NewPostsService(
	log *logger.Logger,
	cfg *config.Config,
	repo repository.Posts,
) Posts {
	return &posts{
		log:  log,
		cfg:  cfg,
		repo: repo,
	}
}

func (s *posts) CreatePost(
	ctx context.Context,
	post *model.NewPost,
) (int, error) {
	return s.repo.CreatePost(ctx, post)
}

func (s *posts) GetPosts(
	ctx context.Context,
	getPost *model.Pagination,
) ([]*model.Post, error) {
	posts, err := s.repo.GetPosts(ctx, getPost.Offset, getPost.Limit)
	return posts, err
}
func (s *posts) GetPost(
	ctx context.Context,
	postID int,
) (*model.Post, error) {
	return s.repo.GetPost(ctx, postID)
}

func (s *posts) ClosePost(
	ctx context.Context,
	postID int,
) (bool, error) {
	return s.repo.ClosePost(ctx, postID)
}
