package reverseimagegenerator

import (
	"fmt"
	"sync"

	"github.com/truly-indian/reverseImageSearch/internal/config"
	"github.com/truly-indian/reverseImageSearch/internal/crawler"
	"github.com/truly-indian/reverseImageSearch/internal/logger"
	"github.com/truly-indian/reverseImageSearch/internal/serviceclients"
	"github.com/truly-indian/reverseImageSearch/internal/types"
)

var db map[string][]types.VisualMatch = make(map[string][]types.VisualMatch)

type Service interface {
	ReverseImageGenerator(
		types.ReverseImageGeneratorRequest,
		chan<- []types.Product,
	) error
}

type serviceImpl struct {
	config         *config.Config
	serviceClients serviceclients.SerpAPIClient
	crawler        crawler.Crawler
	logger         logger.Logger
}

func NewService(
	config *config.Config,
	serviceClients serviceclients.SerpAPIClient,
	crawler crawler.Crawler,
	logger logger.Logger,
) Service {
	return &serviceImpl{
		config:         config,
		serviceClients: serviceClients,
		crawler:        crawler,
		logger:         logger,
	}
}

func (s *serviceImpl) ReverseImageGenerator(req types.ReverseImageGeneratorRequest, productBatchChan chan<- []types.Product) error {
	imageUrl := req.ImageUrl
	limit := 5

	if len(db[imageUrl]) == 0 {
		serpResp, err := s.serviceClients.GetReverseImageData(imageUrl)
		if err != nil {
			s.logger.LogError(fmt.Sprintf("error while fetching serp api details for imgUrl: %v", imageUrl), err)
			return err
		}
		db[imageUrl] = serpResp.VisualMatches
	}

	visualMatches := db[imageUrl]
	processBatch := func(start, end int) ([]types.Product, error) {
		var wg sync.WaitGroup
		results := make([]types.Product, end-start)

		for i := start; i < end; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				match := visualMatches[i]
				product := types.Product{
					Name:       match.Title,
					Price:      match.Price.Extracted_value,
					UserRating: match.Rating,
				}
				if product.Name == "" || product.Price == 0.0 || product.UserRating == 0.0 {
					p, err := s.crawler.CrawlUrl(match.Link)
					if err != nil {
						s.logger.LogError(fmt.Sprintf("error while crawling link: %v", match.Link), err)
						if product.Name != "" || product.Price != 0.0 || product.UserRating != 0.0 {
							results[i-start] = product
						}
					}
					if product.Name == "" && p.Name != "" {
						product.Name = p.Name
					}
					if product.Price == 0.0 && p.Price != 0.0 {
						product.Price = p.Price
					}
					if product.UserRating == 0.0 && p.UserRating != 0.0 {
						product.UserRating = p.UserRating
					}
					if product.Name != "" || product.Price != 0.0 || product.UserRating != 0.0 {
						results[i-start] = product
					}
				}
			}(i)
		}
		wg.Wait()

		return results, nil
	}

	for i := 0; i < len(visualMatches); i += limit {
		start := i
		end := i + limit
		if end > len(visualMatches) {
			end = len(visualMatches)
		}

		batchProducts, err := processBatch(start, end)
		var filteredResults []types.Product
		for _, p := range batchProducts {
			if p.Name != "" || p.Price != 0.0 || p.UserRating != 0.0 {
				filteredResults = append(filteredResults, p)
			}
		}
		if err != nil {
			s.logger.LogError(fmt.Sprintf("error while processing batch from: %v to %v", start, end), err)
			return err
		}

		if len(filteredResults) > 0 {
			productBatchChan <- filteredResults
		}
	}

	return nil
}
