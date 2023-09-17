package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Stock Trading Engine (Go Matching Core) Initialized")
	ob := NewOrderBook("RELIANCE")

	// Test cross matching
	ob.ProcessLimitOrder(&Order{ID: 1, Symbol: "RELIANCE", Side: Sell, Price: 2450.00, Quantity: 100, Timestamp: time.Now()})
	trades := ob.ProcessLimitOrder(&Order{ID: 2, Symbol: "RELIANCE", Side: Buy, Price: 2450.00, Quantity: 60, Timestamp: time.Now()})

	for _, t := range trades {
		fmt.Printf("Executed Trade: %d shares @ ₹%.2f\n", t.Quantity, t.Price)
	}
}
