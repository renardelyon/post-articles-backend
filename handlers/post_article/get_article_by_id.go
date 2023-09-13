package postarticle_handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *handler) GetArticleById(c *gin.Context) {
	h.log.Info("handler.GetArticleById")

	id := c.Param("id")

	h.log.WithFields(logrus.Fields{
		"id": id,
	}).Info()

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.repo.GetArticleById(h.ctx, idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
