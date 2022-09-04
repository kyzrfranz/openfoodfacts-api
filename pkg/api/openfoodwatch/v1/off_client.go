package v1

import "stock/pkg/rest"

type OFFV1Interface interface {
	RESTClient() rest.OFFCRestClient
	CategoriesGetter
	ProductsGetter
}

type OFFV1Client struct {
	restClient rest.OFFCRestClient
}

func (c *OFFV1Client) Categories() CategoriesInterface {
	return newCategories(c.restClient)
}

func (c *OFFV1Client) Products() ProductsInterface {
	return newProducts(c.restClient)
}

func (c *OFFV1Client) RESTClient() rest.OFFCRestClient {
	if c == nil {
		return nil
	}
	return c.restClient
}

func NewForOpts(opts ...rest.OFFClientOption) *OFFV1Client {
	return &OFFV1Client{restClient: rest.NewOFFRestClient(opts...)}
}
