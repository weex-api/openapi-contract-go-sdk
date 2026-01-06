// Package types provides common type definitions used across the WEEX Contract API SDK.
package types

import (
	"fmt"
	"strconv"
)

// MarginMode represents the margin mode for positions
type MarginMode int

const (
	MarginModeUnknown  MarginMode = 0
	MarginModeShared   MarginMode = 1 // Cross margin (全仓)
	MarginModeIsolated MarginMode = 3 // Isolated margin (逐仓)
)

// String returns the string representation of MarginMode
func (m MarginMode) String() string {
	switch m {
	case MarginModeShared:
		return "SHARED"
	case MarginModeIsolated:
		return "ISOLATED"
	default:
		return "UNKNOWN"
	}
}

// PositionMode represents the position mode
type PositionMode int

const (
	PositionModeUnknown PositionMode = 0
	PositionModeHedge   PositionMode = 2 // Bidirectional mode (双向持仓)
)

// String returns the string representation of PositionMode
func (p PositionMode) String() string {
	switch p {
	case PositionModeHedge:
		return "HEDGE"
	default:
		return "UNKNOWN"
	}
}

// SplitPositionMode represents the separated mode
type SplitPositionMode int

const (
	SplitPositionModeUnknown   SplitPositionMode = 0
	SplitPositionModeCombined  SplitPositionMode = 1 // Combined (合并)
	SplitPositionModeSeparated SplitPositionMode = 2 // Separated (分离)
)

// String returns the string representation of SplitPositionMode
func (s SplitPositionMode) String() string {
	switch s {
	case SplitPositionModeCombined:
		return "COMBINED"
	case SplitPositionModeSeparated:
		return "SEPARATED"
	default:
		return "UNKNOWN"
	}
}

// OrderType represents the type of order
type OrderType int

const (
	OrderTypeOpenLong   OrderType = 1 // Open long position (开多)
	OrderTypeOpenShort  OrderType = 2 // Open short position (开空)
	OrderTypeCloseLong  OrderType = 3 // Close long position (平多)
	OrderTypeCloseShort OrderType = 4 // Close short position (平空)
)

// String returns the string representation of OrderType
func (o OrderType) String() string {
	switch o {
	case OrderTypeOpenLong:
		return "OPEN_LONG"
	case OrderTypeOpenShort:
		return "OPEN_SHORT"
	case OrderTypeCloseLong:
		return "CLOSE_LONG"
	case OrderTypeCloseShort:
		return "CLOSE_SHORT"
	default:
		return "UNKNOWN"
	}
}

// OrderExecutionType represents the order execution type
type OrderExecutionType int

const (
	OrderExecNormal            OrderExecutionType = 0 // Normal order (普通委托)
	OrderExecPostOnly          OrderExecutionType = 1 // Post-only (只做maker)
	OrderExecFillOrKill        OrderExecutionType = 2 // Fill or kill (全部成交或立即取消)
	OrderExecImmediateOrCancel OrderExecutionType = 3 // Immediate or cancel (立即成交并取消剩余)
)

// String returns the string representation of OrderExecutionType
func (o OrderExecutionType) String() string {
	switch o {
	case OrderExecNormal:
		return "NORMAL"
	case OrderExecPostOnly:
		return "POST_ONLY"
	case OrderExecFillOrKill:
		return "FILL_OR_KILL"
	case OrderExecImmediateOrCancel:
		return "IMMEDIATE_OR_CANCEL"
	default:
		return "UNKNOWN"
	}
}

// PriceMatch represents the price matching type
type PriceMatch int

const (
	PriceMatchLimit  PriceMatch = 0 // Limit order (限价)
	PriceMatchMarket PriceMatch = 1 // Market order (市价)
)

// String returns the string representation of PriceMatch
func (p PriceMatch) String() string {
	switch p {
	case PriceMatchLimit:
		return "LIMIT"
	case PriceMatchMarket:
		return "MARKET"
	default:
		return "UNKNOWN"
	}
}

// PositionSide represents the position side
type PositionSide string

const (
	PositionSideLong  PositionSide = "LONG"  // Long position (多头)
	PositionSideShort PositionSide = "SHORT" // Short position (空头)
)

// OrderSide represents the order side
type OrderSide string

const (
	OrderSideBuy  OrderSide = "BUY"  // Buy order (买入)
	OrderSideSell OrderSide = "SELL" // Sell order (卖出)
)

