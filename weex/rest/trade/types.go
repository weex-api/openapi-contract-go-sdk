package trade

// PlaceOrderRequest is the request for PlaceOrder
type PlaceOrderRequest struct {
	Symbol                string `json:"symbol"`                          // Required: Trading pair
	ClientOid             string `json:"client_oid"`                      // Required: Custom order ID (max 40 chars)
	Size                  string `json:"size"`                            // Required: Order quantity
	Type                  string `json:"type"`                            // Required: 1:Open long, 2:Open short, 3:Close long, 4:Close short
	OrderType             string `json:"order_type"`                      // Required: 0:Normal, 1:Post-Only, 2:FOK, 3:IOC
	MatchPrice            string `json:"match_price"`                     // Required: 0:Limit price, 1:Market price
	Price                 string `json:"price"`                           // Required: Order price (for limit orders)
	PresetTakeProfitPrice string `json:"presetTakeProfitPrice,omitempty"` // Optional: Preset take-profit price
	PresetStopLossPrice   string `json:"presetStopLossPrice,omitempty"`   // Optional: Preset stop-loss price
	MarginMode            int    `json:"marginMode,omitempty"`            // Optional: 1:Cross, 3:Isolated (default 1)
}

// PlaceOrderResponse is the response for PlaceOrder
type PlaceOrderResponse struct {
	ClientOid string `json:"client_oid"` // Client-generated order identifier
	OrderId   string `json:"order_id"`   // Order ID
}

// BatchOrderRequest is the individual order in batch request
type BatchOrderRequest struct {
	ClientOid             string `json:"client_oid"`                      // Required: Custom order ID
	Size                  string `json:"size"`                            // Required: Order quantity
	Type                  string `json:"type"`                            // Required: 1-4
	OrderType             string `json:"order_type"`                      // Required: 0-3
	MatchPrice            string `json:"match_price"`                     // Required: 0:Limit, 1:Market
	Price                 string `json:"price"`                           // Required: Order price
	PresetTakeProfitPrice string `json:"presetTakeProfitPrice,omitempty"` // Optional
	PresetStopLossPrice   string `json:"presetStopLossPrice,omitempty"`   // Optional
}

// PlaceBatchOrdersRequest is the request for batch orders
type PlaceBatchOrdersRequest struct {
	Symbol        string              `json:"symbol"`               // Required: Trading pair
	MarginMode    int                 `json:"marginMode,omitempty"` // Optional: 1:Cross, 3:Isolated
	OrderDataList []BatchOrderRequest `json:"orderDataList"`        // Required: Max 20 orders
}

// BatchOrderInfo represents single order result in batch response
type BatchOrderInfo struct {
	OrderId      string `json:"order_id"`      // Order ID
	ClientOid    string `json:"client_oid"`    // Client order ID
	Result       bool   `json:"result"`        // Order status
	ErrorCode    string `json:"error_code"`    // Error code if failed
	ErrorMessage string `json:"error_message"` // Error message if failed
}

// PlaceBatchOrdersResponse is the response for batch orders
type PlaceBatchOrdersResponse struct {
	OrderInfo []BatchOrderInfo `json:"order_info"` // Order list
	Result    bool             `json:"result"`     // Request result
}

// CancelOrderRequest is the request for CancelOrder
type CancelOrderRequest struct {
	OrderId   string `json:"orderId,omitempty"`   // Order ID (either orderId or clientOid required)
	ClientOid string `json:"clientOid,omitempty"` // Client customized ID
}

// CancelOrderResponse is the response for CancelOrder
type CancelOrderResponse struct {
	OrderId   string `json:"order_id"`   // Order ID
	ClientOid string `json:"client_oid"` // Client identifier
	Result    bool   `json:"result"`     // Cancellation status
	ErrMsg    string `json:"err_msg"`    // Error message if cancellation failed
}

// CancelBatchOrdersRequest is the request for batch cancel
type CancelBatchOrdersRequest struct {
	Ids  []string `json:"ids,omitempty"`  // Order IDs (either ids or cids required)
	Cids []string `json:"cids,omitempty"` // Client order IDs
}

// CancelOrderResult represents cancellation result for one order
type CancelOrderResult struct {
	ErrMsg    string `json:"err_msg"`    // Error message if cancellation failed
	OrderId   string `json:"order_id"`   // Order ID
	ClientOid string `json:"client_oid"` // Client order ID
	Result    bool   `json:"result"`     // Whether cancellation succeeded
}

