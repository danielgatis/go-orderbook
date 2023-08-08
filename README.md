# Go - Orderbook

[![Go Report Card](https://goreportcard.com/badge/github.com/danielgatis/go-orderbook?style=flat-square)](https://goreportcard.com/report/github.com/danielgatis/go-orderbook)
[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/danielgatis/go-orderbook/master/LICENSE)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/danielgatis/go-orderbook)

The pkg `go-orderbook` implements a limit order book for high-frequency trading (HFT), as described by WK Selph.

Based on WK Selph's Blogpost:
https://goo.gl/KF1SRm

## Install

```bash
go get -u github.com/danielgatis/go-orderbook
```

And then import the package in your code:

```go
import "github.com/danielgatis/go-orderbook"
```

### Example

An example described below is one of the use cases.

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/danielgatis/go-orderbook"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func main() {
	book := orderbook.NewOrderBook("BTC/BRL")

	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		side := []orderbook.Side{orderbook.Buy, orderbook.Sell}[rand.Intn(2)]

		book.ProcessPostOnlyOrder(uuid.New().String(), uuid.New().String(), side, decimal.NewFromInt(rand.Int63n(1000)), decimal.NewFromInt(rand.Int63n(1000)))
	}

	depth, _ := json.Marshal(book.Depth())
	var buf bytes.Buffer
	json.Indent(&buf, depth, "", "  ")
	fmt.Println(buf.String())
}
```

```
❯ go run main.go
{
  "bids": [
    {
      "amount": "392",
      "price": "930"
    },
    {
      "amount": "872",
      "price": "907"
    },
    {
      "amount": "859",
      "price": "790"
    },
    {
      "amount": "643",
      "price": "424"
    },
    {
      "amount": "269",
      "price": "244"
    },
    {
      "amount": "160",
      "price": "83"
    },
    {
      "amount": "74",
      "price": "65"
    }
  ],
  "asks": [
    {
      "amount": "178",
      "price": "705"
    },
    {
      "amount": "253",
      "price": "343"
    },
    {
      "amount": "805",
      "price": "310"
    }
  ]
}
```

## License

Copyright (c) 2020-present [Daniel Gatis](https://github.com/danielgatis)

Licensed under [MIT License](./LICENSE)

### Buy me a coffee

Liked some of my work? Buy me a coffee (or more likely a beer)

<a href="https://www.buymeacoffee.com/danielgatis" target="_blank"><img src="https://bmc-cdn.nyc3.digitaloceanspaces.com/BMC-button-images/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;"></a>
