package coingecko

import (
	"encoding/json"
	"fmt"
	"github.com/superoo7/go-gecko/format"
	"github.com/superoo7/go-gecko/v3/types"
	"net/url"
	"strings"
)

// CoinsList /coins/list
func (c *Client) CoinsList() (*types.CoinList, error) {
	coinsListURL := fmt.Sprintf("%s/coins/list", c.baseURL)
	resp, err := c.makeHTTPRequest(coinsListURL)
	if err != nil {
		return nil, err
	}

	var data *types.CoinList
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}

	return data, nil
}

type CoinsMarketParams struct {
	VsCurrency            string               // Required. The target currency of market data (usd, eur, jpy, etc.)
	CoinIds               []string             // The ids of the coin, comma separated crytocurrency symbols (base). refers to /coins/list.
	Category              string               // filter by coin category. Refer to /coin/categories/list
	Order                 types.CoinsOrderType // When blank will be set to "market_cap_desc"
	PageSize              int                  // Starts from 1 - 250, when invalid will be set to 100
	PageNo                int                  // Starts from 1, when < 1, will be set to 1
	Sparkline             bool                 // Include sparkline 7 days data (eg. true, false)
	PriceChangePercentage []types.PriceChangePercentage
}

func (p CoinsMarketParams) Validate() error {
	if p.VsCurrency == "" {
		return fmt.Errorf("VsCurrency is required")
	}

	return nil
}

func (p CoinsMarketParams) encodeQueryParams() string {
	params := url.Values{}
	// vs_currency
	params.Add("vs_currency", p.VsCurrency)

	// order
	if p.Order < 0 || p.Order > 5 {
		p.Order = types.CoinsOrderTypeMarketCapDesc
	}
	params.Add("order", p.Order.String())

	// ids
	if len(p.CoinIds) != 0 {
		idsParam := strings.Join(p.CoinIds, ",")
		params.Add("ids", idsParam)
	}

	// per_page
	if p.PageSize <= 0 || p.PageSize > 250 {
		p.PageSize = 100
	}
	params.Add("per_page", format.Int2String(p.PageSize))

	// PageNo
	if p.PageNo <= 1 {
		p.PageNo = 1
	}
	params.Add("page", format.Int2String(p.PageNo))

	// sparkline
	if p.Sparkline {
		params.Add("sparkline", format.Bool2String(p.Sparkline))
	}

	// price_change_percentage
	pcpLen := len(p.PriceChangePercentage)
	if pcpLen != 0 {
		sb := strings.Builder{}
		lastPcpIndex := pcpLen - 1
		for i, pcp := range p.PriceChangePercentage {
			sb.WriteString(pcp.String())
			if i != lastPcpIndex {
				sb.WriteString(",")
			}
		}
		params.Add("price_change_percentage", sb.String())
	}

	return params.Encode()
}

// CoinsMarket /coins/market
func (c *Client) CoinsMarket(params CoinsMarketParams) (types.CoinsMarket, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	coinsMarketsURL := fmt.Sprintf("%s/coins/markets?%s", c.baseURL, params.encodeQueryParams())
	resp, err := c.makeHTTPRequest(coinsMarketsURL)
	if err != nil {
		return nil, err
	}

	var data types.CoinsMarket
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}

	return data, nil
}

type CoinsIDParams struct {
	Id            string
	Localization  bool
	Tickers       bool
	MarketData    bool
	CommunityData bool
	DeveloperData bool
	Sparkline     bool
}

func (c CoinsIDParams) Validate() error {
	if c.Id == "" {
		return fmt.Errorf("id is required")
	}

	return nil
}

func (c CoinsIDParams) encodeNonIdParams() string {
	params := url.Values{}

	params.Add("localization", format.Bool2String(c.Localization))
	params.Add("tickers", format.Bool2String(c.Tickers))
	params.Add("market_data", format.Bool2String(c.MarketData))
	params.Add("community_data", format.Bool2String(c.CommunityData))
	params.Add("developer_data", format.Bool2String(c.DeveloperData))
	params.Add("sparkline", format.Bool2String(c.Sparkline))

	return params.Encode()
}

// CoinsID /coins/{id}
func (c *Client) CoinsID(params CoinsIDParams) (*types.CoinsID, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	coinsURL := fmt.Sprintf("%s/coins/%s?%s", c.baseURL, params.Id, params.encodeNonIdParams())
	resp, err := c.makeHTTPRequest(coinsURL)
	if err != nil {
		return nil, err
	}

	var data *types.CoinsID
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}

	return data, nil
}

// CoinsIDTickers /coins/{id}/tickers
func (c *Client) CoinsIDTickers(id string, page int) (*types.CoinsIDTickers, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	params := url.Values{}
	if page > 0 {
		params.Add("page", format.Int2String(page))
	}

	coinsIDURL := fmt.Sprintf("%s/coins/%s/tickers?%s", c.baseURL, id, params.Encode())
	resp, err := c.makeHTTPRequest(coinsIDURL)
	if err != nil {
		return nil, err
	}

	var data *types.CoinsIDTickers
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}

	return data, nil
}

// CoinsIDHistory /coins/{id}/history?date={date}&localization=false
func (c *Client) CoinsIDHistory(id string, date string, localization bool) (*types.CoinsIDHistory, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	if date == "" {
		return nil, fmt.Errorf("date is required")
	}

	params := url.Values{}
	params.Add("date", date)
	params.Add("localization", format.Bool2String(localization))

	coinsIDHistoryURL := fmt.Sprintf("%s/coins/%s/history?%s", c.baseURL, id, params.Encode())
	resp, err := c.makeHTTPRequest(coinsIDHistoryURL)
	if err != nil {
		return nil, err
	}

	var data *types.CoinsIDHistory
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}

	return data, nil
}

// CoinsIDMarketChart /coins/{id}/market_chart?vs_currency={usd, eur, jpy, etc.}&days={1,14,30,max}
func (c *Client) CoinsIDMarketChart(id string, vsCurrency string, days string) (*types.CoinsIDMarketChart, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	if vsCurrency == "" {
		return nil, fmt.Errorf("vsCurrency is required")
	}

	if days == "" {
		return nil, fmt.Errorf("days is required")
	}

	params := url.Values{}
	params.Add("vs_currency", vsCurrency)
	params.Add("days", days)

	coinsIDMarketChartURL := fmt.Sprintf("%s/coins/%s/market_chart?%s", c.baseURL, id, params.Encode())
	resp, err := c.makeHTTPRequest(coinsIDMarketChartURL)
	if err != nil {
		return nil, err
	}

	data := types.CoinsIDMarketChart{}
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

// CoinsIDStatusUpdates

// CoinsIDContractAddress https://api.coingecko.com/api/v3/coins/{id}/contract/{contract_address}
// func CoinsIDContractAddress(id string, address string) (nil, error) {
// 	url := fmt.Sprintf("%s/coins/%s/contract/%s", c.baseURL, id, address)
// 	resp, err := request.makeHTTPRequest(url)
// 	if err != nil {
// 		return nil, err
// 	}
// }
