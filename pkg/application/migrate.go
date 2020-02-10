package application

import (
	log "github.com/sirupsen/logrus"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Migrate applies migrations
func (app *App) Migrate() {
	url := app.config.Database.URL
	path := app.config.Database.MigrationsPath
	m, err := migrate.New(path, url)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal(err)
	} else {
		log.WithFields(log.Fields{
			"url":  url,
			"path": path,
		}).Debug("Automigrate finished successfully")
	}
}
