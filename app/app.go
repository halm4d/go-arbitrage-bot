package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/halm4d/arbitragecli/color"
	"github.com/halm4d/arbitragecli/constants"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

func Run() {
	//fmt.Printf("Running arbbot with Fee: %v, BasePrice: %v, Verbose: %v\n", constants.Fee, constants.BasePrice, constants.Verbose)
	//fmt.Printf("Calculating possible routes...\n")
	//symbols, usdtSymbols := NewSymbols()
	//
	//routes := symbols.calculateAllRoutes()
	//fmt.Printf("Found routes: %v\n", len(*routes))
	//
	//for {
	//	symbols.updatePrices()
	//	usdtSymbols.updatePrices()
	//
	//	startOfCalculation := time.Now()
	//	profitableRoutes, loosedRoutes := routes.getProfitableRoutes(symbols, usdtSymbols)
	//	endOfCalculation := time.Now()
	//
	//	go func(profitableRoutes *RoutesWithProfit, loosedRoutes *RoutesWithProfit) {
	//		lenOfProfitableRoutes := len(*profitableRoutes)
	//		if constants.Verbose {
	//			if lenOfProfitableRoutes != 0 {
	//				fmt.Printf("%sFound profitable routes: %v\n", color.Green, lenOfProfitableRoutes)
	//				fmt.Printf("%sBest possible routes: \n", color.Green)
	//				profitableRoutes.print(10)
	//			} else {
	//				fmt.Printf("%sProfitable route not found yet. Best possible route was: %s\n", color.Red, loosedRoutes.getBestRouteString())
	//				loosedRoutes.print(10)
	//			}
	//		} else {
	//			if lenOfProfitableRoutes != 0 {
	//				fmt.Printf("%sFound profitable routes: %v\n", color.Green, lenOfProfitableRoutes)
	//				fmt.Printf("%sBest possible route was: %s\n", color.Green, profitableRoutes.getBestRouteString())
	//			} else {
	//				fmt.Printf("%sProfitable route not found yet. Best possible route was: %s\n", color.Red, loosedRoutes.getBestRouteString())
	//			}
	//		}
	//	}(profitableRoutes, loosedRoutes)
	//
	//	fmt.Printf("%sCalculation time: %v ms\n", color.Cyan, endOfCalculation.UnixMilli()-startOfCalculation.UnixMilli())
	//	time.Sleep(time.Second * 1)
	//}
}

var symbols, usdtSymbols = NewSymbols()

func RunWebSocket() {
	fmt.Printf("Running arbbot with WEBSOCKET and Fee: %v, BasePrice: %v, Verbose: %v\n", constants.Fee, constants.BasePrice, constants.Verbose)
	//fmt.Printf("Calculating possible routes...\n")
	//symbols, usdtSymbols := NewSymbols()
	//
	//routes := symbols.calculateAllRoutes()
	//fmt.Printf("Found routes: %v\n", len(*routes))
	done = make(chan interface{})    // Channel to indicate that the receiverHandler is done
	interrupt = make(chan os.Signal) // Channel to listen for interrupt signal to terminate gracefully

	signal.Notify(interrupt, os.Interrupt) // Notify the interrupt channel for SIGINT

	socketUrl := "wss://stream.binance.com:9443" + "/ws/!bookTicker"
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		log.Fatal("Error connecting to Websocket Server:", err)
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	routes := CalculateAllRoutes(&symbols)
	fmt.Printf("Found routes: %v\n", len(*routes))

	go receiveHandler(conn)
	go findArbitrates(routes)
	go printMap()

	// Our main loop for the client
	// We send our relevant packets here
	for {
		select {
		case <-interrupt:
			// We received a SIGINT (Ctrl + C). Terminate gracefully...
			log.Println("Received SIGINT interrupt signal. Closing all pending connections")
			// Close our websocket connection
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Error during closing websocket:", err)
				return
			}
			select {
			case <-done:
				log.Println("Receiver Channel Closed! Exiting....")
			case <-time.After(time.Duration(1) * time.Second):
				log.Println("Timeout in closing receiving channel. Exiting....")
			}
			return
		}
	}
}

func findArbitrates(routes *Routes) {
	for {
		time.Sleep(time.Second * 1)
		startOfCalculation := time.Now()
		mu.Lock()
		var bookTickerMapCopy = make(map[string]BookTicker)
		for key, value := range bookTickerMap {
			bookTickerMapCopy[key] = value
		}
		mu.Unlock()
		profitableRoutes, loosedRoutes := routes.getProfitableRoutes(&bookTickerMapCopy)
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
	}
}

var done chan interface{}
var interrupt chan os.Signal

func printMap() {
	for {
		time.Sleep(time.Second)
		fmt.Printf("len: %v btcusdt: %+v btceth: %+v\n", len(bookTickerMap), bookTickerMap["BTCUSDT"], bookTickerMap["ETHBTC"])
	}
}

func receiveHandler(connection *websocket.Conn) {
	defer close(done)
	for {
		_, msg, err := connection.ReadMessage()
		if err != nil {
			log.Println("Error in receive:", err)
			return
		}
		//log.Printf("Received: %s\n", msg)

		var target BookTickerSocketResp
		err = json.Unmarshal(msg, &target)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		//log.Printf("Received: %+v\n", target)
		symbol, ok := symbols[target.S]
		if !ok {
			symbol, ok = usdtSymbols[target.S]
			if !ok {
				continue
			}
		}
		bidPrice, _ := strconv.ParseFloat(target.B, 64)
		bidQty, _ := strconv.ParseFloat(target.B1, 64)
		askPrice, _ := strconv.ParseFloat(target.A, 64)
		askQty, _ := strconv.ParseFloat(target.A1, 64)
		mu.Lock()
		bookTickerMap[target.S] = BookTicker{
			Symbol:     target.S,
			BaseAsset:  symbol.BaseAsset,
			QuoteAsset: symbol.QuoteAsset,
			BidPrice:   bidPrice,
			BidQty:     bidQty,
			AskPrice:   askPrice,
			AskQty:     askQty,
		}
		mu.Unlock()
	}
}

var mu = sync.Mutex{}
var bookTickerMap = make(map[string]BookTicker)

type BookTicker struct {
	Symbol     string
	BaseAsset  string
	QuoteAsset string
	BidPrice   float64
	BidQty     float64
	AskPrice   float64
	AskQty     float64
}

type BookTickerSocketResp struct {
	U  int    `json:"u"`
	S  string `json:"s"`
	B  string `json:"b"`
	B1 string `json:"B"`
	A  string `json:"a"`
	A1 string `json:"A"`
}
