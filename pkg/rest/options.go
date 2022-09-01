package rest

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
