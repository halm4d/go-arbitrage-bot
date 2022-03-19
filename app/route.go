package app

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

const (
	Buy Type = iota
	Sell
)

type Type int8

type Routes []Trades
type Trades []Trade
type RoutesWithProfit []RouteWithProfit

type Trade struct {
	From   string
	To     string
	Symbol string
	Price  float64
	Type   Type
}

//  Symbols{
// 		Symbol{Symbol: "ETHBTC", BaseAsset: "ETH", QuoteAsset: "BTC"},
// 		Symbol{Symbol: "LTCETH", BaseAsset: "LTC", QuoteAsset: "ETH"},
// 		Symbol{Symbol: "LTCBTC", BaseAsset: "LTC", QuoteAsset: "BTC"},
//  }

// 	BTC -> ETH -> LTC -> BTC
// 	BTC -> LTC -> ETH -> BTC
// 	LTC -> BTC -> ETH -> LTC
// 	LTC -> ETH -> BTC -> LTC
// 	ETH -> LTC -> BTC -> ETH
// 	ETH -> BTC -> LTC -> ETH

func CalculateAllRoutes(symbols Symbols, busdSymbols Symbols) RoutesWithProfit {
	var routeWithProfits RoutesWithProfit
	for _, startEndAsset := range AllCryptoCurrency() { // BTC ->
		for _, asset1 := range symbols.findAllByAsset(startEndAsset) { // BTC -> ETH ->
			var targetAsset1 = getTargetAsset(asset1, startEndAsset)
			for _, asset2 := range symbols.findAllByAsset(targetAsset1) { // BTC -> ETH -> LTC ->
				targetAsset2 := getTargetAsset(asset2, targetAsset1)
				if targetAsset2 == startEndAsset {
					continue
				}
				pair, err := symbols.findByAssetPairs(targetAsset2, startEndAsset)
				if err != nil {
					continue
				}
				trades := Trades{
					{
						From:   startEndAsset,
						To:     targetAsset1,
						Symbol: asset1.Symbol,
						Price:  asset1.Price,
						Type:   Buy,
					},
					{
						From:   targetAsset1,
						To:     targetAsset2,
						Symbol: asset2.Symbol,
						Price:  asset2.Price,
						Type:   Buy,
					},
					{
						From:   targetAsset2,
						To:     startEndAsset,
						Symbol: pair.Symbol,
						Price:  pair.Price,
						Type:   0,
					},
				}
				profit, err := trades.calculateProfit(symbols, busdSymbols)
				if err != nil {
					continue
				}
				routeWithProfits = append(routeWithProfits, profit)
			}
		}
	}
	return routeWithProfits
}

func getTargetAsset(symbol Symbol, ignore string) string {
	var targetAsset2 string
	if symbol.QuoteAsset != ignore {
		targetAsset2 = symbol.QuoteAsset
	} else {
		targetAsset2 = symbol.BaseAsset
	}
	return targetAsset2
}

func (t RouteWithProfit) print() {
	var readableTrade string
	for i, trade := range t.Trades {
		if i == 0 {
			readableTrade = fmt.Sprintf("%s %s -> %s", readableTrade, trade.From, trade.To)
		} else {
			readableTrade = fmt.Sprintf("%s -> %s", readableTrade, trade.To)
		}
	}
	fmt.Printf("%s Profit: %f USD\n", readableTrade, t.Profit)
}

func (r RoutesWithProfit) print(top int) {
	sort.Slice(r, func(i, j int) bool {
		return r[i].Profit > r[j].Profit
	})
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

type RouteWithProfit struct {
	Trades
	Profit float64
}

func (t Trades) calculateProfit(symbols Symbols, busdSymbols Symbols) (RouteWithProfit, error) {
	baseBudget := 100.0
	var previousPrice = baseBudget
	for i, trade := range t {
		if i == 0 {
			busdSymbol, err := busdSymbols.findByAssetPairs("BUSD", trade.From)
			if err != nil {
				return RouteWithProfit{}, err
			}
			previousPrice = convertPrice(busdSymbol, previousPrice, "BUSD", trade.From)
		}
		symbol, err := symbols.findByAssetPairs(trade.From, trade.To)
		if err != nil {
			return RouteWithProfit{}, err
		}
		previousPrice = convertPrice(symbol, previousPrice, trade.From, trade.To)

		if i+1 == len(t) {
			busdSymbol, err := busdSymbols.findByAssetPairs(trade.To, "BUSD")
			if err != nil {
				return RouteWithProfit{}, err
			}
			previousPrice = convertPrice(busdSymbol, previousPrice, trade.To, "BUSD")
			return RouteWithProfit{
				Trades: t,
				Profit: previousPrice - baseBudget,
			}, nil
		}
	}
	return RouteWithProfit{}, errors.New("cannot calculate price")
}

//func (r Routes) CalculateProfits(symbols Symbols, busdSymbols Symbols) []RouteWithProfit {
//	var routesWithProfit []RouteWithProfit
//	for _, trades := range r {
//		var previousPrice = 100.0
//		for i, trade := range trades {
//			if i == 0 {
//				busdSymbol, err := busdSymbols.findByAssetPairs("BUSD", trade.From)
//				if err != nil {
//					break
//				}
//				previousPrice = convertPrice(busdSymbol, previousPrice, "BUSD", trade.From)
//			}
//			symbol, err := symbols.findByAssetPairs(trade.From, trade.To)
//			if err != nil {
//				break
//			}
//			previousPrice = convertPrice(symbol, previousPrice, trade.From, trade.To)
//
//			if i+1 == len(trades) {
//				busdSymbol, err := busdSymbols.findByAssetPairs(trade.To, "BUSD")
//				if err != nil {
//					break
//				}
//				previousPrice = convertPrice(busdSymbol, previousPrice, trade.To, "BUSD")
//
//				routesWithProfit = append(routesWithProfit, RouteWithProfit{
//					Trades: trades,
//					Profit: previousPrice,
//				})
//			}
//		}
//	}
//	return routesWithProfit
//}

func convertPrice(symbol Symbol, basePrice float64, from string, to string) float64 {
	fee := 1 - 0.75/100
	if strings.EqualFold(symbol.BaseAsset, from) && strings.EqualFold(symbol.QuoteAsset, to) {
		return symbol.Price * basePrice * fee
	} else {
		return (1 / symbol.Price) * basePrice * fee
	}
}