// CancelBatchOrdersResponse is the response for batch cancel
type CancelBatchOrdersResponse struct {
	Result                bool                `json:"result"`                // Processing result
	OrderIds              []string            `json:"orderIds"`              // List of order IDs to be cancelled
	ClientOids            []string            `json:"clientOids"`            // List of client order IDs
	CancelOrderResultList []CancelOrderResult `json:"cancelOrderResultList"` // List of cancellation results
	FailInfos             []CancelOrderResult `json:"failInfos"`             // List of failed cancellation info
}

// CancelAllOrdersRequest is the request for CancelAllOrders
type CancelAllOrdersRequest struct {
	Symbol          string `json:"symbol,omitempty"` // Trading pair (optional, if not provided, cancels all)
	CancelOrderType string `json:"cancelOrderType"`  // Required: "normal" or "plan"
}

// CancelAllOrdersResultItem represents single cancellation result
type CancelAllOrdersResultItem struct {
	OrderId int64 `json:"orderId"` // Order ID
	Success bool  `json:"success"` // Whether the order was cancelled successfully
}

// PlacePendingOrderRequest is the request for PlacePendingOrder (trigger order)
type PlacePendingOrderRequest struct {
	Symbol       string `json:"symbol"`               // Required: Trading pair
	ClientOid    string `json:"client_oid"`           // Required: Custom order ID (≤40 chars)
	Size         string `json:"size"`                 // Required: Order quantity
	Type         string `json:"type"`                 // Required: 1:Open long, 2:Open short, 3:Close long, 4:Close short
	MatchType    string `json:"match_type"`           // Required: 0:Limit price, 1:Market price
	ExecutePrice string `json:"execute_price"`        // Required: Execution price
	TriggerPrice string `json:"trigger_price"`        // Required: Trigger price
	MarginMode   int    `json:"marginMode,omitempty"` // Optional: 1:Cross, 3:Isolated
}

// CancelPendingOrderRequest is the request for CancelPendingOrder
type CancelPendingOrderRequest struct {
	OrderId string `json:"orderId"` // Required: Order ID
}

// PlaceTpSlOrderRequest is the request for PlaceTpSlOrder
type PlaceTpSlOrderRequest struct {
	Symbol        string `json:"symbol"`                 // Required: Trading pair
	ClientOrderId string `json:"clientOrderId"`          // Required: Custom order ID (max 40 chars)
	PlanType      string `json:"planType"`               // Required: "profit_plan" or "loss_plan"
	TriggerPrice  string `json:"triggerPrice"`           // Required: Trigger price
	ExecutePrice  string `json:"executePrice,omitempty"` // Optional: Execution price (0 or empty = market)
	Size          string `json:"size"`                   // Required: Order quantity
	PositionSide  string `json:"positionSide"`           // Required: "long" or "short"
	MarginMode    int    `json:"marginMode,omitempty"`   // Optional: 1:Cross, 3:Isolated
}

// PlaceTpSlOrderResultItem represents single TP/SL order result
type PlaceTpSlOrderResultItem struct {
	OrderId int64 `json:"orderId"` // Order ID (0 if failed)
	Success bool  `json:"success"` // Whether the TP/SL order was placed successfully
}

// ModifyTpSlOrderRequest is the request for ModifyTpSlOrder
type ModifyTpSlOrderRequest struct {
	OrderId          int64  `json:"orderId"`                    // Required: Order ID
	TriggerPrice     string `json:"triggerPrice"`               // Required: New trigger price
	ExecutePrice     string `json:"executePrice,omitempty"`     // Optional: New execution price
	TriggerPriceType int    `json:"triggerPriceType,omitempty"` // Optional: 1:Last price, 3:Mark price
}

// ModifyTpSlOrderResponse is the response for ModifyTpSlOrder
type ModifyTpSlOrderResponse struct {
	Code        string `json:"code"`        // Response code, "00000" indicates success
	Msg         string `json:"msg"`         // Response message
	RequestTime int64  `json:"requestTime"` // Return time (Unix millisecond timestamp)
	Data        string `json:"data"`        // Response data
}

// ClosePositionsRequest is the request for ClosePositions
type ClosePositionsRequest struct {
	Symbol string `json:"symbol,omitempty"` // Trading pair (optional, if not provided, closes all)
}

