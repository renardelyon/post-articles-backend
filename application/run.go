package application

import (
	"article/config"
	route "article/routes"
	"fmt"

	"github.com/gin-gonic/gin"
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
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.ErrorLogger())

	route.SetupRouterPostArticle(app.Context, app.Logger, r)

	app.Logger.Info("Starting server " + cfg.Application.ServerPort)
	if err := r.Run(fmt.Sprintf(":%s", cfg.Application.ServerPort)); err != nil {
		return err
	}
	return nil
}
