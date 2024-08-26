// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/truly-indian/reverseImageSearch/internal/config"
	"github.com/truly-indian/reverseImageSearch/internal/reverseimagegenerator"
	"github.com/truly-indian/reverseImageSearch/internal/server"
	"github.com/truly-indian/reverseImageSearch/internal/serviceclients"
	"github.com/truly-indian/reverseImageSearch/internal/utils"
)

// Injectors from di.go:

func InitDependencies() (ServerDependencies, error) {
	configConfig := config.GetConfig()
	serverServer := server.NewServer(configConfig)
	client := server.NewHTTPClient()
	httpClient := utils.GetHTTPClient()
	serpAPIClient := serviceclients.NewSerpAPIClient(client, configConfig, httpClient)
	service := reverseimagegenerator.NewService(configConfig, serpAPIClient)
	handler := reverseimagegenerator.NewHandler(configConfig, service)
	handlers := server.Handlers{
		ReverseImageGenerator: handler,
	}
	serverDependencies := ServerDependencies{
		config:   configConfig,
		server:   serverServer,
		handlers: handlers,
	}
	return serverDependencies, nil
}

// di.go:

type ServerDependencies struct {
	config   *config.Config
	server   *server.Server
	handlers server.Handlers
}