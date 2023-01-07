package coingecko

import (
	"encoding/json"
	"fmt"
	"github.com/edward-yakop/go-gecko/v3/types"
)

// Ping /ping endpoint
func (c *Client) Ping() (*types.Ping, error) {
	pingURL := fmt.Sprintf("%s/ping", c.baseURL)

	resp, err := c.makeHTTPRequest(pingURL)
	if err != nil {
		return nil, err
	}

	var data *types.Ping
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}

	return data, nil
}
