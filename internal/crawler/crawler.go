package crawler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/truly-indian/reverseImageSearch/internal/config"
	"github.com/truly-indian/reverseImageSearch/internal/logger"
	"github.com/truly-indian/reverseImageSearch/internal/types"
)

type Crawler interface {
	CrawlUrl(string) (types.Product, error)
}

type crawlerImpl struct {
	config     *config.Config
	logger     logger.Logger
	llmCrawler LLMCrawler
}

func NewCrawler(
	c *config.Config,
	logger logger.Logger,
	llmCrawler LLMCrawler,
) Crawler {
	return &crawlerImpl{
		config:     c,
		logger:     logger,
		llmCrawler: llmCrawler,
	}
}

func (cr *crawlerImpl) CrawlUrl(link string) (types.Product, error) {
	c := colly.NewCollector()
	product := types.Product{}

	var htmlContent string

	c.OnResponse(func(r *colly.Response) {
		htmlContent = string(r.Body)
	})

	c.SetRequestTimeout(10 * time.Second)

	c.OnHTML("title, h1, meta[property='og:title']", func(e *colly.HTMLElement) {
		product.Name = e.Text
	})

	c.OnHTML("span.price, div.price, meta[property='og:price:amount", func(e *colly.HTMLElement) {
		price, _ := strconv.ParseFloat(e.Text, 32)
		product.Price = float32(price)
	})

	c.OnHTML("span.rating, div.rating, meta[property='og:rating:user-rating']", func(e *colly.HTMLElement) {
		rating, _ := strconv.ParseFloat(e.Text, 32)
		product.UserRating = float32(rating)
	})

	c.OnError(func(_ *colly.Response, err error) {
		cr.logger.LogError(fmt.Sprintf("something went wrong while crawling: %v", link), err)
	})

	err := c.Visit(link)

	if err != nil {
		return types.Product{}, err
	}

	if product.Name == "" || product.Price == 0.0 || product.UserRating == 0.0 {
		cr.logger.LogInfo("Fetching Product Details from LLM Crawler")
		product, _ = cr.llmCrawler.LLMCrawl(htmlContent)
	}

	return product, nil
}
