package coingecko

import (
	"encoding/json"
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
func doReq(req *http.Request, client *http.Client) ([]byte, http.Header, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	if http.StatusOK != resp.StatusCode {
		return nil, nil, fmt.Errorf("%s", body)
	}

	//dumpResponse(resp, body)

	return body, resp.Header, nil
}

func dumpResponse(resp *http.Response, body []byte) {
	_ = os.WriteFile("resp.json", body, os.ModePerm)

	cp := headerToMap(resp.Header, "cache-control", "expires", "age")

	headerBA, _ := json.Marshal(cp)
	_ = os.WriteFile("resp_header.json", headerBA, os.ModePerm)
}

func headerToMap(header http.Header, keys ...string) map[string][]string {
	r := make(map[string][]string)
	for k, v := range header {
		for _, key := range keys {
			if strings.EqualFold(k, key) {
				r[k] = v
			}
		}
	}

	return r
}

// makeHTTPRequest HTTP request helper
func (c *Client) makeHTTPRequestWithHeader(url string) ([]byte, http.Header, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}

	if c.httpRequestModifier != nil {
		c.httpRequestModifier(req)
	}

	resp, header, err := doReq(req, c.httpClient)
	if err != nil {
		return nil, nil, err
	}

	return resp, header, err
}

func (c *Client) makeHTTPRequest(url string) ([]byte, error) {
	resp, _, err := c.makeHTTPRequestWithHeader(url)

	return resp, err
}

func firstError(fst, snd error) error {
	if fst != nil {
		return fst
	}

	return snd
}
