package route

import (
	postarticle "article/handlers/post_article"
	postarticle_repo "article/repositories/post_articles"
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SetupRouterPostArticle(
	ctx context.Context,
	log *logrus.Logger,
	dbRead *sql.DB,
	dbWrite *gorm.DB,
	r *gin.Engine) {
	// ROUTING
	repo := postarticle_repo.NewRepo(log, dbRead, dbWrite)

	handler := postarticle.NewHandler(ctx, log, dbWrite, repo)

	r.POST("/article", handler.InsertArticle)
}
