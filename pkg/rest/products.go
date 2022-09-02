package rest

import (
	"context"
	"fmt"
	"log"
	v1 "stock/pkg/meta/openfoodwatch/v1"
)

func (c ofwClient) GetProductsInCategory(ctx context.Context, categoryId string) (*v1.ProductsResponse, error) {
	requestUrl := fmt.Sprintf("%s/%s/%s.json", c.baseURL, "category", categoryId)

	log.Println(fmt.Sprintf("calling %s", requestUrl))
	resp, err := c.httpClient.Get(requestUrl)
	if err != nil {
		return nil, err
	}

	responseEntity, err := JSON[v1.ProductsResponse](resp, &v1.ProductsResponse{})

	var restResponse = &v1.ProductsResponse{Count: responseEntity.Count}
	for _, k := range responseEntity.Products {
		k.Links = []v1.Link{v1.Link{
			Rel:  "self",
			Href: fmt.Sprintf("/products/%s", k.Id),
		}}
		restResponse.Products = append(restResponse.Products, k)
	}

	return restResponse, nil
}

func (c ofwClient) Get(ctx context.Context, productId string) (*v1.ProductDetailResponse, error) {
	requestUrl := fmt.Sprintf("%s/%s/%s", c.baseURL, "api/v2/product", productId)

	log.Println(fmt.Sprintf("calling %s", requestUrl))
	resp, err := c.httpClient.Get(requestUrl)
	if err != nil {
		return nil, err
	}

	return JSON[v1.ProductDetailResponse](resp, &v1.ProductDetailResponse{})
}
