package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/truly-indian/reverseImageSearch/internal/config"
)

type Handlers struct {
}

func (h *Handlers) HelloWorldHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

func (s *Server) InitRoutes(h Handlers, c *config.Config) {
	router := s.routerGroups.rootRouter
	router.GET("/hello-world", h.HelloWorldHandler)
}
