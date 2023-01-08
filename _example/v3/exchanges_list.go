package main

import (
	"fmt"
	"log"

	gecko "github.com/edward-yakop/go-gecko/v3"
)

func main() {
	cg := gecko.NewClient(nil)
	exchanges, err := cg.ExchangesList()
	if err != nil {
		log.Fatal(err)
	}

	for id, name := range exchanges {
		fmt.Println(id, name)
	}
}
