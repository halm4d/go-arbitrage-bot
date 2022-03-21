package app

import (
	"fmt"
	"github.com/halm4d/arbitragecli/color"
	"github.com/halm4d/arbitragecli/constants"
	"time"
)

func Run() {
	fmt.Printf("Running arbbot with Fee: %v, BasePrice: %v, Verbose: %v\n", constants.Fee, constants.BasePrice, constants.Verbose)
	fmt.Printf("Calculating possible routes...\n")
	symbols, usdtSymbols := NewSymbols()

	routes := symbols.CalculateAllRoutes()
	fmt.Printf("Found routes: %v\n", len(*routes))

	for {
		symbols.updatePrices()
		usdtSymbols.updatePrices()

		startOfCalculation := time.Now()
		profitableRoutes, loosedRoutes := routes.getProfitableRoutes(*symbols, *usdtSymbols)
		endOfCalculation := time.Now()

		go func(profitableRoutes *RoutesWithProfit, loosedRoutes *RoutesWithProfit) {
			lenOfProfitableRoutes := len(*profitableRoutes)
			if constants.Verbose {
				if lenOfProfitableRoutes != 0 {
					fmt.Printf("%sFound profitable routes: %v\n", color.Green, lenOfProfitableRoutes)
					fmt.Printf("%sBest possible routes: \n", color.Green)
					profitableRoutes.print(10)
				} else {
					fmt.Printf("%sProfitable route not found yet. Best possible route was: %s\n", color.Red, loosedRoutes.getBestRouteString())
					loosedRoutes.print(10)
				}
			} else {
				if lenOfProfitableRoutes != 0 {
					fmt.Printf("%sFound profitable routes: %v\n", color.Green, lenOfProfitableRoutes)
					fmt.Printf("%sBest possible route was: %s\n", color.Green, profitableRoutes.getBestRouteString())
				} else {
					fmt.Printf("%sProfitable route not found yet. Best possible route was: %s\n", color.Red, loosedRoutes.getBestRouteString())
				}
			}
		}(profitableRoutes, loosedRoutes)

		fmt.Printf("%sCalculation time: %v ms\n", color.Cyan, endOfCalculation.UnixMilli()-startOfCalculation.UnixMilli())
		time.Sleep(time.Second * 1)
	}

}
