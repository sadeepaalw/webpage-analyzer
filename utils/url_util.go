package utils

import "net/url"

func GetBaseURL(rawUrl string) (*url.URL, error) {
	parsedURL, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}
	parsedURL.Path = ""
	parsedURL.RawQuery = ""
	parsedURL.Fragment = ""
	return parsedURL, nil
}
