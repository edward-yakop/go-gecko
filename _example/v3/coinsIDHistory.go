package main

import (
	"fmt"
	"log"

	gecko "github.com/edward-yakop/go-gecko/v3"
)

func main() {
	cg := gecko.NewClient(nil)
	btc, err := cg.CoinsIDHistory(gecko.CoinsIDHistoryParams{
		CoinID:       "bitcoin",
		SnapshotDate: "30-12-2018",
		Localization: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(*btc)
}
