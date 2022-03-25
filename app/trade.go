package app

import (
	"errors"
	"github.com/halm4d/arbitragecli/constants"
)

type Trades []Trade

type Trade struct {
	From   string
	To     string
	Symbol string
	Type   Type
}

func (t *Trades) calculateProfit(bookTickerMap *map[string]BookTicker) (*RouteWithProfit, error) {
	baseBudget := constants.BasePrice
	var previousPrice = baseBudget
	for i, trade := range *t {
		if i == 0 {
			usdtSymbol, err := FindBookTickerByAssetPair(*bookTickerMap, constants.USDT, trade.From)
			if err != nil {
				return &RouteWithProfit{}, err
			}
			previousPrice = ConvertPrice(usdtSymbol, previousPrice, constants.USDT, trade.From, false)
		}
		symbol, err := FindBookTickerByAssetPair(*bookTickerMap, trade.From, trade.To)
		if err != nil {
			return &RouteWithProfit{}, err
		}
		previousPrice = ConvertPrice(symbol, previousPrice, trade.From, trade.To, true)

		if i+1 == len(*t) {
			usdtSymbol, err := FindBookTickerByAssetPair(*bookTickerMap, trade.To, constants.USDT)
			if err != nil {
				return &RouteWithProfit{}, err
			}
			previousPrice = ConvertPrice(usdtSymbol, previousPrice, trade.To, constants.USDT, false)
			return &RouteWithProfit{
				Trades: *t,
				Profit: previousPrice - baseBudget,
			}, nil
		}
	}
	return &RouteWithProfit{}, errors.New("cannot calculate price")
}
