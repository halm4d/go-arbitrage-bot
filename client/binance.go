package client

import (
	"encoding/json"
	"github.com/halm4d/arbitragecli/arb"
	"github.com/halm4d/arbitragecli/error"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var client = &http.Client{Timeout: 10 * time.Second}

func GetExchangeInfo() *arb.ExchangeInfoResp {
	return Get[arb.ExchangeInfoResp]("https://api.binance.com/api/v3/exchangeInfo")
}

func Get[T any](url string) *T {
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

	return &target
}
