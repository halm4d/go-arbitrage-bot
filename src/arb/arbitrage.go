package arb

import (
	"github.com/halm4d/go-arbitrage-bot/src/color"
	"github.com/halm4d/go-arbitrage-bot/src/constants"
	"log"
	"sort"
	"sync"
	"time"
)

const (
	Buy Type = iota
	Sell
)

type Type int8

type Arbitrages []*Arbitrage

func New(symbols *Symbols) *Arbitrages {
	var arbs = make(Arbitrages, 0)

	var wg sync.WaitGroup
	mu := &sync.Mutex{}

	for _, startEndAsset := range AllCryptoCurrency() {
		wg.Add(1)
		go func(symbols *Symbols, startEndAsset string) {
			defer wg.Done()
			arbitrages := symbols.calculateArbsForSymbol(startEndAsset)
			mu.Lock()
			arbs = append(arbs, *arbitrages...)
			mu.Unlock()
		}(symbols, startEndAsset)
	}
	wg.Wait()
	return &arbs
}

func (a *Arbitrages) Run(bt *BookTickers) (profitableArbs Arbitrages) {
	for {
		time.Sleep(time.Second)
		startOfCalculation := time.Now()
		bt.MU.Lock()
		var cbt = make(BookTickerMap)
		for key, value := range bt.CryptoBookTickers {
			cbt[key] = value
		}
		var ubt = make(BookTickerMap)
		for key, value := range bt.USDTBookTickers {
			ubt[key] = value
		}
		bt.MU.Unlock()
		profitableArbs = a.CalculateProfits(&cbt, &ubt)
		logArbs(&profitableArbs, startOfCalculation)
	}
}

func logArbs(pr *Arbitrages, startOfCalculation time.Time) {
	bestRoute := pr.GetBestRoute()
	if constants.Verbose {
		if bestRoute.ProfitPercentage > 0 {
			log.Printf("%sCalculation time: %vms %sMost profitable route was:%s%s\n", color.Cyan, time.Now().UnixMilli()-startOfCalculation.UnixMilli(), color.Green, bestRoute.GetRouteString(), color.Reset)
			pr.Print(10)
		} else {
			log.Printf("%sCalculation time: %vms %sMost profitable route was:%s%s\n", color.Cyan, time.Now().UnixMilli()-startOfCalculation.UnixMilli(), color.Purple, bestRoute.GetRouteString(), color.Reset)
			pr.Print(10)
		}
	} else {
		if bestRoute.ProfitPercentage > 0 {
			log.Printf("%sCalculation time: %vms %sMost profitable route was:%s%s\n", color.Cyan, time.Now().UnixMilli()-startOfCalculation.UnixMilli(), color.Green, bestRoute.GetRouteString(), color.Reset)
		} else {
			log.Printf("%sCalculation time: %-3vms %sMost profitable route was:%s%s\n", color.Cyan, time.Now().UnixMilli()-startOfCalculation.UnixMilli(), color.Purple, bestRoute.GetRouteString(), color.Reset)
		}
	}
}

func (a *Arbitrages) CalculateProfits(bookTickerMap *BookTickerMap, usdtBookTicker *BookTickerMap) (profitableArbs Arbitrages) {
	var wg sync.WaitGroup
	mu := &sync.Mutex{}

	profitableArbs = make(Arbitrages, 0)
	for _, trades := range *a {
		wg.Add(1)
		go func(trades *Arbitrage, bookTickerMap *BookTickerMap, usdtBookTicker *BookTickerMap) {
			defer wg.Done()
			trades.Profit = trades.CalculateProfit(bookTickerMap, usdtBookTicker)
			trades.ProfitPercentage = trades.Profit * (100 / constants.BasePrice)
			mu.Lock()
			profitableArbs = append(profitableArbs, trades)
			mu.Unlock()
		}(trades, bookTickerMap, usdtBookTicker)
	}
	wg.Wait()
	sortArbs(profitableArbs)
	return profitableArbs
}

func sortArbs(profitableArbs Arbitrages) {
	var wg sync.WaitGroup
	wg.Add(1)
	go sortArb(profitableArbs, &wg)
	wg.Wait()
}

func sortArb(a Arbitrages, wg *sync.WaitGroup) {
	defer wg.Done()
	sort.Slice(a, func(i, j int) bool {
		return a[i].ProfitPercentage > a[j].ProfitPercentage
	})
	if len(a) > 10 {
		a = a[:10]
	}
}
