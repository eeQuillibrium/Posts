package repository

import (
	"context"

	"github.com/eeQuillibrium/posts/config"
	"github.com/eeQuillibrium/posts/graph/model"
	"github.com/eeQuillibrium/posts/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type topics struct {
	log *logger.Logger
	cfg *config.Config
	db  *sqlx.DB
}

func NewTopicsRepository(
	log *logger.Logger,
	cfg *config.Config,
	db *sqlx.DB,
) Topics {
	return &topics{
		log: log,
		cfg: cfg,
		db:  db,
	}
}

func (r *topics) CreateTopic(
	ctx context.Context,
	topic *model.NewTopic,
) (int, error) {
	row := r.db.QueryRowxContext(ctx, "INSERT INTO Topic (user_id, text, header) VALUES ($1, $2, $3) RETURNING id",
		topic.UserID, topic.Text, topic.Header)

	var topicID int
	if err := row.Scan(&topicID); err != nil {
		return 0, err
	}

	return topicID, nil
}
func (r *topics) GetTopics(
	ctx context.Context,
	offset int,
	limit int,
) ([]*model.Topic, error) {
	topics := []*model.Topic{}

	if err := r.db.SelectContext(ctx, &topics, "SELECT * FROM Topic ORDER BY id desc LIMIT $1 OFFSET $2",
		limit, offset); err != nil {
		return nil, err
	}

	return topics, nil
}
