package rest

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"github.com/jellydator/ttlcache/v2"
	"net/http"
	v1 "stock/pkg/meta/openfoodwatch/v1"
	"time"
)

type OFWCRestClient interface {
	GetCategories(ctx context.Context) (*v1.CategoriesResponse, error)
	GetProductsInCategory(ctx context.Context, categoryId string) (*v1.ProductsResponse, error)
	Get(ctx context.Context, productId string) (*v1.ProductDetailResponse, error)
}

type ofwClient struct {
	httpClient *http.Client
	baseURL    string
	username   string
	password   string
	baseUrl    string
	cache      ttlcache.SimpleCache
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

func JSON[R any](resp *http.Response, responseEntity *R) (*R, error) {
	defer resp.Body.Close()
	err := json.NewDecoder(resp.Body).Decode(responseEntity)
	if err != nil {
		return nil, err
	}
	return responseEntity, nil
}
