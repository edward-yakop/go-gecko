package coingecko

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSimpleSinglePrice(t *testing.T) {
	err := setupGock("json/simple_single_price.json", "/simple/price")
	require.NoError(t, err)

	simplePrice, err := c.SimpleSinglePrice("bitcoin", "usd")
	require.NoError(t, err)

	assert.Equal(t, "bitcoin", simplePrice.ID, "simplePrice.ID")
	assert.Equal(t, "usd", simplePrice.Currency, "simplePrice.Currency")
	assert.Equal(t, 5013.61, simplePrice.MarketPrice, "simplePrice.MarketPrice")
}

func TestSimplePrice(t *testing.T) {
	err := setupGock("json/simple_price.json", "/simple/price")
	require.NoError(t, err)

	ids := []string{"bitcoin", "ethereum"}
	vc := []string{"usd", "myr"}
	sp, err := c.SimplePrice(ids, vc)
	require.NoError(t, err)

	bitcoin := (*sp)["bitcoin"]
	eth := (*sp)["ethereum"]

	assert.Equal(t, 5005.73, bitcoin["usd"], "bitcoin['usd']")
	assert.Equal(t, float64(20474), bitcoin["myr"], "bitcoin['myr']")

	assert.Equal(t, 163.58, eth["usd"], "eth['usd']")
	assert.Equal(t, 669.07, eth["myr"], "eth['myr']")
}

func TestSimpleSupportedVSCurrencies(t *testing.T) {
	err := setupGock("json/simple_supported_vs_currencies.json", "/simple/supported_vs_currencies")
	s, err := c.SimpleSupportedVSCurrencies()
	require.NoError(t, err)

	assert.Len(t, *s, 54)
}
