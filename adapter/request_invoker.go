package adapter

import (
	"io"
	"net/http"
	"time"
)

func InvokeRequest(url string, method string) ([]byte, int, error) {

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, 0, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return body, resp.StatusCode, nil
}
