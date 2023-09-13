package postarticle_repo

import (
	postarticle_model "article/model/post_article"
	"context"
)

func (r *repo) GetArticleById(ctx context.Context, id int) (res postarticle_model.Post, err error) {
	r.logger.Info("repo.GetArticleById")

	query := `
	SELECT 
		id,
		title,
		content,
		category,
		created_date,
		updated_date,
		status
	FROM posts p 
	WHERE id = ?`

	stmt, err := r.dbRead.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	err = stmt.QueryRowContext(ctx, id).Scan(
		&res.Id,
		&res.Title,
		&res.Content,
		&res.Category,
		&res.CreatedDate,
		&res.UpdatedDate,
		&res.Status,
	)
	if err != nil {
		return
	}

	return
}
