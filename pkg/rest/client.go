package rest

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	v1 "stock/pkg/meta/openfoodwatch/v1"
	"time"
)

type OFWCRestClient interface {
	GetCategories(ctx context.Context) (*v1.CategoryResponse, error)
}

type ofwClient struct {
	httpClient *http.Client
	baseURL    string
	username   string
	password   string
	baseUrl    string
}

func NewOFWRestClient(opts ...OFWClientOption) OFWCRestClient {

	owc := &ofwClient{
		httpClient: &http.Client{
			Timeout: time.Second * 60,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
				TLSHandshakeTimeout: time.Second * 15,
			},
		},
		username: "",
		password: "",
	}

	for _, opt := range opts {
		opt(owc)
	}

	return owc
}

func (c ofwClient) GetCategories(ctx context.Context) (*v1.CategoryResponse, error) {
	requestUrl := fmt.Sprintf("%s/%s", c.baseURL, "categories.json")

	log.Println(fmt.Sprintf("calling %s", requestUrl))
	resp, err := c.httpClient.Get(requestUrl)
	if err != nil {
		return nil, err
	}

	var responseEntity = &v1.CategoryResponse{}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(responseEntity)
	return responseEntity, nil
}
