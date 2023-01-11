package coingecko

import (
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/edward-yakop/go-gecko/format"
	"github.com/edward-yakop/go-gecko/v3/types"
	"net/url"
	"strings"
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
func (c *Client) Exchanges(params ExchangesParam) (*types.Exchanges, error) {
	exchangesURL := fmt.Sprintf("%s/exchanges?%s", c.baseURL, params.encodeQueryParams())

	resp, header, err := c.makeHTTPRequestWithHeader(exchangesURL)
	if err != nil {
		return nil, err
	}

	m := make(map[string]types.Exchange)
	r := &types.Exchanges{
		BasePageResult: types.NewBasePageResult(header),
		Entries:        m,
	}

	_, _ = jsonparser.ArrayEach(resp, func(ba []byte, _ jsonparser.ValueType, _ int, pErr error) {
		hasError := err != nil || pErr != nil
		if hasError {
			err = firstError(err, pErr)
			return
		}

		var e types.Exchange
		if err = json.Unmarshal(ba, &e); err == nil {
			m[e.Id] = e
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

func (c *Client) ExchangesID(exchangeID string) (*types.ExchangeDetail, error) {
	if exchangeID == "" {
		return nil, fmt.Errorf("exchangeID is required")
	}

	exchangesListURL := fmt.Sprintf("%s/exchanges/%s", c.baseURL, exchangeID)

	resp, err := c.makeHTTPRequest(exchangesListURL)
	if err != nil {
		return nil, err
	}

	data := &types.ExchangeDetail{}
	if err = json.Unmarshal(resp, data); err != nil {
		return nil, err
	}

	return data, nil
}

type ExchangesIDTickersParams struct {
	ExchangeID             string   // ExchangeID, can be obtained from ExchangesList.
	CoinIds                []string // filter tickers by CoinIds (ref: CoinList)
	IncludeExchangeLogo    bool     // Flag to show ExchangeLogo
	PageNo                 int      // Page through results
	Show2PctOrderBookDepth bool     // flag to show 2% orderbook depth i.e., CostToMoveUpUsd and CostToMoveDownUsd
	Order                  types.TickerOrder
}

func (p ExchangesIDTickersParams) Valid() error {
	if p.ExchangeID == "" {
		return fmt.Errorf("ExchangeID is required")
	}

	return nil
}

func (p ExchangesIDTickersParams) encodeQueryParamsWithoutExchangeID() string {
	params := url.Values{}

	if len(p.CoinIds) > 0 {
		params.Add("coin_ids", strings.Join(p.CoinIds, ","))
	}

	params.Add("include_exchange_logo", format.Bool2String(p.IncludeExchangeLogo))

	if p.PageNo < 1 {
		p.PageNo = 1
	}
	params.Add("page", format.Int2String(p.PageNo))

	if p.Show2PctOrderBookDepth {
		params.Add("depth", format.Bool2String(p.Show2PctOrderBookDepth))
	}

	params.Add("order", p.Order.String())

	return params.Encode()

}

func (c *Client) ExchangesIDTickers(params ExchangesIDTickersParams) (*types.ExchangeTickers, error) {
	if err := params.Valid(); err != nil {
		return nil, err
	}

	exchangesListURL := fmt.Sprintf("%s/exchanges/%s/tickers?%s", c.baseURL, params.ExchangeID, params.encodeQueryParamsWithoutExchangeID())

	resp, err := c.makeHTTPRequest(exchangesListURL)
	if err != nil {
		return nil, err
	}

	data := &types.ExchangeTickers{}
	if err = json.Unmarshal(resp, data); err != nil {
		return nil, err
	}

	return data, nil
}
