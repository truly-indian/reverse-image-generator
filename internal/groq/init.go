package groq

import (
	"github.com/jpoz/groq"
	"github.com/truly-indian/reverseImageSearch/internal/config"
	"github.com/truly-indian/reverseImageSearch/internal/logger"
)

func NewGroq(config *config.Config, logger logger.Logger) *groq.Client {
	logger.LogInfo("Initialising groqClient")
	client := groq.NewClient(groq.WithAPIKey(config.GetSecrets()["groqAIKey"]))
	logger.LogInfo("Initiased groq client successfully")
	return client
}
