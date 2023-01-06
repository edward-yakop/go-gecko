package coingecko

import (
	"fmt"
	"io"
	"net/http"
)

var baseURL = "https://api.coingecko.com/api/v3"

type HttpRequestModifier func(r *http.Request)

// Client struct
type Client struct {
	httpClient          *http.Client
	baseURL             string
	httpRequestModifier HttpRequestModifier
}

type ClientOption func(client *Client)

func WithHttpRequestModifier(f HttpRequestModifier) ClientOption {
	return func(c *Client) {
		c.httpRequestModifier = f
	}
}

func WithAPIKey(apiKey string) ClientOption {
	return func(c *Client) {
		c.baseURL = "https://pro-api.coingecko.com/api/v3"
		c.httpRequestModifier = func(r *http.Request) {
			r.Header.Set("x-cg-pro-api-key", apiKey)
		}
	}
}

// NewClient create new client object
func NewClient(httpClient *http.Client, options ...ClientOption) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		baseURL:    baseURL,
		httpClient: httpClient,
	}

	for _, option := range options {
		option(c)
	}

	return c
}

// helper
// doReq HTTP client
func doReq(req *http.Request, client *http.Client) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if http.StatusOK != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}

	return body, nil
}

// makeHTTPRequest HTTP request helper
func (c *Client) makeHTTPRequest(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	if c.httpRequestModifier != nil {
		c.httpRequestModifier(req)
	}

	resp, err := doReq(req, c.httpClient)
	if err != nil {
		return nil, err
	}

	return resp, err
}
