package coingecko

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCoinsList(t *testing.T) {
	err := setupGock("json/coins_list.json", "/coins/list")
	require.NoError(t, err)

	list, err := c.CoinsList()
	require.NoError(t, err)

	item := (*list)[0]
	assert.Equal(t, "01coin", item.ID, "item.ID")
}
