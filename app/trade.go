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

func (t *Trades) calculateProfit(symbols *Symbols, usdtSymbols *Symbols) (*RouteWithProfit, error) {
	baseBudget := constants.BasePrice
	var previousPrice = baseBudget
	for i, trade := range *t {
		if i == 0 {
			usdtSymbol, err := usdtSymbols.findByAssetPair(constants.USDT, trade.From)
			if err != nil {
				return &RouteWithProfit{}, err
			}
			previousPrice = usdtSymbol.convertPrice(previousPrice, constants.USDT, trade.From, false)
		}
		symbol, err := symbols.findByAssetPair(trade.From, trade.To)
		if err != nil {
			return &RouteWithProfit{}, err
		}
		previousPrice = symbol.convertPrice(previousPrice, trade.From, trade.To, true)

		if i+1 == len(*t) {
			usdtSymbol, err := usdtSymbols.findByAssetPair(trade.To, constants.USDT)
			if err != nil {
				return &RouteWithProfit{}, err
			}
			previousPrice = usdtSymbol.convertPrice(previousPrice, trade.To, constants.USDT, false)
			return &RouteWithProfit{
				Trades: *t,
				Profit: previousPrice - baseBudget,
			}, nil
		}
	}
	return &RouteWithProfit{}, errors.New("cannot calculate price")
}
