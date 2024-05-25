package graph

import (
	"github.com/eeQuillibrium/posts/graph/model"
	"github.com/eeQuillibrium/posts/graph/storage"
	"github.com/eeQuillibrium/posts/internal/service"
	"github.com/eeQuillibrium/posts/pkg/logger"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	service    *service.Service
	log        *logger.Logger
	ps         *storage.PostStorage
	notifyChan chan *model.Notification
}

func NewResolver(
	service *service.Service,
	log *logger.Logger,
	notifyChan chan *model.Notification,
) *Resolver {
	return &Resolver{
		service:    service,
		log:        log,
		ps:         storage.NewPostStorage(),
		notifyChan: notifyChan,
	}
}
