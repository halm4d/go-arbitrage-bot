package app

import (
	"fmt"
	"github.com/halm4d/arbitragecli/color"
	"time"
)

func Run() {
	//symbols := NewSymbols()
	//symbol, _ := symbols.findBySymbol("BTCBUSD")
	//fmt.Printf("%+v\n", symbol)
	//
	//time.Sleep(time.Second * 1)
	//
	//symbols.updatePrices()
	//symbol, _ = symbols.findBySymbol("BTCBUSD")
	//fmt.Printf("%+v\n", symbol)

	// BTC -> ETH -> LTC -> BTC
	// BTC -> LTC -> ETH -> BTC
	// LTC -> BTC -> ETH -> LTC
	// LTC -> ETH -> BTC -> LTC
	// ETH -> LTC -> BTC -> ETH
	// ETH -> BTC -> LTC -> ETH

	//symbols := Symbols{
	//	Symbol{Symbol: "ETHBTC", BaseAsset: "ETH", QuoteAsset: "BTC"},
	//	Symbol{Symbol: "LTCETH", BaseAsset: "LTC", QuoteAsset: "ETH"},
	//	Symbol{Symbol: "LTCBTC", BaseAsset: "LTC", QuoteAsset: "BTC"},
	//}
	//usdtSymbols := Symbols{
	//	Symbol{
	//		Symbol:     "BTCBUSD",
	//		BaseAsset:  "BTC",
	//		QuoteAsset: "BUSD",
	//	},
	//	Symbol{
	//		Symbol:     "LTCBUSD",
	//		BaseAsset:  "LTC",
	//		QuoteAsset: "BUSD",
	//	},
	//	Symbol{
	//		Symbol:     "ETHBUSD",
	//		BaseAsset:  "ETH",
	//		QuoteAsset: "BUSD",
	//	},
	//}
	//symbols.updatePrices()
	//usdtSymbols.updatePrices()

	symbols, usdtSymbols := NewSymbols()
	fmt.Println(len(symbols))
	fmt.Println(len(usdtSymbols))

	routes := CalculateAllRoutes(symbols)
	fmt.Printf("Found routes: %v\n", len(routes))

	for {
		symbols.updatePrices()
		usdtSymbols.updatePrices()

		profitableRoutes, lossedRoutes := routes.getProfitableRoutes(symbols, usdtSymbols)
		if len(profitableRoutes) != 0 {
			fmt.Printf("%sFound profitable routes: %v\n", color.Green, len(profitableRoutes))
			profitableRoutes.print(10)
		} else {
			fmt.Printf("%sProfitable route not found.\n", color.Red)
			lossedRoutes.print(2)
		}

		fmt.Printf("%sTime: %s\n", color.Reset, time.Now())
		time.Sleep(time.Second * 1)
	}

}
