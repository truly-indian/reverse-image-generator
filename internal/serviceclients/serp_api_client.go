package serviceclients

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/truly-indian/reverseImageSearch/internal/config"
	"github.com/truly-indian/reverseImageSearch/internal/logger"
	"github.com/truly-indian/reverseImageSearch/internal/types"
	"github.com/truly-indian/reverseImageSearch/internal/utils"
)

type SerpAPIClient interface {
	GetReverseImageData(string) (types.SerpAPIResponse, error)
}

type serpAPIClient struct {
	client     *http.Client
	config     *config.Config
	httpClient utils.HTTPClient
	logger     logger.Logger
}

func NewSerpAPIClient(
	client *http.Client,
	cfg *config.Config,
	httpClient utils.HTTPClient,
	logger logger.Logger,
) SerpAPIClient {
	return &serpAPIClient{
		client:     client,
		config:     cfg,
		httpClient: httpClient,
		logger:     logger,
	}
}

func (s *serpAPIClient) GetReverseImageData(imageUrl string) (types.SerpAPIResponse, error) {
	apiPath := getUrl(s, imageUrl)
	resp, err := s.httpClient.Get(utils.HTTPPayload{
		Client:  s.client,
		URL:     apiPath,
		Timeout: s.config.GetSerpAPITimeOutInMs(),
	})

	if err != nil {
		s.logger.LogError(fmt.Sprintf("error while making serp api call for imageUrl: %v", imageUrl), err)
		return types.SerpAPIResponse{}, err
	}

	if resp.StatusCode == http.StatusOK {
		var serpAPIResponse types.SerpAPIResponse
		unMarshalErr := json.Unmarshal(resp.Body, &serpAPIResponse)
		if unMarshalErr != nil {
			s.logger.LogError(fmt.Sprintf("error while unmarshalleing serp response for imageUrl: %v", imageUrl), unMarshalErr)
			return types.SerpAPIResponse{}, unMarshalErr
		}
		return serpAPIResponse, nil
	}

	s.logger.LogError(fmt.Sprintf("serp API response is not 200, it is for imageUrl: %v", imageUrl), err)
	return types.SerpAPIResponse{}, errors.New("internal server error")
}

func getUrl(s *serpAPIClient, imgUrl string) string {
	encodedImageURL := url.QueryEscape(imgUrl)
	path := s.config.GetSerpAPI()
	path = strings.Replace(path, "{googleEngine}", "google_lens", 1)
	path = strings.Replace(path, "{imageUrl}", encodedImageURL, 1)
	path = strings.Replace(path, "{key}", s.config.GetSecrets()["serpAPIKey"], 1)
	return path
}
