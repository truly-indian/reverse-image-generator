package reverseimagegenerator

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/truly-indian/reverseImageSearch/internal/config"
)

type Handler struct {
	config *config.Config
}

func NewHandler(c *config.Config) *Handler {
	return &Handler{
		config: c,
	}
}

func (h *Handler) HellowWorldHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello from reverse image generator")
}
