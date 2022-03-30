package app

import (
	"fmt"
	"github.com/halm4d/go-arbitrage-bot/src/arb"
	"github.com/halm4d/go-arbitrage-bot/src/client"
	"github.com/halm4d/go-arbitrage-bot/src/constants"
)

func RunWebSocket() {
	fmt.Printf("Running arbbot with WEBSOCKET and Fee: %v, BasePrice: %v, Verbose: %v\n", constants.Fee, constants.BasePrice, constants.Verbose)

	symbols := arb.NewSymbols()
	symbols.Init(client.GetExchangeInfo())

	arbs := arb.New(symbols)
	fmt.Printf("Found arbs: %v\n", len(*arbs))

	client.RunWebSocket(symbols, func(bt *arb.BookTickers) {
		go arbs.Run(bt)
	})
}
