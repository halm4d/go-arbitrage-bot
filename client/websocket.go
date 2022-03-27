package client

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/halm4d/arbitragecli/arb"
	"github.com/halm4d/arbitragecli/constants"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

var done chan interface{}
var interrupt chan os.Signal

var bt *arb.BookTickers

func RunWebSocket(symbols *arb.Symbols, fn func(bt *arb.BookTickers)) {
	done = make(chan interface{})    // Channel to indicate that the receiverHandler is done
	interrupt = make(chan os.Signal) // Channel to listen for interrupt signal to terminate gracefully
	bt = &arb.BookTickers{
		MU:                sync.Mutex{},
		USDTBookTickers:   make(arb.BookTickerMap),
		CryptoBookTickers: make(arb.BookTickerMap),
	}

	signal.Notify(interrupt, os.Interrupt) // Notify the interrupt channel for SIGINT

	socketUrl := "wss://stream.binance.com:9443/ws/!bookTicker"
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

	go receiveHandler(conn, symbols)
	go fn(bt)

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

func receiveHandler(connection *websocket.Conn, symbols *arb.Symbols) {
	defer close(done)
	for {
		_, msg, err := connection.ReadMessage()
		if err != nil {
			log.Println("Error in receive:", err)
			return
		}

		var target arb.BookTickerSocketResp
		err = json.Unmarshal(msg, &target)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		symbol, ok := symbols.Symbols[target.S]
		if !ok {
			continue
		}
		if symbol.QuoteAsset == constants.USDT || symbol.BaseAsset == constants.USDT {
			bidPrice, _ := strconv.ParseFloat(target.B, 64)
			bidQty, _ := strconv.ParseFloat(target.B1, 64)
			askPrice, _ := strconv.ParseFloat(target.A, 64)
			askQty, _ := strconv.ParseFloat(target.A1, 64)
			bt.MU.Lock()
			bt.USDTBookTickers[target.S] = &arb.BookTicker{
				Symbol:     target.S,
				BaseAsset:  symbol.BaseAsset,
				QuoteAsset: symbol.QuoteAsset,
				BidPrice:   bidPrice,
				BidQty:     bidQty,
				AskPrice:   askPrice,
				AskQty:     askQty,
			}
			bt.MU.Unlock()
		} else {
			bidPrice, _ := strconv.ParseFloat(target.B, 64)
			bidQty, _ := strconv.ParseFloat(target.B1, 64)
			askPrice, _ := strconv.ParseFloat(target.A, 64)
			askQty, _ := strconv.ParseFloat(target.A1, 64)
			bt.MU.Lock()
			bt.CryptoBookTickers[target.S] = &arb.BookTicker{
				Symbol:     target.S,
				BaseAsset:  symbol.BaseAsset,
				QuoteAsset: symbol.QuoteAsset,
				BidPrice:   bidPrice,
				BidQty:     bidQty,
				AskPrice:   askPrice,
				AskQty:     askQty,
			}
			bt.MU.Unlock()
		}
	}
}
