package types

// AllCurrencies map all currencies (USD, BTC) to float64
type AllCurrencies map[string]float64

// CoinsListItem item in CoinList
type CoinsListItem struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

// EventCountryItem item in EventsCountries
type EventCountryItem struct {
	Country string `json:"country"`
	Code    string `json:"code"`
}

// ExchangeRatesItem item in ExchangeRate
type ExchangeRatesItem map[string]ExchangeRatesItemStruct

// ExchangeRatesItemStruct struct in ExchangeRateItem
type ExchangeRatesItemStruct struct {
	Name  string  `json:"name"`
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
	Type  string  `json:"type"`
}

// Global for data of /global
type Global struct {
	ActiveCryptocurrencies          uint16        `json:"active_cryptocurrencies"`
	UpcomingICOs                    uint16        `json:"upcoming_icos"`
	EndedICOs                       uint16        `json:"ended_icos"`
	Markets                         uint16        `json:"markets"`
	MarketCapChangePercentage24hUSD float32       `json:"market_cap_change_percentage_24h_usd"`
	TotalMarketCap                  AllCurrencies `json:"total_market_cap"`
	TotalVolume                     AllCurrencies `json:"total_volume"`
	MarketCapPercentage             AllCurrencies `json:"market_cap_percentage"`
	UpdatedAt                       int64         `json:"updated_at"`
}