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

func TestCoinsID(t *testing.T) {
	err := setupGock("json/coins_id.json", "/coins/dogecoin")
	require.NoError(t, err)

	coin, err := c.CoinsID("dogecoin", true, true, true, true, true, true)
	require.NoError(t, err)
	require.NotNil(t, coin)

	require.Equal(t, "doge", coin.Symbol, "coin.Symbol")
	if assert.Len(t, coin.Categories, 4) {
		assert.Equal(t, "EY specific change", coin.Categories[3], "coin.Categories[4]")
	}
	assert.Equal(t, 3, int(coin.BlockTimeInMin), "coin.BlockTimeInMin")
	assert.Equal(t, "Scrypt", coin.HashingAlgorithm, "coin.HashingAlgorithm")
	assert.Equal(t, "الدوجكوين", coin.Localization["ar"], "coin.Localization[\"ar\"]")
	assert.Equal(t, "처음에 \"joke currency\"라고 불리기도 하면서 장난처럼 시작한 도지코인은 일본 개인 시바 이누를 마스코트로 사용합니다. 이 시바견은 인터넷에서 재미로 사용되던 그림이며, 같은 그림이 코인의 로고로 이용되고 있습니다. 이를 통해 도지코인이 재밌고 친근한 가상화폐라는 점을 강조합니다. 실제로 당장 홈페이지만 접속해도, 아주 쉽게 도지코인 지갑을 설치할 수 있습니다. \r\n\r\n개발자 빌리 마커스는 불법 마약 거래 사이트 실크로드에서 사용되는 비트코인과는 달리, 악용되지 않으면서 더 넓은 인구들의 사용을 위해 도지코인을 만든 것이라고 합니다.\r\n\r\n코인 특징\r\n1. 도지코인은 라이트코인에서 포크된 럭키코인에서 포크되었습니다. 그래서 처음에는 럭키코인처럼 채굴보상이 랜덤이었는데, 이후 정해진 보상으로 정책을 바꿨습니다. \r\n\r\n2. 도지코인은 빠른 코인 생산 속도를 가지고 있습니다. 처음에는 생산량이 1,000억 개로 고정돼있었는데, 무제한 생산으로 바뀌었습니다. 현재 10,000개의 코인이 1분마다 생겨나는 중이고, 1년에는 52억 개의 새로운 도지코인이 생겨납니다. 2015년 6월 30일 1,000억 개의 코인이 이미 생산되었습니다. \r\n\r\n3. 도지코인은 SNS에서 팁을 줄 수 있는 시스템을 통해 인기를 얻었습니다. 즉, 도지코인을 이용해 사용자들이 흥미롭거나 가치 있는 콘텐츠를 제공한 사람에게 팁을 주는 것입니다. 레딧, 트위터, 트위치티비(Twitch.TV)등에서 이런 서비스를 제공하는 도지팁봇(Dogetipbot)이 등장하기도 했으나 현재 사용 가능한 팁봇은 제한적입니다. ", coin.Description["ko"], "coin.Description[\"ko\"]")
	assert.Equal(t, "dogecoin", (*coin.Links)["twitter_screen_name"], "coin.Links[\"twitter_screen_name\"]")
	assert.Equal(t, "2013-12-08", coin.GenesisDate, "coin.GenesisDate")
	assert.Equal(t, 8, int(coin.MarketCapRank), "coin.MarketCapRank")
	assert.Equal(t, 6, int(coin.CoinGeckoRank), "coin.CoinGeckoRank")
}

func TestClient_CoinsIDHistory(t *testing.T) {
	err := setupGock("json/coins_id_history.json", "/coins/bitcoin/history")
	require.NoError(t, err)

	history, err := c.CoinsIDHistory("bitcoin", "06-01-2022", true)
	require.NoError(t, err)
	require.NotNil(t, history)

	assert.Equal(t, "btc", history.Symbol, "history.Symbol")
	assert.Equal(t, "bitcoin", history.ID, "history.ID")
	assert.Equal(t, "ビットコイン", history.Localization["ja"], "history.Localization['ja']")
	assert.Equal(t, "https://assets.coingecko.com/coins/images/1/thumb/bitcoin.png?1547033579", history.Image.Thumb, "history.Image.Thumb")
	assert.Equal(t, 1.0, history.MarketData.CurrentPrice["btc"], "history.MarketData.CurrentPrice[\"btc\"]")
	assert.Equal(t, 6451969929451.801, history.MarketData.MarketCap["hkd"], "history.MarketData.MarketCap[\"hkd\"]")
	assert.Equal(t, 66819045.223431624, history.MarketData.TotalVolume["bnb"], "history.MarketData.TotalVolume[\"bnb\"]")
	assert.Equal(t, 4231416, int(*history.CommunityData.TwitterFollowers), "history.CommunityData.TwitterFollowers")
	assert.Equal(t, 31111, int(*history.DeveloperData.Forks), "history.DeveloperData.Forks")
}
