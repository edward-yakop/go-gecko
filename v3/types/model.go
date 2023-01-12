package types

import (
	"encoding/json"
	"time"
)

// OrderType

type CoinsMarketOrder int

const (
	CoinMarketOrderMarketCapDesc = iota
	CoinMarketOrderMarketCapAsc
	CoinMarketOrderGeckoDesc
	CoinMarketOrderGeckoAsc
	CoinMarketOrderVolumeAsc
	CoinMarketOrderVolumeDesc
)

func (cmo CoinsMarketOrder) String() string {
	return []string{
		"market_cap_desc",
		"market_cap_asc",
		"gecko_desc",
		"gecko_asc",
		"volume_asc",
		"volume_desc",
	}[cmo]
}

// PriceChangePercentage

type PriceChangePercentage int

const (
	PriceChangePercentage1H = iota
	PriceChangePercentage24H
	PriceChangePercentage7D
	PriceChangePercentage14D
	PriceChangePercentage30D
	PriceChangePercentage200D
	PriceChangePercentage1Y
)

func (pcp PriceChangePercentage) String() string {
	return []string{
		"1h",
		"24h",
		"7d",
		"14d",
		"30d",
		"200d",
		"1y",
	}[pcp]
}

type TickerOrder int

const (
	TickerOrderTrustScoreDesc = iota
	TickerOrderTrustScoreAsc
	TickerOrderVolumeDesc
)

func (cto TickerOrder) String() string {
	return []string{
		"trust_score_desc",
		"trust_score_asc",
		"volume_desc",
	}[cto]
}

// SHARED
// coinBaseStruct [private]
type coinBaseStruct struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

// AllCurrencies map all currencies (USD, BTC) to float64
type AllCurrencies map[string]float64

// LocalizationItem map all locale (en, zh) into respective string
type LocalizationItem map[string]string

// TYPES

// DescriptionItem map all description (in locale) into respective string
type DescriptionItem map[string]string

// LinksItem map all links
type LinksItem map[string]interface{}

// ChartItem

type ChartItem struct {
	Time  time.Time
	Value float64
}

func (ci *ChartItem) UnmarshalJSON(data []byte) error {
	var content [2]float64
	if err := json.Unmarshal(data, &content); err != nil {
		return err
	}

	ci.Time = time.Unix(0, int64(content[0])*int64(time.Millisecond))
	ci.Value = content[1]

	return nil
}

// MarketDataItem map all market data item
type MarketDataItem struct {
	CurrentPrice                           AllCurrencies     `json:"current_price"`
	ROI                                    *ROIItem          `json:"roi"`
	ATH                                    AllCurrencies     `json:"ath"`
	ATHChangePercentage                    AllCurrencies     `json:"ath_change_percentage"`
	ATHDate                                map[string]string `json:"ath_date"`
	ATL                                    AllCurrencies     `json:"atl"`
	ATLChangePercentage                    AllCurrencies     `json:"atl_change_percentage"`
	ATLDate                                map[string]string `json:"atl_date"`
	MarketCap                              AllCurrencies     `json:"market_cap"`
	MarketCapRank                          uint16            `json:"market_cap_rank"`
	TotalVolume                            AllCurrencies     `json:"total_volume"`
	High24                                 AllCurrencies     `json:"high_24h"`
	Low24                                  AllCurrencies     `json:"low_24h"`
	PriceChange24h                         float64           `json:"price_change_24h"`
	PriceChangePercentage24h               float64           `json:"price_change_percentage_24h"`
	PriceChangePercentage7d                float64           `json:"price_change_percentage_7d"`
	PriceChangePercentage14d               float64           `json:"price_change_percentage_14d"`
	PriceChangePercentage30d               float64           `json:"price_change_percentage_30d"`
	PriceChangePercentage60d               float64           `json:"price_change_percentage_60d"`
	PriceChangePercentage200d              float64           `json:"price_change_percentage_200d"`
	PriceChangePercentage1y                float64           `json:"price_change_percentage_1y"`
	MarketCapChange24h                     float64           `json:"market_cap_change_24h"`
	MarketCapChangePercentage24h           float64           `json:"market_cap_change_percentage_24h"`
	PriceChange24hInCurrency               AllCurrencies     `json:"price_change_24h_in_currency"`
	PriceChangePercentage1hInCurrency      AllCurrencies     `json:"price_change_percentage_1h_in_currency"`
	PriceChangePercentage24hInCurrency     AllCurrencies     `json:"price_change_percentage_24h_in_currency"`
	PriceChangePercentage7dInCurrency      AllCurrencies     `json:"price_change_percentage_7d_in_currency"`
	PriceChangePercentage14dInCurrency     AllCurrencies     `json:"price_change_percentage_14d_in_currency"`
	PriceChangePercentage30dInCurrency     AllCurrencies     `json:"price_change_percentage_30d_in_currency"`
	PriceChangePercentage60dInCurrency     AllCurrencies     `json:"price_change_percentage_60d_in_currency"`
	PriceChangePercentage200dInCurrency    AllCurrencies     `json:"price_change_percentage_200d_in_currency"`
	PriceChangePercentage1yInCurrency      AllCurrencies     `json:"price_change_percentage_1y_in_currency"`
	MarketCapChange24hInCurrency           AllCurrencies     `json:"market_cap_change_24h_in_currency"`
	MarketCapChangePercentage24hInCurrency AllCurrencies     `json:"market_cap_change_percentage_24h_in_currency"`
	TotalSupply                            *float64          `json:"total_supply"`
	CirculatingSupply                      float64           `json:"circulating_supply"`
	Sparkline                              *SparklineItem    `json:"sparkline_7d"`
	LastUpdated                            string            `json:"last_updated"`
}

