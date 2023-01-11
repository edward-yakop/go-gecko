package coingecko

import (
	"github.com/edward-yakop/go-gecko/v3/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestClient_Exchanges(t *testing.T) {
	err := setupGock("json/exchanges.json", "/exchanges")
	require.NoError(t, err)

	got, err := c.Exchanges(ExchangesParam{
		PageSize: 250,
	})

	require.NoError(t, err)
	require.NotNil(t, got)

	require.Len(t, got.Entries, 250)

	gdax := got.Entries["gdax"]
	assert.Equal(t, "gdax", gdax.Id, "gdax.Id")
	assert.Equal(t, "Coinbase Exchange", gdax.Name, "gdax.Name")
	assert.Equal(t, 2012, *gdax.YearEstablished, "gdax.YearEstablished")
	assert.Equal(t, "United States", *gdax.Country, "gdax.Country")
	assert.Equal(t, "https://www.coinbase.com", gdax.Url, "gdax.Url")
	assert.Equal(t, false, *gdax.HasTradingIncentive, "gdax.HasTradingIncentive")
	assert.Equal(t, 10, *gdax.TrustScore, "gdax.TrustScore")
	assert.Equal(t, 1, *gdax.TrustScoreRank, "gdax.TrustScoreRank")
	assert.Equal(t, 27215.167728668854, gdax.TradeVolume24HBtc, "gdax.TradeVolume24HBtc")
	assert.Equal(t, 27215.167728668854, gdax.TradeVolume24HBtcNormalized, "gdax.TradeVolume24HBtcNormalized")
}

func TestClient_ExchangesList(t *testing.T) {
	err := setupGock("json/exchanges_list.json", "/exchanges/list")
	require.NoError(t, err)

	got, err := c.ExchangesList()
	require.NoError(t, err)
	require.NotEmpty(t, got)

	binance, ok := got["binance"]
	assert.True(t, ok)
	assert.Equal(t, "Binance", binance)
}

func TestClient_ExchangesID(t *testing.T) {
	err := setupGock("json/exchanges_id.json", "/exchanges/binance")
	require.NoError(t, err)

	got, err := c.ExchangesID("binance")
	require.NoError(t, err)
	require.NotNil(t, got)

	assert.Equal(t, "Binance", got.Name, "got.Name")
	assert.Equal(t, 2017, got.YearEstablished, "got.YearEstablished")
	assert.Equal(t, "Cayman Islands", got.Country, "got.Country")
	assert.Equal(t, "https://www.binance.com/", got.Url, "got.Url")
	assert.Equal(t, false, got.HasTradingIncentive, "got.HasTradingIncentive")
	assert.Equal(t, true, got.Centralized, "got.Centralized")
	assert.Equal(t, 9, got.TrustScore, "got.TrustScore")
	assert.Equal(t, 9, got.TrustScoreRank, "got.TrustScoreRank")
	assert.Equal(t, 268336.72482915886, got.TradeVolume24HBtc, "got.TradeVolume24HBtc")
	assert.Equal(t, 99259.77539045323, got.TradeVolume24HBtcNormalized, "got.TradeVolume24HBtcNormalized")
	assert.Len(t, got.Tickers, 100, "len(got.Tickers)")
	assert.Len(t, got.StatusUpdates, 7, "len(got.StatusUpdates)")
}

func TestClient_ExchangesIDTickers(t *testing.T) {
	err := setupGock("json/exchanges_tickers.json", "/exchanges/binance/tickers")
	require.NoError(t, err)

	got, err := c.ExchangesIDTickers(ExchangesIDTickersParams{
		ExchangeID:             "binance",
		CoinIds:                []string{"bitcoin"},
		IncludeExchangeLogo:    true,
		PageNo:                 1,
		Show2PctOrderBookDepth: true,
		Order:                  types.TickerOrderVolumeDesc,
	})
	require.NoError(t, err)
	require.NotNil(t, got)

	assert.Equal(t, "Binance", got.Name, "got.Name")
	assert.Len(t, got.Tickers, 13, "len(got.Tickers)")

	btcUsdt := got.Tickers[0]
	assert.Equal(t, "BTC", btcUsdt.Base, "got.Tickers[0].Base")
	assert.Equal(t, "USDT", btcUsdt.Target, "got.Tickers[0].Target")
	assert.Equal(t, "Binance", btcUsdt.Market.Name, "got.Tickers[0].Market.Name")
	assert.Equal(t, "binance", btcUsdt.Market.Identifier, "got.Tickers[0].Market.Identifier")
	assert.Equal(t, false, btcUsdt.Market.TradingIncentive, "got.Tickers[0].Market.TradingIncentive")
	assert.Equal(t, 17426.43, btcUsdt.Last, "got.Tickers[0].Last")
	assert.Equal(t, 218640.30858082982, btcUsdt.Volume, "got.Tickers[0].Volume")
	assert.Equal(t, 16946250.5177944, *btcUsdt.CostToMoveUpUsd, "got.Tickers[0].CostToMoveUpUsd")
	assert.Equal(t, 17557635.6421924, *btcUsdt.CostToMoveDownUsd, "got.Tickers[0].CostToMoveDownUsd")
	assert.Equal(t, 17423.43, btcUsdt.ConvertedLast["usd"], "got.Tickers[0].ConvertedLast[\"usd\"]")
	assert.Equal(t, 3809463925, int(btcUsdt.ConvertedVolume["usd"]), "got.Tickers[0].ConvertedVolume[\"usd\"]")
	assert.Equal(t, "green", btcUsdt.TrustScore, "got.Tickers[0].TrustScore")
	assert.Equal(t, 0.011952, btcUsdt.BidAskSpreadPercentage, "got.Tickers[0].BidAskSpreadPercentage")
	assert.Equal(t, "2023-01-11T02:46:38Z", btcUsdt.Timestamp.Format(time.RFC3339), "got.Tickers[0].Timestamp")
	assert.Equal(t, "2023-01-11T02:46:38Z", btcUsdt.LastTradedAt.Format(time.RFC3339), "got.Tickers[0].LastTradedAt")
	assert.Equal(t, "2023-01-11T02:46:38Z", btcUsdt.LastFetchAt.Format(time.RFC3339), "got.Tickers[0].LastFetchAt")
	assert.Equal(t, false, btcUsdt.IsAnomaly, "got.Tickers[0].IsAnomaly")
	assert.Equal(t, false, btcUsdt.IsStale, "got.Tickers[0].IsStale")
	assert.Equal(t, "https://www.binance.com/en/trade/BTC_USDT?ref=37754157", btcUsdt.TradeUrl, "got.Tickers[0].TradeUrl")
	assert.Equal(t, "bitcoin", btcUsdt.CoinID, "got.Tickers[0].CoinID")
	assert.Equal(t, "tether", btcUsdt.TargetCoinID, "got.Tickers[0].TargetCoinID")
}
