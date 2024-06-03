package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/eeQuillibrium/posts/config"
	"github.com/eeQuillibrium/posts/graph/model"
	"github.com/eeQuillibrium/posts/pkg/logger"

	"github.com/jmoiron/sqlx"
)

var (
	errPostClosed = errors.New("post is closed by author")
)

type commentsRepository struct {
	log *logger.Logger
	cfg *config.Config
	db  *sqlx.DB
}

func NewCommentsRepository(
	log *logger.Logger,
	cfg *config.Config,
	db *sqlx.DB,
) Comments {
	return &commentsRepository{
		log: log,
		cfg: cfg,
		db:  db,
	}
}

func (cr *commentsRepository) CreateComment(
	ctx context.Context,
	comment *model.NewComment,
) (int, error) {
	var isClosed bool

	if err := cr.db.GetContext(ctx, &isClosed, "SELECT is_closed FROM Posts WHERE id = $1",
		comment.PostID); err != nil {
		return 0, fmt.Errorf("commentsRepository.CreateComment(): %w", err)
	}
	if isClosed {
		return 0, fmt.Errorf("commentsRepository.CreateComment(): %w", errPostClosed)
	}

	row := cr.db.QueryRowxContext(ctx, "INSERT INTO Comments (user_id, post_id, parent_id, text, level, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		comment.UserID, comment.PostID, comment.ParentID, comment.Text, comment.Level, time.Now())

	var postID int
	if err := row.Scan(&postID); err != nil {
		return 0, fmt.Errorf("commentsRepository.CreateComment(): %w", err)
	}

	return postID, nil
}

func (r *commentsRepository) GetComments(
	ctx context.Context,
	postID int,
) ([]*model.Comment, error) {
	var comments []*model.Comment

	if err := r.db.SelectContext(ctx, &comments, "SELECT * FROM Comments WHERE post_id = $1",
		postID); err != nil {
		return nil, fmt.Errorf("commentsRepository.GetComments(): %w", err)
	}

	return comments, nil
}

func (r *commentsRepository) GetByParentComment(
	ctx context.Context,
	commentID int,
) ([]*model.Comment, error) {
	var comments []*model.Comment

	if err := r.db.SelectContext(ctx, &comments, "SELECT * FROM Comments WHERE parent_id = $1",
		commentID); err != nil {
		return nil, fmt.Errorf("commentsRepository.GetByComment(): %w", err)
	}

	return comments, nil
}
