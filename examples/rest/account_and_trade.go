// Trading Example
//
// This example demonstrates how to use the WEEX Contract API SDK
// for trading operations (placing orders, managing positions, etc.).
//
// ⚠️ WARNING: This example uses real trading functions!
// Make sure to use test credentials or very small amounts.
//
// Run: go run examples/rest/account_and_trade.go

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/weex/openapi-contract-go-sdk/weex"
	"github.com/weex/openapi-contract-go-sdk/weex/rest/account"
	"github.com/weex/openapi-contract-go-sdk/weex/rest/trade"
	"github.com/weex/openapi-contract-go-sdk/weex/types"
)

func main() {
	// ⚠️ IMPORTANT: Replace with your actual API credentials
	config := weex.NewDefaultConfig().
		WithAPIKey("your-api-key").
		WithSecretKey("your-secret-key").
		WithPassphrase("your-passphrase").
		WithLogLevel(weex.LogLevelInfo)

	// Create client
	client, err := weex.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// ===== ACCOUNT EXAMPLES =====

	fmt.Println("=== Example 1: Get Account Balance ===")
	assets, err := client.Account().GetAccountBalance(ctx)
	if err != nil {
		log.Printf("Failed to get account balance: %v", err)
	} else {
		fmt.Printf("Account Assets (%d currencies):\n", len(assets))
		for _, asset := range assets {
			fmt.Printf("  %s: Available=%s, Frozen=%s, Equity=%s, UnrealizedPnL=%s\n",
				asset.CoinName, asset.Available, asset.Frozen, asset.Equity, asset.UnrealizedPnl)
		}
	}

	fmt.Println("\n=== Example 2: Get All Positions ===")
	positions, err := client.Account().GetAllPositions(ctx, &account.GetAllPositionsRequest{})
	if err != nil {
		log.Printf("Failed to get positions: %v", err)
	} else {
		fmt.Printf("Positions (%d):\n", len(positions))
		for _, pos := range positions {
			fmt.Printf("  %s [%s]: Size=%s, AvgPrice=%s, UnrealizedPnL=%s, Leverage=%sx\n",
				pos.Symbol, pos.PositionSide, pos.Size, pos.AverageOpenPrice, pos.UnrealizedPnl, pos.Leverage)
			fmt.Printf("    LiquidatePrice=%s, MarkPrice=%s\n", pos.LiquidatePrice, pos.MarkPrice)
		}
	}

	fmt.Println("\n=== Example 3: Get User Config ===")
	config1, err := client.Account().GetUserConfig(ctx, &account.GetUserConfigRequest{
		Symbol: "cmt_btcusdt",
	})
	if err != nil {
		log.Printf("Failed to get user config: %v", err)
	} else {
		fmt.Printf("Config for %s:\n", config1.Symbol)
		fmt.Printf("  MarginMode: %d, PositionMode: %d\n", config1.MarginMode, config1.PositionMode)
		fmt.Printf("  Leverage: %s (Long: %s, Short: %s)\n",
			config1.Leverage, config1.LongLeverage, config1.ShortLeverage)
	}

	// ===== TRADING EXAMPLES =====

	fmt.Println("\n=== Example 4: Get Current Orders ===")
	orders, err := client.Trade().GetCurrentOrderStatus(ctx, &trade.GetOrdersRequest{
		Symbol: "cmt_btcusdt",
		Limit:  10,
	})
	if err != nil {
		log.Printf("Failed to get current orders: %v", err)
	} else {
		fmt.Printf("Current Orders (%d):\n", len(orders.Orders))
		for _, order := range orders.Orders {
			fmt.Printf("  OrderID: %s, ClientOID: %s\n", order.OrderId, order.ClientOid)
			fmt.Printf("    Type: %d, Price: %s, Size: %s, Filled: %s\n",
				order.Type, order.Price, order.Size, order.FilledSize)
			fmt.Printf("    State: %d, CreateTime: %v\n",
				order.State, time.UnixMilli(order.CreateTime).Format("2006-01-02 15:04:05"))
		}
	}

	fmt.Println("\n=== Example 5: Get Trade Fills ===")
	fills, err := client.Trade().GetTradeDetails(ctx, &trade.GetFillsRequest{
		Symbol: "cmt_btcusdt",
		Limit:  10,
	})
	if err != nil {
		log.Printf("Failed to get fills: %v", err)
	} else {
		fmt.Printf("Recent Fills (%d):\n", len(fills.Fills))
		for _, fill := range fills.Fills {
			fmt.Printf("  FillID: %s, OrderID: %s\n", fill.FillId, fill.OrderId)
			fmt.Printf("    Price: %s, Size: %s, Fee: %s %s\n",
				fill.Price, fill.Size, fill.Fee, fill.FeeCoin)
			fmt.Printf("    Side: %s, Liquidity: %s, RealizedPnL: %s\n",
				fill.Side, fill.Liquidity, fill.RealizedPnl)
		}
	}

	// ===== TRADING OPERATIONS (COMMENTED OUT FOR SAFETY) =====

	/*
		⚠️ WARNING: The following examples will place REAL ORDERS!
		Uncomment only if you understand the risks and want to test with real money.

		fmt.Println("\n=== Example 6: Place Limit Order ===")
		// Generate unique client order ID
		clientOid := fmt.Sprintf("sdk_test_%d", time.Now().UnixMilli())

		placeReq := &trade.PlaceOrderRequest{
			Symbol:     "cmt_btcusdt",
			ClientOid:  clientOid,
			Size:       types.NewDecimalFromString("0.001"), // Very small size for testing
			Type:       types.OrderTypeOpenLong,             // Open long
			OrderType:  types.OrderExecNormal,               // Normal order
			MatchPrice: types.PriceMatchLimit,               // Limit order
			Price:      types.NewDecimalFromString("40000"), // Set your price
			MarginMode: int(types.MarginModeShared),         // Cross margin
		}

		placeResp, err := client.Trade().PlaceOrder(ctx, placeReq)
		if err != nil {
			log.Printf("Failed to place order: %v", err)
		} else {
			fmt.Printf("Order placed successfully!\n")
			fmt.Printf("  OrderID: %s\n", placeResp.OrderId)
			fmt.Printf("  ClientOID: %s\n", placeResp.ClientOid)
		}

		// Wait a bit before canceling
		time.Sleep(2 * time.Second)

		fmt.Println("\n=== Example 7: Cancel Order ===")
		cancelReq := &trade.CancelOrderRequest{
			OrderId: placeResp.OrderId,
			Symbol:  "cmt_btcusdt",
		}

		cancelResp, err := client.Trade().CancelOrder(ctx, cancelReq)
		if err != nil {
			log.Printf("Failed to cancel order: %v", err)
		} else {
			fmt.Printf("Order canceled successfully!\n")
			fmt.Printf("  OrderID: %s\n", cancelResp.OrderId)
		}
	*/

	/*
		fmt.Println("\n=== Example 8: Adjust Leverage ===")
		leverageReq := &account.AdjustLeverageRequest{
			Symbol:     "cmt_btcusdt",
			MarginMode: int(types.MarginModeShared),
			Leverage:   types.NewDecimalFromString("10"), // 10x leverage
		}

		leverageResp, err := client.Account().AdjustLeverage(ctx, leverageReq)
		if err != nil {
			log.Printf("Failed to adjust leverage: %v", err)
		} else {
			fmt.Printf("Leverage adjusted successfully!\n")
			fmt.Printf("  Symbol: %s\n", leverageResp.Symbol)
			fmt.Printf("  New Leverage: %s\n", leverageResp.Leverage)
		}
	*/

	/*
		fmt.Println("\n=== Example 9: Place Batch Orders ===")
		batchReq := &trade.PlaceOrdersBatchRequest{
			Orders: []trade.PlaceOrderRequest{
				{
					Symbol:     "cmt_btcusdt",
					ClientOid:  fmt.Sprintf("batch1_%d", time.Now().UnixMilli()),
					Size:       types.NewDecimalFromString("0.001"),
					Type:       types.OrderTypeOpenLong,
					OrderType:  types.OrderExecNormal,
					MatchPrice: types.PriceMatchLimit,
					Price:      types.NewDecimalFromString("40000"),
				},
				{
					Symbol:     "cmt_btcusdt",
					ClientOid:  fmt.Sprintf("batch2_%d", time.Now().UnixMilli()),
					Size:       types.NewDecimalFromString("0.001"),
					Type:       types.OrderTypeOpenLong,
					OrderType:  types.OrderExecNormal,
					MatchPrice: types.PriceMatchLimit,
					Price:      types.NewDecimalFromString("39900"),
				},
			},
		}

		batchResp, err := client.Trade().PlaceOrdersBatch(ctx, batchReq)
		if err != nil {
			log.Printf("Failed to place batch orders: %v", err)
		} else {
			fmt.Printf("Batch orders placed:\n")
			fmt.Printf("  Success: %d\n", len(batchResp.Success))
			fmt.Printf("  Failed: %d\n", len(batchResp.Failed))
		}
	*/

	fmt.Println("\n=== All examples completed! ===")
	fmt.Println("\n⚠️ Note: Trading examples are commented out for safety.")
	fmt.Println("Uncomment them in the source code if you want to test real trading.")
}
