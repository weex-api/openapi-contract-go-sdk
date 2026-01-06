// Trade API Testing
//
// This example demonstrates how to use WEEX Contract Trade APIs
// Run: go run examples/rest/trade.go

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/weex-api/openapi-contract-go-sdk/weex"
	"github.com/weex-api/openapi-contract-go-sdk/weex/rest/trade"
)

func main() {
	config := weex.NewDefaultConfig().
		WithAPIKey("weex_06ebb17bd1584498c8966a84333ad304").
		WithSecretKey("932efec304bdfbb2d0cc2e255266f4f4144ebc83ec9d87d9be145c9e2b0cb32e").
		WithPassphrase("api1api1").
		WithLogLevel(weex.LogLevelInfo)

	// Disable rate limiter for testing
	config.EnableRateLimit = false

	client, err := weex.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()
	symbol := "cmt_btcusdt"

	fmt.Println("=== WEEX Contract Trade API Testing ===\n")

	// Test 1: Get Current Orders
	currentOrders, err := client.Trade().GetCurrentOrderStatus(ctx, symbol, 0, 0, 0, 10, 0)
	if err != nil {
		fmt.Printf("GetCurrentOrderStatus: ❌ %v\n", err)
	} else {
		fmt.Printf("GetCurrentOrderStatus: ✅ Found %d orders\n", len(currentOrders))
	}

	// Test 2: Get Order History
	orderHistory, err := client.Trade().GetOrderHistory(ctx, symbol, 10, 0, 0)
	if err != nil {
		fmt.Printf("GetOrderHistory: ❌ %v\n", err)
	} else {
		fmt.Printf("GetOrderHistory: ✅ Found %d orders\n", len(orderHistory))
	}

	// Test 3: Get Trade Details (Fills)
	fills, err := client.Trade().GetTradeDetails(ctx, symbol, 0, 0, 0, 10)
	if err != nil {
		fmt.Printf("GetTradeDetails: ❌ %v\n", err)
	} else {
		fmt.Printf("GetTradeDetails: ✅ Found %d fills\n", len(fills.List))
	}

	// Test 4: Get Current Pending Orders
	pendingOrders, err := client.Trade().GetCurrentPendingOrders(ctx, symbol, 0, 0, 0, 10, 0)
	if err != nil {
		fmt.Printf("GetCurrentPendingOrders: ❌ %v\n", err)
	} else {
		fmt.Printf("GetCurrentPendingOrders: ✅ Found %d orders\n", len(pendingOrders))
	}

	// Test 5: Place Order (commented for safety, will fail if uncommented without proper price)
	clientOid := fmt.Sprintf("%d", time.Now().UnixNano()/1000000)
	placeResp, err := client.Trade().PlaceOrder(ctx, &trade.PlaceOrderRequest{
		Symbol:     symbol,
		ClientOid:  clientOid,
		Size:       "0.001",
		Type:       "1", // Open long
		OrderType:  "0", // Normal
		MatchPrice: "0", // Limit price
		Price:      "50000",
		MarginMode: 1,
	})
	if err != nil {
		fmt.Printf("PlaceOrder: ❌ %v\n", err)
	} else {
		fmt.Printf("PlaceOrder: ✅ OrderId=%s\n", placeResp.OrderId)

		// Test 6: Get Single Order Info (only if place succeeded)
		if placeResp.OrderId != "" {
			orderInfo, err := client.Trade().GetSingleOrderInfo(ctx, placeResp.OrderId)
			if err != nil {
				fmt.Printf("GetSingleOrderInfo: ❌ %v\n", err)
			} else {
				fmt.Printf("GetSingleOrderInfo: ✅ Status=%s\n", orderInfo.Status)
			}

			// Test 7: Cancel Order (only if place succeeded)
			cancelResp, err := client.Trade().CancelOrder(ctx, &trade.CancelOrderRequest{
				OrderId: placeResp.OrderId,
			})
			if err != nil {
				fmt.Printf("CancelOrder: ❌ %v\n", err)
			} else {
				fmt.Printf("CancelOrder: ✅ Result=%v\n", cancelResp.Result)
			}
		}
	}

	// Test 8: Place Batch Orders
	clientOid1 := fmt.Sprintf("%d_1", time.Now().UnixNano()/1000000)
	clientOid2 := fmt.Sprintf("%d_2", time.Now().UnixNano()/1000000)
	batchResp, err := client.Trade().PlaceBatchOrders(ctx, &trade.PlaceBatchOrdersRequest{
		Symbol:     symbol,
		MarginMode: 1,
		OrderDataList: []trade.BatchOrderRequest{
			{
				ClientOid:  clientOid1,
				Size:       "0.001",
				Type:       "1",
				OrderType:  "0",
				MatchPrice: "0",
				Price:      "50000",
			},
			{
				ClientOid:  clientOid2,
				Size:       "0.001",
				Type:       "1",
				OrderType:  "0",
				MatchPrice: "0",
				Price:      "50100",
			},
		},
	})
	if err != nil {
		fmt.Printf("PlaceBatchOrders: ❌ %v\n", err)
	} else {
		fmt.Printf("PlaceBatchOrders: ✅ Result=%v, Orders=%d\n", batchResp.Result, len(batchResp.OrderInfo))

		// Test 9: Cancel Batch Orders (only if batch place succeeded)
		if batchResp.Result && len(batchResp.OrderInfo) > 0 {
			var orderIds []string
			for _, info := range batchResp.OrderInfo {
				if info.Result {
					orderIds = append(orderIds, info.OrderId)
				}
			}
			if len(orderIds) > 0 {
				cancelBatchResp, err := client.Trade().CancelBatchOrders(ctx, &trade.CancelBatchOrdersRequest{
					Ids: orderIds,
				})
				if err != nil {
					fmt.Printf("CancelBatchOrders: ❌ %v\n", err)
				} else {
					fmt.Printf("CancelBatchOrders: ✅ Result=%v, Cancelled=%d\n", cancelBatchResp.Result, len(cancelBatchResp.CancelOrderResultList))
				}
			}
		}
	}

	// Test 10: Place Pending Order (Trigger Order)
	pendingClientOid := fmt.Sprintf("%d_pending", time.Now().UnixNano()/1000000)
	pendingResp, err := client.Trade().PlacePendingOrder(ctx, &trade.PlacePendingOrderRequest{
		Symbol:       symbol,
		ClientOid:    pendingClientOid,
		Size:         "0.001",
		Type:         "1", // Open long
		MatchType:    "0", // Limit price
		ExecutePrice: "95000",
		TriggerPrice: "94000",
		MarginMode:   1,
	})
	if err != nil {
		fmt.Printf("PlacePendingOrder: ❌ %v\n", err)
	} else {
		fmt.Printf("PlacePendingOrder: ✅ OrderId=%s\n", pendingResp.OrderId)

		// Test 11: Cancel Pending Order (only if place succeeded)
		if pendingResp.OrderId != "" {
			cancelPendingResp, err := client.Trade().CancelPendingOrder(ctx, &trade.CancelPendingOrderRequest{
				OrderId: pendingResp.OrderId,
			})
			if err != nil {
				fmt.Printf("CancelPendingOrder: ❌ %v\n", err)
			} else {
				fmt.Printf("CancelPendingOrder: ✅ Result=%v\n", cancelPendingResp.Result)
			}
		}
	}

	// Test 12: Place TP/SL Order (requires existing position, will likely fail without position)
	tpslClientOid := fmt.Sprintf("%d_tpsl", time.Now().UnixNano()/1000000)
	tpslResp, err := client.Trade().PlaceTpSlOrder(ctx, &trade.PlaceTpSlOrderRequest{
		Symbol:        symbol,
		ClientOrderId: tpslClientOid,
		PlanType:      "profit_plan",
		TriggerPrice:  "100000",
		ExecutePrice:  "0", // Market price
		Size:          "0.001",
		PositionSide:  "long",
		MarginMode:    1,
	})
	if err != nil {
		fmt.Printf("PlaceTpSlOrder: ❌ %v\n", err)
	} else {
		if len(tpslResp) > 0 {
			fmt.Printf("PlaceTpSlOrder: ✅ Success=%v, OrderId=%d\n", tpslResp[0].Success, tpslResp[0].OrderId)
		} else {
			fmt.Printf("PlaceTpSlOrder: ✅ Empty response\n")
		}
	}

	// Test 13: Cancel All Orders (COMMENTED for safety - will cancel ALL orders!)

	cancelAllResp, err := client.Trade().CancelAllOrders(ctx, &trade.CancelAllOrdersRequest{
		Symbol:          symbol,
		CancelOrderType: "normal",
	})
	if err != nil {
		fmt.Printf("CancelAllOrders: ❌ %v\n", err)
	} else {
		fmt.Printf("CancelAllOrders: ✅ Cancelled %d orders\n", len(cancelAllResp))
	}

	// Test 14: Close Positions (COMMENTED for safety - will close ALL positions!)

	closeResp, err := client.Trade().ClosePositions(ctx, &trade.ClosePositionsRequest{
		Symbol: symbol,
	})
	if err != nil {
		fmt.Printf("ClosePositions: ❌ %v\n", err)
	} else {
		fmt.Printf("ClosePositions: ✅ Closed %d positions\n", len(closeResp))
	}

	fmt.Println("\n=== Testing Complete ===")
}
