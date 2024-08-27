package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/truly-indian/reverseImageSearch/internal/config"
	"github.com/truly-indian/reverseImageSearch/internal/logger"
)

type Server struct {
	config       *config.Config
	engine       *gin.Engine
	routerGroups RouterGroups
	logger       logger.Logger
}

type RouterGroups struct {
	rootRouter *gin.Engine
}

func NewServer(c *config.Config, logger logger.Logger) *Server {
	engine := gin.New()

	return &Server{
		config: c,
		engine: engine,
		routerGroups: RouterGroups{
			rootRouter: engine,
		},
		logger: logger,
	}
}

func (s *Server) Run(h Handlers) {
	s.InitRoutes(h, s.config)
	srv := &http.Server{
		Addr:    s.config.ListenAddress(),
		Handler: s.engine,
	}

	go listenServer(srv)
	s.logger.LogInfo("Server is up and running Successfully :) !!")
	waitForShutdown(srv)
}

func listenServer(server *http.Server) {
	err := server.ListenAndServe()
	if err != http.ErrServerClosed {
		panic(err)
	}
}

func waitForShutdown(server *http.Server) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig,
		syscall.SIGINT,
		syscall.SIGTERM)
	_ = <-sig

	fmt.Println("server shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Shutdown(ctx)
}
