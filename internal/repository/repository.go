package repository

import (
	"context"

	"github.com/eeQuillibrium/posts/config"
	"github.com/eeQuillibrium/posts/graph/model"
	"github.com/eeQuillibrium/posts/pkg/logger"
	"github.com/jmoiron/sqlx"
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
		parentID int,
	) ([]*model.Comment, error)
	PaginationComment(
		ctx context.Context,
		postID int,
		offset int,
		limit int,
	) ([]*model.Comment, error)
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
		offset int,
		limit int,
	) ([]*model.Post, error)
	GetPost(
		ctx context.Context,
		postID int,
	) (*model.Post, error)
}

type Auth interface {
	CreateUser(
		ctx context.Context,
		login string,
		passhash []byte,
		name string,
	) (int, error)
}

type Repository struct {
	Posts
	Auth
	Comments
}

func NewRepository(
	log *logger.Logger,
	cfg *config.Config,
	db *sqlx.DB,
) *Repository {
	return &Repository{
		Posts:    NewPostsRepository(log, cfg, db),
		Auth:     NewAuthRepository(log, cfg, db),
		Comments: NewCommentsRepository(log, cfg, db),
	}
}
