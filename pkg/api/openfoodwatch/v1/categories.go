package v1

import (
	"context"
	v1 "stock/pkg/meta/openfoodwatch/v1"
	"stock/pkg/rest"
)

type CategoriesGetter interface {
	Categories() CategoriesInterface
}

type CategoriesInterface interface {
	List(ctx context.Context) (*v1.CategoriesResponse, error)
}

type categories struct {
	client rest.OFWCRestClient
}

func newCategories(client rest.OFWCRestClient) categories {
	return categories{
		client: client,
	}
}

func (v categories) List(ctx context.Context) (result *v1.CategoriesResponse, err error) {
	result, err = v.client.GetCategories(ctx)
	return
}
