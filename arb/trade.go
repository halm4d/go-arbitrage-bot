package arb

import (
	"fmt"
	"github.com/halm4d/arbitragecli/constants"
	"log"
	"math"
)

type Trades struct {
	Trades []Trade
	Profit float64
}

type Trade struct {
	From   string
	To     string
	Symbol string
	Type   Type
}

func (t *Trades) CalculateProfit(bookTickerMap *BookTickerMap, usdtBookTicker *BookTickerMap) float64 {
	var previousPrice = constants.BasePrice
	for i, trade := range t.Trades {
		if i == 0 {
			usdtSymbol, err := usdtBookTicker.FindBookTickerByAssetPair(constants.USDT, trade.From)
			if err != nil {
				return -math.MaxFloat64
			}
			previousPrice = usdtSymbol.ConvertPrice(previousPrice, constants.USDT, trade.From, false)
		}
		symbol, err := bookTickerMap.FindBookTickerByAssetPair(trade.From, trade.To)
		if err != nil {
			return -math.MaxFloat64
		}
		previousPrice = symbol.ConvertPrice(previousPrice, trade.From, trade.To, true)

		if i+1 == len(t.Trades) {
			usdtSymbol, err := usdtBookTicker.FindBookTickerByAssetPair(trade.To, constants.USDT)
			if err != nil {
				return -math.MaxFloat64
			}
			previousPrice = usdtSymbol.ConvertPrice(previousPrice, trade.To, constants.USDT, false)
			return previousPrice - constants.BasePrice
		}
	}
	return -math.MaxFloat64
}

func (t *Trades) getRouteString() string {
	readableTrade := ""
	for i, trade := range t.Trades {
		if i == 0 {
			readableTrade = fmt.Sprintf("%s %s -> %s", readableTrade, trade.From, trade.To)
		} else {
			readableTrade = fmt.Sprintf("%s -> %s", readableTrade, trade.To)
		}
	}
	return fmt.Sprintf("%s Profit: %f USD", readableTrade, t.Profit)
}

func (t *Trades) print() {
	readableTrade := ""
	for i, trade := range t.Trades {
		if i == 0 {
			readableTrade = fmt.Sprintf("%s %s -> %s", readableTrade, trade.From, trade.To)
		} else {
			readableTrade = fmt.Sprintf("%s -> %s", readableTrade, trade.To)
		}
	}
	log.Printf("%s Profit: %f USD\n", readableTrade, t.Profit)
}

func (r Routes) GetBestRouteString() string {
	if len(r) == 0 {
		return ""
	}
	return r[0].getRouteString()
}

func (r Routes) Print(top int) {
	var limit int
	if len(r) > top {
		limit = top
	} else {
		limit = len(r)
	}
	for _, route := range r[:limit] {
		route.print()
	}
}
