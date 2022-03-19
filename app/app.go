package app

import (
	"fmt"
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
	//busdSymbols := Symbols{
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
	//busdSymbols.updatePrices()

	symbols, busdSymbols := NewSymbols()
	fmt.Println(len(symbols))
	fmt.Println(len(busdSymbols))

	routes := CalculateAllRoutes(symbols, busdSymbols)
	fmt.Println(len(routes))
	routes.print(50)

	//for _, profit := range routes.CalculateProfits(testSymbols, busdSymbols) {
	//	fmt.Printf("%+v\n", profit)
	//}

}
