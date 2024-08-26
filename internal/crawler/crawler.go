package crawler

import (
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/truly-indian/reverseImageSearch/internal/config"
	"github.com/truly-indian/reverseImageSearch/internal/types"
)

type Crawler interface {
	CrawlUrl(string) (types.Product, error)
}

type crawlerImpl struct {
	config *config.Config
}

func NewCrawler(c *config.Config) Crawler {
	return &crawlerImpl{
		config: c,
	}
}

func (cr *crawlerImpl) CrawlUrl(link string) (types.Product, error) {
	c := colly.NewCollector()
	product := types.Product{}

	c.SetRequestTimeout(1 * time.Second)
	c.OnHTML("title, h1, meta[property='og:title']", func(e *colly.HTMLElement) {
		product.Name = e.Text
	})

	c.OnHTML("span.price, div.price, meta[property='og:price:amount", func(e *colly.HTMLElement) {
		product.Price = e.Text
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("something went wrong while crawling: ", err, link)
	})

	err := c.Visit(link)

	if err != nil {
		return types.Product{}, err
	}

	return product, nil
}
