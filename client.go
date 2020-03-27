package groupme

// Client is a GroupMe API client.
type Client struct {
	BaseURL     string
	AccessToken string
}

// NewClient returns a new GroupMe API client.
func NewClient(baseURL, accessToken string) Client {
	return Client{
		BaseURL:     baseURL,
		AccessToken: accessToken,
	}
}
