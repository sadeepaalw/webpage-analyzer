package adapter

import (
	"io"
	"net/http"
	"time"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type RequestInvoker struct {
	Client HTTPClient
}

func NewRequestInvoker() *RequestInvoker {
	return &RequestInvoker{
		Client: &http.Client{Timeout: 5 * time.Second},
	}
}

func (ri *RequestInvoker) InvokeRequest(url string, method string) ([]byte, int, error) {

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, 0, err
	}

	resp, err := ri.Client.Do(req)
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
