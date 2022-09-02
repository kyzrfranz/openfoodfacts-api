package rest

import (
	"context"
	"fmt"
	"github.com/jellydator/ttlcache/v2"
	"log"
	v1 "stock/pkg/meta/openfoodwatch/v1"
	"strings"
)

func (c ofwClient) GetCategories(ctx context.Context) (*v1.CategoriesResponse, error) {
	requestUrl := fmt.Sprintf("%s/%s", c.baseURL, "categories.json")

	val, err := c.cache.Get("categories")
	if err != ttlcache.ErrNotFound {
		return val.(*v1.CategoriesResponse), nil
	}

	log.Println(fmt.Sprintf("calling %s", requestUrl))
	resp, err := c.httpClient.Get(requestUrl)
	if err != nil {
		return nil, err
	}

	responseEntity, err := JSON[v1.CategoriesResponse](resp, &v1.CategoriesResponse{})

	var restResponse = &v1.CategoriesResponse{Count: responseEntity.Count}
	for _, k := range responseEntity.Tags {
		k.Links = []v1.Link{{
			Rel:  "self",
			Href: fmt.Sprintf("/category/%s", strings.Replace(k.Id, "en:", "", 1)),
		}}
		restResponse.Tags = append(restResponse.Tags, k)
	}

	c.cache.Set("categories", restResponse)
	return restResponse, nil
}
