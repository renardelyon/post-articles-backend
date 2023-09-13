package postarticle_repo

import (
	postarticle_model "article/model/post_article"
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type repo struct {
	logger  *logrus.Logger
	dbRead  *sql.DB
	dbWrite *gorm.DB
}

type Repo interface {
	InsertArticle(ctx context.Context, tx *gorm.DB, post *postarticle_model.Post) (err error)
	GetArticles(ctx context.Context, limit, offset int) (res []postarticle_model.Post, err error)
}

func NewRepo(logger *logrus.Logger, dbRead *sql.DB, dbWrite *gorm.DB) Repo {
	return &repo{
		logger:  logger,
		dbRead:  dbRead,
		dbWrite: dbWrite,
	}
}
