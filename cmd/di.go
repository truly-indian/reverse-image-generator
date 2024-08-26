//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/truly-indian/reverseImageSearch/internal/config"
	"github.com/truly-indian/reverseImageSearch/internal/reverseimagegenerator"
	"github.com/truly-indian/reverseImageSearch/internal/server"
)

type ServerDependencies struct {
	config   *config.Config
	server   *server.Server
	handlers server.Handlers
}

func InitDependencies() (ServerDependencies, error) {
	wire.Build(
		wire.Struct(new(ServerDependencies), "*"),
		wire.Struct(new(server.Handlers), "*"),
		server.WireSet,
		config.GetConfig,
		reverseimagegenerator.WireSet,
	)

	return ServerDependencies{}, nil
}
