package repository

import (
	"context"
	"errors"

	"time"

	"github.com/eeQuillibrium/posts/config"
	"github.com/eeQuillibrium/posts/graph/model"
	"github.com/eeQuillibrium/posts/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type postsRepository struct {
	log *logger.Logger
	cfg *config.Config
	db  *sqlx.DB
}

func NewPostsRepository(
	log *logger.Logger,
	cfg *config.Config,
	db *sqlx.DB,
) Posts {
	return &postsRepository{
		log: log,
		cfg: cfg,
		db:  db,
	}
}

func (r *postsRepository) CreatePost(
	ctx context.Context,
	Post *model.NewPost,
) (int, error) {

	row := r.db.QueryRowxContext(ctx, "INSERT INTO Posts (user_id, text, header, created_at, is_closed) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		Post.UserID, Post.Text, Post.Header, time.Now(), false)

	var PostID int
	if err := row.Scan(&PostID); err != nil {
		return 0, errors.New("postsRepository.CreatePost(): " + err.Error())
	}

	return PostID, nil
}

func (r *postsRepository) GetPosts(
	ctx context.Context,
	offset int,
	limit int,
) ([]*model.Post, error) {
	posts := []*model.Post{}

	if err := r.db.SelectContext(ctx, &posts, "SELECT * FROM Posts ORDER BY id desc LIMIT $1 OFFSET $2",
		limit, offset); err != nil {
		return nil, errors.New("postsRepository.GetPosts(): " + err.Error())
	}

	return posts, nil
}
func (r *postsRepository) ClosePost(
	ctx context.Context,
	postID int,
) (bool, error) {
	q := `
	UPDATE Posts
	SET is_closed = $2
	WHERE id = $1;
	`

	_, err := r.db.ExecContext(ctx, q, postID, true)
	if err != nil {
		return false, errors.New("postsRepository.ClosePost(): " + err.Error())
	}

	return true, nil
}
func (r *postsRepository) GetPost(
	ctx context.Context,
	postID int,
) (*model.Post, error) {
	post := model.Post{}

	if err := r.db.GetContext(ctx, &post, "SELECT * FROM Posts WHERE id = $1",
		postID); err != nil {
		return nil, errors.New("postsRepository.GetPost(): " + err.Error())
	}

	return &post, nil
}
