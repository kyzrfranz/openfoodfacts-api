package v1

import "stock/pkg/rest"

type OFWV1Interface interface {
	RESTClient() rest.OFWCRestClient
	CategoriesGetter
}

type OFWV1Client struct {
	restClient rest.OFWCRestClient
}

func (c *OFWV1Client) Categories() CategoriesInterface {
	return newCategories(c.restClient)
}

func (c *OFWV1Client) RESTClient() rest.OFWCRestClient {
	if c == nil {
		return nil
	}
	return c.restClient
}

func NewForOpts() *OFWV1Client {
	return &OFWV1Client{restClient: rest.NewOFWRestClient()}
}
