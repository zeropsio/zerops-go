package sdkBase

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func NewEnvironment(config Config, httpClient *http.Client) Environment {
	return newEnvironment(config, httpClient, http.Header{})
}

func newEnvironment(config Config, httpClient *http.Client, headers http.Header) Environment {
	return Environment{
		config:     config,
		httpClient: httpClient,
		headers:    headers,
	}
}

type Environment struct {
	config     Config
	httpClient *http.Client
	headers    http.Header
}

func (e Environment) Request(ctx context.Context, method string, url string, body io.Reader) (req *http.Request, err error) {
	url = strings.TrimRight(e.config.Endpoint, "/") + "/" + strings.TrimLeft(url, "/")

	req, err = http.NewRequest(method, url, body)
	if err != nil {
		return
	}
	req = req.WithContext(ctx)
	req.Header = e.headers.Clone()
	req.Header.Add("Content-Type", "application/json")
	if e.headers.Get("host") != "" {
		req.Host = e.headers.Get("host")
	}
	return
}

func (e Environment) Do(req *http.Request) (*http.Response, error) {
	return e.httpClient.Do(req)
}

func (e Environment) Authorize(in string) Environment {
	headers := e.headers.Clone()
	headers.Add("Authorization", fmt.Sprintf("Bearer %s", in))
	return newEnvironment(e.config, e.httpClient, headers)
}
