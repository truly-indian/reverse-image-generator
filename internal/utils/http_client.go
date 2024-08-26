package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const ApplicationJSON = "application/json"

const (
	PUT  string = "PUT"
	GET  string = "GET"
	POST string = "POST"
)

type HTTPPayload struct {
	Client  *http.Client
	URL     string
	Body    interface{}
	Headers map[string]string
	Timeout time.Duration
}

type HTTPResponse struct {
	StatusCode int
	Body       []byte
}

type HTTPClient interface {
	Put(hp HTTPPayload) (HTTPResponse, error)
	Post(hp HTTPPayload) (HTTPResponse, error)
	Get(hp HTTPPayload) (HTTPResponse, error)
}

type httpClientImpl struct {
}

func (h *httpClientImpl) Put(hp HTTPPayload) (HTTPResponse, error) {
	body := new(bytes.Buffer)
	if hp.Body != nil {
		_ = json.NewEncoder(body).Encode(hp.Body)
	}
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), hp.Timeout)
	defer cancel()

	httpRequest, _ := http.NewRequestWithContext(ctxWithTimeout, PUT, hp.URL, body)
	setRequestHeaders(hp.Headers, httpRequest)
	response, err := hp.Client.Do(httpRequest)
	if err != nil {
		return HTTPResponse{StatusCode: 500}, err
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	return HTTPResponse{
		StatusCode: response.StatusCode,
		Body:       responseBody,
	}, err
}

func (h *httpClientImpl) Post(hp HTTPPayload) (HTTPResponse, error) {
	body := new(bytes.Buffer)
	if hp.Body != nil {
		_ = json.NewEncoder(body).Encode(hp.Body)
	}
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), hp.Timeout)
	defer cancel()

	httpRequest, _ := http.NewRequestWithContext(ctxWithTimeout, POST, hp.URL, body)
	//httpRequest.Header.Set("Content-Type", ApplicationJSON)
	setRequestHeaders(hp.Headers, httpRequest)
	response, err := hp.Client.Do(httpRequest)
	if err != nil {
		return HTTPResponse{StatusCode: 500}, err
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	return HTTPResponse{
		StatusCode: response.StatusCode,
		Body:       responseBody,
	}, err
}

func (h *httpClientImpl) Get(hp HTTPPayload) (HTTPResponse, error) {
	body := new(bytes.Buffer)
	if hp.Body != nil {
		_ = json.NewEncoder(body).Encode(hp.Body)
	}
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), hp.Timeout)
	defer cancel()

	httpRequest, _ := http.NewRequestWithContext(ctxWithTimeout, GET, hp.URL, body)
	//httpRequest.Header.Set("Content-Type", ApplicationJSON)
	setRequestHeaders(hp.Headers, httpRequest)
	response, err := hp.Client.Do(httpRequest)
	if err != nil {
		return HTTPResponse{StatusCode: 500}, err
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	return HTTPResponse{
		StatusCode: response.StatusCode,
		Body:       responseBody,
	}, err
}

func GetHTTPClient() HTTPClient {
	return &httpClientImpl{}
}

func setRequestHeaders(headers map[string]string, req *http.Request) {
	req.Header.Set("Content-Type", ApplicationJSON)
	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}
}