// CommunityDataItem map all community data item
type CommunityDataItem struct {
	FacebookLikes            *uint        `json:"facebook_likes"`
	TwitterFollowers         *uint        `json:"twitter_followers"`
	RedditAveragePosts48h    *float64     `json:"reddit_average_posts_48h"`
	RedditAverageComments48h *float64     `json:"reddit_average_comments_48h"`
	RedditSubscribers        *uint        `json:"reddit_subscribers"`
	RedditAccountsActive48h  *interface{} `json:"reddit_accounts_active_48h"`
	TelegramChannelUserCount *uint        `json:"telegram_channel_user_count"`
}

// DeveloperDataItem map all developer data item
type DeveloperDataItem struct {
	Forks              *uint `json:"forks"`
	Stars              *uint `json:"stars"`
	Subscribers        *uint `json:"subscribers"`
	TotalIssues        *uint `json:"total_issues"`
	ClosedIssues       *uint `json:"closed_issues"`
	PRMerged           *uint `json:"pull_requests_merged"`
	PRContributors     *uint `json:"pull_request_contributors"`
	CommitsCount4Weeks *uint `json:"commit_count_4_weeks"`
}

// PublicInterestItem map all public interest item
type PublicInterestItem struct {
	AlexaRank   uint `json:"alexa_rank"`
	BingMatches uint `json:"bing_matches"`
}

// ImageItem struct for all sizes of image
type ImageItem struct {
	Thumb string `json:"thumb"`
	Small string `json:"small"`
	Large string `json:"large"`
}

// ROIItem ROI Item
type ROIItem struct {
	Times      float64 `json:"times"`
	Currency   string  `json:"currency"`
	Percentage float64 `json:"percentage"`
}

// SparklineItem for sparkline
type SparklineItem struct {
	Price []float64 `json:"price"`
}

// TickerItem for ticker
type TickerItem struct {
	Base   string `json:"base"`
	Target string `json:"target"`
	Market struct {
		Name             string `json:"name"`
		Identifier       string `json:"identifier"`
		TradingIncentive bool   `json:"has_trading_incentive"`
	} `json:"market"`
	Last                   float64            `json:"last"`
	Volume                 float64            `json:"volume"`
	CostToMoveUpUsd        *float64           `json:"cost_to_move_up_usd"`
	CostToMoveDownUsd      *float64           `json:"cost_to_move_down_usd"`
	ConvertedLast          map[string]float64 `json:"converted_last"`
	ConvertedVolume        map[string]float64 `json:"converted_volume"`
	TrustScore             string             `json:"trust_score"`
	BidAskSpreadPercentage float64            `json:"bid_ask_spread_percentage"`
	Timestamp              time.Time          `json:"timestamp"`
	LastTradedAt           time.Time          `json:"last_traded_at"`
	LastFetchAt            time.Time          `json:"last_fetch_at"`
	IsAnomaly              bool               `json:"is_anomaly"`
	IsStale                bool               `json:"is_stale"`
	TradeUrl               string             `json:"trade_url"`
	CoinID                 string             `json:"coin_id"`
	TargetCoinID           string             `json:"target_coin_id,omitempty"`
}

