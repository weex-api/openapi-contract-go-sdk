// Package trade provides trading API endpoints
package trade

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/weex-api/openapi-contract-go-sdk/weex/rest"
)

// Service provides access to trading API endpoints
type Service struct {
	client *rest.Client
}

// NewService creates a new trade service
func NewService(client *rest.Client) *Service {
	return &Service{client: client}
}

// PlaceOrder places a new order
// POST /trade/order
// Weight(IP): 10, Weight(UID): 5
//
// Reference: /contract/Transaction_API/PlaceOrder.md
func (s *Service) PlaceOrder(ctx context.Context, req *PlaceOrderRequest) (*PlaceOrderResponse, error) {
	path := "/trade/order"

	var response PlaceOrderResponse
	err := s.client.Post(ctx, path, req, &response, 10, 5)
	return &response, err
}

// PlaceOrdersBatch places multiple orders in a batch
// POST /trade/orders/batch
// Weight(IP): 20, Weight(UID): 10
//
// Reference: /contract/Transaction_API/PlaceOrdersBatch.md
func (s *Service) PlaceOrdersBatch(ctx context.Context, req *PlaceOrdersBatchRequest) (*PlaceOrdersBatchResponse, error) {
	path := "/trade/orders/batch"

	// Validate: max 10 orders
	if len(req.Orders) > 10 {
		return nil, fmt.Errorf("maximum 10 orders allowed in batch, got %d", len(req.Orders))
	}

	var response PlaceOrdersBatchResponse
	err := s.client.Post(ctx, path, req, &response, 20, 10)
	return &response, err
}

// PlacePendingOrder places a pending/trigger order
// POST /trade/order/pending
// Weight(IP): 10, Weight(UID): 5
//
// Reference: /contract/Transaction_API/PlacePendingOrder.md
func (s *Service) PlacePendingOrder(ctx context.Context, req *PlacePendingOrderRequest) (*PlaceOrderResponse, error) {
	path := "/trade/order/pending"

	var response PlaceOrderResponse
	err := s.client.Post(ctx, path, req, &response, 10, 5)
	return &response, err
}

// PlaceTpSlOrder places a take profit/stop loss order
// POST /trade/order/tpsl
// Weight(IP): 10, Weight(UID): 5
//
// Reference: /contract/Transaction_API/PlaceTpSlOrder.md
func (s *Service) PlaceTpSlOrder(ctx context.Context, req *PlaceTpSlOrderRequest) (*PlaceOrderResponse, error) {
	path := "/trade/order/tpsl"

	var response PlaceOrderResponse
	err := s.client.Post(ctx, path, req, &response, 10, 5)
	return &response, err
}

// CancelOrder cancels an order
// POST /trade/order/cancel
// Weight(IP): 10, Weight(UID): 5
//
// Reference: /contract/Transaction_API/CancelOrder.md
func (s *Service) CancelOrder(ctx context.Context, req *CancelOrderRequest) (*CancelOrderResponse, error) {
	path := "/trade/order/cancel"

	// Validate: either orderId or clientOid required
	if req.OrderId == "" && req.ClientOid == "" {
		return nil, fmt.Errorf("either orderId or clientOid is required")
	}

	var response CancelOrderResponse
	err := s.client.Post(ctx, path, req, &response, 10, 5)
	return &response, err
}

// CancelOrdersBatch cancels multiple orders in a batch
// POST /trade/orders/cancel/batch
// Weight(IP): 20, Weight(UID): 10
//
// Reference: /contract/Transaction_API/CancelOrdersBatch.md
func (s *Service) CancelOrdersBatch(ctx context.Context, req *CancelOrdersBatchRequest) (*PlaceOrdersBatchResponse, error) {
	path := "/trade/orders/cancel/batch"

	// Validate: max 10 orders
	if len(req.Orders) > 10 {
		return nil, fmt.Errorf("maximum 10 orders allowed in batch, got %d", len(req.Orders))
	}

	var response PlaceOrdersBatchResponse
	err := s.client.Post(ctx, path, req, &response, 20, 10)
	return &response, err
}

// CancelAllOrders cancels all orders for a symbol
// POST /trade/orders/cancel/all
// Weight(IP): 20, Weight(UID): 10
//
// Reference: /contract/Transaction_API/CancelAllOrders.md
func (s *Service) CancelAllOrders(ctx context.Context, req *CancelAllOrdersRequest) (*CancelAllOrdersResponse, error) {
	path := "/trade/orders/cancel/all"

	var response CancelAllOrdersResponse
	err := s.client.Post(ctx, path, req, &response, 20, 10)
	return &response, err
}

