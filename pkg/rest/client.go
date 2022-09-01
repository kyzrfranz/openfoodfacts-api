package rest

import (
	"context"
	"crypto/tls"
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
