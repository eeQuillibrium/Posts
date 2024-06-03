package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/eeQuillibrium/posts/config"
	"github.com/eeQuillibrium/posts/graph/model"
	"github.com/eeQuillibrium/posts/internal/repository"
	"github.com/eeQuillibrium/posts/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

var (
	errEmptyInput = errors.New("Register(): empty password or login")
)

type authService struct {
	log  *logger.Logger
	cfg  *config.Config
	repo repository.Auth
}

func NewAuthService(
	log *logger.Logger,
	cfg *config.Config,
	repo repository.Auth,
) Auth {
	return &authService{log: log, cfg: cfg, repo: repo}
}
func (as *authService) CreateUser(
	ctx context.Context,
	user *model.NewUser,
) (int, error) {
	if user.Password == "" || user.Login == "" {
		return 0, errEmptyInput
	}

	passhash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
	if err != nil {
		return 0, fmt.Errorf("authService.Register():\n%w", err)
	}

	userID, err := as.repo.CreateUser(ctx, user.Login, passhash, user.Name)
	if err != nil {
		return 0, fmt.Errorf("authService.Register():\n%w", err)
	}

	return userID, nil
}
