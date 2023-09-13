package postarticle

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) InsertArticle(c *gin.Context) {
	c.JSON(http.StatusOK, "hello world")
}
