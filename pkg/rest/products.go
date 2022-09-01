package rest

import (
	"context"
	"encoding/json"
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

	var responseEntity = &v1.ProductsResponse{}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(responseEntity)

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

	var responseEntity = &v1.ProductDetailResponse{}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(responseEntity)
	if err != nil {
		return nil, err
	}

	return responseEntity, nil
}
