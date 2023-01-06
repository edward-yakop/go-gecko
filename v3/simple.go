package coingecko

import (
	"encoding/json"
	"fmt"
	"github.com/superoo7/go-gecko/v3/types"
	"net/url"
	"strings"
)

// SimpleSinglePrice /simple/price  Single ID and Currency (ids, vs_currency)
func (c *Client) SimpleSinglePrice(id string, vsCurrency string) (*types.SimpleSinglePrice, error) {
	idParam := []string{strings.ToLower(id)}
	vcParam := []string{strings.ToLower(vsCurrency)}

	t, err := c.SimplePrice(idParam, vcParam)
	if err != nil {
		return nil, err
	}

	curr := (*t)[id]
	if len(curr) == 0 {
		return nil, fmt.Errorf("id or vsCurrency not existed")
	}

	return &types.SimpleSinglePrice{
		ID:          id,
		Currency:    vsCurrency,
		MarketPrice: curr[vsCurrency],
	}, nil
}

// SimplePrice /simple/price Multiple ID and Currency (ids, vs_currencies)
func (c *Client) SimplePrice(ids []string, vsCurrencies []string) (*map[string]map[string]float64, error) {
	params := url.Values{}
	idsParam := strings.Join(ids, ",")
	vsCurrenciesParam := strings.Join(vsCurrencies, ",")

	params.Add("ids", idsParam)
	params.Add("vs_currencies", vsCurrenciesParam)

	simplePriceURL := fmt.Sprintf("%s/simple/price?%s", c.baseURL, params.Encode())
	resp, err := c.makeHTTPRequest(simplePriceURL)
	if err != nil {
		return nil, err
	}

	t := make(map[string]map[string]float64)
	if err = json.Unmarshal(resp, &t); err != nil {
		return nil, err
	}

	return &t, nil
}

// SimpleSupportedVSCurrencies /simple/supported_vs_currencies
func (c *Client) SimpleSupportedVSCurrencies() (*types.SimpleSupportedVSCurrencies, error) {
	simpleURL := fmt.Sprintf("%s/simple/supported_vs_currencies", c.baseURL)
	resp, err := c.makeHTTPRequest(simpleURL)
	if err != nil {
		return nil, err
	}

	var data *types.SimpleSupportedVSCurrencies
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}

	return data, nil
}
