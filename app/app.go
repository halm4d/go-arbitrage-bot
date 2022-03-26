package app

import (
	"fmt"
	"github.com/halm4d/arbitragecli/arb"
	"github.com/halm4d/arbitragecli/client"
	"github.com/halm4d/arbitragecli/color"
	"github.com/halm4d/arbitragecli/constants"
	"log"
	"time"
)

func RunWebSocket() {
	fmt.Printf("Running arbbot with WEBSOCKET and Fee: %v, BasePrice: %v, Verbose: %v\n", constants.Fee, constants.BasePrice, constants.Verbose)

	symbols := arb.NewSymbols()
	symbols.Init(client.GetExchangeInfo())

	routes := arb.CalculateAllRoutes(&symbols.CS)
	fmt.Printf("Found routes: %v\n", len(*routes))

	client.RunWebSocket(symbols, func(bt *arb.BookTickers) {
		go findArbitrates(bt, routes)
	})
}

func findArbitrates(bt *arb.BookTickers, r *arb.Routes) {
	for {
		time.Sleep(time.Second * 1)
		startOfCalculation := time.Now()
		bt.MU.Lock()
		var cbt = make(arb.BookTickerMap)
		for key, value := range bt.CryptoBookTickers {
			cbt[key] = value
		}
		var ubt = make(arb.BookTickerMap)
		for key, value := range bt.USDTBookTickers {
			ubt[key] = value
		}
		bt.MU.Unlock()
		profitableRoutes, unProfitableRoutes := r.GetProfitableRoutes(&cbt, &ubt)
		logRoutes(&profitableRoutes, &unProfitableRoutes, startOfCalculation)
	}
}

func logRoutes(pr *arb.Routes, upr *arb.Routes, startOfCalculation time.Time) {
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
