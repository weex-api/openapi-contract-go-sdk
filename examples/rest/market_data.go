// Market Data Example
//
// This example demonstrates how to use the WEEX Contract API SDK
// to fetch market data (public endpoints).
//
// Run: go run examples/rest/market_data.go

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/weex-api/openapi-contract-go-sdk/weex"
	"github.com/weex-api/openapi-contract-go-sdk/weex/rest/market"
	"github.com/weex-api/openapi-contract-go-sdk/weex/types"
)

func main() {
	// Create configuration for public endpoints
	// No API credentials required for public market data
	config := weex.NewDefaultConfig()
	config.LogLevel = weex.LogLevelInfo

	// Create client
	client, err := weex.NewPublicClient(config)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// Example 1: Get server time
	fmt.Println("=== Example 1: Get Server Time ===")
	serverTime, err := client.Market().GetServerTime(ctx)
	if err != nil {
		log.Printf("Failed to get server time: %v", err)
	} else {
		fmt.Printf("Server Time: %v (%d)\n", time.UnixMilli(serverTime.Timestamp), serverTime.Timestamp)
		fmt.Printf("ISO Time: %s\n", serverTime.ISO)
	}

	// Example 2: Get contract information
	fmt.Println("\n=== Example 2: Get Contract Information ===")
	contracts, err := client.Market().GetContracts(ctx, &market.GetContractsRequest{
		Symbol: "cmt_btcusdt",
	})
	if err != nil {
		log.Printf("Failed to get contracts: %v", err)
	} else {
		for _, contract := range contracts {
			fmt.Printf("Contract: %s\n", contract.Symbol)
			fmt.Printf("  Underlying Index: %s\n", contract.UnderlyingIndex)
			fmt.Printf("  Quote Currency: %s\n", contract.QuoteCurrency)
			fmt.Printf("  Coin: %s\n", contract.Coin)
			fmt.Printf("  Max Leverage: %d\n", contract.MaxLeverage)
			fmt.Printf("  Maker Fee: %s\n", contract.MakerFeeRate)
			fmt.Printf("  Taker Fee: %s\n", contract.TakerFeeRate)
		}
	}

	// Example 3: Get ticker
	fmt.Println("\n=== Example 3: Get Ticker ===")
	ticker, err := client.Market().GetTicker(ctx, "cmt_btcusdt")
	if err != nil {
		log.Printf("Failed to get ticker: %v", err)
	} else {
		fmt.Printf("Ticker for %s:\n", ticker.Symbol)
		fmt.Printf("  Last Price: %s\n", ticker.Last)
		fmt.Printf("  Mark Price: %s\n", ticker.MarkPrice)
		fmt.Printf("  Index Price: %s\n", ticker.IndexPrice)
		fmt.Printf("  24h Change: %s%%\n", ticker.PriceChangePercent)
		fmt.Printf("  24h High: %s\n", ticker.High24h)
		fmt.Printf("  24h Low: %s\n", ticker.Low24h)
		fmt.Printf("  24h Volume: %s\n", ticker.Volume24h)
		fmt.Printf("  Best Bid: %s\n", ticker.BestBid)
		fmt.Printf("  Best Ask: %s\n", ticker.BestAsk)
	}

	// Example 4: Get all tickers
	fmt.Println("\n=== Example 4: Get All Tickers ===")
	allTickers, err := client.Market().GetAllTickers(ctx)
	if err != nil {
		log.Printf("Failed to get all tickers: %v", err)
	} else {
		fmt.Printf("Total contracts: %d\n", len(allTickers))
		fmt.Println("Top 5 contracts:")
		for i, t := range allTickers {
			if i >= 5 {
				break
			}
			fmt.Printf("  %s: %s (Change: %s%%)\n", t.Symbol, t.Last, t.PriceChangePercent)
		}
	}

	// Example 5: Get order book depth
	fmt.Println("\n=== Example 5: Get Order Book Depth ===")
	depth, err := client.Market().GetDepth(ctx, &market.GetDepthRequest{
		Symbol: "cmt_btcusdt",
		Limit:  15, // Must be 15 or 200
	})
	if err != nil {
		log.Printf("Failed to get depth: %v", err)
	} else {
		fmt.Printf("Order Book for cmt_btcusdt:\n")
		fmt.Println("  Bids (Buy Orders):")
		for i, bid := range depth.Bids {
			if i >= 5 {
				break
			}
			if len(bid) >= 2 {
				fmt.Printf("    Price: %s, Quantity: %s\n", bid[0], bid[1])
			}
		}
		fmt.Println("  Asks (Sell Orders):")
		for i, ask := range depth.Asks {
			if i >= 5 {
				break
			}
			if len(ask) >= 2 {
				fmt.Printf("    Price: %s, Quantity: %s\n", ask[0], ask[1])
			}
		}
	}

	// Example 6: Get recent trades
	fmt.Println("\n=== Example 6: Get Recent Trades ===")
	trades, err := client.Market().GetTrades(ctx, &market.GetTradesRequest{
		Symbol: "cmt_btcusdt",
		Limit:  5,
	})
	if err != nil {
		log.Printf("Failed to get trades: %v", err)
	} else {
		fmt.Printf("Recent trades for cmt_btcusdt:\n")
		for i, trade := range trades {
			if i >= 5 {
				break
			}
			side := "SELL"
			if trade.IsBuyerMaker {
				side = "BUY"
			}
			fmt.Printf("  [%s] Price: %s, Qty: %s, Time: %v\n",
				side, trade.Price, trade.Size, time.UnixMilli(trade.Time).Format("15:04:05"))
		}
	}

	// Example 7: Get kline/candlestick data
	fmt.Println("\n=== Example 7: Get Kline Data ===")
	klines, err := client.Market().GetKlines(ctx, &market.GetKlinesRequest{
		Symbol:   "cmt_btcusdt",
		Interval: types.Interval1Hour,
		Limit:    5,
	})
	if err != nil {
		log.Printf("Failed to get klines: %v", err)
	} else {
		fmt.Printf("Recent 1h klines for cmt_btcusdt:\n")
		for _, kline := range klines {
			if len(kline) >= 7 {
				fmt.Printf("  Time: %s, O: %s, H: %s, L: %s, C: %s, V: %s\n",
					kline[0], kline[1], kline[2], kline[3], kline[4], kline[5])
			}
		}
	}

	// Example 8: Get funding rate
	fmt.Println("\n=== Example 8: Get Funding Rate ===")
	fundingRate, err := client.Market().GetFundingRate(ctx, "cmt_btcusdt")
	if err != nil {
		log.Printf("Failed to get funding rate: %v", err)
	} else {
		fmt.Printf("Funding Rate for %s:\n", fundingRate.Symbol)
		fmt.Printf("  Current Rate: %s\n", fundingRate.FundingRate)
		fmt.Printf("  Collect Cycle: %d minutes\n", fundingRate.CollectCycle)
		fmt.Printf("  Timestamp: %v\n", time.UnixMilli(fundingRate.Timestamp).Format("2006-01-02 15:04:05"))
	}

	// Example 9: Get index price
	fmt.Println("\n=== Example 9: Get Index Price ===")
	indexPrice, err := client.Market().GetIndexPrice(ctx, "cmt_btcusdt")
	if err != nil {
		log.Printf("Failed to get index price: %v", err)
	} else {
		fmt.Printf("Index Price for %s: %s\n", indexPrice.Symbol, indexPrice.Index)
	}

	// Example 10: Get open interest
	fmt.Println("\n=== Example 10: Get Open Interest ===")
	openInterest, err := client.Market().GetOpenInterest(ctx, "cmt_btcusdt")
	if err != nil {
		log.Printf("Failed to get open interest: %v", err)
	} else {
		fmt.Printf("Open Interest for %s:\n", openInterest.Symbol)
		fmt.Printf("  Base Volume: %s\n", openInterest.BaseVolume)
		fmt.Printf("  Target Volume: %s\n", openInterest.TargetVolume)
	}

	fmt.Println("\n=== All examples completed! ===")
}
