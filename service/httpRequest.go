package service

import (
	"log"
	"net/http"
	"time"
)

var HttpAPIClient HttpInterface = HttpClient{}

type HttpInterface interface {
	Get(string) (*http.Response, error)
}

type HttpClient struct{}

func (h HttpClient) Get(url string) (*http.Response, error) {

	client := http.Client{
		Timeout: time.Duration(2 * time.Second),
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Printf("Error while making request object => %v", err)
		return nil, err
	}
	return client.Do(req)
}
