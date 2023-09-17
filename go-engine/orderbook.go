package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type Side string

const (
	Buy  Side = "BUY"
	Sell Side = "SELL"
)

type Order struct {
	ID        int64     `json:"id"`
	Symbol    string    `json:"symbol"`
	Side      Side      `json:"side"`
	Price     float64   `json:"price"`
	Quantity  int64     `json:"quantity"`
	Timestamp time.Time `json:"timestamp"`
}

type Trade struct {
	BuyOrderID  int64     `json:"buy_order_id"`
	SellOrderID int64     `json:"sell_order_id"`
	Symbol      string    `json:"symbol"`
	Price       float64   `json:"price"`
	Quantity    int64     `json:"quantity"`
	Timestamp   time.Time `json:"timestamp"`
}

type OrderBook struct {
	mu     sync.RWMutex
	Symbol string
	Bids   []*Order // Sorted descending by price
	Asks   []*Order // Sorted ascending by price
}

func NewOrderBook(symbol string) *OrderBook {
	return &OrderBook{
		Symbol: symbol,
		Bids:   make([]*Order, 0),
		Asks:   make([]*Order, 0),
	}
}

// ProcessLimitOrder matches incoming order against opposite book or rests it
func (ob *OrderBook) ProcessLimitOrder(order *Order) []Trade {
	ob.mu.Lock()
	defer ob.mu.Unlock()

	trades := make([]Trade, 0)

	if order.Side == Buy {
		for len(ob.Asks) > 0 && order.Quantity > 0 {
			bestAsk := ob.Asks[0]
			if order.Price < bestAsk.Price {
				break // Cannot cross spread
			}

			matchedQty := min(order.Quantity, bestAsk.Quantity)
			trades = append(trades, Trade{
				BuyOrderID:  order.ID,
				SellOrderID: bestAsk.ID,
				Symbol:      ob.Symbol,
				Price:       bestAsk.Price,
				Quantity:    matchedQty,
				Timestamp:   time.Now(),
			})

			order.Quantity -= matchedQty
			bestAsk.Quantity -= matchedQty

			if bestAsk.Quantity == 0 {
				ob.Asks = ob.Asks[1:]
			}
		}

		if order.Quantity > 0 {
			ob.Bids = append(ob.Bids, order)
			sort.Slice(ob.Bids, func(i, j int) bool {
				if ob.Bids[i].Price == ob.Bids[j].Price {
					return ob.Bids[i].Timestamp.Before(ob.Bids[j].Timestamp)
				}
				return ob.Bids[i].Price > ob.Bids[j].Price
			})
		}
	} else {
		for len(ob.Bids) > 0 && order.Quantity > 0 {
			bestBid := ob.Bids[0]
			if order.Price > bestBid.Price {
				break
			}

			matchedQty := min(order.Quantity, bestBid.Quantity)
			trades = append(trades, Trade{
				BuyOrderID:  bestBid.ID,
				SellOrderID: order.ID,
				Symbol:      ob.Symbol,
				Price:       bestBid.Price,
				Quantity:    matchedQty,
				Timestamp:   time.Now(),
			})

			order.Quantity -= matchedQty
			bestBid.Quantity -= matchedQty

			if bestBid.Quantity == 0 {
				ob.Bids = ob.Bids[1:]
			}
		}

		if order.Quantity > 0 {
			ob.Asks = append(ob.Asks, order)
			sort.Slice(ob.Asks, func(i, j int) bool {
				if ob.Asks[i].Price == ob.Asks[j].Price {
					return ob.Asks[i].Timestamp.Before(ob.Asks[j].Timestamp)
				}
				return ob.Asks[i].Price < ob.Asks[j].Price
			})
		}
	}

	return trades
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
