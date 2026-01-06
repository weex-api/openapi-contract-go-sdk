package trade

import (
	"github.com/weex/openapi-contract-go-sdk/weex/types"
)

// Order represents an order
type Order struct {
	OrderId               string        `json:"orderId"`               // Order ID
	ClientOid             string        `json:"clientOid"`             // Client order ID
	Symbol                string        `json:"symbol"`                // Contract symbol
	Type                  int           `json:"type"`                  // Order type (1:开多,2:开空,3:平多,4:平空)
	OrderType             int           `json:"orderType"`             // Order execution type (0:普通,1:只做maker,2:FOK,3:IOC)
	MatchPrice            int           `json:"matchPrice"`            // Price match (0:限价,1:市价)
	Price                 types.Decimal `json:"price"`                 // Order price
	Size                  types.Decimal `json:"size"`                  // Order size
	FilledSize            types.Decimal `json:"filledSize"`            // Filled size
	AvgPrice              types.Decimal `json:"avgPrice"`              // Average fill price
	State                 int           `json:"state"`                 // Order state
	MarginMode            int           `json:"marginMode"`            // Margin mode
	Leverage              types.Decimal `json:"leverage"`              // Leverage
	PresetTakeProfitPrice types.Decimal `json:"presetTakeProfitPrice"` // Take profit price
	PresetStopLossPrice   types.Decimal `json:"presetStopLossPrice"`   // Stop loss price
	CreateTime            int64         `json:"createTime"`            // Create time
	UpdateTime            int64         `json:"updateTime"`            // Update time
}

// PlaceOrderRequest is the request for PlaceOrder
type PlaceOrderRequest struct {
	Symbol                string        `json:"symbol"`                          // Required: contract symbol
	ClientOid             string        `json:"clientOid"`                       // Required: client order ID (unique)
	Size                  types.Decimal `json:"size"`                            // Required: order size
	Type                  int           `json:"type"`                            // Required: order type (1-4)
	OrderType             int           `json:"orderType"`                       // Required: order execution type (0-3)
	MatchPrice            int           `json:"matchPrice"`                      // Required: price match (0:limit,1:market)
	Price                 types.Decimal `json:"price,omitempty"`                 // Price (required for limit orders)
	PresetTakeProfitPrice types.Decimal `json:"presetTakeProfitPrice,omitempty"` // Optional: take profit price
	PresetStopLossPrice   types.Decimal `json:"presetStopLossPrice,omitempty"`   // Optional: stop loss price
	MarginMode            int           `json:"marginMode,omitempty"`            // Optional: margin mode
}

// PlaceOrderResponse is the response for PlaceOrder
type PlaceOrderResponse struct {
	OrderId   string `json:"orderId"`   // Order ID
	ClientOid string `json:"clientOid"` // Client order ID
}

// PlaceOrdersBatchRequest is the request for PlaceOrdersBatch
type PlaceOrdersBatchRequest struct {
	Orders []PlaceOrderRequest `json:"orders"` // Required: array of orders (max 10)
}

// PlaceOrdersBatchResponse is the response for PlaceOrdersBatch
type PlaceOrdersBatchResponse struct {
	Success []PlaceOrderResponse `json:"success"` // Successfully placed orders
	Failed  []struct {
		ClientOid string `json:"clientOid"` // Client order ID
		ErrorCode string `json:"errorCode"` // Error code
		ErrorMsg  string `json:"errorMsg"`  // Error message
	} `json:"failed"` // Failed orders
}

// PendingOrder represents a pending/trigger order
type PendingOrder struct {
	OrderId      string        `json:"orderId"`      // Order ID
	ClientOid    string        `json:"clientOid"`    // Client order ID
	Symbol       string        `json:"symbol"`       // Contract symbol
	Type         int           `json:"type"`         // Order type
	OrderType    int           `json:"orderType"`    // Order execution type
	TriggerPrice types.Decimal `json:"triggerPrice"` // Trigger price
	TriggerType  int           `json:"triggerType"`  // Trigger type (1:mark,2:last)
	ExecutePrice types.Decimal `json:"executePrice"` // Execute price
	Size         types.Decimal `json:"size"`         // Order size
	State        int           `json:"state"`        // Order state
	MarginMode   int           `json:"marginMode"`   // Margin mode
	CreateTime   int64         `json:"createTime"`   // Create time
	UpdateTime   int64         `json:"updateTime"`   // Update time
}

// PlacePendingOrderRequest is the request for PlacePendingOrder
type PlacePendingOrderRequest struct {
	Symbol       string        `json:"symbol"`               // Required: contract symbol
	ClientOid    string        `json:"clientOid"`            // Required: client order ID
	Size         types.Decimal `json:"size"`                 // Required: order size
	Type         int           `json:"type"`                 // Required: order type
	TriggerPrice types.Decimal `json:"triggerPrice"`         // Required: trigger price
	TriggerType  int           `json:"triggerType"`          // Required: trigger type (1:mark,2:last)
	ExecutePrice types.Decimal `json:"executePrice"`         // Required: execute price (0 for market)
	MarginMode   int           `json:"marginMode,omitempty"` // Optional: margin mode
}

// TpSlOrder represents a take profit/stop loss order
type TpSlOrder struct {
	OrderId      string        `json:"orderId"`      // Order ID
	Symbol       string        `json:"symbol"`       // Contract symbol
	PositionSide string        `json:"positionSide"` // Position side
	PlanType     int           `json:"planType"`     // Plan type (1:TP,2:SL)
	TriggerPrice types.Decimal `json:"triggerPrice"` // Trigger price
	TriggerType  int           `json:"triggerType"`  // Trigger type
	Size         types.Decimal `json:"size"`         // Size
	State        int           `json:"state"`        // State
	CreateTime   int64         `json:"createTime"`   // Create time
	UpdateTime   int64         `json:"updateTime"`   // Update time
}

