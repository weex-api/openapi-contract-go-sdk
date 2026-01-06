// WebSocket Private Channels Example
//
// This example demonstrates how to subscribe to private WebSocket channels
// for real-time account, position, order, and fill updates.
//
// ⚠️ WARNING: This requires valid API credentials!
//
// Run: go run examples/websocket/private_channels.go

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/weex-api/openapi-contract-go-sdk/weex"
	"github.com/weex-api/openapi-contract-go-sdk/weex/websocket"
	"github.com/weex-api/openapi-contract-go-sdk/weex/websocket/private"
)

func main() {
	// ⚠️ IMPORTANT: Replace with your actual API credentials
	// Or use environment variables for security
	apiKey := os.Getenv("WEEX_API_KEY")
	secretKey := os.Getenv("WEEX_SECRET_KEY")
	passphrase := os.Getenv("WEEX_PASSPHRASE")

	if apiKey == "" || secretKey == "" || passphrase == "" {
		log.Fatal("❌ Please set WEEX_API_KEY, WEEX_SECRET_KEY, and WEEX_PASSPHRASE environment variables")
	}

	// Create configuration with API credentials
	config := weex.NewDefaultConfig().
		WithAPIKey(apiKey).
		WithSecretKey(secretKey).
		WithPassphrase(passphrase).
		WithLogLevel(weex.LogLevelInfo)

	// Create authenticator
	auth := weex.NewAuthenticator(apiKey, secretKey, passphrase)

	// Create private WebSocket client
	client := private.NewClient(config, auth)

	// Set up connection callbacks
	client.SetOnConnect(func() {
		fmt.Println("✅ WebSocket connected and authenticated successfully!")
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

	// Connect to WebSocket (this will also authenticate)
	ctx := context.Background()
	fmt.Println("🔌 Connecting to WEEX Private WebSocket...")
	if err := client.Connect(ctx); err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	// ===== Example 1: Subscribe to Account Balance Updates =====
	fmt.Println("\n=== Example 1: Subscribe to Account Balance Updates ===")
	err := client.SubscribeAccount(func(account *websocket.AccountData) error {
		fmt.Printf("\n💰 [ACCOUNT UPDATE]\n")
		for _, asset := range account.Data {
			fmt.Printf("  %s: Available=%s, Frozen=%s, Equity=%s, UnrealizedPnL=%s\n",
				asset.CoinName, asset.Available, asset.Frozen, asset.Equity, asset.UnrealizedPnl)
		}
		return nil
	})
	if err != nil {
		log.Printf("Failed to subscribe to account: %v", err)
	}

	// ===== Example 2: Subscribe to Position Updates =====
	fmt.Println("\n=== Example 2: Subscribe to Position Updates ===")
	err = client.SubscribePositions(func(position *websocket.PositionData) error {
		fmt.Printf("\n📈 [POSITION UPDATE]\n")
		for _, pos := range position.Data {
			fmt.Printf("  %s [%s]: Size=%s, AvgPrice=%s, MarkPrice=%s\n",
				pos.Symbol, pos.PositionSide, pos.Size, pos.AverageOpenPrice, pos.MarkPrice)
			fmt.Printf("    UnrealizedPnL=%s, Leverage=%sx, LiquidatePrice=%s\n",
				pos.UnrealizedPnl, pos.Leverage, pos.LiquidatePrice)
		}
		return nil
	})
	if err != nil {
		log.Printf("Failed to subscribe to positions: %v", err)
	}

	// ===== Example 3: Subscribe to Order Updates =====
	fmt.Println("\n=== Example 3: Subscribe to Order Updates ===")
	err = client.SubscribeOrders(func(order *websocket.OrderData) error {
		fmt.Printf("\n📋 [ORDER UPDATE]\n")
		for _, o := range order.Data {
			stateStr := getOrderStateString(o.State)
			fmt.Printf("  Order: %s (ClientOID: %s)\n", o.OrderId, o.ClientOid)
			fmt.Printf("    Symbol: %s, Type: %d, Side: %s\n", o.Symbol, o.Type, o.Side)
			fmt.Printf("    Price: %s, Size: %s, Filled: %s, State: %s\n",
				o.Price, o.Size, o.FilledSize, stateStr)
			if !o.RealizedPnl.IsZero() {
				fmt.Printf("    RealizedPnL: %s\n", o.RealizedPnl)
			}
		}
		return nil
	})
	if err != nil {
		log.Printf("Failed to subscribe to orders: %v", err)
	}

	// ===== Example 4: Subscribe to Fill/Execution Updates =====
	fmt.Println("\n=== Example 4: Subscribe to Fill (Execution) Updates ===")
	err = client.SubscribeFills(func(fill *websocket.FillData) error {
		fmt.Printf("\n✅ [FILL/EXECUTION]\n")
		for _, f := range fill.Data {
			fmt.Printf("  FillID: %s (OrderID: %s)\n", f.FillId, f.OrderId)
			fmt.Printf("    Symbol: %s, Side: %s, Liquidity: %s\n", f.Symbol, f.Side, f.Liquidity)
			fmt.Printf("    Price: %s, Size: %s, Fee: %s %s\n",
				f.Price, f.Size, f.Fee, f.FeeCoin)
			if !f.RealizedPnl.IsZero() {
				fmt.Printf("    RealizedPnL: %s\n", f.RealizedPnl)
			}
		}
		return nil
	})
	if err != nil {
		log.Printf("Failed to subscribe to fills: %v", err)
	}

	fmt.Println("\n✅ All private channel subscriptions active.")
	fmt.Println("📡 Listening for real-time updates... (Press Ctrl+C to exit)\n")

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	fmt.Println("\n\n🛑 Shutting down...")

	// Cleanup: Unsubscribe from all channels
	fmt.Println("Unsubscribing from channels...")
	client.UnsubscribeAccount()
	client.UnsubscribePositions()
	client.UnsubscribeOrders()
	client.UnsubscribeFills()

	// Close connection
	if err := client.Close(); err != nil {
		log.Printf("Error closing connection: %v", err)
	}

	fmt.Println("👋 Goodbye!")
}

// getOrderStateString returns a human-readable order state string
func getOrderStateString(state int) string {
	switch state {
	case 0:
		return "PENDING"
	case 1:
		return "OPEN"
	case 2:
		return "PARTIALLY_FILLED"
	case 3:
		return "FILLED"
	case 4:
		return "CANCELED"
	case 5:
		return "FAILED"
	default:
		return fmt.Sprintf("UNKNOWN(%d)", state)
	}
}
