package arb

import (
	"fmt"
	"github.com/halm4d/arbitragecli/constants"
	"log"
	"math"
)

type Arbitrage struct {
	Trades []Trade
	Profit float64
}

type Trade struct {
	From   string
	To     string
	Symbol string
	Type   Type
}

func (t *Arbitrage) CalculateProfit(bookTickerMap *BookTickerMap, usdtBookTicker *BookTickerMap) float64 {
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

func (t *Arbitrage) getRouteString() string {
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

func (t *Arbitrage) print() {
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

func (a Arbitrages) GetBestRouteString() string {
	if len(a) == 0 {
		return ""
	}
	return a[0].getRouteString()
}

func (a Arbitrages) Print(top int) {
	var limit int
	if len(a) > top {
		limit = top
	} else {
		limit = len(a)
	}
	for _, route := range a[:limit] {
		route.print()
	}
}
