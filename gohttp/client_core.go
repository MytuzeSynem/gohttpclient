package gohttp

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	defaultConnectionTimeout  = 1 * time.Second
)

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}, queryParams map[string]string) (*Response, error) {

	allHeaders := c.getRequestHeaders(headers)

	requestBody, err := c.getRequestBody(allHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("unable to create request")
	}

	q := request.URL.Query()
	for  queryParam, value := range queryParams{
		if len(queryParam) > 0 {
			q.Add(queryParam, value)
		}

	}
	request.URL.RawQuery = q.Encode()

	request.Header = allHeaders

	client := c.getHttpClient()

	response, err := client.Do(request)
	if err != nil{
		return nil, err
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	finalResponse := Response{
		status: response.Status,
		statusCode: response.StatusCode,
		headers: response.Header,
		body: responseBody,

	}
	return &finalResponse, nil
}

func (c *httpClient) getHttpClient() *http.Client {
	c.clientOnce.Do(func() {
		c.client = &http.Client{
			Timeout: c.getConnectionTimeout(),
		}
	})

	return c.client
}

func (c *httpClient) getRequestHeaders(requestHeaders http.Header) http.Header  {

	result :=  make(http.Header)

	for header, value := range requestHeaders {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}
	return result
}

func (c *httpClient) getRequestBody (contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return  nil, nil
	}

	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(body)

	default:
		return json.Marshal(body)
	}
}

func (c *httpClient) getConnectionTimeout() time.Duration {
	if c.builder.connectionTimeout > 0 {
		return c.builder.connectionTimeout
	}
	return defaultConnectionTimeout
}

