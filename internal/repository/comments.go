package repository

import (
	"context"
	"errors"
	"time"

	"github.com/eeQuillibrium/posts/config"
	loaders "github.com/eeQuillibrium/posts/graph/loader"
	"github.com/eeQuillibrium/posts/graph/model"
	"github.com/eeQuillibrium/posts/pkg/logger"

	"github.com/jmoiron/sqlx"
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
		return 0, errors.New("commentsRepository.CreateComment(): " + err.Error())
	}
	if isClosed {
		return 0, errors.New("commentsRepository.CreateComment():" +  "post is closed by author")
	}

	row := cr.db.QueryRowxContext(ctx, "INSERT INTO Comments (user_id, post_id, parent_id, text, level, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		comment.UserID, comment.PostID, comment.ParentID, comment.Text, comment.Level, time.Now())

	var postID int
	if err := row.Scan(&postID); err != nil {
		return 0, errors.New("commentsRepository.CreateComment(): " + err.Error())
	}

	return postID, nil
}

func (r *commentsRepository) GetComments(
	ctx context.Context,
	postID int,
) ([]*model.Comment, error) {
	var commentIDs []int

	if err := r.db.SelectContext(ctx, &commentIDs, "SELECT id FROM Comments WHERE post_id = $1",
		postID); err != nil {
		return nil, errors.New("commentsRepository.GetComments(): " + err.Error())
	}

	return loaders.GetComments(ctx, commentIDs)
}

func (r *commentsRepository) GetByParentComment(
	ctx context.Context,
	commentID int,
) ([]*model.Comment, error) {
	var comments []*model.Comment

	if err := r.db.SelectContext(ctx, &comments, "SELECT * FROM Comments WHERE parent_id = $1",
		commentID); err != nil {
		return nil, errors.New("commentsRepository.GetByComment(): " + err.Error())
	}

	return comments, nil
}
