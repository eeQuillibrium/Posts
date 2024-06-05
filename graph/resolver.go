package graph

import (
	"os"

	"github.com/eeQuillibrium/posts/graph/model"
	"github.com/eeQuillibrium/posts/graph/storage"
	"github.com/eeQuillibrium/posts/internal/service"
	"github.com/eeQuillibrium/posts/pkg/logger"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	service     *service.Service
	log         *logger.Logger
	notifyChan  chan *model.Notification
	storageMode string // POSTGRES/INMEMORY
	st          *storage.Storage
}

func NewResolver(
	service *service.Service,
	log *logger.Logger,
	notifyChan chan *model.Notification,
	st *storage.Storage,
) *Resolver {
	return &Resolver{
		service:     service,
		log:         log,
		notifyChan:  notifyChan,
		storageMode: os.Getenv("STORAGE_MODE"),
		st:          st,
	}
}
