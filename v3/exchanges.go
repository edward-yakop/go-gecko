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
	PageSize int `json:"page_size"` // Total results per page, between 1-250. When invalid, default to 100.
	PageNo   int `json:"page_no"`   // Page through results. Valid values bigger than 1. When invalid, default to 1.
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

	resp, header, err := c.makeHTTPRequest(exchangesURL)
	if err != nil {
		return nil, err
	}

	m := make(map[string]types.Exchange)
	r := &types.Exchanges{
		BasePageResult: types.NewBasePageResult(header, params.PageNo),
		Exchanges:      m,
	}

	_, _ = jsonparser.ArrayEach(resp, func(ba []byte, _ jsonparser.ValueType, _ int, pErr error) {
		hasError := err != nil || pErr != nil
		if hasError {
			err = firstError(err, pErr)
			return
		}

		var e types.Exchange
		if err = json.Unmarshal(ba, &e); err == nil {
			m[e.ID] = e
		}
	})

	if err != nil {
		return nil, err
	}

	return r, nil
}

// ExchangesList https://api.coingecko.com/api/v3/exchanges/list
func (c *Client) ExchangesList() (*types.ExchangesList, error) {
	exchangesListURL := fmt.Sprintf("%s/exchanges/list", c.baseURL)

	resp, header, err := c.makeHTTPRequest(exchangesListURL)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	r := &types.ExchangesList{
		BaseResult: types.NewBaseResult(header),
		Exchanges:  m,
	}

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
			m[id] = name
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

	resp, header, err := c.makeHTTPRequest(exchangesListURL)
	if err != nil {
		return nil, err
	}

	data := &types.ExchangeDetail{
		BaseResult: types.NewBaseResult(header),
	}
	if err = json.Unmarshal(resp, data); err != nil {
		return nil, err
	}

	return data, nil
}

type ExchangesIDTickersParams struct {
	ExchangeID             string            `json:"exchange_id"`                 // ExchangeID, can be obtained from ExchangesList. Required.
	CoinIds                []string          `json:"coin_ids"`                    // filter tickers by CoinIds (ref: CoinsList). Optional.
	ExchangeLogo           bool              `json:"exchange_logo"`               // Include ExchangeLogo
	PageNo                 int               `json:"page_no"`                     // Page through results. If < 1, default to 1.
	Show2PctOrderBookDepth bool              `json:"show_2_pct_order_book_depth"` // flag to show 2% orderbook depth i.e., CostToMoveUpUsd and CostToMoveDownUsd
	Order                  types.TickerOrder `json:"order"`                       // Default to TickerOrderTrustScoreDesc
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

	params.Add("include_exchange_logo", format.Bool2String(p.ExchangeLogo))

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

	resp, header, err := c.makeHTTPRequest(exchangesListURL)
	if err != nil {
		return nil, err
	}

	data := &types.ExchangeTickers{
		BasePageResult: types.NewBasePageResult(header, params.PageNo),
	}
	if err = json.Unmarshal(resp, data); err != nil {
		return nil, err
	}

	return data, nil
}
