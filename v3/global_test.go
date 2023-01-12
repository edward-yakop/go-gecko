package coingecko

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_Global(t *testing.T) {
	err := setupGock("json/global.json", "json/common.headers.json", "/global")
	require.NoError(t, err)

	got, err := c.Global()
	require.NoError(t, err)
	require.NotNil(t, got)

	assert.Equal(t, commonBaseResult, got.BaseResult)

	assert.Equal(t, 12733, got.ActiveCryptocurrencies, "got.ActiveCryptocurrencies")
	assert.Equal(t, 0, got.UpcomingICOs, "got.UpcomingICOs")
	assert.Equal(t, 3376, got.EndedICOs, "got.EndedICOs")
	assert.Equal(t, 628, got.Markets, "got.Markets")
	assert.Equal(t, 2496883425557.1143, got.TotalMarketCap["xrp"], "got.TotalMarketCap[\"xrp\"]")
	assert.Equal(t, 64403004430.82966, got.TotalVolume["xrp"], "got.TotalVolume[\"xrp\"]")
	assert.Equal(t, 2.0261520751961872, got.MarketCapPercentage["xrp"], "got.MarketCapPercentage[\"xrp\"]")
	assert.Equal(t, -0.06802310774450489, got.MarketCapChangePercentage24hUSD, "got.MarketCapChangePercentage24hUSD")
	assert.Equal(t, int64(1673138179), got.UpdatedAt, "got.UpdatedAt")
}
