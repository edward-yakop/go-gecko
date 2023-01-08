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
