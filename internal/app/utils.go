package app

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
	_ "github.com/golang-migrate/migrate/v4/source/file"

)

func (a *app) createMigrations(db *sqlx.DB) {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		a.log.Fatal(errors.New("app.CreateMigrations(): " + err.Error()))
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		a.log.Fatal(errors.New("app.CreateMigrations(): " + err.Error()))
	}
	defer m.Close()

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		a.log.Fatal(errors.New("app.CreateMigrations(): " + err.Error()))
	}
}
