package coingecko

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestSimplePrice(t *testing.T) {
	err := setupGockWithHeader("json/simple_price.json", "json/common.headers.json", "/simple/price")
	require.NoError(t, err)

	sp, err := c.SimplePrice(SimplePriceParams{
		CoinIDs:              []string{"bitcoin", "ethereum"},
		VsCurrencies:         []string{"usd", "eur", "btc", "myr", "idr"},
		IncludeMarketCap:     false,
		Include24HrVolume:    false,
		Include24HrChange:    false,
		IncludeLastUpdatedAt: false,
	})
	require.NoError(t, err)
	require.NotNil(t, sp)

	assert.Equal(t, commonBaseResult, sp.BaseResult)

	if bitcoin := sp.Coins["bitcoin"]; assert.NotNil(t, bitcoin, "sp.Coins[\"bitcoin\"]") {
		if usd := bitcoin.Currencies["usd"]; assert.NotNil(t, usd, "bitcoin.Currencies[\"usd\"]") {
			assert.Equal(t, 18109.977835707818, usd.Price, "bitcoin.usd.Price")
			assert.Equal(t, 348916085115.57404, *usd.MarketCap, "bitcoin.usd.MarketCap")
			assert.Equal(t, 33339997882.460022, *usd.Volume24H, "bitcoin.usd.Volume24H")
			assert.Equal(t, 3.88635385851444, *usd.ChangePercentage24H, "bitcoin.usd.ChangePercentage24H")
		}
		if myr := bitcoin.Currencies["myr"]; assert.NotNil(t, myr, "bitcoin.Currencies[\"myr\"]") {
			assert.Equal(t, 79055.48624621533, myr.Price, "bitcoin.myr.Price")
			assert.Equal(t, 1523123386355.013, *myr.MarketCap, "bitcoin.myr.MarketCap")
			assert.Equal(t, 145539092756.30273, *myr.Volume24H, "bitcoin.myr.Volume24H")
			assert.Equal(t, 3.7034302535037207, *myr.ChangePercentage24H, "bitcoin.myr.ChangePercentage24H")
		}
		assert.Equal(t, time.Date(2023, time.January, 12, 8, 36, 57, 0, time.UTC), bitcoin.LastUpdatedAt.UTC())
	}

	if ethereum := sp.Coins["ethereum"]; assert.NotNil(t, ethereum, "sp.Coins[\"ethereum\"]") {
		if usd := ethereum.Currencies["usd"]; assert.NotNil(t, usd, "ethereum.Currencies[\"usd\"]") {
			assert.Equal(t, 1396.4856571417022, usd.Price, "ethereum.usd.Price")
			assert.Equal(t, 168253794309.9934, *usd.MarketCap, "ethereum.usd.MarketCap")
			assert.Equal(t, 10129519929.930872, *usd.Volume24H, "ethereum.usd.Volume24H")
			assert.Equal(t, 4.746932530748111, *usd.ChangePercentage24H, "ethereum.usd.ChangePercentage24H")
		}
		if myr := ethereum.Currencies["myr"]; assert.NotNil(t, myr, "ethereum.Currencies[\"myr\"]") {
			assert.Equal(t, 6096.078839120672, myr.Price, "ethereum.myr.Price")
			assert.Equal(t, 734478288301.4131, *myr.MarketCap, "ethereum.myr.MarketCap")
			assert.Equal(t, 44218393350.127235, *myr.Volume24H, "ethereum.myr.Volume24H")
			assert.Equal(t, 4.562493614560572, *myr.ChangePercentage24H, "ethereum.myr.ChangePercentage24H")
		}
		assert.Equal(t, time.Date(2023, time.January, 12, 8, 36, 34, 0, time.UTC), ethereum.LastUpdatedAt.UTC())
	}
}

func TestClient_parseSimplePriceItem(t *testing.T) {
	got, err := c.parseSimplePriceItem("", []byte(`{
    "usd": 18109.977835707818,
    "usd_market_cap": 348916085115.57404,
    "usd_24h_vol": 33339997882.460022,
    "usd_24h_change": 3.88635385851444,
    "eur": 16840.232959712837,
    "eur_market_cap": 324519523528.20685,
    "eur_24h_vol": 31002430610.927105,
    "eur_24h_change": 3.698892545077078
  }`))
	require.NoError(t, err)
	require.NotNil(t, got)

	if usd := got.Currencies["usd"]; assert.NotNil(t, usd, "usd") {
		assert.Equal(t, 18109.977835707818, usd.Price, "usd.Price")
		assert.Equal(t, 348916085115.57404, *usd.MarketCap, "usd.MarketCap")
		assert.Equal(t, 33339997882.460022, *usd.Volume24H, "usd.Volume24H")
		assert.Equal(t, 3.88635385851444, *usd.ChangePercentage24H, "usd.ChangePercentage24H")
	}
	if eur := got.Currencies["eur"]; assert.NotNil(t, eur, "eur") {
		assert.Equal(t, 16840.232959712837, eur.Price, "eur.Price")
		assert.Equal(t, 324519523528.20685, *eur.MarketCap, "eur.MarketCap")
		assert.Equal(t, 31002430610.927105, *eur.Volume24H, "eur.Volume24H")
		assert.Equal(t, 3.698892545077078, *eur.ChangePercentage24H, "eur.ChangePercentage24H")
	}
}

func TestSimpleSupportedVSCurrencies(t *testing.T) {
	err := setupGock("json/simple_supported_vs_currencies.json", "/simple/supported_vs_currencies")
	s, err := c.SimpleSupportedVSCurrencies()
	require.NoError(t, err)

	assert.Len(t, *s, 54)
}
