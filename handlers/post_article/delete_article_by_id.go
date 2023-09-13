package postarticle_handler

import (
	postarticle_model "article/model/post_article"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *handler) DeleteArticleById(c *gin.Context) {
	h.log.Info("handler.DeleteArticleById")

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

	err = h.repo.DeleteArticleById(h.ctx, tx, &postarticle_model.Post{}, idInt)
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
