package app

import (
	"fmt"
	"github.com/halm4d/arbitragecli/color"
	"time"
)

func Run() {
	fmt.Printf("Calculating possible routes...")
	symbols, usdtSymbols := NewSymbols()
	fmt.Println(len(symbols))
	fmt.Println(len(usdtSymbols))

	routes := CalculateAllRoutes(symbols)
	fmt.Printf("Found routes: %v\n", len(routes))

	for {
		symbols.updatePrices()
		usdtSymbols.updatePrices()

		startOfCalculation := time.Now()
		profitableRoutes, lossedRoutes := routes.getProfitableRoutes(symbols, usdtSymbols)
		endOfCalculation := time.Now()
		if len(profitableRoutes) != 0 {
			fmt.Printf("%sFound profitable routes: %v\n", color.Green, len(profitableRoutes))
			profitableRoutes.print(10)
		} else {
			fmt.Printf("%sProfitable route not found.\n", color.Red)
			lossedRoutes.print(2)
		}

		fmt.Printf("%sCalculation time: %v ms\n", color.Cyan, endOfCalculation.UnixMilli()-startOfCalculation.UnixMilli())
		time.Sleep(time.Second * 1)
	}

}
