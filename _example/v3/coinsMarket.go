package main

import (
	"fmt"
	"log"

	gecko "github.com/superoo7/go-gecko/v3"
	geckoTypes "github.com/superoo7/go-gecko/v3/types"
)

func main() {
	cg := gecko.NewClient(nil)

	priceChangePercentage := []geckoTypes.PriceChangePercentage{
		geckoTypes.PriceChangePercentage1H,
		geckoTypes.PriceChangePercentage24H,
		geckoTypes.PriceChangePercentage7D,
		geckoTypes.PriceChangePercentage14D,
		geckoTypes.PriceChangePercentage30D,
		geckoTypes.PriceChangePercentage200D,
		geckoTypes.PriceChangePercentage1Y,
	}

	market, err := cg.CoinsMarket(gecko.CoinsMarketParams{
		VsCurrency:            "usd",
		CoinIds:               []string{"bitcoin", "ethereum", "steem"},
		Order:                 geckoTypes.CoinMarketOrderMarketCapDesc,
		PageSize:              1,
		PageNo:                1,
		Sparkline:             true,
		PriceChangePercentage: []geckoTypes.PriceChangePercentage{},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total coins: ", len(market))
	fmt.Println(market)

	// with pagination instead
	market, err = cg.CoinsMarket(gecko.CoinsMarketParams{
		VsCurrency:            "usd",
		Order:                 geckoTypes.CoinMarketOrderMarketCapDesc,
		PageSize:              10,
		PageNo:                1,
		Sparkline:             true,
		PriceChangePercentage: priceChangePercentage,
	})
	fmt.Println("Total coins: ", len(market))
	fmt.Println(market)
}
