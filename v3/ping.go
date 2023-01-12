package coingecko

import (
	"encoding/json"
	"fmt"
	"github.com/edward-yakop/go-gecko/v3/types"
)

// Ping /ping endpoint
func (c *Client) Ping() (*types.Ping, error) {
	pingURL := fmt.Sprintf("%s/ping", c.baseURL)

	resp, header, err := c.makeHTTPRequest(pingURL)
	if err != nil {
		return nil, err
	}

	data := &types.Ping{
		BaseResult: types.NewBaseResult(header),
	}
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}

	return data, nil
}
