package postarticle

import (
	postarticle_model "article/model/post_article"
	postarticle_payload "article/payload/post_article"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) InsertArticle(c *gin.Context) {
	h.log.Info("handler.InsertArticle")

	var payload postarticle_payload.PostArticlePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := h.dbWrite.Begin()
	defer tx.Rollback()

	err := h.repo.InsertArticle(h.ctx, tx, &postarticle_model.Post{
		Title:    payload.Title,
		Content:  payload.Content,
		Category: payload.Category,
		Status:   (string)(payload.Status),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := tx.Commit()
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
