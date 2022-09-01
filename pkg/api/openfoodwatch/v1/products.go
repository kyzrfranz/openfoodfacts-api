package v1

import (
	"context"
	v1 "stock/pkg/meta/openfoodwatch/v1"
	"stock/pkg/rest"
)

type ProductsGetter interface {
	Products() ProductsInterface
}

type ProductsInterface interface {
	GetForCategory(ctx context.Context, categoryId string) (*v1.ProductsResponse, error)
	Get(ctx context.Context, productId string) (*v1.ProductDetailResponse, error)
}

type Products struct {
	client rest.OFWCRestClient
}

func newProducts(client rest.OFWCRestClient) Products {
	return Products{
		client: client,
	}
}

func (v Products) GetForCategory(ctx context.Context, categoryId string) (result *v1.ProductsResponse, err error) {
	result, err = v.client.GetProductsInCategory(ctx, categoryId)
	return
}

func (v Products) Get(ctx context.Context, productId string) (result *v1.ProductDetailResponse, err error) {
	result, err = v.client.Get(ctx, productId)
	return
}
