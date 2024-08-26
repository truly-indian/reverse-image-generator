package serviceclients

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/truly-indian/reverseImageSearch/internal/config"
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
}

func NewSerpAPIClient(
	client *http.Client,
	cfg *config.Config,
	httpClient utils.HTTPClient,
) SerpAPIClient {
	return &serpAPIClient{
		client:     client,
		config:     cfg,
		httpClient: httpClient,
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
		fmt.Println("error while making serp api call: ", err)
		return types.SerpAPIResponse{}, err
	}

	if resp.StatusCode == http.StatusOK {
		var serpAPIResponse types.SerpAPIResponse
		unMarshalErr := json.Unmarshal(resp.Body, &serpAPIResponse)
		if unMarshalErr != nil {
			fmt.Println("error while unmarshalleing serp response: ", unMarshalErr)
			return types.SerpAPIResponse{}, unMarshalErr
		}
		return serpAPIResponse, nil
	}

	fmt.Println("serp API response is not 200, it is: ", resp.StatusCode)
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
