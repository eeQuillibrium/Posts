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

func (cr *commentsRepository) GetPostComments(
	ctx context.Context,
	postID int,
	limit int,
) ([]*model.Comment, error) {
	var comments []*model.Comment

	if err := cr.db.SelectContext(ctx, &comments, "SELECT * FROM Comments WHERE post_id = $1 ORDER BY id DESC LIMIT $2 OFFSET $3",
		postID, limit, 0); err != nil {
		return nil, fmt.Errorf("commentsRepository.GetPostComments(): %w", err)
	}

	return comments, nil
}

func (cr *commentsRepository) GetByParentComment(
	ctx context.Context,
	commentID int,
) ([]*model.Comment, error) {
	var comments []*model.Comment

	if err := cr.db.SelectContext(ctx, &comments, "SELECT * FROM Comments WHERE parent_id = $1",
		commentID); err != nil {
		return nil, fmt.Errorf("commentsRepository.GetByParentComment(): %w", err)
	}
	return comments, nil
}
func (cr *commentsRepository) PaginationComment(
	ctx context.Context,
	postID int,
	offset int,
	limit int,
) ([]*model.Comment, error) {
	comments := []*model.Comment{}

	q := `
	SELECT *
	FROM Comments
	WHERE post_id = $1
	ORDER BY id
	LIMIT $2 OFFSET $3
	`

	rows, err := cr.db.QueryxContext(ctx, q, postID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("commentsRepository.PaginationComments(): %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var comment model.Comment
		if err := rows.StructScan(&comment); err != nil {
			return nil, fmt.Errorf("commentsRepository.PaginationComments(): %w", err)
		}
		comments = append(comments, &comment)
	}

	return comments, nil
}
