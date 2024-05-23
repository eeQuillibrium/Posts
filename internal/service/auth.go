package service

import (
	"context"
	"errors"

	"github.com/eeQuillibrium/posts/config"
	"github.com/eeQuillibrium/posts/graph/model"
	"github.com/eeQuillibrium/posts/internal/repository"
	"github.com/eeQuillibrium/posts/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

type auth struct {
	log  *logger.Logger
	cfg  *config.Config
	repo repository.Auth
}

func NewAuthService(
	log *logger.Logger,
	cfg *config.Config,
	repo repository.Auth,
) *auth {
	return &auth{log: log, cfg: cfg, repo: repo}
}
func (s *auth) Register(
	ctx context.Context,
	user *model.NewUser,
) (*model.User, error) {
	if user.Password == "" {
		return nil, errors.New("Register(): empty password")
	}

	passhash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
	if err != nil {
		return nil, errors.New("Register(): " + err.Error())
	}

	newUser, err := s.repo.Register(ctx, user.Login, passhash, user.Name)
	if err != nil {
		return nil, errors.New("repo.Register(): " + err.Error())
	}
	newUser.Password = user.Password

	return newUser, nil
}
func (s *auth) Login(
	ctx context.Context,
	user *model.User,
) (*model.User, error) {

	entityUser, err := s.repo.Login(ctx, user.Login, user.Password)
	if err != nil {
		return nil, errors.New("repo.Login(): " + err.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(entityUser.Password), []byte(user.Password)); err != nil {
		return nil, errors.New("Register() bcrypt.Compare(): " + err.Error())
	}

	//return jwt.GenerateToken(ctx, entityUser.ID, tokenTTL)
	return nil, nil
}
