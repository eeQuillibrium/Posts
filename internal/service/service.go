package service

import (
	"context"

	"github.com/eeQuillibrium/posts/config"
	"github.com/eeQuillibrium/posts/graph/model"
	"github.com/eeQuillibrium/posts/internal/repository"
	"github.com/eeQuillibrium/posts/pkg/logger"
)

type Comments interface {
	CreateComment(
		ctx context.Context,
		comment *model.NewComment,
	) (int, error)
	GetPostComments(
		ctx context.Context,
		postID int,
		limit int,
	) ([]*model.Comment, error)
	GetByParentComment(
		ctx context.Context,
		commentID int,
	) ([]*model.Comment, error)
	PaginationComment(
		ctx context.Context,
		postID int,
		offset int,
		limit int,
	) ([]*model.Comment, error)
}
type Auth interface {
	CreateUser(
		ctx context.Context,
		user *model.NewUser,
	) (int, error)
}

type Posts interface {
	CreatePost(
		ctx context.Context,
		post *model.NewPost,
	) (int, error)
	ClosePost(
		ctx context.Context,
		postID int,
	) (bool, error)
	GetPosts(
		ctx context.Context,
		getPost *model.Pagination,
	) ([]*model.Post, error)
	GetPost(
		ctx context.Context,
		postID int,
	) (*model.Post, error)
}

type Service struct {
	Posts
	Auth
	Comments
}

func NewService(
	log *logger.Logger,
	cfg *config.Config,
	repo *repository.Repository,
) *Service {
	return &Service{
		Posts:    NewPostsService(log, cfg, repo.Posts),
		Auth:     NewAuthService(log, cfg, repo.Auth),
		Comments: NewCommentsService(log, cfg, repo.Comments),
	}
}
