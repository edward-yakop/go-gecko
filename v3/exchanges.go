package coingecko

import (
	"encoding/json"
	"fmt"
	"github.com/edward-yakop/go-gecko/format"
	"github.com/edward-yakop/go-gecko/v3/types"
	"net/url"
)

type ExchangesParam struct {
	PageSize int
	PageNo   int
}

func (p ExchangesParam) encodeQueryParams() string {
	urlValues := url.Values{}

	if p.PageSize < 1 || p.PageSize > 250 {
		p.PageSize = 100
	}
	urlValues.Add("per_page", format.Int2String(p.PageSize))

	if p.PageNo < 1 {
		p.PageNo = 1
	}
	urlValues.Add("page", format.Int2String(p.PageNo))

	return urlValues.Encode()
}

// Exchanges https://api.coingecko.com/api/v3/exchanges
func (c *Client) Exchanges(params ExchangesParam) ([]types.Exchange, error) {
	exchangesURL := fmt.Sprintf("%s/exchanges?%s", c.baseURL, params.encodeQueryParams())

	resp, err := c.makeHTTPRequest(exchangesURL)
	if err != nil {
		return nil, err
	}

	var data []types.Exchange
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}

	return data, nil
}
