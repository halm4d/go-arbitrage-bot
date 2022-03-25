package app

import (
	"fmt"
	"log"
	"sort"
)

type RoutesWithProfit []RouteWithProfit

type RouteWithProfit struct {
	Trades
	Profit float64
}

func (t *RouteWithProfit) getRouteString() string {
	var readableTrade string
	for i, trade := range t.Trades {
		if i == 0 {
			readableTrade = fmt.Sprintf("%s %s -> %s", readableTrade, trade.From, trade.To)
		} else {
			readableTrade = fmt.Sprintf("%s -> %s", readableTrade, trade.To)
		}
	}
	return fmt.Sprintf("%s Profit: %f USD", readableTrade, t.Profit)
}

func (t *RouteWithProfit) print() {
	var readableTrade string
	for i, trade := range t.Trades {
		if i == 0 {
			readableTrade = fmt.Sprintf("%s %s -> %s", readableTrade, trade.From, trade.To)
		} else {
			readableTrade = fmt.Sprintf("%s -> %s", readableTrade, trade.To)
		}
	}
	log.Printf("%s Profit: %f USD\n", readableTrade, t.Profit)
}

func (r RoutesWithProfit) getBestRouteString() string {
	sort.Slice(r, func(i, j int) bool {
		return r[i].Profit > r[j].Profit
	})
	return r[0].getRouteString()
}

func (r RoutesWithProfit) print(top int) {
	sort.Slice(r, func(i, j int) bool {
		return r[i].Profit > r[j].Profit
	})
	var limit int
	if len(r) > top {
		limit = top
	} else {
		limit = len(r)
	}
	for _, route := range r[:limit] {
		route.print()
	}
}
