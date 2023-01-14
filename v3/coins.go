package coingecko

import (
	"encoding/json"
	"fmt"
	"github.com/edward-yakop/go-gecko/format"
	"github.com/edward-yakop/go-gecko/v3/types"
	"net/url"
	"strings"
)

// CoinsList /coins/list
func (c *Client) CoinsList() (*types.CoinsList, error) {
	coinsListURL := fmt.Sprintf("%s/coins/list", c.baseURL)
	resp, header, err := c.makeHTTPRequest(coinsListURL)
	if err != nil {
		return nil, err
	}

	var data = &types.CoinsList{
		BaseResult: types.NewBaseResult(header),
		Coins:      []types.CoinsListItem{},
	}
	if err = json.Unmarshal(resp, &data.Coins); err != nil {
		return nil, err
	}

	return data, nil
}

type CoinsMarketParams struct {
	VsCurrency            string                        `json:"vs_currency"` // Required. The target currency of market data (usd, eur, jpy, etc.)
	CoinIDs               []string                      `json:"coin_ids"`    // The ids of the coin, crytocurrency symbols (base). refers to /coins/list.
	Category              string                        `json:"category"`    // filter by coin category. Refer to /coin/categories/list
	Order                 types.CoinsMarketOrder        `json:"order"`       // When blank will be set to "market_cap_desc"
	PageSize              int                           `json:"page_size"`   // Starts from 1 - 250, when invalid will be set to 100
	PageNo                int                           `json:"page_no"`     // Starts from 1, when < 1, will be set to 1
	Sparkline             bool                          `json:"sparkline"`   // Include sparkline 7 days data (eg. true, false)
	PriceChangePercentage []types.PriceChangePercentage `json:"price_change_percentage"`
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
		p.Order = types.CoinMarketOrderMarketCapDesc
	}
	params.Add("order", p.Order.String())

	// ids
	if len(p.CoinIDs) != 0 {
		idsParam := strings.Join(p.CoinIDs, ",")
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

// CoinsMarkets /coins/markets
func (c *Client) CoinsMarkets(params CoinsMarketParams) (*types.CoinsMarkets, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	coinsMarketsURL := fmt.Sprintf("%s/coins/markets?%s", c.baseURL, params.encodeQueryParams())
	resp, header, err := c.makeHTTPRequest(coinsMarketsURL)
	if err != nil {
		return nil, err
	}

	data := &types.CoinsMarkets{
		BaseResult: types.NewBaseResult(header),
		Markets:    []types.CoinsMarketItem{},
	}
	if err = json.Unmarshal(resp, &data.Markets); err != nil {
		return nil, err
	}

	return data, nil
}

type CoinsIDParams struct {
	CoinID        string `json:"coin_id"`        // CoinID (can be obtained from /coins)
	Localization  bool   `json:"localization"`   // Include all localized languages in response
	Tickers       bool   `json:"tickers"`        // Include tickers data. If true returns up to 100 entries. Use CoinsIDTickers
	MarketData    bool   `json:"market_data"`    // Include market data
	CommunityData bool   `json:"community_data"` // Include community data
	DeveloperData bool   `json:"developer_data"` // Include developer data
	Sparkline     bool   `json:"sparkline"`      // Include sparkline 7 days data
}

func (c CoinsIDParams) Validate() error {
	if c.CoinID == "" {
		return fmt.Errorf("id is required")
	}

	return nil
}

func (c CoinsIDParams) encodeNonIDQueryParams() string {
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

	coinsURL := fmt.Sprintf("%s/coins/%s?%s", c.baseURL, params.CoinID, params.encodeNonIDQueryParams())
	resp, header, err := c.makeHTTPRequest(coinsURL)
	if err != nil {
		return nil, err
	}

	data := &types.CoinsID{
		BaseResult: types.NewBaseResult(header),
	}
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}

	return data, nil
}

type CoinsIDTickersParam struct {
	CoinsID                string            `json:"coins_id"`                    // CoinID (can be obtained from /coins)
	ExchangeIDs            []string          `json:"exchange_ids"`                // filter results by exchange_ids ExchangesList
	ExchangeLogo           bool              `json:"exchange_logo"`               // flag to include exchange logo
	PageNo                 int               `json:"page_no"`                     // Page through results
	Order                  types.TickerOrder `json:"order"`                       // If not set default to trust_score_desc
	Show2PctOrderBookDepth bool              `json:"show_2_pct_order_book_depth"` // flag to show 2% order book depth
}

