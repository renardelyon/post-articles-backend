package postarticle_repo

import (
	postarticle_model "article/model/post_article"
	"context"
)

func (r *repo) GetArticles(ctx context.Context, limit, offset int) (res []postarticle_model.Post, err error) {
	r.logger.Info("repo.GetArticles")

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
	LIMIT ? OFFSET ?`

	stmt, err := r.dbRead.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	rows, err := stmt.QueryContext(ctx, limit, offset)
	if err != nil {
		return
	}

	for rows.Next() {
		var post postarticle_model.Post
		err = rows.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.Category,
			&post.CreatedDate,
			&post.UpdatedDate,
			&post.Status,
		)
		if err != nil {
			return
		}

		res = append(res, post)
	}

	return
}
