package arb

import (
	"errors"
	"github.com/halm4d/arbitragecli/constants"
	"strings"
	"sync"
)

type BookTickers struct {
	MU                sync.Mutex
	USDTBookTickers   BookTickerMap
	CryptoBookTickers BookTickerMap
}

type BookTickerMap map[string]BookTicker

type BookTicker struct {
	Symbol     string
	BaseAsset  string
	QuoteAsset string
	BidPrice   float64
	BidQty     float64
	AskPrice   float64
	AskQty     float64
}

func (bookTickerMap *BookTickerMap) FindBookTickerByAssetPair(asset1 string, asset2 string) (*BookTicker, error) {
	for _, bookTicker := range *bookTickerMap {
		if (bookTicker.BaseAsset == asset1 || bookTicker.QuoteAsset == asset1) && (bookTicker.BaseAsset == asset2 || bookTicker.QuoteAsset == asset2) {
			return &bookTicker, nil
		}
	}
	return &BookTicker{}, errors.New("symbol not found")
}

func (bookTicker *BookTicker) ConvertPrice(basePrice float64, from string, to string, addFee bool) float64 {
	fee := .0
	if addFee {
		fee = constants.Fee
	}
	if strings.EqualFold(bookTicker.BaseAsset, from) && strings.EqualFold(bookTicker.QuoteAsset, to) { // SELL
		return bookTicker.BidPrice * basePrice * (1 - (fee / 100))
	} else {
		return (1 / bookTicker.AskPrice) * basePrice * (1 - (fee / 100))
	}
}
