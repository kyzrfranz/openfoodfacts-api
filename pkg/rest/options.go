package rest

import "github.com/jellydator/ttlcache/v2"

type OFWClientOption func(client *ofwClient)

func WithBaseURL(baseURL string) OFWClientOption {
	return func(c *ofwClient) {
		c.baseURL = baseURL
	}
}

func WithBasicAuth(username, password string) OFWClientOption {
	return func(c *ofwClient) {
		c.username = username
		c.password = password
	}
}

func WithCache(cache ttlcache.SimpleCache) OFWClientOption {
	return func(c *ofwClient) {
		c.cache = cache
	}
}
