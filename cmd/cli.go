package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/truly-indian/reverseImageSearch/internal/config"
)

func initCLI() *cobra.Command {
	var cliCmd = &cobra.Command{
		Use:   "reverse-image-generator",
		Short: "service for generating reverse image",
	}

	cliCmd.AddCommand(startCommand())
	return cliCmd
}

func startCommand() *cobra.Command {
	var startCmd = &cobra.Command{
		Use:   "start",
		Short: "Starts the service",
		Run: func(cmd *cobra.Command, args []string) {
			configFile := "application"
			if len(args) > 0 {
				configFile = args[0]
			}

			configConfig := config.InitConfig(configFile)
			fmt.Println("configs: ", configConfig)
			serverDependencies, _ := InitDependencies()
			serverDependencies.server.Run(serverDependencies.handlers)
		},
	}

	return startCmd
}
