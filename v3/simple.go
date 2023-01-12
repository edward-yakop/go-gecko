package coingecko

import (
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/edward-yakop/go-gecko/format"
	"github.com/edward-yakop/go-gecko/v3/types"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type SimplePriceParams struct {
	CoinIDs              []string // CoinIds refers to CoinsList
	VsCurrencies         []string // VsCurrencies refers to SimpleSupportedVSCurrencies
	IncludeMarketCap     bool     // Sets to true to include market cap. Default false.
	Include24HrVolume    bool     // Sets to true to include 24 Hr volume. Default false.
	Include24HrChange    bool     // Sets to true to include 24 Hr change. Default false.
	IncludeLastUpdatedAt bool     // Sets to true to include last updated at. Default false.
	Precision            string   // valid values: empty string, "full", an int [0,18]
}

func (params SimplePriceParams) Valid() error {
	if len(params.CoinIDs) < 1 {
		return fmt.Errorf("CoinIDs is required and must contain at least 1 item")
	}

	if len(params.VsCurrencies) < 1 {
		return fmt.Errorf("VsCurrencies is required and must contain at least 1 item")
	}

	if p := params.toInt(params.Precision); params.Precision != "" && params.Precision != "full" && (p < 0 || p > 18) {
		return fmt.Errorf("Precision must either be empty string, or \"full\", or an int [0,18]")
	}

	return nil
}

func (params SimplePriceParams) toInt(v string) int {
	if asInt, err := strconv.Atoi(v); err == nil {
		return asInt
	}

	return -1
}

func (params SimplePriceParams) encode() string {
	values := url.Values{}

	values.Add("ids", strings.Join(params.CoinIDs, ","))
	values.Add("vs_currencies", strings.Join(params.VsCurrencies, ","))
	if params.IncludeMarketCap {
		values.Add("include_market_cap", format.Bool2String(params.IncludeMarketCap))
	}
	if params.Include24HrVolume {
		values.Add("include_24hr_vol", format.Bool2String(params.Include24HrVolume))
	}
	if params.Include24HrChange {
		values.Add("include_24hr_change", format.Bool2String(params.Include24HrChange))
	}
	if params.IncludeLastUpdatedAt {
		values.Add("include_last_updated_at", format.Bool2String(params.IncludeLastUpdatedAt))
	}
	if params.Precision != "" {
		values.Add("precision", params.Precision)
	}

	return values.Encode()
}

// SimplePrice /simple/price Multiple ID and Currency (ids, vs_currencies)
func (c *Client) SimplePrice(params SimplePriceParams) (*types.SimplePrice, error) {
	if err := params.Valid(); err != nil {
		return nil, err
	}

	simplePriceURL := fmt.Sprintf("%s/simple/price?%s", c.baseURL, params.encode())
	resp, header, err := c.makeHTTPRequestWithHeader(simplePriceURL)
	if err != nil {
		return nil, err
	}

	coins := make(map[string]*types.SimplePriceItem)
	r := &types.SimplePrice{
		BaseResult: types.NewBaseResult(header),
		Coins:      coins,
	}

	err = jsonparser.ObjectEach(resp, func(coinIDBA []byte, ba []byte, _ jsonparser.ValueType, offset int) error {
		coinID := string(coinIDBA)
		item, iErr := c.parseSimplePriceItem(coinID, ba)
		if iErr == nil {
			coins[coinID] = item
		}

		return iErr
	})

	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) parseSimplePriceItem(coinID string, ba []byte) (*types.SimplePriceItem, error) {
	r := &types.SimplePriceItem{
		Currencies: make(map[string]*types.SimplePriceCurrencyItem),
	}

	err := jsonparser.ObjectEach(ba, func(keyBA []byte, ba []byte, dataType jsonparser.ValueType, offset int) error {
		key := string(keyBA)
		if key == "last_updated_at" {
			timeAsInt, pErr := jsonparser.ParseInt(ba)
			if pErr == nil {
				r.LastUpdatedAt = time.Unix(timeAsInt, 0)
			} else {
				pErr = fmt.Errorf("error parsing %s.last_updated_at = %s: %v", coinID, string(ba), pErr)
			}

			return pErr
		}

		if underscoreIndex := strings.Index(key, "_"); underscoreIndex != -1 {
			currCode := key[:underscoreIndex]
			currItem := r.Currencies[currCode]
			if currItem == nil {
				currItem = &types.SimplePriceCurrencyItem{}
				r.Currencies[currCode] = currItem
			}

			suffix := key[underscoreIndex+1:]
			v, pErr := jsonparser.ParseFloat(ba)
			if pErr == nil {
				switch suffix {
				case "market_cap":
					currItem.MarketCap = &v
				case "24h_vol":
					currItem.Volume24H = &v
				case "24h_change":
					currItem.ChangePercentage24H = &v
				}
			} else {
				pErr = fmt.Errorf("error parsing %s.%s = %s: %v", coinID, key, string(ba), pErr)
			}

			return pErr
		} else {
			currItem := r.Currencies[key]
			if currItem == nil {
				currItem = &types.SimplePriceCurrencyItem{}
				r.Currencies[key] = currItem
			}
			v, pErr := jsonparser.ParseFloat(ba)
			if pErr == nil {
				currItem.Price = v
			} else {
				pErr = fmt.Errorf("error parsing %s.%s.Price = %s: %v", coinID, key, string(ba), pErr)
			}

			return pErr
		}
	})

	if err != nil {
		return nil, err
	}

	return r, err
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
