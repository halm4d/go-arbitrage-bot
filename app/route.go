package app

import (
	"errors"
	"fmt"
	"github.com/halm4d/arbitragecli/constants"
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

func CalculateAllRoutes(symbols Symbols) Routes {
	var routes Routes
	for _, startEndAsset := range AllCryptoCurrency() { // BTC ->
		for _, asset1 := range symbols.findAllByAsset(startEndAsset) { // BTC -> ETH ->
			var targetAsset1 = getTargetAsset(asset1, startEndAsset)
			for _, asset2 := range symbols.findAllByAsset(targetAsset1) { // BTC -> ETH -> LTC ->
				targetAsset2 := getTargetAsset(asset2, targetAsset1)
				if targetAsset2 == startEndAsset {
					continue
				}
				pair, err := symbols.findByAssetPair(targetAsset2, startEndAsset)
				if err != nil {
					continue
				}
				trades := Trades{
					{
						From:   startEndAsset,
						To:     targetAsset1,
						Symbol: asset1.Symbol,
						Type:   Buy,
					},
					{
						From:   targetAsset1,
						To:     targetAsset2,
						Symbol: asset2.Symbol,
						Type:   Buy,
					},
					{
						From:   targetAsset2,
						To:     startEndAsset,
						Symbol: pair.Symbol,
						Type:   0,
					},
				}
				routes = append(routes, trades)
			}
		}
	}
	return routes
}

func (r Routes) getProfitableRoutes(symbols Symbols, usdtSymbols Symbols) (RoutesWithProfit, RoutesWithProfit) {
	var profitableRoutes RoutesWithProfit
	var routesWithLoss RoutesWithProfit
	for _, trades := range r {
		profit, err := trades.calculateProfit(symbols, usdtSymbols)
		if err != nil {
			continue
		}
		if profit.Profit > 0 {
			profitableRoutes = append(profitableRoutes, RouteWithProfit{
				Trades: trades,
				Profit: profit.Profit,
			})
		} else {
			routesWithLoss = append(routesWithLoss, RouteWithProfit{
				Trades: trades,
				Profit: profit.Profit,
			})
		}
	}
	sort.Slice(profitableRoutes, func(i, j int) bool {
		return profitableRoutes[i].Profit > profitableRoutes[j].Profit
	})
	sort.Slice(routesWithLoss, func(i, j int) bool {
		return routesWithLoss[i].Profit > routesWithLoss[j].Profit
	})
	return profitableRoutes, routesWithLoss
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

func (t Trades) calculateProfit(symbols Symbols, usdtSymbols Symbols) (RouteWithProfit, error) {
	baseBudget := 100.0
	var previousPrice = baseBudget
	for i, trade := range t {
		if i == 0 {
			usdtSymbols, err := usdtSymbols.findByAssetPair(constants.USDT, trade.From)
			if err != nil {
				return RouteWithProfit{}, err
			}
			previousPrice = convertPrice(usdtSymbols, previousPrice, constants.USDT, trade.From, 0)
		}
		symbol, err := symbols.findByAssetPair(trade.From, trade.To)
		if err != nil {
			return RouteWithProfit{}, err
		}
		previousPrice = convertPrice(symbol, previousPrice, trade.From, trade.To, 0.075)

		if i+1 == len(t) {
			usdtSymbols, err := usdtSymbols.findByAssetPair(trade.To, constants.USDT)
			if err != nil {
				return RouteWithProfit{}, err
			}
			previousPrice = convertPrice(usdtSymbols, previousPrice, trade.To, constants.USDT, 0)
			return RouteWithProfit{
				Trades: t,
				Profit: previousPrice - baseBudget,
			}, nil
		}
	}
	return RouteWithProfit{}, errors.New("cannot calculate price")
}

func convertPrice(symbol Symbol, basePrice float64, from string, to string, fee float64) float64 {
	if strings.EqualFold(symbol.BaseAsset, from) && strings.EqualFold(symbol.QuoteAsset, to) { // SELL
		return symbol.BidPrice * basePrice * (1 - (fee / 100))
	} else { // BUY
		return (1 / symbol.AskPrice) * basePrice * (1 - (fee / 100))
	}
}
