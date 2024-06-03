package service_mocks

import (
	"context"

	"github.com/eeQuillibrium/posts/graph/model"
)

type MockAuthRepository struct{}

func (r *MockAuthRepository) CreateUser(
	ctx context.Context,
	user *model.NewUser,
) (int, error) {
	return 1, nil
}

type MockPostsRepository struct{}

func (s *MockPostsRepository) CreatePost(
	ctx context.Context,
	post *model.NewPost,
) (int, error) {
	return 1, nil
}
func (s *MockPostsRepository) ClosePost(
	ctx context.Context,
	postID int,
) (bool, error) {
	return true, nil
}
func (s *MockPostsRepository) GetPosts(
	ctx context.Context,
	getPost *model.Pagination,
) ([]*model.Post, error) {
	return nil, nil
}
func (s *MockPostsRepository) GetPost(
	ctx context.Context,
	postID int,
) (*model.Post, error) {
	return nil, nil
}

type CommentsNode struct {
	comments []*model.Comment
	child *CommentsNode
}
type MockCommentsRepository struct{
	root *CommentsNode
}


func (s *MockCommentsRepository) CreateComment(
	ctx context.Context,
	comment *model.NewComment,
) (int, error) {
	return 1, nil
}
func (s *MockCommentsRepository) GetComments(
	ctx context.Context,
	postID int,
) ([]*model.Comment, error) {
	return nil, nil
}
func (s *MockCommentsRepository) GetByParentComment(
	ctx context.Context,
	commentID int,
) ([]*model.Comment, error) {
	return nil, nil
}