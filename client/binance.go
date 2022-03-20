package client

import (
	"encoding/json"
	"github.com/halm4d/arbitragecli/error"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var client = &http.Client{Timeout: 10 * time.Second}

func GetPrices(rc chan *[]TickerPriceResp) {
	var target = Get[[]TickerPriceResp]("https://api.binance.com/api/v3/ticker/bookTicker")
	rc <- &target
}

func GetExchangeInfo(rc chan *[]SymbolResp) {
	var target = Get[ExchangeInfoResp]("https://api.binance.com/api/v3/exchangeInfo")
	rc <- &target.Symbols
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
	//Price    string `json:"price"`
	AskPrice string `json:"askPrice"`
	BidPrice string `json:"bidPrice"`
}

type SymbolResp struct {
	Symbol     string `json:"symbol"`
	BaseAsset  string `json:"baseAsset"`
	QuoteAsset string `json:"quoteAsset"`
	Status     string `json:"status"`
}

type ExchangeInfoResp struct {
	Symbols []SymbolResp `json:"symbols"`
}
