package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// MarketTick represents incoming tick from NSE/BSE binary feed
type MarketTick struct {
	Token         uint32  `json:"token"`
	LastPrice     float64 `json:"last_price"` // Price in Rupees (converted from paise)
	LastQuantity  uint32  `json:"last_quantity"`
	Volume        uint64  `json:"volume"`
	Open          float64 `json:"open"`
	High          float64 `json:"high"`
	Low           float64 `json:"low"`
	Close         float64 `json:"close"`
	TimestampSecs uint32  `json:"timestamp"`
}

// DecodeBinaryTick decodes raw 32-byte UDP packet
func DecodeBinaryTick(data []byte) (*MarketTick, error) {
	if len(data) < 32 {
		return nil, fmt.Errorf("insufficient packet length: %d", len(data))
	}

	reader := bytes.NewReader(data)

	var token uint32
	var pricePaise uint32
	var lastQty uint32
	var volume uint32
	var openPaise uint32
	var highPaise uint32
	var lowPaise uint32
	var timeSecs uint32

	binary.Read(reader, binary.BigEndian, &token)
	binary.Read(reader, binary.BigEndian, &pricePaise)
	binary.Read(reader, binary.BigEndian, &lastQty)
	binary.Read(reader, binary.BigEndian, &volume)
	binary.Read(reader, binary.BigEndian, &openPaise)
	binary.Read(reader, binary.BigEndian, &highPaise)
	binary.Read(reader, binary.BigEndian, &lowPaise)
	binary.Read(reader, binary.BigEndian, &timeSecs)

	return &MarketTick{
		Token:         token,
		LastPrice:     float64(pricePaise) / 100.0,
		LastQuantity:  lastQty,
		Volume:        uint64(volume),
		Open:          float64(openPaise) / 100.0,
		High:          float64(highPaise) / 100.0,
		Low:           float64(lowPaise) / 100.0,
		TimestampSecs: timeSecs,
	}, nil
}
