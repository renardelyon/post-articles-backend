package route

import (
	postarticle "article/handlers/post_article"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetupRouterPostArticle(ctx context.Context, log *logrus.Logger, r *gin.Engine) {
	// ROUTING
	handler := postarticle.NewHandler(ctx, log)

	r.GET("/article", handler.InsertArticle)
}
