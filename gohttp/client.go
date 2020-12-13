package gohttp

import (
	"net/http"
	"sync"
)

type httpClient struct {
	builder *clientBuilder

	client     *http.Client
	clientOnce sync.Once

}

type Client interface {
	Get(url string, headers http.Header, queryParams map[string]string) (*Response, error)
	Post(url string, body interface{}, headers http.Header, queryParams map[string]string) (*Response, error)
	Put(url string, body interface{}, headers http.Header, queryParams map[string]string) (*Response, error)
	Delete(url string, headers http.Header, queryParams map[string]string) (*Response, error)
}

func (c *httpClient) Get(url string, headers http.Header, queryParams map[string]string) (*Response, error) {
	return c.do(http.MethodGet, url, headers, nil, queryParams)
}

func (c *httpClient) Post(url string, body interface{}, headers http.Header, queryParams map[string]string) (*Response, error) {
	return c.do(http.MethodPost, url, headers, body, queryParams)
}

func (c *httpClient) Put(url string, body interface{}, headers http.Header, queryParams map[string]string) (*Response, error) {
	return c.do(http.MethodPut, url, headers, body, queryParams)
}

func (c *httpClient) Delete(url string, headers http.Header ,queryParams map[string]string) (*Response, error) {
	return c.do(http.MethodDelete, url, headers, nil, queryParams)
}
