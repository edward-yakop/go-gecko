package main

import (
	"fmt"
	"log"

	gecko "github.com/edward-yakop/go-gecko/v3"
)

func main() {
	cg := gecko.NewClient(nil)
	exchanges, err := cg.Exchanges(gecko.ExchangesParam{
		PageSize: 1000,
		PageNo:   1,
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range exchanges {
		fmt.Println(r.Id)
		fmt.Println(r.Name)
	}
}