// OrderStatus represents the status of an order
type OrderStatus int

const (
	OrderStatusNotTriggered OrderStatus = -1 // Not triggered (未触发)
	OrderStatusPending      OrderStatus = 0  // Pending (待成交)
	OrderStatusPartial      OrderStatus = 1  // Partially filled (部分成交)
	OrderStatusFilled       OrderStatus = 2  // Fully filled (完全成交)
	OrderStatusCanceling    OrderStatus = 3  // Canceling (撤单中)
	OrderStatusCanceled     OrderStatus = 4  // Canceled (已撤单)
)

// String returns the string representation of OrderStatus
func (o OrderStatus) String() string {
	switch o {
	case OrderStatusNotTriggered:
		return "NOT_TRIGGERED"
	case OrderStatusPending:
		return "PENDING"
	case OrderStatusPartial:
		return "PARTIAL"
	case OrderStatusFilled:
		return "FILLED"
	case OrderStatusCanceling:
		return "CANCELING"
	case OrderStatusCanceled:
		return "CANCELED"
	default:
		return "UNKNOWN"
	}
}

// Decimal represents a decimal number as a string to avoid precision loss.
// All price and quantity fields use this type.
type Decimal string

// Float64 converts the Decimal to float64
func (d Decimal) Float64() (float64, error) {
	return strconv.ParseFloat(string(d), 64)
}

// MustFloat64 converts the Decimal to float64, panics on error
func (d Decimal) MustFloat64() float64 {
	f, err := d.Float64()
	if err != nil {
		panic(fmt.Sprintf("failed to convert %s to float64: %v", d, err))
	}
	return f
}

// IsZero returns true if the decimal is zero or empty
func (d Decimal) IsZero() bool {
	return d == "" || d == "0" || d == "0.0"
}

// String returns the string representation of Decimal
func (d Decimal) String() string {
	return string(d)
}

// NewDecimal creates a new Decimal from a float64
func NewDecimal(f float64) Decimal {
	return Decimal(strconv.FormatFloat(f, 'f', -1, 64))
}

// NewDecimalFromString creates a new Decimal from a string
func NewDecimalFromString(s string) Decimal {
	return Decimal(s)
}

// PriceQty represents a price-quantity pair used in order book depth data
type PriceQty struct {
	Price    Decimal `json:"price"`    // Price level
	Quantity Decimal `json:"quantity"` // Quantity at this price level
}

// KlineInterval represents the candlestick interval
type KlineInterval string

const (
	Interval1Min   KlineInterval = "1m"
	Interval3Min   KlineInterval = "3m"
	Interval5Min   KlineInterval = "5m"
	Interval15Min  KlineInterval = "15m"
	Interval30Min  KlineInterval = "30m"
	Interval1Hour  KlineInterval = "1h"
	Interval2Hour  KlineInterval = "2h"
	Interval4Hour  KlineInterval = "4h"
	Interval6Hour  KlineInterval = "6h"
	Interval8Hour  KlineInterval = "8h"
	Interval12Hour KlineInterval = "12h"
	Interval1Day   KlineInterval = "1d"
	Interval3Day   KlineInterval = "3d"
	Interval1Week  KlineInterval = "1w"
	Interval1Month KlineInterval = "1M"
)

// Constants for API base URLs
const (
	DefaultBaseURL       = "https://api-contract.weex.com"
	DefaultWSPublicURL   = "wss://ws-contract.weex.com/v2/ws/public"
	DefaultWSPrivateURL  = "wss://ws-contract.weex.com/v2/ws/private"
	DefaultAPIPathPrefix = "/capi/v2"
)

// HTTP headers
const (
	HeaderAccessKey        = "ACCESS-KEY"
	HeaderAccessSign       = "ACCESS-SIGN"
	HeaderAccessPassphrase = "ACCESS-PASSPHRASE"
	HeaderAccessTimestamp  = "ACCESS-TIMESTAMP"
	HeaderContentType      = "Content-Type"
	HeaderLocale           = "locale"
	HeaderUserAgent        = "User-Agent"
)

// Content types
const (
	ContentTypeJSON = "application/json"
)

// Default values
const (
	DefaultLocale    = "en"
	DefaultUserAgent = "weex-contract-go-sdk/1.0.0"
)
