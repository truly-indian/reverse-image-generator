package reverseimagegenerator

import (
	"encoding/json"
	"fmt"

	"github.com/truly-indian/reverseImageSearch/internal/config"
	"github.com/truly-indian/reverseImageSearch/internal/serviceclients"
	"github.com/truly-indian/reverseImageSearch/internal/types"
)

type Service interface {
	ReverseImageGenerator(types.ReverseImageGeneratorRequest) ([]types.ReverseImageGeneratorResponse, error)
}

type serviceImpl struct {
	config         *config.Config
	serviceClients serviceclients.SerpAPIClient
}

func NewService(config *config.Config, serviceClients serviceclients.SerpAPIClient) Service {
	service := &serviceImpl{
		config:         config,
		serviceClients: serviceClients,
	}
	return service
}

func (s *serviceImpl) ReverseImageGenerator(req types.ReverseImageGeneratorRequest) ([]types.ReverseImageGeneratorResponse, error) {

	imageUrl := req.ImageUrl

	serpResp, err := s.serviceClients.GetReverseImageData(imageUrl)

	if err != nil {
		fmt.Println("error while fetching serp api details for imgUrl: ", imageUrl, err)
		return []types.ReverseImageGeneratorResponse{}, err
	}

	jsonData, _ := json.Marshal(serpResp)
	fmt.Println("jsonResp: ", string(jsonData))

	fmt.Println("serpREsp: ", serpResp)

	return []types.ReverseImageGeneratorResponse{
		{
			ProductName: "Deepak Car",
			Title:       "Deepak Car Title",
			UserRating:  "1",
		},
	}, nil
}
