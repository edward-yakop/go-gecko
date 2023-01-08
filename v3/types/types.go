package types

// Ping https://api.coingecko.com/api/v3/ping
type Ping struct {
	GeckoSays string `json:"gecko_says"`
}

// SimpleSinglePrice https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd
type SimpleSinglePrice struct {
	ID          string
	Currency    string
	MarketPrice float64
}

// SimpleSupportedVSCurrencies https://api.coingecko.com/api/v3/simple/supported_vs_currencies
type SimpleSupportedVSCurrencies []string

// CoinList https://api.coingecko.com/api/v3/coins/list
type CoinList []CoinsListItem

// CoinsMarket https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=100&page=1&sparkline=false
type CoinsMarket []CoinsMarketItem

// CoinsID https://api.coingecko.com/api/v3/coins/bitcoin
type CoinsID struct {
	coinBaseStruct
	BlockTimeInMin      int32               `json:"block_time_in_minutes"`
	HashingAlgorithm    string              `json:"hashing_algorithm"`
	Categories          []string            `json:"categories"`
	Localization        LocalizationItem    `json:"localization"`
	Description         DescriptionItem     `json:"description"`
	Links               *LinksItem          `json:"links"`
	Image               ImageItem           `json:"image"`
	CountryOrigin       string              `json:"country_origin"`
	GenesisDate         string              `json:"genesis_date"`
	MarketCapRank       uint16              `json:"market_cap_rank"`
	CoinGeckoRank       uint16              `json:"coingecko_rank"`
	CoinGeckoScore      float64             `json:"coingecko_score"`
	DeveloperScore      float64             `json:"developer_score"`
	CommunityScore      float64             `json:"community_score"`
	LiquidityScore      float64             `json:"liquidity_score"`
	PublicInterestScore float64             `json:"public_interest_score"`
	MarketData          *MarketDataItem     `json:"market_data"`
	CommunityData       *CommunityDataItem  `json:"community_data"`
	DeveloperData       *DeveloperDataItem  `json:"developer_data"`
	PublicInterestStats *PublicInterestItem `json:"public_interest_stats"`
	StatusUpdates       []StatusUpdateItem  `json:"status_updates"`
	LastUpdated         string              `json:"last_updated"`
	Tickers             []TickerItem        `json:"tickers"`
}

// CoinsIDTickers https://api.coingecko.com/api/v3/coins/steem/tickers?page=1
type CoinsIDTickers struct {
	Name    string       `json:"name"`
	Tickers []TickerItem `json:"tickers"`
}

// CoinsIDHistory https://api.coingecko.com/api/v3/coins/steem/history?date=30-12-2018
type CoinsIDHistory struct {
	coinBaseStruct
	Localization   LocalizationItem             `json:"localization"`
	Image          ImageItem                    `json:"image"`
	MarketData     *CoinIDHistoryMarketDataItem `json:"market_data"`
	CommunityData  *CommunityDataItem           `json:"community_data"`
	DeveloperData  *DeveloperDataItem           `json:"developer_data"`
	PublicInterest *PublicInterestItem          `json:"public_interest_stats"`
}

type CoinIDHistoryMarketDataItem struct {
	CurrentPrice map[string]float64 `json:"current_price"`
	MarketCap    map[string]float64 `json:"market_cap"`
	TotalVolume  map[string]float64 `json:"total_volume"`
}

// CoinsIDMarketChart https://api.coingecko.com/api/v3/coins/bitcoin/market_chart?vs_currency=usd&days=1
type CoinsIDMarketChart struct {
	Prices       []ChartItem `json:"prices"`
	MarketCaps   []ChartItem `json:"market_caps"`
	TotalVolumes []ChartItem `json:"total_volumes"`
}

// CoinsIDStatusUpdates

// CoinsIDContractAddress https://api.coingecko.com/api/v3/coins/{id}/contract/{contract_address}
// type CoinsIDContractAddress struct {
// 	ID                  string           `json:"id"`
// 	Symbol              string           `json:"symbol"`
// 	Name                string           `json:"name"`
// 	BlockTimeInMin      uint16           `json:"block_time_in_minutes"`
// 	Categories          []string         `json:"categories"`
// 	Localization        LocalizationItem `json:"localization"`
// 	Description         DescriptionItem  `json:"description"`
// 	Links               LinksItem        `json:"links"`
// 	Image               ImageItem        `json:"image"`
// 	CountryOrigin       string           `json:"country_origin"`
// 	GenesisDate         string           `json:"genesis_date"`
// 	ContractAddress     string           `json:"contract_address"`
// 	MarketCapRank       uint16           `json:"market_cap_rank"`
// 	CoinGeckoRank       uint16           `json:"coingecko_rank"`
// 	CoinGeckoScore      float64          `json:"coingecko_score"`
// 	DeveloperScore      float64          `json:"developer_score"`
// 	CommunityScore      float64          `json:"community_score"`
// 	LiquidityScore      float64          `json:"liquidity_score"`
// 	PublicInterestScore float64          `json:"public_interest_score"`
// 	MarketData          `json:"market_data"`
// }

// Exchange https://api.coingecko.com/api/v3/exchanges
type Exchange struct {
	Id                          string  `json:"id"`
	Name                        string  `json:"name"`
	YearEstablished             *int    `json:"year_established"`
	Country                     *string `json:"country"`
	Description                 *string `json:"description"`
	Url                         string  `json:"url"`
	Image                       string  `json:"image"`
	HasTradingIncentive         *bool   `json:"has_trading_incentive"`
	TrustScore                  *int    `json:"trust_score"`
	TrustScoreRank              *int    `json:"trust_score_rank"`
	TradeVolume24HBtc           float64 `json:"trade_volume_24h_btc"`
	TradeVolume24HBtcNormalized float64 `json:"trade_volume_24h_btc_normalized"`
}

// ExchangeDetail https://api.coingecko.com/api/v3/exchanges/{id}
type ExchangeDetail struct {
	Name                        string             `json:"name"`
	YearEstablished             int                `json:"year_established"`
	Country                     string             `json:"country"`
	Description                 string             `json:"description"`
	Url                         string             `json:"url"`
	Image                       string             `json:"image"`
	FacebookUrl                 string             `json:"facebook_url"`
	RedditUrl                   string             `json:"reddit_url"`
	TelegramUrl                 string             `json:"telegram_url"`
	SlackUrl                    string             `json:"slack_url"`
	OtherUrl1                   string             `json:"other_url_1"`
	OtherUrl2                   string             `json:"other_url_2"`
	TwitterHandle               string             `json:"twitter_handle"`
	HasTradingIncentive         bool               `json:"has_trading_incentive"`
	Centralized                 bool               `json:"centralized"`
	PublicNotice                string             `json:"public_notice"`
	AlertNotice                 string             `json:"alert_notice"`
	TrustScore                  int                `json:"trust_score"`
	TrustScoreRank              int                `json:"trust_score_rank"`
	TradeVolume24HBtc           float64            `json:"trade_volume_24h_btc"`
	TradeVolume24HBtcNormalized float64            `json:"trade_volume_24h_btc_normalized"`
	Tickers                     []TickerItem       `json:"tickers"`
	StatusUpdates               []StatusUpdateItem `json:"status_updates"`
}

// ExchangeRatesResponse https://api.coingecko.com/api/v3/exchange_rates
type ExchangeRatesResponse struct {
	Rates ExchangeRates `json:"rates"`
}

// GlobalResponse https://api.coingecko.com/api/v3/global
type GlobalResponse struct {
	Data Global `json:"data"`
}