// PlaceTpSlOrderRequest is the request for PlaceTpSlOrder
type PlaceTpSlOrderRequest struct {
	Symbol       string        `json:"symbol"`                 // Required: contract symbol
	PositionSide string        `json:"positionSide"`           // Required: position side
	PlanType     int           `json:"planType"`               // Required: plan type (1:TP,2:SL)
	TriggerPrice types.Decimal `json:"triggerPrice"`           // Required: trigger price
	TriggerType  int           `json:"triggerType"`            // Required: trigger type (1:mark,2:last)
	Size         types.Decimal `json:"size,omitempty"`         // Optional: size (0 for全部)
	ExecutePrice types.Decimal `json:"executePrice,omitempty"` // Optional: execute price
	MarginMode   int           `json:"marginMode,omitempty"`   // Optional: margin mode
}

// CancelOrderRequest is the request for CancelOrder
type CancelOrderRequest struct {
	OrderId   string `json:"orderId,omitempty"`   // Order ID (either orderId or clientOid required)
	ClientOid string `json:"clientOid,omitempty"` // Client order ID
	Symbol    string `json:"symbol"`              // Required: contract symbol
}

// CancelOrderResponse is the response for CancelOrder
type CancelOrderResponse struct {
	OrderId   string `json:"orderId"`   // Order ID
	ClientOid string `json:"clientOid"` // Client order ID
}

// CancelOrdersBatchRequest is the request for CancelOrdersBatch
type CancelOrdersBatchRequest struct {
	Orders []CancelOrderRequest `json:"orders"` // Required: array of orders to cancel (max 10)
}

// CancelAllOrdersRequest is the request for CancelAllOrders
type CancelAllOrdersRequest struct {
	Symbol string `json:"symbol"` // Required: contract symbol
}

// CancelAllOrdersResponse is the response for CancelAllOrders
type CancelAllOrdersResponse struct {
	Success []string `json:"success"` // Successfully canceled order IDs
}

// ClosePositionsRequest is the request for ClosePositions
type ClosePositionsRequest struct {
	Symbol       string        `json:"symbol"`               // Required: contract symbol
	PositionSide string        `json:"positionSide"`         // Required: position side
	Size         types.Decimal `json:"size,omitempty"`       // Optional: size to close (0 for all)
	MarginMode   int           `json:"marginMode,omitempty"` // Optional: margin mode
}

// ClosePositionsResponse is the response for ClosePositions
type ClosePositionsResponse struct {
	OrderId string `json:"orderId"` // Order ID of the close order
}

// GetOrdersRequest is the request for GetCurrentOrderStatus and GetOrderHistory
type GetOrdersRequest struct {
	Symbol    string // Optional: contract symbol
	OrderId   string // Optional: order ID
	ClientOid string // Optional: client order ID
	Type      int    // Optional: order type
	State     int    // Optional: order state
	StartTime int64  // Optional: start time
	EndTime   int64  // Optional: end time
	Limit     int    // Optional: page size (default 100, max 500)
}

// OrdersResponse represents paginated orders response
type OrdersResponse struct {
	Orders   []Order `json:"orders"`   // Orders list
	NextFlag bool    `json:"nextFlag"` // Has next page
	Totals   int     `json:"totals"`   // Total count
}

// GetPendingOrdersRequest is the request for GetCurrentPendingOrders and GetHistoricalPendingOrders
type GetPendingOrdersRequest struct {
	Symbol    string // Optional: contract symbol
	StartTime int64  // Optional: start time
	EndTime   int64  // Optional: end time
	Limit     int    // Optional: page size
}

// PendingOrdersResponse represents paginated pending orders response
type PendingOrdersResponse struct {
	Orders   []PendingOrder `json:"orders"`   // Pending orders list
	NextFlag bool           `json:"nextFlag"` // Has next page
	Totals   int            `json:"totals"`   // Total count
}

// Fill represents a trade fill
type Fill struct {
	FillId      string        `json:"fillId"`      // Fill ID
	OrderId     string        `json:"orderId"`     // Order ID
	Symbol      string        `json:"symbol"`      // Contract symbol
	Side        string        `json:"side"`        // Side (BUY/SELL)
	Price       types.Decimal `json:"price"`       // Fill price
	Size        types.Decimal `json:"size"`        // Fill size
	Fee         types.Decimal `json:"fee"`         // Fee
	FeeCoin     string        `json:"feeCoin"`     // Fee coin
	Liquidity   string        `json:"liquidity"`   // Liquidity (maker/taker)
	RealizedPnl types.Decimal `json:"realizedPnl"` // Realized PnL
	CreateTime  int64         `json:"createTime"`  // Create time
}

// GetFillsRequest is the request for GetTradeDetails
type GetFillsRequest struct {
	Symbol    string // Optional: contract symbol
	OrderId   string // Optional: order ID
	StartTime int64  // Optional: start time
	EndTime   int64  // Optional: end time
	Limit     int    // Optional: page size
}

// FillsResponse represents paginated fills response
type FillsResponse struct {
	Fills    []Fill `json:"fills"`    // Fills list
	NextFlag bool   `json:"nextFlag"` // Has next page
	Totals   int    `json:"totals"`   // Total count
}
