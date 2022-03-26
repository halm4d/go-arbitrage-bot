package arb

import (
	"github.com/halm4d/arbitragecli/constants"
	"math"
	"sort"
	"sync"
)

const (
	Buy Type = iota
	Sell
)

type Type int8

type Routes []Trades

func CalculateAllRoutes(symbols *SymbolsMap) *Routes {
	var routes = make(Routes, 0)

	var wg sync.WaitGroup
	mu := &sync.Mutex{}

	for _, startEndAsset := range AllCryptoCurrency() {
		wg.Add(1)
		go func(symbols *SymbolsMap, startEndAsset string) {
			defer wg.Done()
			OneOfMyStructs := calculateRoutesForSymbol(symbols, startEndAsset)
			mu.Lock()
			routes = append(routes, *OneOfMyStructs...)
			mu.Unlock()
		}(symbols, startEndAsset)
	}
	wg.Wait()
	return &routes
}

func calculateRoutesForSymbol(symbols *SymbolsMap, startEndAsset string) *Routes {
	var routes Routes
	for _, asset1 := range *symbols.FindAllSymbolByAsset(startEndAsset) {
		var targetAsset1 = asset1.GetTargetAsset(startEndAsset)
		if asset1.BaseAsset == constants.USDT || asset1.QuoteAsset == constants.USDT {
			continue
		}
		for _, asset2 := range *symbols.FindAllSymbolByAsset(targetAsset1) {
			if asset2.BaseAsset == constants.USDT || asset2.QuoteAsset == constants.USDT {
				continue
			}
			targetAsset2 := asset2.GetTargetAsset(targetAsset1)
			if targetAsset2 == startEndAsset {
				continue
			}
			pair, err := symbols.FindByAssetPair(targetAsset2, startEndAsset)
			if err != nil {
				continue
			}
			trades := Trades{
				Trades: []Trade{
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
				},
				Profit: -math.MaxFloat64,
			}
			routes = append(routes, trades)
		}
	}
	return &routes
}

func (r *Routes) GetProfitableRoutes(bookTickerMap *BookTickerMap, usdtBookTicker *BookTickerMap) (profitableRoutes Routes, unProfitableRoutes Routes) {
	var wg sync.WaitGroup
	mu := &sync.Mutex{}

	profitableRoutes = []Trades{}
	unProfitableRoutes = []Trades{}
	for _, trades := range *r {
		wg.Add(1)
		go func(trades Trades, bookTickerMap *BookTickerMap, usdtBookTicker *BookTickerMap) {
			defer wg.Done()
			trades.Profit = trades.CalculateProfit(bookTickerMap, usdtBookTicker)
			if trades.Profit > 0 {
				mu.Lock()
				profitableRoutes = append(profitableRoutes, trades)
				mu.Unlock()
			} else {
				mu.Lock()
				unProfitableRoutes = append(unProfitableRoutes, trades)
				mu.Unlock()
			}
		}(trades, bookTickerMap, usdtBookTicker)
	}
	wg.Wait()
	sortRoutes(profitableRoutes, unProfitableRoutes)
	return profitableRoutes, unProfitableRoutes
}

func sortRoutes(profitableRoutes Routes, unProfitableRoutes Routes) {
	var wg sync.WaitGroup
	wg.Add(2)
	go sortRoute(profitableRoutes, &wg)
	go sortRoute(unProfitableRoutes, &wg)
	wg.Wait()
}

func sortRoute(routes Routes, wg *sync.WaitGroup) {
	defer wg.Done()
	sort.Slice(routes, func(i, j int) bool {
		return routes[i].Profit > routes[j].Profit
	})
	if len(routes) > 10 {
		routes = routes[:10]
	}
}
