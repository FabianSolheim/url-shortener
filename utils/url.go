package utils

import (
	"fmt"
	"net/url"
	"strings"
)

func ParseAlias(alias string) (string, error) {
	alias = strings.TrimSpace(alias)

	if alias == "" {
		return "", fmt.Errorf("alias cannot be empty")
	}

	alias = strings.ReplaceAll(alias, " ", "-")

	return alias, nil
}

func ParseAndValidateUrl(inputURL string) (*url.URL, error) {
    inputURL = strings.TrimSpace(inputURL)

    if inputURL == "" {
        return nil, fmt.Errorf("URL cannot be empty")
    }

    parsedURL, err := url.Parse(inputURL)
    if err != nil {
        return nil, err
    }

    if !parsedURL.IsAbs() {
        return nil, fmt.Errorf("invalid URL: must be an absolute URL")
    }

    return parsedURL, nil
}