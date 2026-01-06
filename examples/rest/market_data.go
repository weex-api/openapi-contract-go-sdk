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
		fmt.Printf("Server Time: %v (%d)\n", time.UnixMilli(serverTime.ServerTime), serverTime.ServerTime)
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
			fmt.Printf("  Base Coin: %s\n", contract.BaseCoin)
			fmt.Printf("  Quote Coin: %s\n", contract.QuoteCoin)
			fmt.Printf("  Contract Type: %s\n", contract.ContractType)
			fmt.Printf("  Max Leverage: %s\n", contract.MaxLeverage)
			fmt.Printf("  Maker Fee: %s\n", contract.MakerFee)
			fmt.Printf("  Taker Fee: %s\n", contract.TakerFee)
		}
	}

	// Example 3: Get ticker
	fmt.Println("\n=== Example 3: Get Ticker ===")
	ticker, err := client.Market().GetTicker(ctx, "cmt_btcusdt")
	if err != nil {
		log.Printf("Failed to get ticker: %v", err)
	} else {
		fmt.Printf("Ticker for %s:\n", ticker.Symbol)
		fmt.Printf("  Last Price: %s\n", ticker.LastPrice)
		fmt.Printf("  Mark Price: %s\n", ticker.MarkPrice)
		fmt.Printf("  Index Price: %s\n", ticker.IndexPrice)
		fmt.Printf("  24h Change: %s (%s%%)\n", ticker.PriceChange, ticker.PriceChangePercent)
		fmt.Printf("  24h High: %s\n", ticker.HighPrice)
		fmt.Printf("  24h Low: %s\n", ticker.LowPrice)
		fmt.Printf("  24h Volume: %s\n", ticker.Volume)
		fmt.Printf("  Best Bid: %s @ %s\n", ticker.BidQty, ticker.BidPrice)
		fmt.Printf("  Best Ask: %s @ %s\n", ticker.AskQty, ticker.AskPrice)
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
			fmt.Printf("  %s: %s (Change: %s%%)\n", t.Symbol, t.LastPrice, t.PriceChangePercent)
		}
	}

	// Example 5: Get order book depth
	fmt.Println("\n=== Example 5: Get Order Book Depth ===")
	depth, err := client.Market().GetDepth(ctx, &market.GetDepthRequest{
		Symbol: "cmt_btcusdt",
		Limit:  5, // Get top 5 levels
	})
	if err != nil {
		log.Printf("Failed to get depth: %v", err)
	} else {
		fmt.Printf("Order Book for %s:\n", depth.Symbol)
		fmt.Println("  Bids (Buy Orders):")
		for i, bid := range depth.Bids {
			if i >= 5 {
				break
			}
			fmt.Printf("    Price: %s, Quantity: %s\n", bid.Price, bid.Quantity)
		}
		fmt.Println("  Asks (Sell Orders):")
		for i, ask := range depth.Asks {
			if i >= 5 {
				break
			}
			fmt.Printf("    Price: %s, Quantity: %s\n", ask.Price, ask.Quantity)
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
				side, trade.Price, trade.Qty, time.UnixMilli(trade.Time).Format("15:04:05"))
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
			fmt.Printf("  Time: %v, O: %s, H: %s, L: %s, C: %s, V: %s\n",
				time.UnixMilli(kline.OpenTime).Format("2006-01-02 15:04"),
				kline.Open, kline.High, kline.Low, kline.Close, kline.Volume)
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
		fmt.Printf("  Funding Time: %v\n", time.UnixMilli(fundingRate.FundingTime).Format("2006-01-02 15:04:05"))
		fmt.Printf("  Next Rate: %s\n", fundingRate.NextFundingRate)
		fmt.Printf("  Next Time: %v\n", time.UnixMilli(fundingRate.NextFundingTime).Format("2006-01-02 15:04:05"))
	}

	// Example 9: Get index price
	fmt.Println("\n=== Example 9: Get Index Price ===")
	indexPrice, err := client.Market().GetIndexPrice(ctx, "cmt_btcusdt")
	if err != nil {
		log.Printf("Failed to get index price: %v", err)
	} else {
		fmt.Printf("Index Price for %s: %s\n", indexPrice.Symbol, indexPrice.IndexPrice)
	}

	// Example 10: Get open interest
	fmt.Println("\n=== Example 10: Get Open Interest ===")
	openInterest, err := client.Market().GetOpenInterest(ctx, "cmt_btcusdt")
	if err != nil {
		log.Printf("Failed to get open interest: %v", err)
	} else {
		fmt.Printf("Open Interest for %s:\n", openInterest.Symbol)
		fmt.Printf("  Open Interest: %s\n", openInterest.OpenInterest)
		fmt.Printf("  Open Interest Value: %s\n", openInterest.OpenInterestValue)
	}

	fmt.Println("\n=== All examples completed! ===")
}
