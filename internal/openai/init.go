package openai

import (
	"fmt"

	"github.com/sashabaranov/go-openai"
	"github.com/truly-indian/reverseImageSearch/internal/config"
	"github.com/truly-indian/reverseImageSearch/internal/logger"
)

func NewOpenAIClient(config *config.Config, logger logger.Logger) *openai.Client {
	logger.LogInfo("Initialising openAIClient")
	c := openai.DefaultConfig(config.GetSecrets()["openAIKey"])
	c.BaseURL = "https://api.pawan.krd/v1"
	openAIClient := openai.NewClientWithConfig(c)
	fmt.Println("initialised openAICLineT: ", openAIClient)
	return openAIClient
}
