package arb

import (
	"errors"
	"github.com/halm4d/arbitragecli/constants"
	"sync"
)

type BookTickers struct {
	MU                sync.Mutex
	USDTBookTickers   BookTickerMap
	CryptoBookTickers BookTickerMap
}

type BookTickerMap map[string]*BookTicker

type BookTicker struct {
	Symbol     string
	BaseAsset  string
	QuoteAsset string
	BidPrice   float64
	BidQty     float64
	AskPrice   float64
	AskQty     float64
}

func (bookTickerMap BookTickerMap) FindBookTickerByAssetPair(asset1 string, asset2 string) (*BookTicker, error) {
	bookTicker, ok := bookTickerMap[asset1+asset2]
	if ok {
		return bookTicker, nil
	}
	bookTicker, ok = bookTickerMap[asset2+asset1]
	if ok {
		return bookTicker, nil
	}
	return &BookTicker{}, errors.New("symbol not found")
}

func (bookTicker *BookTicker) ConvertPrice(basePrice float64, from string, to string, addFee bool) float64 {
	fee := .0
	if addFee {
		fee = constants.Fee
	}
	if bookTicker.Symbol == from+to { // SELL
		return bookTicker.BidPrice * basePrice * (1 - (fee / 100))
	} else {
		return (1 / bookTicker.AskPrice) * basePrice * (1 - (fee / 100))
	}
}
