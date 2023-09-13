package postarticle_repo

import (
	postarticle_model "article/model/post_article"
	"article/pkg/error_pkg"
	"context"

	"gorm.io/gorm"
)

func (r *repo) InsertArticle(ctx context.Context, tx *gorm.DB, post *postarticle_model.Post) (err error) {
	r.logger.Info("repo.InsertArticle")
	defer error_pkg.CheckErr(r.logger, err)
	result := tx.WithContext(ctx).Create(post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