// CancelPendingOrder cancels a pending/trigger order
// POST /trade/order/pending/cancel
// Weight(IP): 10, Weight(UID): 5
//
// Reference: /contract/Transaction_API/CancelPendingOrder.md
func (s *Service) CancelPendingOrder(ctx context.Context, req *CancelOrderRequest) (*CancelOrderResponse, error) {
	path := "/trade/order/pending/cancel"

	var response CancelOrderResponse
	err := s.client.Post(ctx, path, req, &response, 10, 5)
	return &response, err
}

// ModifyTpSlOrder modifies a take profit/stop loss order
// PUT /trade/order/tpsl
// Weight(IP): 10, Weight(UID): 5
//
// Reference: /contract/Transaction_API/ModifyTpSlOrder.md
func (s *Service) ModifyTpSlOrder(ctx context.Context, req *PlaceTpSlOrderRequest) (*PlaceOrderResponse, error) {
	path := "/trade/order/tpsl"

	var response PlaceOrderResponse
	err := s.client.Put(ctx, path, req, &response, 10, 5)
	return &response, err
}

// ClosePositions closes positions
// POST /trade/position/close
// Weight(IP): 10, Weight(UID): 5
//
// Reference: /contract/Transaction_API/ClosePositions.md
func (s *Service) ClosePositions(ctx context.Context, req *ClosePositionsRequest) (*ClosePositionsResponse, error) {
	path := "/trade/position/close"

	var response ClosePositionsResponse
	err := s.client.Post(ctx, path, req, &response, 10, 5)
	return &response, err
}

// GetCurrentOrderStatus gets current order status (open orders)
// GET /trade/orders/current
// Weight(IP): 20, Weight(UID): 10
//
// Reference: /contract/Transaction_API/GetCurrentOrderStatus.md
func (s *Service) GetCurrentOrderStatus(ctx context.Context, req *GetOrdersRequest) (*OrdersResponse, error) {
	params := url.Values{}

	if req != nil {
		if req.Symbol != "" {
			params.Set("symbol", req.Symbol)
		}
		if req.OrderId != "" {
			params.Set("orderId", req.OrderId)
		}
		if req.ClientOid != "" {
			params.Set("clientOid", req.ClientOid)
		}
		if req.Type > 0 {
			params.Set("type", strconv.Itoa(req.Type))
		}
		if req.State >= 0 {
			params.Set("state", strconv.Itoa(req.State))
		}
		if req.Limit > 0 {
			params.Set("limit", strconv.Itoa(req.Limit))
		}
	}

	path := "/trade/orders/current"
	if len(params) > 0 {
		path = path + "?" + params.Encode()
	}

	var response OrdersResponse
	err := s.client.Get(ctx, path, &response, 20, 10)
	return &response, err
}

// GetSingleOrderInfo gets single order information
// GET /trade/order
// Weight(IP): 5, Weight(UID): 2
//
// Reference: /contract/Transaction_API/GetSingleOrderInfo.md
func (s *Service) GetSingleOrderInfo(ctx context.Context, orderId, clientOid, symbol string) (*Order, error) {
	params := url.Values{}
	params.Set("symbol", symbol)

	if orderId != "" {
		params.Set("orderId", orderId)
	} else if clientOid != "" {
		params.Set("clientOid", clientOid)
	} else {
		return nil, fmt.Errorf("either orderId or clientOid is required")
	}

	path := "/trade/order?" + params.Encode()

	var order Order
	err := s.client.Get(ctx, path, &order, 5, 2)
	return &order, err
}

// GetOrderHistory gets order history (completed orders)
// GET /trade/orders/history
// Weight(IP): 20, Weight(UID): 10
//
// Reference: /contract/Transaction_API/GetOrderHistory.md
func (s *Service) GetOrderHistory(ctx context.Context, req *GetOrdersRequest) (*OrdersResponse, error) {
	params := url.Values{}

	if req != nil {
		if req.Symbol != "" {
			params.Set("symbol", req.Symbol)
		}
		if req.Type > 0 {
			params.Set("type", strconv.Itoa(req.Type))
		}
		if req.State >= 0 {
			params.Set("state", strconv.Itoa(req.State))
		}
		if req.StartTime > 0 {
			params.Set("startTime", strconv.FormatInt(req.StartTime, 10))
		}
		if req.EndTime > 0 {
			params.Set("endTime", strconv.FormatInt(req.EndTime, 10))
		}
		if req.Limit > 0 {
			params.Set("limit", strconv.Itoa(req.Limit))
		}
	}

	path := "/trade/orders/history"
	if len(params) > 0 {
		path = path + "?" + params.Encode()
	}

	var response OrdersResponse
	err := s.client.Get(ctx, path, &response, 20, 10)
	return &response, err
}

