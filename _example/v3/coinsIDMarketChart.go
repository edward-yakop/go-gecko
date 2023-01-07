package main

import (
	"fmt"
	gecko "github.com/edward-yakop/go-gecko/v3"
	"log"
)

func main() {
	cg := gecko.NewClient(nil)
	m, err := cg.CoinsIDMarketChart(gecko.CoinsIDMarketChartParams{
		CoinsID:    "bitcoin",
		VsCurrency: "usd",
		Days:       "1",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Prices\n")
	for _, v := range m.Prices {
		fmt.Printf("%s:%.04f\n", v.Time.String(), v.Value)
	}

	fmt.Printf("MarketCaps\n")
	for _, v := range m.MarketCaps {
		fmt.Printf("%s:%.04f\n", v.Time.String(), v.Value)
	}

	fmt.Printf("TotalVolumes\n")
	for _, v := range m.TotalVolumes {
		fmt.Printf("%s:%.04f\n", v.Time.String(), v.Value)
	}
}
