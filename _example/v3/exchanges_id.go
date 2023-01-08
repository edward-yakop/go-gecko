package main

import (
	"fmt"
	"log"

	gecko "github.com/edward-yakop/go-gecko/v3"
)

func main() {
	cg := gecko.NewClient(nil)
	binance, err := cg.ExchangesID("binance")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(binance.Name)
}
