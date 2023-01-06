package coingecko

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPing(t *testing.T) {
	err := setupGock("json/ping.json", "/ping")
	ping, err := c.Ping()

	require.NoError(t, err)
	assert.Equal(t, "(V3) To the Moon!", ping.GeckoSays)
}
