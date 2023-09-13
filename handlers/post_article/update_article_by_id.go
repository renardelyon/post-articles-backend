package postarticle_handler

import (
	postarticle_model "article/model/post_article"
	postarticle_payload "article/payload/post_article"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *handler) UpdateArticleById(c *gin.Context) {
	h.log.Info("handler.UpdateArticleById")

	var payload postarticle_payload.PostArticlePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")

	h.log.WithFields(logrus.Fields{
		"id": id,
	}).Info()

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := h.dbWrite.Begin()
	defer tx.Rollback()

	err = h.repo.UpdateArticleById(h.ctx, tx, &postarticle_model.Post{
		Title:    payload.Title,
		Content:  payload.Content,
		Category: payload.Category,
		Status:   (string)(payload.Status),
	}, idInt)
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
