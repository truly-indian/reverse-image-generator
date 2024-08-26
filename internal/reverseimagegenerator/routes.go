package reverseimagegenerator

import "github.com/gin-gonic/gin"

func (h *Handler) InitRoutes(router *gin.Engine) {
	router.GET("/sanity", h.HellowWorldHandler)
}
