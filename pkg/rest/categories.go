package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	v1 "stock/pkg/meta/openfoodwatch/v1"
	"strings"
)

func (c ofwClient) GetCategories(ctx context.Context) (*v1.CategoriesResponse, error) {
	requestUrl := fmt.Sprintf("%s/%s", c.baseURL, "categories.json")

	log.Println(fmt.Sprintf("calling %s", requestUrl))
	resp, err := c.httpClient.Get(requestUrl)
	if err != nil {
		return nil, err
	}

	var responseEntity = &v1.CategoriesResponse{}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(responseEntity)

	var restResponse = &v1.CategoriesResponse{Count: responseEntity.Count}
	for _, k := range responseEntity.Tags {
		k.Links = []v1.Link{v1.Link{
			Rel:  "self",
			Href: fmt.Sprintf("/category/%s", strings.Replace(k.Id, "en:", "", 1)),
		}}
		restResponse.Tags = append(restResponse.Tags, k)
	}

	return restResponse, nil
}
