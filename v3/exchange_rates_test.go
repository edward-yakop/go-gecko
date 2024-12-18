package coingecko

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_ExchangeRates(t *testing.T) {
	err := setupGock("json/exchange_rates.json", "json/common.headers.json", "/exchange_rates")
	require.NoError(t, err)

	got, err := c.ExchangeRates()
	require.NoError(t, err)
	require.NotNil(t, got)

	assert.Equal(t, commonBaseResult, got.BaseResult)

	if eth, ok := got.Rates["eth"]; assert.True(t, ok, "got[\"eth\"]") {
		assert.Equal(t, "Ether", eth.Name, "eth.Name")
		assert.Equal(t, "ETH", eth.Unit, "eth.Unit")
		assert.Equal(t, 13.404, eth.Value, "eth.Value")
		assert.Equal(t, "crypto", eth.Type, "eth.Type")
	}
}
