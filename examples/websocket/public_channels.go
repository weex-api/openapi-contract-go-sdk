// WebSocket Public Channels Example
//
// This example demonstrates how to subscribe to public WebSocket channels
// for real-time market data (ticker, depth, candlesticks, trades).
//
// Run: go run examples/websocket/public_channels.go

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/weex/openapi-contract-go-sdk/weex"
	"github.com/weex/openapi-contract-go-sdk/weex/websocket"
	"github.com/weex/openapi-contract-go-sdk/weex/websocket/public"
)

func main() {
	// Create configuration for public channels
	config := weex.NewDefaultConfig().
		WithLogLevel(weex.LogLevelInfo)

	// Create public WebSocket client
	client := public.NewClient(config)

	// Set up connection callbacks
	client.SetOnConnect(func() {
		fmt.Println("✅ WebSocket connected successfully!")
	})

	client.SetOnDisconnect(func(err error) {
		if err != nil {
			fmt.Printf("⚠️  WebSocket disconnected: %v\n", err)
		} else {
			fmt.Println("WebSocket disconnected")
		}
	})

	client.SetOnError(func(err error) {
		fmt.Printf("❌ WebSocket error: %v\n", err)
	})

	// Connect to WebSocket
	ctx := context.Background()
	if err := client.Connect(ctx); err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	fmt.Println("🔌 Connecting to WEEX WebSocket...")

	// ===== Example 1: Subscribe to Ticker =====
	fmt.Println("\n=== Example 1: Subscribe to Ticker (BTC/USDT) ===")
	err := client.SubscribeTicker("cmt_btcusdt", func(ticker *websocket.TickerData) error {
		if len(ticker.Data) > 0 {
			t := ticker.Data[0]
			fmt.Printf("📊 [TICKER] %s: Price=%s, 24h Change=%s%%, Volume=%s\n",
				t.Symbol, t.LastPrice, t.PriceChangePercent, t.Volume24h)
		}
		return nil
	})
	if err != nil {
		log.Printf("Failed to subscribe to ticker: %v", err)
	}

	// ===== Example 2: Subscribe to Order Book Depth =====
	fmt.Println("\n=== Example 2: Subscribe to Depth (BTC/USDT) ===")
	err = client.SubscribeDepth("cmt_btcusdt", func(depth *websocket.DepthData) error {
		if len(depth.Data) > 0 {
			d := depth.Data[0]
			fmt.Printf("📖 [DEPTH] %s: Best Bid=%s @ %s, Best Ask=%s @ %s\n",
				d.Symbol,
				d.Bids[0].Quantity, d.Bids[0].Price,
				d.Asks[0].Quantity, d.Asks[0].Price)
		}
		return nil
	})
	if err != nil {
		log.Printf("Failed to subscribe to depth: %v", err)
	}

	// ===== Example 3: Subscribe to Candlesticks (1 minute) =====
	fmt.Println("\n=== Example 3: Subscribe to Candlesticks (BTC/USDT, 1m) ===")
	err = client.SubscribeCandlestick("cmt_btcusdt", "1m", func(kline *websocket.CandlestickData) error {
		if len(kline.Data) > 0 {
			k := kline.Data[0]
			fmt.Printf("🕯️  [KLINE] %s %s: O=%s H=%s L=%s C=%s V=%s\n",
				k.Symbol, k.Interval, k.Open, k.High, k.Low, k.Close, k.Volume)
		}
		return nil
	})
	if err != nil {
		log.Printf("Failed to subscribe to candlestick: %v", err)
	}

	// ===== Example 4: Subscribe to Recent Trades =====
	fmt.Println("\n=== Example 4: Subscribe to Trades (BTC/USDT) ===")
	err = client.SubscribeTrades("cmt_btcusdt", func(trades *websocket.TradesData) error {
		for _, trade := range trades.Data {
			fmt.Printf("💱 [TRADE] %s: %s %s @ %s (ID: %s)\n",
				trade.Symbol, trade.Side, trade.Size, trade.Price, trade.TradeId)
		}
		return nil
	})
	if err != nil {
		log.Printf("Failed to subscribe to trades: %v", err)
	}

	// ===== Example 5: Subscribe to Multiple Symbols =====
	fmt.Println("\n=== Example 5: Subscribe to ETH Ticker ===")
	err = client.SubscribeTicker("cmt_ethusdt", func(ticker *websocket.TickerData) error {
		if len(ticker.Data) > 0 {
			t := ticker.Data[0]
			fmt.Printf("📊 [TICKER] %s: Price=%s, 24h Change=%s%%\n",
				t.Symbol, t.LastPrice, t.PriceChangePercent)
		}
		return nil
	})
	if err != nil {
		log.Printf("Failed to subscribe to ETH ticker: %v", err)
	}

	fmt.Println("\n✅ All subscriptions active. Press Ctrl+C to exit.")
	fmt.Println("📡 Listening for real-time updates...\n")

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	fmt.Println("\n\n🛑 Shutting down...")

	// Cleanup: Unsubscribe from all channels
	fmt.Println("Unsubscribing from channels...")
	client.UnsubscribeTicker("cmt_btcusdt")
	client.UnsubscribeDepth("cmt_btcusdt")
	client.UnsubscribeCandlestick("cmt_btcusdt", "1m")
	client.UnsubscribeTrades("cmt_btcusdt")
	client.UnsubscribeTicker("cmt_ethusdt")

	// Close connection
	if err := client.Close(); err != nil {
		log.Printf("Error closing connection: %v", err)
	}

	fmt.Println("👋 Goodbye!")
}
