package storage

import (
	"context"

	"github.com/eeQuillibrium/posts/graph/model"
)

type Comments interface {
	CreateComment(
		ctx context.Context,
		comment *model.NewComment,
	) (int, error)
	GetPostComments(
		ctx context.Context,
		postID int,
	) ([]*model.Comment, error)
	GetChildLevel(
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
type Users interface {
	CreateUser(
		ctx context.Context,
		user *model.NewUser,
	) (int, error)
}

type Storage struct {
	Comments
	Posts
	Users
	Mediator
}

func NewStorage() *Storage {
	comments := NewCommentsStorage()
	posts := NewPostsStorage()
	users := NewUsersStorage()
	mediator := NewStorageMediator(comments, posts, users)

	return &Storage{
		Comments: comments,
		Posts:    posts,
		Users:    users,
		Mediator: mediator,
	}
}
