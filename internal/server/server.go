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
)

type Server struct {
	config       *config.Config
	engine       *gin.Engine
	routerGroups RouterGroups
}

type RouterGroups struct {
	rootRouter *gin.Engine
}

func NewServer(c *config.Config) *Server {
	engine := gin.New()
	loggerConfig := gin.LoggerConfig{
		SkipPaths: []string{"/sanity", "/health"},
	}
	engine.Use(LoggerWithConfig(loggerConfig), gin.Recovery())

	return &Server{
		config: c,
		engine: engine,
		routerGroups: RouterGroups{
			rootRouter: engine,
		},
	}
}

func (s *Server) Run(h Handlers) {
	s.InitRoutes(h, s.config)
	srv := &http.Server{
		Addr:    s.config.ListenAddress(),
		Handler: s.engine,
	}

	go listenServer(srv)
	fmt.Println("Server is up and running Successfully :) !!")
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

// LoggerWithConfig instance a Logger middleware with config.
func LoggerWithConfig(conf gin.LoggerConfig) gin.HandlerFunc {

	notlogged := conf.SkipPaths

	var skip map[string]struct{}

	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notlogged {
			skip[path] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Log only when path is not being skipped
		if _, ok := skip[path]; !ok {
			param := gin.LogFormatterParams{
				Request: c.Request,
				Keys:    c.Keys,
			}

			// Stop timer
			param.TimeStamp = time.Now()
			param.Latency = param.TimeStamp.Sub(start)

			param.ClientIP = c.ClientIP()
			param.Method = c.Request.Method
			param.StatusCode = c.Writer.Status()
			param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
			//userId := c.Request.Header.Get("userId")
			param.BodySize = c.Writer.Size()

			if raw != "" {
				path = path + "?" + raw
			}

			// logger.Info(logger.Format{
			// 	Message: fmt.Sprintf("Accessing %s", path),
			// 	Data: map[string]string{
			// 		"method":     param.Method,
			// 		"clientIP":   param.ClientIP,
			// 		"statusCode": strconv.Itoa(param.StatusCode),
			// 		"error":      param.ErrorMessage,
			// 		"latency":    param.Latency.String(),
			// 		"size":       strconv.Itoa(param.BodySize),
			// 		"userId":     userId,
			// 	},
			// })
		}
	}
}
