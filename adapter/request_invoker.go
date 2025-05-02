package adapter

import (
	"net/http"
	"time"
)

func InvokeRequest(url string, method string) (*http.Response, error) {

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	//defer resp.Body.Close()

	return resp, nil
}
