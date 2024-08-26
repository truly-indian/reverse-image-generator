package reverseimagegenerator

import (
	"fmt"

	"github.com/truly-indian/reverseImageSearch/internal/config"
	"github.com/truly-indian/reverseImageSearch/internal/crawler"
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
}

func NewService(
	config *config.Config,
	serviceClients serviceclients.SerpAPIClient,
	crawler crawler.Crawler,
) Service {
	service := &serviceImpl{
		config:         config,
		serviceClients: serviceClients,
		crawler:        crawler,
	}
	return service
}

func (s *serviceImpl) ReverseImageGenerator(req types.ReverseImageGeneratorRequest) (types.ReverseImageGeneratorResponse, error) {
	limit := 5
	imageUrl := req.ImageUrl
	page := req.Page

	if len(db[imageUrl]) == 0 {
		serpResp, err := s.serviceClients.GetReverseImageData(imageUrl)
		if err != nil {
			fmt.Println("error while fetching serp api details for imgUrl: ", imageUrl, err)
			return types.ReverseImageGeneratorResponse{}, err
		}
		db[imageUrl] = serpResp.VisualMatches
	}
	startIndex := (page - 1) * limit
	endIndex := startIndex + limit
	if startIndex >= len(db[imageUrl]) {
		return types.ReverseImageGeneratorResponse{
			Products: []types.Product{},
		}, nil
	}
	relevantMatches := db[imageUrl][startIndex:endIndex]
	productList := getProductList(s, relevantMatches)
	return types.ReverseImageGeneratorResponse{
		Products: productList,
	}, nil
}

func getProductList(s *serviceImpl, matches []types.VisualMatch) []types.Product {
	productList := []types.Product{}
	for _, vismatch := range matches {
		product := types.Product{}
		if vismatch.Title != "" {
			product.Name = vismatch.Title
		}
		if vismatch.Price.Value != "" {
			product.Price = vismatch.Price.Value
		}
		if product.Name == "" || product.Price == "" {
			p, err := s.crawler.CrawlUrl(vismatch.Link)
			if err != nil {
				fmt.Println("error while crawling link: ", vismatch.Link, err)
				continue
			}
			if product.Name == "" && p.Name != "" {
				product.Name = p.Name
			}
			if product.Price == "" && p.Price != "" {
				product.Price = p.Price
			}
		}
		productList = append(productList, product)
	}
	return productList
}
