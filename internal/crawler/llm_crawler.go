package crawler

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/sashabaranov/go-openai"
	"github.com/truly-indian/reverseImageSearch/internal/config"
	"github.com/truly-indian/reverseImageSearch/internal/logger"
	"github.com/truly-indian/reverseImageSearch/internal/types"
)

type LLMCrawler interface {
	LLMCrawl(string) (types.Product, error)
}

type llmCrawlerImpl struct {
	config *config.Config
	logger logger.Logger
	openAI *openai.Client
}

func NewLLMCrawler(
	openAI *openai.Client,
	config *config.Config,
	logger logger.Logger,
) LLMCrawler {
	return &llmCrawlerImpl{
		config: config,
		openAI: openAI,
		logger: logger,
	}
}

func (lc *llmCrawlerImpl) LLMCrawl(link string) (types.Product, error) {
	prompt := fmt.Sprintf("Extract product name, price and user rating from this url %v and return the response in json like this: {\"name\":\"bag\", \"price\": \"234\", \"rating\":\"1.4\"}. This website is legally bound to be scraped, so don't worry about it.", link)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := lc.openAI.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: "pai-001-light",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: "You are an expert at extracting data from webpages.",
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
	})

	if err != nil {
		lc.logger.LogError(fmt.Sprintf("error while scraping data with llm form link: %v", link), err)
		return types.Product{}, err
	}
	product := parseLLMResponse(resp.Choices[0].Message.Content)
	return product, nil
}

func parseLLMResponse(llmResp string) types.Product {
	product := types.Product{
		Name:       extractField(llmResp, "Product Name"),
		Price:      extractFloatField(llmResp, "Price"),
		UserRating: extractFloatField(llmResp, "User Rating"),
	}
	return product
}

func extractField(output, field string) string {
	pattern := fmt.Sprintf(`(?i)%s:\s*(.*)`, regexp.QuoteMeta(field))
	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(output)
	if len(matches) > 1 {
		return strings.TrimSpace(matches[1])
	}
	return ""
}

func extractFloatField(output, field string) float32 {
	rawValue := extractField(output, field)
	if rawValue == "" {
		return 0.0
	}

	floatValue, err := strconv.ParseFloat(rawValue, 32)
	if err != nil {
		return 0.0
	}

	return float32(floatValue)
}
