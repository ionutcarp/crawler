package main

import (
	"fmt"
	"io"
	"net/http"
)

func getHTML(rawURL string) (string, error) {
	rawPage, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("error getting raw page: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(rawPage.Body)
	if rawPage.StatusCode > 399 {
		return "", fmt.Errorf("got HTTP error: %s", rawPage.Status)
	}
	if rawPage.Header.Get("content-type") != "text/html" {
		return "", fmt.Errorf("got non-HTML response")
	}

	html, err := io.ReadAll(rawPage.Body)
	if err != nil {
		return "", fmt.Errorf("couldn't read response body: %v", err)
	}
	return string(html), nil
}
