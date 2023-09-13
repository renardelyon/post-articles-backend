package application

import (
	"article/config"
	"errors"
)

func (app *Application) Run(cfg *config.Config) error {
	if app.IsMigration {
		return runMigration(app)
	}
	return runApp(cfg, app)
}

func runMigration(app *Application) error {
	switch app.MigrationFlag {
	case "up":
		return app.MigrationRunner.Up()
	case "down":
		return app.MigrationRunner.Rollback()
	}
	return app.MigrationRunner.Up()
}

func runApp(cfg *config.Config, app *Application) error {
	return errors.New("unrecognized mode")
}
