package app

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/eeQuillibrium/posts/config"
	"github.com/eeQuillibrium/posts/internal/repository"
	"github.com/eeQuillibrium/posts/internal/service"
	"github.com/eeQuillibrium/posts/pkg/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type app struct {
	log *logger.Logger
	cfg *config.Config
}

func NewApp(
	log *logger.Logger,
	cfg *config.Config,
) *app {
	return &app{
		log: log,
		cfg: cfg,
	}
}

func (a *app) Run() error {
	ctx, cancel := signal.NotifyContext(context.TODO(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		a.cfg.PostgresDB.Host, a.cfg.PostgresDB.Port, a.cfg.PostgresDB.Username, a.cfg.PostgresDB.DBName, os.Getenv("DB_PASSWORD"), a.cfg.PostgresDB.SSLMode)
	db, err := sqlx.ConnectContext(ctx, "postgres", dsn)
	if err != nil {
		return errors.New("postgres db connect: " + err.Error())
	}

	repo := repository.NewRepository(a.log, a.cfg, db)
	services := service.NewService(a.log, a.cfg, repo)

	go func() {
		if err := a.runHttpServer(services, db); err != nil {
			a.log.Warnf("runHttpServer(): %v", err)
		}
		cancel()
	}()

	<-ctx.Done()

	return nil
}