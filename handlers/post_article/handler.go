package postarticle

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type handler struct {
	ctx context.Context
	log *logrus.Logger
}

type Handler interface {
	InsertArticle(c *gin.Context)
}

func NewHandler(ctx context.Context, log *logrus.Logger) Handler {
	return &handler{
		ctx: ctx,
		log: log,
	}
}
