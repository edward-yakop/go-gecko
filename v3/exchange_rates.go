package coingecko

import (
	"encoding/json"
	"fmt"
	"github.com/edward-yakop/go-gecko/v3/types"
)

// ExchangeRates https://api.coingecko.com/api/v3/exchange_rates
func (c *Client) ExchangeRates() (*types.ExchangeRates, error) {
	exchangeRatesURL := fmt.Sprintf("%s/exchange_rates", c.baseURL)

	resp, header, err := c.makeHTTPRequest(exchangeRatesURL)
	if err != nil {
		return nil, err
	}

	data := &types.ExchangeRates{
		BaseResult: types.NewBaseResult(header),
		Rates:      make(map[string]types.ExchangeRatesItem),
	}
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}

	return data, nil
}
