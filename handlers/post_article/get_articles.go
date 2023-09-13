package postarticle_handler

import (
	postarticle_payload "article/payload/post_article"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *handler) GetArticles(c *gin.Context) {
	h.log.Info("handler.GetArticles")

	var payload postarticle_payload.GetArticlesPaylod
	err := c.ShouldBindQuery(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.log.WithFields(logrus.Fields{
		"limit":  payload.Limit,
		"offset": payload.Offset,
	}).Info()

	res, err := h.repo.GetArticles(h.ctx, payload.Limit, payload.Offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
