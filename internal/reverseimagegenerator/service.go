package reverseimagegenerator

import (
	"fmt"

	"github.com/truly-indian/reverseImageSearch/internal/config"
	"github.com/truly-indian/reverseImageSearch/internal/crawler"
	"github.com/truly-indian/reverseImageSearch/internal/logger"
	"github.com/truly-indian/reverseImageSearch/internal/serviceclients"
	"github.com/truly-indian/reverseImageSearch/internal/types"
)

var db map[string][]types.VisualMatch = make(map[string][]types.VisualMatch)

type Service interface {
	ReverseImageGenerator(types.ReverseImageGeneratorRequest) (types.ReverseImageGeneratorResponse, error)
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
	service := &serviceImpl{
		config:         config,
		serviceClients: serviceClients,
		crawler:        crawler,
		logger:         logger,
	}
	return service
}

func (s *serviceImpl) ReverseImageGenerator(req types.ReverseImageGeneratorRequest) (types.ReverseImageGeneratorResponse, error) {
	imageUrl := req.ImageUrl
	page := req.Page

	if len(db[imageUrl]) == 0 {
		serpResp, err := s.serviceClients.GetReverseImageData(imageUrl)
		if err != nil {
			s.logger.LogError(fmt.Sprintf("error while fetching serp api details for imgUrl: %v", imageUrl), err)
			return types.ReverseImageGeneratorResponse{}, err
		}
		db[imageUrl] = serpResp.VisualMatches
	}
	productList := getProductList(s, imageUrl, page, 5)

	return types.ReverseImageGeneratorResponse{
		Products: productList,
	}, nil
}

func getProductList(s *serviceImpl, imageUrl string, pageToken int, limit int) []types.Product {
	productList := []types.Product{}
	startIndex := (pageToken - 1) * limit
	if startIndex >= len(db[imageUrl]) {
		return productList
	}

	for i := startIndex; i < len(db[imageUrl]); i++ {
		vismatch := db[imageUrl][i]
		product := types.Product{}

		if vismatch.Title != "" {
			product.Name = vismatch.Title
		}
		if vismatch.Price.Extracted_value != 0.0 {
			product.Price = vismatch.Price.Extracted_value
		}
		if vismatch.Rating != 0.0 {
			product.UserRating = vismatch.Rating
		}
		if product.Name == "" || product.Price == 0.0 || product.UserRating == 0.0 {
			p, err := s.crawler.CrawlUrl(vismatch.Link)
			if err != nil {
				s.logger.LogError(fmt.Sprintf("error while crawling link: %v", vismatch.Link), err)
				if product.Name != "" {
					productList = append(productList, product)
				}
				continue
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
		}

		if product.Name != "" || product.Price != 0.0 || product.UserRating != 0.0 {
			productList = append(productList, product)
		}

		if len(productList) >= limit {
			return productList
		}
	}

	return productList
}
