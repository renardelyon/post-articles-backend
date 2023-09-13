package postarticle_handler

import (
	postarticle_repo "article/repositories/post_articles"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type handler struct {
	ctx     context.Context
	log     *logrus.Logger
	dbWrite *gorm.DB
	repo    postarticle_repo.Repo
}

type Handler interface {
	InsertArticle(c *gin.Context)
	GetArticles(c *gin.Context)
	GetArticleById(c *gin.Context)
	UpdateArticleById(c *gin.Context)
	DeleteArticleById(c *gin.Context)
}

func NewHandler(ctx context.Context, log *logrus.Logger, dbWrite *gorm.DB, repo postarticle_repo.Repo) Handler {
	return &handler{
		ctx:     ctx,
		log:     log,
		dbWrite: dbWrite,
		repo:    repo,
	}
}
