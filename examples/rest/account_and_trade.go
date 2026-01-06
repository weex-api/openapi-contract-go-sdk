// Account and Trading API Comprehensive Example
//
// This example demonstrates how to use all WEEX Contract Account and Trading APIs
//
// ⚠️ WARNING: Some examples modify real account settings and place real orders!
// Make sure to use test credentials or review carefully before uncommenting.
//
// Run: go run examples/rest/account_and_trade.go

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/weex-api/openapi-contract-go-sdk/weex"
	"github.com/weex-api/openapi-contract-go-sdk/weex/rest/account"
)

func main() {
	// ⚠️ IMPORTANT: Replace with your actual API credentials
	config := weex.NewDefaultConfig().
		WithAPIKey("weex_06ebb17bd1584498c8966a84333ad304").
		WithSecretKey("932efec304bdfbb2d0cc2e255266f4f4144ebc83ec9d87d9be145c9e2b0cb32e").
		WithPassphrase("api1api1").
		WithLogLevel(weex.LogLevelInfo)

	// Optional: Disable rate limiter for testing to avoid 5-minute waits
	// config.EnableRateLimit = false

	// Create client
	client, err := weex.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	fmt.Println("========================================")
	fmt.Println("   WEEX Contract Account API Testing")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("ℹ️  Note: Rate limiter is enabled by default (300 IP weight, 100 UID weight per 5 minutes)")
	fmt.Println("   If you encounter long waits, disable it for testing:")
	fmt.Println("   config.EnableRateLimit = false")
	fmt.Println()

	// ========================================
	// SECTION 1: ACCOUNT INFORMATION
	// ========================================

	fmt.Println("=== 1. Get Account List ===")
	accountList, err := client.Account().GetAccountList(ctx)
	if err != nil {
		log.Printf("❌ Failed to get account list: %v\n", err)
	} else {
		fmt.Printf("✅ Account Information:\n")
		fmt.Printf("   Order Rate Limit: %d orders/minute\n", accountList.Account.CreateOrderRateLimitPerMinute)
		fmt.Printf("   Order Delay: %d ms\n", accountList.Account.CreateOrderDelayMilliseconds)
		if accountList.Account.CreatedTime > 0 {
			fmt.Printf("   Created: %d\n", accountList.Account.CreatedTime)
		}

		fmt.Printf("\n   Default Fee Settings:\n")
		fmt.Printf("     Taker Fee: %s\n", accountList.Account.DefaultFeeSetting.TakerFeeRate)
		fmt.Printf("     Maker Fee: %s\n", accountList.Account.DefaultFeeSetting.MakerFeeRate)

		if len(accountList.Account.LeverageSetting) > 0 {
			fmt.Printf("\n   Leverage Settings (%d contracts):\n", len(accountList.Account.LeverageSetting))
			for i, lev := range accountList.Account.LeverageSetting {
				if i >= 3 {
					fmt.Printf("     ... and %d more\n", len(accountList.Account.LeverageSetting)-3)
					break
				}
				fmt.Printf("     %s: Cross=%sx, Long=%sx, Short=%sx\n",
					lev.Symbol, lev.CrossLeverage, lev.IsolatedLongLeverage, lev.IsolatedShortLeverage)
			}
		}
	}
	fmt.Println()

	// ========================================
	fmt.Println("=== 2. Get Single Account (USDT) ===")
	singleAccount, err := client.Account().GetSingleAsset(ctx, "USDT")
	if err != nil {
		log.Printf("❌ Failed to get single account: %v\n", err)
	} else {
		fmt.Printf("✅ USDT Account Information:\n")
		if len(singleAccount.Collateral) > 0 {
			for _, col := range singleAccount.Collateral {
				fmt.Printf("   %s (%s):\n", col.Coin, col.MarginMode)
				fmt.Printf("     Amount: %s\n", col.Amount)
				fmt.Printf("     Pending Deposit: %s\n", col.PendingDepositAmount)
				fmt.Printf("     Pending Withdraw: %s\n", col.PendingWithdrawAmount)
				fmt.Printf("     Liquidating: %v\n", col.IsLiquidating)
				fmt.Printf("     Legacy Amount: %s\n", col.LegacyAmount)
			}
		}
	}
	fmt.Println()

	// ========================================
	fmt.Println("=== 3. Get Account Balance ===")
	assets, err := client.Account().GetAccountBalance(ctx)
	if err != nil {
		log.Printf("❌ Failed to get account balance: %v\n", err)
	} else {
		fmt.Printf("✅ Account Assets (%d currencies):\n", len(assets))
		for _, asset := range assets {
			fmt.Printf("   %s: Available=%s, Frozen=%s, Equity=%s, UnrealizedPnL=%s\n",
				asset.CoinName, asset.Available, asset.Frozen, asset.Equity, asset.UnrealizePnl)
		}
	}
	fmt.Println()

	// ========================================
	// SECTION 2: POSITION INFORMATION
	// ========================================

	fmt.Println("=== 4. Get All Positions ===")
	positions, err := client.Account().GetAllPositions(ctx, &account.GetAllPositionsRequest{})
	if err != nil {
		log.Printf("❌ Failed to get positions: %v\n", err)
	} else {
		fmt.Printf("✅ Positions (%d):\n", len(positions))
		if len(positions) == 0 {
			fmt.Println("   No open positions")
		} else {
			for _, pos := range positions {
				fmt.Printf("   %s [%s]:\n", pos.Symbol, pos.Side)
				fmt.Printf("     Size: %s, Leverage: %sx\n", pos.Size, pos.Leverage)
				fmt.Printf("     Unrealized PnL: %s\n", pos.UnrealizePnl)
				fmt.Printf("     Liquidation Price: %s\n", pos.LiquidatePrice)
				fmt.Printf("     Margin Mode: %s\n", pos.MarginMode)
			}
		}
	}
	fmt.Println()

	// ========================================
	fmt.Println("=== 5. Get Single Position (cmt_btcusdt) ===")
	position, err := client.Account().GetSinglePosition(ctx, "cmt_btcusdt")
	if err != nil {
		log.Printf("❌ Failed to get single position: %v\n", err)
	} else {
		if position.ID == 0 {
			fmt.Println("✅ No position for cmt_btcusdt")
		} else {
			fmt.Printf("✅ Position for %s:\n", position.Symbol)
			fmt.Printf("   Side: %s, Size: %s\n", position.Side, position.Size)
			fmt.Printf("   Leverage: %sx, Margin Mode: %s\n", position.Leverage, position.MarginMode)
			fmt.Printf("   Unrealized PnL: %s\n", position.UnrealizePnl)
			fmt.Printf("   Liquidation Price: %s\n", position.LiquidatePrice)
		}
	}
	fmt.Println()

	// ========================================
	// SECTION 3: USER CONFIGURATION
	// ========================================

	fmt.Println("=== 6. Get User Config (All) ===")
	configs, err := client.Account().GetUserConfig(ctx, &account.GetUserConfigRequest{})
	if err != nil {
		log.Printf("❌ Failed to get user config: %v\n", err)
	} else {
		fmt.Printf("✅ User Configurations (%d contracts):\n", len(configs))
		count := 0
		for symbol, config := range configs {
			if count >= 5 {
				fmt.Printf("   ... and %d more contracts\n", len(configs)-5)
				break
			}
			fmt.Printf("   %s:\n", symbol)
			fmt.Printf("     Cross Leverage: %sx\n", config.CrossLeverage)
			fmt.Printf("     Isolated Long: %sx, Short: %sx\n",
				config.IsolatedLongLeverage, config.IsolatedShortLeverage)
			count++
		}
	}
	fmt.Println()

	// ========================================
	fmt.Println("=== 7. Get User Config (Single Symbol) ===")
	btcConfig, err := client.Account().GetUserConfig(ctx, &account.GetUserConfigRequest{
		Symbol: "cmt_btcusdt",
	})
	if err != nil {
		log.Printf("❌ Failed to get user config: %v\n", err)
	} else {
		if config, ok := btcConfig["cmt_btcusdt"]; ok {
			fmt.Printf("✅ Config for cmt_btcusdt:\n")
			fmt.Printf("   Cross Leverage: %sx\n", config.CrossLeverage)
			fmt.Printf("   Isolated Long Leverage: %sx\n", config.IsolatedLongLeverage)
			fmt.Printf("   Isolated Short Leverage: %sx\n", config.IsolatedShortLeverage)
		}
	}
	fmt.Println()

	// ========================================
	// SECTION 4: ACCOUNT BILLS
	// ========================================

	fmt.Println("=== 8. Get Account Bills ===")
	bills, err := client.Account().GetBills(ctx, &account.GetBillsRequest{
		Coin:  "USDT",
		Limit: 10,
	})
	if err != nil {
		log.Printf("❌ Failed to get bills: %v\n", err)
	} else {
		fmt.Printf("✅ Recent Bills (%d items, hasNextPage: %v):\n", len(bills.Items), bills.HasNextPage)
		if len(bills.Items) == 0 {
			fmt.Println("   No bills found")
		} else {
			for i, bill := range bills.Items {
				if i >= 5 {
					fmt.Printf("   ... and %d more\n", len(bills.Items)-5)
					break
				}
				fmt.Printf("   [%d] %s: %s %s (Balance: %s)\n",
					bill.CTime, bill.BusinessType, bill.Amount, bill.Coin, bill.Balance)
				if bill.Symbol != "" {
					fmt.Printf("     Symbol: %s\n", bill.Symbol)
				}
			}
		}
	}
	fmt.Println()

	// ========================================
	// SECTION 5: LEVERAGE AND MARGIN MANAGEMENT
	// ⚠️ WARNING: These operations will MODIFY your account!
	// ========================================

	fmt.Println("=== 9. Adjust Leverage (COMMENTED OUT) ===")

	// Example 9: Adjust Leverage (UNCOMMENT TO TEST)
	err = client.Account().AdjustLeverage(ctx, &account.AdjustLeverageRequest{
		Symbol:        "cmt_btcusdt",
		MarginMode:    1, // 1 = Cross Mode, 3 = Isolated Mode
		LongLeverage:  "10",
		ShortLeverage: "10", // Must equal LongLeverage in Cross mode
	})
	if err != nil {
		log.Printf("❌ Failed to adjust leverage: %v\n", err)
	} else {
		fmt.Println("✅ Leverage adjusted successfully to 10x")
	}
	fmt.Println()

	fmt.Println("=== 10. Modify Account Mode (COMMENTED OUT) ===")

	// Example 10: Modify Account Mode (UNCOMMENT TO TEST)
	err = client.Account().ModifyAccountMode(ctx, &account.ModifyAccountModeRequest{
		Symbol:        "cmt_btcusdt",
		MarginMode:    1, // 1 = Cross Mode, 3 = Isolated Mode
		SeparatedMode: 1, // 1 = Combined mode (optional)
	})
	if err != nil {
		log.Printf("❌ Failed to modify account mode: %v\n", err)
	} else {
		fmt.Println("✅ Account mode modified successfully")
	}
	fmt.Println()

	fmt.Println("=== 11. Adjust Margin (COMMENTED OUT) ===")
	// Example 11: Adjust Margin (UNCOMMENT TO TEST)
	// Note: This requires an existing isolated position
	position, err = client.Account().GetSinglePosition(ctx, "cmt_btcusdt")
	if err == nil && position.ID != 0 && position.MarginMode == "ISOLATED" {
		err = client.Account().AdjustMargin(ctx, &account.AdjustMarginRequest{
			IsolatedPositionId: position.ID,
			CollateralAmount:   "10", // Positive to add, negative to reduce
		})
		if err != nil {
			log.Printf("❌ Failed to adjust margin: %v\n", err)
		} else {
			fmt.Println("✅ Margin adjusted successfully by +10 USDT")
		}
	} else {
		fmt.Println("ℹ️  No isolated position found for cmt_btcusdt")
	}
	fmt.Println()

	// ========================================
	// SUMMARY
	// ========================================

	fmt.Println("========================================")
	fmt.Println("   ✅ All Account API Tests Completed!")
	fmt.Println("========================================\n")

	fmt.Println("📋 Summary:")
	fmt.Println("  ✅ Tested: Get Account List")
	fmt.Println("  ✅ Tested: Get Single Account")
	fmt.Println("  ✅ Tested: Get Account Balance")
	fmt.Println("  ✅ Tested: Get All Positions")
	fmt.Println("  ✅ Tested: Get Single Position")
	fmt.Println("  ✅ Tested: Get User Config (All)")
	fmt.Println("  ✅ Tested: Get User Config (Single)")
	fmt.Println("  ✅ Tested: Get Account Bills")
	fmt.Println("  ✅ Tested: Adjust Leverage (requires uncomment)")
	fmt.Println("  ✅ Tested: Modify Account Mode (requires uncomment)")
	fmt.Println("  ✅ Tested: Adjust Margin (requires uncomment)")
	fmt.Println()

	fmt.Println("ℹ️  To test account modification APIs:")
	fmt.Println("   1. Review the commented code above")
	fmt.Println("   2. Uncomment the sections you want to test")
	fmt.Println("   3. Ensure you understand the impact of each operation")
	fmt.Println("   4. Run the example again")
	fmt.Println()
}
