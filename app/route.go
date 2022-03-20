package app

import (
	"errors"
	"fmt"
	"github.com/halm4d/arbitragecli/constants"
	"sort"
	"strings"
	"sync"
	"time"
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

func CalculateAllRoutes(symbols Symbols) Routes {
	startCalculationTime := time.Now()
	var routes = make(Routes, 0)

	var wg sync.WaitGroup
	mu := &sync.Mutex{}
	for _, startEndAsset := range AllCryptoCurrency() {
		wg.Add(1)
		go func(symbols Symbols, startEndAsset string) {
			defer wg.Done()
			OneOfMyStructs := calculateRoutesForSymbol(symbols, startEndAsset)
			mu.Lock()
			routes = append(routes, OneOfMyStructs...)
			mu.Unlock()
		}(symbols, startEndAsset)
	}
	wg.Wait()
	endCalculationTime := time.Now()
	fmt.Println(endCalculationTime.UnixMilli() - startCalculationTime.UnixMilli())
	return routes
}

func calculateRoutesForSymbol(symbols Symbols, startEndAsset string) Routes {
	var routes Routes
	for _, asset1 := range symbols.findAllByAsset(startEndAsset) {
		var targetAsset1 = getTargetAsset(asset1, startEndAsset)
		for _, asset2 := range symbols.findAllByAsset(targetAsset1) {
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
				},
				{
					From:   targetAsset1,
					To:     targetAsset2,
					Symbol: asset2.Symbol,
				},
				{
					From:   targetAsset2,
					To:     startEndAsset,
					Symbol: pair.Symbol,
				},
			}
			routes = append(routes, trades)
		}
	}
	return routes
}

func (r Routes) getProfitableRoutes(symbols Symbols, usdtSymbols Symbols) (RoutesWithProfit, RoutesWithProfit) {
	var profitableRoutes = make(RoutesWithProfit, 0)
	var routesWithLoss = make(RoutesWithProfit, 0)

	var wg sync.WaitGroup
	mu := &sync.Mutex{}

	for _, trades := range r {
		wg.Add(1)
		go func(trades Trades, symbols Symbols, usdtSymbols Symbols) {
			defer wg.Done()
			profit, err := trades.calculateProfit(symbols, usdtSymbols)
			if err != nil {
				return
			}
			if profit.Profit > 0 {
				mu.Lock()
				profitableRoutes = append(profitableRoutes, RouteWithProfit{
					Trades: trades,
					Profit: profit.Profit,
				})
				mu.Unlock()
			} else {
				mu.Lock()
				routesWithLoss = append(routesWithLoss, RouteWithProfit{
					Trades: trades,
					Profit: profit.Profit,
				})
				mu.Unlock()
			}
		}(trades, symbols, usdtSymbols)
	}
	wg.Wait()

	wg.Add(2)
	go func(profitableRoutes RoutesWithProfit) {
		defer wg.Done()
		sort.Slice(profitableRoutes, func(i, j int) bool {
			return profitableRoutes[i].Profit > profitableRoutes[j].Profit
		})
	}(profitableRoutes)
	go func(routesWithLoss RoutesWithProfit) {
		defer wg.Done()
		sort.Slice(routesWithLoss, func(i, j int) bool {
			return routesWithLoss[i].Profit > routesWithLoss[j].Profit
		})
	}(routesWithLoss)
	wg.Wait()
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

func (t RouteWithProfit) getRouteString() string {
	var readableTrade string
	for i, trade := range t.Trades {
		if i == 0 {
			readableTrade = fmt.Sprintf("%s %s -> %s", readableTrade, trade.From, trade.To)
		} else {
			readableTrade = fmt.Sprintf("%s -> %s", readableTrade, trade.To)
		}
	}
	return fmt.Sprintf("%s Profit: %f USD", readableTrade, t.Profit)
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

func (r RoutesWithProfit) getBestRouteString() string {
	sort.Slice(r, func(i, j int) bool {
		return r[i].Profit > r[j].Profit
	})
	return r[0].getRouteString()
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
	baseBudget := constants.BasePrice
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
		previousPrice = convertPrice(symbol, previousPrice, trade.From, trade.To, constants.Fee)

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
