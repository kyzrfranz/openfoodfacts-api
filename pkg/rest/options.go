package rest

import "github.com/jellydator/ttlcache/v2"

type OFFClientOption func(client *offClient)

func WithBaseURL(baseURL string) OFFClientOption {
	return func(c *offClient) {
		c.baseURL = baseURL
	}
}

func WithBasicAuth(username, password string) OFFClientOption {
	return func(c *offClient) {
		c.username = username
		c.password = password
	}
}

func WithCache(cache ttlcache.SimpleCache) OFFClientOption {
	return func(c *offClient) {
		c.cache = cache
	}
}
