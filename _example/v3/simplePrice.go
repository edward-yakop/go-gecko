package main

import (
	"fmt"
	"log"

	gecko "github.com/edward-yakop/go-gecko/v3"
)

func main() {
	cg := gecko.NewClient(nil)

	sp, err := cg.SimplePrice(gecko.SimplePriceParams{
		CoinIDs:           []string{"bitcoin", "ethereum"},
		VsCurrencies:      []string{"usd", "myr"},
		MarketCap:         true,
		Include24HrVolume: true,
		Include24HrChange: true,
		LastUpdatedAt:     true,
		Precision:         "full",
	})
	if err != nil {
		log.Fatal(err)
	}
	bitcoin := sp.Coins["bitcoin"]
	eth := sp.Coins["ethereum"]
	fmt.Println(fmt.Sprintf("Bitcoin is worth %f usd (myr %f)", bitcoin.Currencies["usd"].Price, bitcoin.Currencies["myr"].Price))
	fmt.Println(fmt.Sprintf("Ethereum is worth %f usd (myr %f)", eth.Currencies["usd"].Price, eth.Currencies["myr"].Price))
}
