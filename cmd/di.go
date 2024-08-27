//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/truly-indian/reverseImageSearch/internal/config"
	"github.com/truly-indian/reverseImageSearch/internal/crawler"
	"github.com/truly-indian/reverseImageSearch/internal/groq"
	"github.com/truly-indian/reverseImageSearch/internal/logger"
	"github.com/truly-indian/reverseImageSearch/internal/openai"
	"github.com/truly-indian/reverseImageSearch/internal/reverseimagegenerator"
	"github.com/truly-indian/reverseImageSearch/internal/server"
	"github.com/truly-indian/reverseImageSearch/internal/serviceclients"
	"github.com/truly-indian/reverseImageSearch/internal/utils"
)

type ServerDependencies struct {
	config   *config.Config
	server   *server.Server
	handlers server.Handlers
}

func InitDependencies() (ServerDependencies, error) {
	wire.Build(
		logger.WireSet,
		wire.Struct(new(ServerDependencies), "*"),
		wire.Struct(new(server.Handlers), "*"),
		server.WireSet,
		config.GetConfig,
		reverseimagegenerator.WireSet,
		serviceclients.WireSet,
		utils.WireSet,
		crawler.WireSet,
		openai.WireSet,
		groq.WireSet,
	)

	return ServerDependencies{}, nil
}
