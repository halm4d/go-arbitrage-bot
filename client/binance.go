package client

import (
	"encoding/json"
	"fmt"
	"github.com/halm4d/arbitragecli/error"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var client = &http.Client{Timeout: 10 * time.Second}

func GetPrices(rc chan *[]TickerPriceResp) {
	fmt.Println("Requesting prices...")
	var target []TickerPriceResp = Get[[]TickerPriceResp]("https://api.binance.com/api/v3/ticker/price")
	rc <- &target
	fmt.Println("Requesting prices done.")
}

func GetExchangeInfo(rc chan *[]SymbolResp) {
	fmt.Println("Requesting exchange info...")
	var target exchangeInfoResp = Get[exchangeInfoResp]("https://api.binance.com/api/v3/exchangeInfo")
	rc <- &target.Symbols
	fmt.Println("Requesting exchange info done.")
}

func Get[T any](url string) T {
	resp, err := client.Get(url)
	if err != nil {
		error.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			error.Fatal(err)
		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		error.Fatal(err)
	}

	var target T
	err = json.Unmarshal(body, &target)
	if err != nil {
		error.Fatal(err)
	}

	return target
}

type TickerPriceResp struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type SymbolResp struct {
	Symbol     string `json:"symbol"`
	BaseAsset  string `json:"baseAsset"`
	QuoteAsset string `json:"quoteAsset"`
}

type exchangeInfoResp struct {
	Symbols []SymbolResp `json:"symbols"`
}
