package main

import (
	"fmt"
	"log"

	gecko "github.com/superoo7/go-gecko/v3"
)

func main() {
	cg := gecko.NewClient(nil)

	tickers, err := cg.CoinsIDTickers(gecko.CoinsIDTickersParam{
		CoinsID: "bitcoin",
		PageNo:  1,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(tickers.Tickers))
	tickers, err = cg.CoinsIDTickers(gecko.CoinsIDTickersParam{
		CoinsID: "bitcoin",
		PageNo:  2,
	})
	fmt.Println(len(tickers.Tickers))
}
