package coingecko

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPing(t *testing.T) {
	err := setupGockWithHeader("json/ping.json", "json/common.headers.json", "/ping")

	ping, err := c.Ping()
	require.NoError(t, err)
	require.NotNil(t, ping)

	assert.Equal(t, commonBaseResult, ping.BaseResult)

	assert.Equal(t, "(V3) To the Moon!", ping.GeckoSays)
}
