package postarticle_repo

import (
	postarticle_model "article/model/post_article"
	"article/pkg/error_pkg"
	"context"

	"gorm.io/gorm"
)

func (r *repo) DeleteArticleById(ctx context.Context, tx *gorm.DB, post *postarticle_model.Post, id int) (err error) {
	r.logger.Info("repo.DeleteArticleById")
	defer error_pkg.CheckErr(r.logger, err)
	result := tx.WithContext(ctx).Where("id = ?", id).Delete(post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
