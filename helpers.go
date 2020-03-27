package groupme

import "net/url"

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
