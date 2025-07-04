package apiclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func setQueries(url *url.URL, queries map[string]string) {
	q := url.Query()
	for key, value := range queries {
		q.Set(key, value)
	}
	url.RawQuery = q.Encode()
}

func fetch(url string) (*http.Response, error) {
	res, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	return res, nil
}

func decodeJSONBody[T any](body io.ReadCloser, target *T) error {
	defer body.Close()

	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(bodyBytes, target); err != nil {
		return fmt.Errorf("failed to decode JSON: %w", err)
	}

	return nil
}
