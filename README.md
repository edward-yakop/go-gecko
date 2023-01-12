# CoinGecko API Client for Go

[![Build Status](https://github.com/edward-yakop/go-gecko/actions/workflows/go.yml/badge.svg)](https://github.com/edward-yakop/go-gecko/actions/workflows/go.yml) [![GoDoc](https://godoc.org/github.com/edward-yakop/go-gecko?status.svg)](https://godoc.org/github.com/edward-yakop/go-gecko)

Simple API Client for CoinGecko written in Go

<p align="center">
    <img src="gogecko.png" alt="gogecko" height="200" />
</p>

gopher resources from [free-gophers-pack](https://github.com/MariaLetta/free-gophers-pack)

## Available endpoint

[Refer to CoinGecko official API](https://www.coingecko.com/api)

|            Endpoint             |       Status       |      Testing       |          Function           |
|:-------------------------------:|:------------------:|:------------------:|:---------------------------:|
|              /ping              | :heavy_check_mark: | :heavy_check_mark: |            Ping             |
|          /simple/price          | :heavy_check_mark: | :heavy_check_mark: |         SimplePrice         |
| /simple/supported_vs_currencies | :heavy_check_mark: | :heavy_check_mark: | SimpleSupportedVSCurrencies |
|           /coins/list           | :heavy_check_mark: | :heavy_check_mark: |          CoinsList          |
|         /coins/markets          | :heavy_check_mark: | :heavy_check_mark: |        CoinsMarkets         |
|           /coins/{id}           | :heavy_check_mark: | :heavy_check_mark: |           CoinsID           |
|       /coins/{id}/tickers       | :heavy_check_mark: | :heavy_check_mark: |       CoinsIDTickers        |
|       /coins/{id}/history       | :heavy_check_mark: | :heavy_check_mark: |       CoinsIDHistory        |
|    /coins/{id}/market_chart     | :heavy_check_mark: | :heavy_check_mark: |     CoinsIDMarketChart      |
|           /exchanges            | :heavy_check_mark: | :heavy_check_mark: |          Exchanges          |
|         /exchanges/{id}         | :heavy_check_mark: | :heavy_check_mark: |         ExchangesID         |
|         /exchanges/list         | :heavy_check_mark: | :heavy_check_mark: |        ExchangesList        |
|     /exchanges/{id}/tickers     | :heavy_check_mark: | :heavy_check_mark: |      ExchangesTickers       |
|         /exchange_rates         | :heavy_check_mark: | :heavy_check_mark: |        ExchangeRate         |
|             /global             | :heavy_check_mark: | :heavy_check_mark: |           Global            |

## Usage

Installation with go get.

```
go get -u github.com/edward-yakop/go-gecko
```

For usage, checkout [Example folder for v3](/_example/v3)

For production, you might need to set time out for httpClient, here's a sample code:

```go
package main

import (
	"net/http"
	"time"

	coingecko "github.com/edward-yakop/go-gecko/v3"
)

func main() {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	CG := coingecko.NewClient(httpClient)
}
```

## Convention

refer to https://medium.com/@marcus.olsson/writing-a-go-client-for-your-restful-api-c193a2f4998c

## License

MIT