// StatusUpdateItem for BEAM
type StatusUpdateItem struct {
	Description string `json:"description"`
	Category    string `json:"category"`
	CreatedAt   string `json:"created_at"`
	User        string `json:"user"`
	UserTitle   string `json:"user_title"`
	Pin         bool   `json:"pin"`
	Project     struct {
		Type  string    `json:"type"`
		Id    string    `json:"id"`
		Name  string    `json:"name"`
		Image ImageItem `json:"image"`
	} `json:"project"`
}

// CoinsListItem item in CoinList
type CoinsListItem struct {
	coinBaseStruct
}

// CoinsMarketItem item in CoinMarket
type CoinsMarketItem struct {
	coinBaseStruct
	Image                               string         `json:"image"`
	CurrentPrice                        float64        `json:"current_price"`
	MarketCap                           float64        `json:"market_cap"`
	MarketCapRank                       int            `json:"market_cap_rank"`
	TotalVolume                         float64        `json:"total_volume"`
	High24                              float64        `json:"high_24h"`
	Low24                               float64        `json:"low_24h"`
	PriceChange24h                      float64        `json:"price_change_24h"`
	PriceChangePercentage24h            float64        `json:"price_change_percentage_24h"`
	MarketCapChange24h                  float64        `json:"market_cap_change_24h"`
	MarketCapChangePercentage24h        float64        `json:"market_cap_change_percentage_24h"`
	CirculatingSupply                   float64        `json:"circulating_supply"`
	TotalSupply                         float64        `json:"total_supply"`
	ATH                                 float64        `json:"ath"`
	ATHChangePercentage                 float64        `json:"ath_change_percentage"`
	ATHDate                             string         `json:"ath_date"`
	ROI                                 *ROIItem       `json:"roi"`
	LastUpdated                         string         `json:"last_updated"`
	SparklineIn7d                       *SparklineItem `json:"sparkline_in_7d"`
	PriceChangePercentage1hInCurrency   *float64       `json:"price_change_percentage_1h_in_currency"`
	PriceChangePercentage24hInCurrency  *float64       `json:"price_change_percentage_24h_in_currency"`
	PriceChangePercentage7dInCurrency   *float64       `json:"price_change_percentage_7d_in_currency"`
	PriceChangePercentage14dInCurrency  *float64       `json:"price_change_percentage_14d_in_currency"`
	PriceChangePercentage30dInCurrency  *float64       `json:"price_change_percentage_30d_in_currency"`
	PriceChangePercentage200dInCurrency *float64       `json:"price_change_percentage_200d_in_currency"`
	PriceChangePercentage1yInCurrency   *float64       `json:"price_change_percentage_1y_in_currency"`
}

// EventCountryItem item in EventsCountries
type EventCountryItem struct {
	Country string `json:"country"`
	Code    string `json:"code"`
}

// ExchangeRates item in ExchangeRate
type ExchangeRates map[string]ExchangeRatesItem

type ExchangeRatesItem struct {
	Name  string  `json:"name"`
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
	Type  string  `json:"type"`
}

// Global for data of /global
type Global struct {
	BaseResult
	ActiveCryptocurrencies          int           `json:"active_cryptocurrencies"`
	UpcomingICOs                    int           `json:"upcoming_icos"`
	EndedICOs                       int           `json:"ended_icos"`
	Markets                         int           `json:"markets"`
	TotalMarketCap                  AllCurrencies `json:"total_market_cap"`
	TotalVolume                     AllCurrencies `json:"total_volume"`
	MarketCapPercentage             AllCurrencies `json:"market_cap_percentage"`
	MarketCapChangePercentage24hUSD float64       `json:"market_cap_change_percentage_24h_usd"`
	UpdatedAt                       int64         `json:"updated_at"`
}
