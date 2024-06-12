package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/eeQuillibrium/posts/config"
	"github.com/eeQuillibrium/posts/internal/repository"
	"github.com/eeQuillibrium/posts/internal/service"
	"github.com/eeQuillibrium/posts/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	db, err := sqlx.ConnectContext(ctx, "postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		a.cfg.PostgresDB.Host, a.cfg.PostgresDB.Port, a.cfg.PostgresDB.Username, a.cfg.PostgresDB.DBName, os.Getenv("DB_PASSWORD"), a.cfg.PostgresDB.SSLMode))
	if err != nil {
		return fmt.Errorf("app.Run(): %w", err)
	}

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		a.log.Fatal(fmt.Errorf("app.CreateMigrations(): %w", err))
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		a.log.Fatal(fmt.Errorf("app.CreateMigrations(): %w", err))
	}
	defer m.Close()

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		a.log.Fatal(fmt.Errorf("app.CreateMigrations(): %w", err))
	}

	repo := repository.NewRepository(a.log, a.cfg, db)
	services := service.NewService(a.log, a.cfg, repo)

	go func() {
		if err := a.runHttpServer(services, db); err != nil {
			a.log.Fatalf("runHttpServer(): %v", err)
		}
		cancel()
	}()

	<-ctx.Done()

	return nil
}
