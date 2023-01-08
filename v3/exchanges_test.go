package coingecko

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_Exchanges(t *testing.T) {
	err := setupGock("json/exchanges.json", "/exchanges")
	require.NoError(t, err)

	got, err := c.Exchanges(ExchangesParam{
		PageSize: 250,
	})

	require.NoError(t, err)
	require.Len(t, got, 250)

	gdax := got["gdax"]
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
