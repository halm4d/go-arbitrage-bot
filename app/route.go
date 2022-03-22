package app

import (
	"sort"
	"sync"
)

const (
	Buy Type = iota
	Sell
)

type Type int8

type Routes []Trades

func (s *Symbols) calculateAllRoutes() *Routes {
	var routes = make(Routes, 0)

	var wg sync.WaitGroup
	mu := &sync.Mutex{}

	for _, startEndAsset := range AllCryptoCurrency() {
		wg.Add(1)
		go func(symbols *Symbols, startEndAsset string) {
			defer wg.Done()
			OneOfMyStructs := symbols.calculateRoutesForSymbol(startEndAsset)
			mu.Lock()
			routes = append(routes, *OneOfMyStructs...)
			mu.Unlock()
		}(s, startEndAsset)
	}
	wg.Wait()
	return &routes
}

func (s *Symbols) calculateRoutesForSymbol(startEndAsset string) *Routes {
	var routes Routes
	for _, asset1 := range *s.findAllByAsset(startEndAsset) {
		var targetAsset1 = asset1.getTargetAsset(startEndAsset)
		for _, asset2 := range *s.findAllByAsset(targetAsset1) {
			targetAsset2 := asset2.getTargetAsset(targetAsset1)
			if targetAsset2 == startEndAsset {
				continue
			}
			pair, err := s.findByAssetPair(targetAsset2, startEndAsset)
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
	return &routes
}

func (r *Routes) getProfitableRoutes(symbols *Symbols, usdtSymbols *Symbols) (*RoutesWithProfit, *RoutesWithProfit) {
	var profitableRoutes = make(RoutesWithProfit, 0)
	var routesWithLoss = make(RoutesWithProfit, 0)

	var wg sync.WaitGroup
	mu := &sync.Mutex{}

	for _, trades := range *r {
		wg.Add(1)
		go func(trades Trades, symbols *Symbols, usdtSymbols *Symbols) {
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

	sortRoutes(profitableRoutes, routesWithLoss)
	return &profitableRoutes, &routesWithLoss
}

func sortRoutes(profitableRoutes RoutesWithProfit, routesWithLoss RoutesWithProfit) {
	var wg sync.WaitGroup
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
}
