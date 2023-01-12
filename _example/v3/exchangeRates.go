package main

import (
	"fmt"
	"log"

	gecko "github.com/edward-yakop/go-gecko/v3"
)

func main() {
	cg := gecko.NewClient(nil)
	exchangeRates, err := cg.ExchangeRates()
	if err != nil {
		log.Fatal(err)
	}
	r := exchangeRates.Rates["btc"]
	fmt.Println(r.Name)
	fmt.Println(r.Unit)
	fmt.Println(r.Value)
	fmt.Println(r.Type)
}