// ClosePositionsResultItem represents single close position result
type ClosePositionsResultItem struct {
	PositionId     int64  `json:"positionId"`     // Position ID
	SuccessOrderId int64  `json:"successOrderId"` // Order ID if successful (0 if failed)
	ErrorMessage   string `json:"errorMessage"`   // Error message if failed
	Success        bool   `json:"success"`        // Whether the position was successfully closed
}

// Order represents an order (for current/history queries)
type Order struct {
	Symbol                string `json:"symbol"`                // Trading pair
	Size                  string `json:"size"`                  // Order amount
	ClientOid             string `json:"client_oid"`            // Client identifier
	CreateTime            string `json:"createTime"`            // Creation time (Unix millisecond timestamp)
	FilledQty             string `json:"filled_qty"`            // Filled quantity
	Fee                   string `json:"fee"`                   // Transaction fee
	OrderId               string `json:"order_id"`              // Order ID
	Price                 string `json:"price"`                 // Order price
	PriceAvg              string `json:"price_avg"`             // Average filled price
	Status                string `json:"status"`                // Order status
	Type                  string `json:"type"`                  // Order type
	OrderType             string `json:"order_type"`            // Order type
	TotalProfits          string `json:"totalProfits"`          // Total PnL
	Contracts             int    `json:"contracts"`             // Order size in contract units
	FilledQtyContracts    int    `json:"filledQtyContracts"`    // Filled quantity in contract units
	PresetTakeProfitPrice string `json:"presetTakeProfitPrice"` // Preset take-profit price
	PresetStopLossPrice   string `json:"presetStopLossPrice"`   // Preset stop-loss price
}

// PlanOrder represents a plan/trigger order
type PlanOrder struct {
	Symbol                string `json:"symbol"`                // Trading pair
	Size                  string `json:"size"`                  // Order amount
	ClientOid             string `json:"client_oid"`            // Client identifier
	CreateTime            string `json:"createTime"`            // Creation time (Unix millisecond timestamp)
	FilledQty             string `json:"filled_qty"`            // Filled quantity
	Fee                   string `json:"fee"`                   // Transaction fee
	OrderId               string `json:"order_id"`              // Order ID
	Price                 string `json:"price"`                 // Order price
	PriceAvg              string `json:"price_avg"`             // Average filled price
	Status                string `json:"status"`                // Order status
	Type                  string `json:"type"`                  // Order type
	OrderType             string `json:"order_type"`            // Order type
	TotalProfits          string `json:"totalProfits"`          // Total PnL
	TriggerPrice          string `json:"triggerPrice"`          // Trigger price
	TriggerPriceType      string `json:"triggerPriceType"`      // Trigger price type
	TriggerTime           string `json:"triggerTime"`           // Trigger time (Unix millisecond timestamp)
	PresetTakeProfitPrice string `json:"presetTakeProfitPrice"` // Preset take-profit price
	PresetStopLossPrice   string `json:"presetStopLossPrice"`   // Preset stop-loss price
}

// Fill represents a trade fill
type Fill struct {
	TradeId             int64  `json:"tradeId"`             // Filled order ID
	OrderId             int64  `json:"orderId"`             // Associated order ID
	Symbol              string `json:"symbol"`              // Trading pair name
	MarginMode          string `json:"marginMode"`          // Margin mode
	SeparatedMode       string `json:"separatedMode"`       // Separated mode
	PositionSide        string `json:"positionSide"`        // Position direction
	OrderSide           string `json:"orderSide"`           // Order direction
	FillSize            string `json:"fillSize"`            // Actual filled quantity
	FillValue           string `json:"fillValue"`           // Actual filled value
	FillFee             string `json:"fillFee"`             // Actual trading fee
	LiquidateFee        string `json:"liquidateFee"`        // Closing fee
	RealizePnl          string `json:"realizePnl"`          // Actual realized PnL
	Direction           string `json:"direction"`           // Actual execution direction
	LiquidateType       string `json:"liquidateType"`       // Liquidation order type
	LegacyOrdeDirection string `json:"legacyOrdeDirection"` // Compatible with legacy order direction types
	CreatedTime         int64  `json:"createdTime"`         // Timestamp (Unix millisecond timestamp)
}

// FillsResponse is the response for GetTradeDetails
type FillsResponse struct {
	List     []Fill `json:"list"`     // Transaction details
	NextFlag bool   `json:"nextFlag"` // Whether more pages exist
	Totals   int    `json:"totals"`   // Total entries
}
