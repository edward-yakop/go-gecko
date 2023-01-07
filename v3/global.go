package coingecko

import (
	"encoding/json"
	"fmt"
	"github.com/edward-yakop/go-gecko/v3/types"
)

// Global https://api.coingecko.com/api/v3/global
func (c *Client) Global() (*types.Global, error) {
	globalURL := fmt.Sprintf("%s/global", c.baseURL)
	resp, err := c.makeHTTPRequest(globalURL)
	if err != nil {
		return nil, err
	}

	var data *types.GlobalResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}

	return &data.Data, nil
}
