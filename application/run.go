package application

import (
	"article/config"
	"article/pkg"
	route "article/routes"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("enum", pkg.ValidateEnum)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.ErrorLogger())

	articleDBClient := app.DbClients["article"]

	route.SetupRouterPostArticle(app.Context, app.Logger, articleDBClient.SqlAdapter, articleDBClient.OrmAdapter, r)

	app.Logger.Info("Starting server " + cfg.Application.ServerPort)
	if err := r.Run(fmt.Sprintf(":%s", cfg.Application.ServerPort)); err != nil {
		return err
	}
	return nil
}
