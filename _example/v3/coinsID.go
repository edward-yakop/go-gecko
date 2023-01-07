package main

import (
	"fmt"
	"log"

	gecko "github.com/superoo7/go-gecko/v3"
)

func main() {
	cg := gecko.NewClient(nil)
	coin, err := cg.CoinsID(gecko.CoinsIDParams{
		Id:            "dogecoin",
		Localization:  true,
		Tickers:       true,
		MarketData:    true,
		CommunityData: true,
		DeveloperData: true,
		Sparkline:     true,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(coin)
}