func (p CoinsIDTickersParam) Validate() error {
	if p.CoinsID == "" {
		return fmt.Errorf("CoinsID is required")
	}

	return nil
}

func (p CoinsIDTickersParam) encodeNonIDQueryParams() string {
	params := url.Values{}

	if len(p.ExchangeIDs) > 0 {
		params.Add("exchange_ids", strings.Join(p.ExchangeIDs, ","))
	}

	params.Add("include_exchange_logo", format.Bool2String(p.ExchangeLogo))

	if p.PageNo < 1 {
		p.PageNo = 1
	}
	params.Add("page", format.Int2String(p.PageNo))

	params.Add("order", p.Order.String())
	params.Add("depth", format.Bool2String(p.Show2PctOrderBookDepth))

	return params.Encode()
}

// CoinsIDTickers /coins/{id}/tickers
func (c *Client) CoinsIDTickers(params CoinsIDTickersParam) (*types.CoinsIDTickers, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	coinsIDURL := fmt.Sprintf("%s/coins/%s/tickers?%s", c.baseURL, params.CoinsID, params.encodeNonIDQueryParams())
	resp, header, err := c.makeHTTPRequest(coinsIDURL)
	if err != nil {
		return nil, err
	}

	data := &types.CoinsIDTickers{
		BasePageResult: types.NewBasePageResult(header, params.PageNo),
	}
	if err = json.Unmarshal(resp, data); err != nil {
		return nil, err
	}

	return data, nil
}

type CoinsIDHistoryParams struct {
	CoinID       string `json:"coin_id"`       // CoinID (can be obtained from /coins)
	SnapshotDate string `json:"snapshot_date"` // The date of data snapshot in dd-mm-yyyy eg. 30-12-2017
	Localization bool   `json:"localization"`  // Set to false to exclude localized languages in response
}

func (p CoinsIDHistoryParams) Validate() error {
	if p.CoinID == "" {
		return fmt.Errorf("CoinID is required")
	}

	if p.SnapshotDate == "" {
		return fmt.Errorf("SnapshotDate is required")
	}

	return nil
}

func (p CoinsIDHistoryParams) encodeNonIDQueryParams() string {
	params := url.Values{}

	params.Add("date", p.SnapshotDate)
	params.Add("localization", format.Bool2String(p.Localization))

	return params.Encode()
}

// CoinsIDHistory /coins/{id}/history?date={date}&localization=false
func (c *Client) CoinsIDHistory(params CoinsIDHistoryParams) (*types.CoinsIDHistory, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	coinsIDHistoryURL := fmt.Sprintf("%s/coins/%s/history?%s", c.baseURL, params.CoinID, params.encodeNonIDQueryParams())
	resp, header, err := c.makeHTTPRequest(coinsIDHistoryURL)
	if err != nil {
		return nil, err
	}

	data := &types.CoinsIDHistory{
		BaseResult: types.NewBaseResult(header),
	}
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}

	return data, nil
}

type CoinsIDMarketChartParams struct {
	CoinsID    string `json:"coins_id"`    // CoinID (can be obtained from /coins)
	VsCurrency string `json:"vs_currency"` // The target currency of market data (usd, eur, jpy, etc.)
	Days       string `json:"days"`        // Data up to number of days ago (eg. 1,14,30,max)
}

func (p CoinsIDMarketChartParams) Validate() error {
	if p.CoinsID == "" {
		return fmt.Errorf("CoinsID is required")
	}

	if p.VsCurrency == "" {
		return fmt.Errorf("VsCurrency is required")
	}

	if p.Days == "" {
		return fmt.Errorf("Days is required")
	}

	return nil
}

func (p CoinsIDMarketChartParams) encodeNonIDQueryParams() string {
	params := url.Values{}

	params.Add("vs_currency", p.VsCurrency)
	params.Add("days", p.Days)

	return params.Encode()
}

// CoinsIDMarketChart /coins/{id}/market_chart?vs_currency={usd, eur, jpy, etc.}&days={1,14,30,max}
func (c *Client) CoinsIDMarketChart(params CoinsIDMarketChartParams) (*types.CoinsIDMarketChart, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	coinsIDMarketChartURL := fmt.Sprintf("%s/coins/%s/market_chart?%s", c.baseURL, params.CoinsID, params.encodeNonIDQueryParams())
	resp, header, err := c.makeHTTPRequest(coinsIDMarketChartURL)
	if err != nil {
		return nil, err
	}

	data := &types.CoinsIDMarketChart{
		BaseResult: types.NewBaseResult(header),
	}
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}

	return data, nil
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
