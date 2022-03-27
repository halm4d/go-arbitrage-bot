package arb

import (
	"github.com/halm4d/arbitragecli/color"
	"github.com/halm4d/arbitragecli/constants"
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

type Arbitrages []Arbitrage

func New(symbols *Symbols) *Arbitrages {
	var arbs = make(Arbitrages, 0)

	var wg sync.WaitGroup
	mu := &sync.Mutex{}

	for _, startEndAsset := range AllCryptoCurrency() {
		wg.Add(1)
		go func(symbols *Symbols, startEndAsset string) {
			defer wg.Done()
			OneOfMyStructs := symbols.CS.calculateArbsForSymbol(startEndAsset)
			mu.Lock()
			arbs = append(arbs, *OneOfMyStructs...)
			mu.Unlock()
		}(symbols, startEndAsset)
	}
	wg.Wait()
	return &arbs
}

func (a *Arbitrages) Run(bt *BookTickers, d time.Duration) (profitableArbs Arbitrages, unProfitableArbs Arbitrages) {
	for {
		time.Sleep(d)
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
		profitableArbs, unProfitableArbs = a.CalculateProfits(&cbt, &ubt)
		logArbs(&profitableArbs, &unProfitableArbs, startOfCalculation)
	}
}

func logArbs(pr *Arbitrages, upr *Arbitrages, startOfCalculation time.Time) {
	if constants.Verbose {
		if len(*pr) > 0 {
			log.Printf("%sCalculation time: %v ms%s %sFound profitable route: %v%s\n", color.Cyan, time.Now().UnixMilli()-startOfCalculation.UnixMilli(), color.Reset, color.Green, len(*pr), color.Reset)
			log.Printf("%sMost profitable route was:%s%s\n", color.Green, pr.GetBestRouteString(), color.Reset)
			pr.Print(10)
		} else {
			log.Printf("%sCalculation time: %v ms%s %sProfitable route not found yet. Best possible route was:%s%s\n", color.Cyan, time.Now().UnixMilli()-startOfCalculation.UnixMilli(), color.Reset, color.Purple, upr.GetBestRouteString(), color.Reset)
			pr.Print(10)
		}
	} else {
		if len(*pr) > 0 {
			log.Printf("%sCalculation time: %v ms%s %sFound profitable route: %v%s\n", color.Cyan, time.Now().UnixMilli()-startOfCalculation.UnixMilli(), color.Reset, color.Green, len(*pr), color.Reset)
			log.Printf("%sMost profitable route was:%s%s\n", color.Green, pr.GetBestRouteString(), color.Reset)
		} else {
			log.Printf("%sCalculation time: %v ms%s %sProfitable route not found yet. Most profitable route was:%s%s\n", color.Cyan, time.Now().UnixMilli()-startOfCalculation.UnixMilli(), color.Reset, color.Purple, upr.GetBestRouteString(), color.Reset)
		}
	}
}

func (a *Arbitrages) CalculateProfits(bookTickerMap *BookTickerMap, usdtBookTicker *BookTickerMap) (profitableArbs Arbitrages, unProfitableArbs Arbitrages) {
	var wg sync.WaitGroup
	mu := &sync.Mutex{}

	profitableArbs = []Arbitrage{}
	unProfitableArbs = []Arbitrage{}
	for _, trades := range *a {
		wg.Add(1)
		go func(trades Arbitrage, bookTickerMap *BookTickerMap, usdtBookTicker *BookTickerMap) {
			defer wg.Done()
			trades.Profit = trades.CalculateProfit(bookTickerMap, usdtBookTicker)
			if trades.Profit > 0 {
				mu.Lock()
				profitableArbs = append(profitableArbs, trades)
				mu.Unlock()
			} else {
				mu.Lock()
				unProfitableArbs = append(unProfitableArbs, trades)
				mu.Unlock()
			}
		}(trades, bookTickerMap, usdtBookTicker)
	}
	wg.Wait()
	sortArbs(profitableArbs, unProfitableArbs)
	return profitableArbs, unProfitableArbs
}

func sortArbs(profitableArbs Arbitrages, unProfitableArbs Arbitrages) {
	var wg sync.WaitGroup
	wg.Add(2)
	go sortArb(profitableArbs, &wg)
	go sortArb(unProfitableArbs, &wg)
	wg.Wait()
}

func sortArb(a Arbitrages, wg *sync.WaitGroup) {
	defer wg.Done()
	sort.Slice(a, func(i, j int) bool {
		return a[i].Profit > a[j].Profit
	})
	if len(a) > 10 {
		a = a[:10]
	}
}
