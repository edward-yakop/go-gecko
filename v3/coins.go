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

// CoinsMarket /coins/market
func (c *Client) CoinsMarket(
	vsCurrency string, ids []string, order string, perPage int, page int, sparkline bool, priceChangePercentage []string,
) (*types.CoinsMarket, error) {
	if vsCurrency == "" {
		return nil, fmt.Errorf("vsCurrency is required")
	}

	params := url.Values{}
	// vs_currency
	params.Add("vs_currency", vsCurrency)

	// order
	if order == "" {
		order = types.OrderTypeObject.MarketCapDesc
	}

	params.Add("order", order)
	// ids
	if len(ids) != 0 {
		idsParam := strings.Join(ids, ",")
		params.Add("ids", idsParam)
	}

	// per_page
	if perPage <= 0 || perPage > 250 {
		perPage = 100
	}

	params.Add("per_page", format.Int2String(perPage))
	params.Add("page", format.Int2String(page))

	// sparkline
	params.Add("sparkline", format.Bool2String(sparkline))

	// price_change_percentage
	if len(priceChangePercentage) != 0 {
		priceChangePercentageParam := strings.Join(priceChangePercentage[:], ",")
		params.Add("price_change_percentage", priceChangePercentageParam)
	}

	coinsMarketsURL := fmt.Sprintf("%s/coins/markets?%s", c.baseURL, params.Encode())
	resp, err := c.makeHTTPRequest(coinsMarketsURL)
	if err != nil {
		return nil, err
	}

	var data *types.CoinsMarket
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}

	return data, nil
}

// CoinsID /coins/{id}
func (c *Client) CoinsID(id string, localization bool, tickers bool, marketData bool, communityData bool, developerData bool, sparkline bool) (*types.CoinsID, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	params := url.Values{}
	params.Add("localization", format.Bool2String(localization))
	params.Add("tickers", format.Bool2String(tickers))
	params.Add("market_data", format.Bool2String(marketData))
	params.Add("community_data", format.Bool2String(communityData))
	params.Add("developer_data", format.Bool2String(developerData))
	params.Add("sparkline", format.Bool2String(sparkline))

	coinsURL := fmt.Sprintf("%s/coins/%s?%s", c.baseURL, id, params.Encode())
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
