package groupme

import "net/url"

// V3BaseURL is GroupMe's v3 API base URL.
const V3BaseURL = "https://api.groupme.com/v3"

func createURL(baseURL, route string, params string) (URL string, err error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}
	base.Path += route

	if params != "" {
		base.RawQuery = params
	}

	return base.String(), nil
}
