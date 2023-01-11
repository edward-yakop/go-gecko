package main

import (
	"fmt"
	"github.com/edward-yakop/go-gecko/v3/types"
	"log"

	gecko "github.com/edward-yakop/go-gecko/v3"
)

func main() {
	cg := gecko.NewClient(nil)
	exchangeTickers, err := cg.ExchangesIDTickers(gecko.ExchangesIDTickersParams{
		ExchangeID:             "binance",
		CoinIds:                []string{"bitcoin"},
		IncludeExchangeLogo:    true,
		PageNo:                 1,
		Show2PctOrderBookDepth: true,
		Order:                  types.TickerOrderVolumeDesc,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(exchangeTickers.Name)
	for _, ticker := range exchangeTickers.Tickers {
		fmt.Println(ticker.Base, ticker.Target)
	}
}
