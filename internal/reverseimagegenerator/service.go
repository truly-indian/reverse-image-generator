package reverseimagegenerator

import (
	"github.com/truly-indian/reverseImageSearch/internal/config"
	"github.com/truly-indian/reverseImageSearch/internal/types"
)

type Service interface {
	ReverseImageGenerator(types.ReverseImageGeneratorRequest) ([]types.ReverseImageGeneratorResponse, error)
}

type serviceImpl struct {
	config *config.Config
}

func NewService(config *config.Config) Service {
	service := &serviceImpl{
		config: config,
	}
	return service
}

func (s *serviceImpl) ReverseImageGenerator(req types.ReverseImageGeneratorRequest) ([]types.ReverseImageGeneratorResponse, error) {
	return []types.ReverseImageGeneratorResponse{
		{
			ProductName: "Deepak Car",
			Title:       "Deepak Car Title",
			UserRating:  "1",
		},
	}, nil
}
