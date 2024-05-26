package service

import (
	"context"
	"errors"

	"github.com/eeQuillibrium/posts/config"
	"github.com/eeQuillibrium/posts/graph/model"
	"github.com/eeQuillibrium/posts/internal/repository"
	"github.com/eeQuillibrium/posts/pkg/logger"
)

type postsService struct {
	log  *logger.Logger
	cfg  *config.Config
	repo repository.Posts
}

func NewPostsService(
	log *logger.Logger,
	cfg *config.Config,
	repo repository.Posts,
) Posts {
	return &postsService{
		log:  log,
		cfg:  cfg,
		repo: repo,
	}
}

func (s *postsService) CreatePost(
	ctx context.Context,
	post *model.NewPost,
) (int, error) {
	postID, err := s.repo.CreatePost(ctx, post)
	if err != nil {
		return 0, errors.New("commentsService.CreatePost():\n" + err.Error())
	}
	return postID, nil
}

func (s *postsService) GetPosts(
	ctx context.Context,
	getPost *model.Pagination,
) ([]*model.Post, error) {
	posts, err := s.repo.GetPosts(ctx, getPost.Offset, getPost.Limit)
	if err != nil {
		return nil, errors.New("postsService.GetPosts():\n" + err.Error())
	}
	return posts, nil
}
func (s *postsService) GetPost(
	ctx context.Context,
	postID int,
) (*model.Post, error) {
	post, err := s.repo.GetPost(ctx, postID)
	if err != nil {
		return nil, errors.New("postsService.GetPost():\n" + err.Error())
	}
	return post, nil
}

func (s *postsService) ClosePost(
	ctx context.Context,
	postID int,
) (bool, error) {
	isClosed, err := s.repo.ClosePost(ctx, postID)
	if err != nil {
		return false, errors.New("postsService.ClosePost():\n" + err.Error())
	}
	return isClosed, nil
}
