package main

import (
	"testing"
	"time"
)

func BenchmarkOrderMatchingThroughput(b *testing.B) {
	ob := NewOrderBook("NIFTY_FUT")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ob.ProcessLimitOrder(&Order{
			ID:        int64(i),
			Symbol:    "NIFTY_FUT",
			Side:      Buy,
			Price:     22500.0,
			Quantity:  50,
			Timestamp: time.Now(),
		})
	}
}