// GetCurrentPendingOrders gets current pending/trigger orders
// GET /trade/orders/pending/current
// Weight(IP): 20, Weight(UID): 10
//
// Reference: /contract/Transaction_API/GetCurrentPendingOrders.md
func (s *Service) GetCurrentPendingOrders(ctx context.Context, req *GetPendingOrdersRequest) (*PendingOrdersResponse, error) {
	params := url.Values{}

	if req != nil {
		if req.Symbol != "" {
			params.Set("symbol", req.Symbol)
		}
		if req.Limit > 0 {
			params.Set("limit", strconv.Itoa(req.Limit))
		}
	}

	path := "/trade/orders/pending/current"
	if len(params) > 0 {
		path = path + "?" + params.Encode()
	}

	var response PendingOrdersResponse
	err := s.client.Get(ctx, path, &response, 20, 10)
	return &response, err
}

// GetHistoricalPendingOrders gets historical pending/trigger orders
// GET /trade/orders/pending/history
// Weight(IP): 20, Weight(UID): 10
//
// Reference: /contract/Transaction_API/GetHistoricalPendingOrders.md
func (s *Service) GetHistoricalPendingOrders(ctx context.Context, req *GetPendingOrdersRequest) (*PendingOrdersResponse, error) {
	params := url.Values{}

	if req != nil {
		if req.Symbol != "" {
			params.Set("symbol", req.Symbol)
		}
		if req.StartTime > 0 {
			params.Set("startTime", strconv.FormatInt(req.StartTime, 10))
		}
		if req.EndTime > 0 {
			params.Set("endTime", strconv.FormatInt(req.EndTime, 10))
		}
		if req.Limit > 0 {
			params.Set("limit", strconv.Itoa(req.Limit))
		}
	}

	path := "/trade/orders/pending/history"
	if len(params) > 0 {
		path = path + "?" + params.Encode()
	}

	var response PendingOrdersResponse
	err := s.client.Get(ctx, path, &response, 20, 10)
	return &response, err
}

// GetTradeDetails gets trade fill details
// GET /trade/fills
// Weight(IP): 20, Weight(UID): 10
//
// Reference: /contract/Transaction_API/GetTradeDetails.md
func (s *Service) GetTradeDetails(ctx context.Context, req *GetFillsRequest) (*FillsResponse, error) {
	params := url.Values{}

	if req != nil {
		if req.Symbol != "" {
			params.Set("symbol", req.Symbol)
		}
		if req.OrderId != "" {
			params.Set("orderId", req.OrderId)
		}
		if req.StartTime > 0 {
			params.Set("startTime", strconv.FormatInt(req.StartTime, 10))
		}
		if req.EndTime > 0 {
			params.Set("endTime", strconv.FormatInt(req.EndTime, 10))
		}
		if req.Limit > 0 {
			params.Set("limit", strconv.Itoa(req.Limit))
		}
	}

	path := "/trade/fills"
	if len(params) > 0 {
		path = path + "?" + params.Encode()
	}

	var response FillsResponse
	err := s.client.Get(ctx, path, &response, 20, 10)
	return &response, err
}

// Validation helpers

// ValidateClientOid validates client order ID
func ValidateClientOid(clientOid string) error {
	if clientOid == "" {
		return fmt.Errorf("clientOid cannot be empty")
	}
	if len(clientOid) > 64 {
		return fmt.Errorf("clientOid cannot exceed 64 characters")
	}
	return nil
}

// ValidateOrderType validates order type
func ValidateOrderType(orderType int) error {
	if orderType < 1 || orderType > 4 {
		return fmt.Errorf("invalid order type: %d (must be 1-4)", orderType)
	}
	return nil
}

// ValidateOrderExecutionType validates order execution type
func ValidateOrderExecutionType(execType int) error {
	if execType < 0 || execType > 3 {
		return fmt.Errorf("invalid order execution type: %d (must be 0-3)", execType)
	}
	return nil
}

// ValidatePriceMatch validates price match type
func ValidatePriceMatch(match int) error {
	if match != 0 && match != 1 {
		return fmt.Errorf("invalid price match: %d (must be 0 or 1)", match)
	}
	return nil
}
