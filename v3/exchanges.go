package coingecko

import (
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
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
func (c *Client) Exchanges(params ExchangesParam) (map[string]types.Exchange, error) {
	exchangesURL := fmt.Sprintf("%s/exchanges?%s", c.baseURL, params.encodeQueryParams())

	resp, err := c.makeHTTPRequest(exchangesURL)
	if err != nil {
		return nil, err
	}

	r := make(map[string]types.Exchange)

	_, _ = jsonparser.ArrayEach(resp, func(ba []byte, _ jsonparser.ValueType, _ int, pErr error) {
		hasError := err != nil || pErr != nil
		if hasError {
			err = firstError(err, pErr)
			return
		}

		var e types.Exchange
		if err = json.Unmarshal(ba, &e); err == nil {
			r[e.Id] = e
		}
	})

	if err != nil {
		return nil, err
	}

	return r, nil
}

// ExchangesList https://api.coingecko.com/api/v3/exchanges/list
func (c *Client) ExchangesList() (map[string]string, error) {
	exchangesListURL := fmt.Sprintf("%s/exchanges/list", c.baseURL)

	resp, err := c.makeHTTPRequest(exchangesListURL)
	if err != nil {
		return nil, err
	}

	r := make(map[string]string)
	itemPaths := [][]string{
		{"id"},
		{"name"},
	}
	_, _ = jsonparser.ArrayEach(resp, func(ba []byte, _ jsonparser.ValueType, _ int, pErr error) {
		hasError := err != nil || pErr != nil
		if hasError {
			err = firstError(err, pErr)
			return
		}

		var id, name string
		jsonparser.EachKey(ba, func(idx int, ba []byte, _ jsonparser.ValueType, pErr error) {
			hasError = err != nil || pErr != nil
			if hasError {
				err = firstError(err, pErr)
				return
			}

			switch idx {
			case 0: // id
				id = string(ba)
			case 1: // name
				name = string(ba)
			}
		}, itemPaths...)
		if err == nil {
			r[id] = name
		}
	})

	if err != nil {
		return nil, err
	}

	return r, nil
}
