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
// POST /capi/v2/order/placeOrder
// Weight(IP): 2, Weight(UID): 5
func (s *Service) PlaceOrder(ctx context.Context, req *PlaceOrderRequest) (*PlaceOrderResponse, error) {
	path := "/order/placeOrder"
	var response PlaceOrderResponse
	err := s.client.Post(ctx, path, req, &response, 2, 5)
	return &response, err
}

// PlaceBatchOrders places multiple orders in a batch
// POST /capi/v2/order/batchOrders
// Weight(IP): 5, Weight(UID): 10
func (s *Service) PlaceBatchOrders(ctx context.Context, req *PlaceBatchOrdersRequest) (*PlaceBatchOrdersResponse, error) {
	path := "/order/batchOrders"
	if len(req.OrderDataList) > 20 {
		return nil, fmt.Errorf("maximum 20 orders allowed in batch, got %d", len(req.OrderDataList))
	}
	var response PlaceBatchOrdersResponse
	err := s.client.Post(ctx, path, req, &response, 5, 10)
	return &response, err
}

// CancelOrder cancels an order
// POST /capi/v2/order/cancel_order
// Weight(IP): 2, Weight(UID): 3
func (s *Service) CancelOrder(ctx context.Context, req *CancelOrderRequest) (*CancelOrderResponse, error) {
	path := "/order/cancel_order"
	if req.OrderId == "" && req.ClientOid == "" {
		return nil, fmt.Errorf("either orderId or clientOid is required")
	}
	var response CancelOrderResponse
	err := s.client.Post(ctx, path, req, &response, 2, 3)
	return &response, err
}

// CancelBatchOrders cancels multiple orders in a batch
// POST /capi/v2/order/cancel_batch_orders
// Weight(IP): 5, Weight(UID): 10
func (s *Service) CancelBatchOrders(ctx context.Context, req *CancelBatchOrdersRequest) (*CancelBatchOrdersResponse, error) {
	path := "/order/cancel_batch_orders"
	if len(req.Ids) == 0 && len(req.Cids) == 0 {
		return nil, fmt.Errorf("either ids or cids is required")
	}
	var response CancelBatchOrdersResponse
	err := s.client.Post(ctx, path, req, &response, 5, 10)
	return &response, err
}

// CancelAllOrders cancels all orders
// POST /capi/v2/order/cancelAllOrders
// Weight(IP): 40, Weight(UID): 50
func (s *Service) CancelAllOrders(ctx context.Context, req *CancelAllOrdersRequest) ([]CancelAllOrdersResultItem, error) {
	path := "/order/cancelAllOrders"
	var response []CancelAllOrdersResultItem
	err := s.client.Post(ctx, path, req, &response, 40, 50)
	return response, err
}

// PlacePendingOrder places a pending/trigger order
// POST /capi/v2/order/plan_order
// Weight(IP): 2, Weight(UID): 5
func (s *Service) PlacePendingOrder(ctx context.Context, req *PlacePendingOrderRequest) (*PlaceOrderResponse, error) {
	path := "/order/plan_order"
	var response PlaceOrderResponse
	err := s.client.Post(ctx, path, req, &response, 2, 5)
	return &response, err
}

// CancelPendingOrder cancels a pending/trigger order
// POST /capi/v2/order/cancel_plan
// Weight(IP): 2, Weight(UID): 3
func (s *Service) CancelPendingOrder(ctx context.Context, req *CancelPendingOrderRequest) (*CancelOrderResponse, error) {
	path := "/order/cancel_plan"
	var response CancelOrderResponse
	err := s.client.Post(ctx, path, req, &response, 2, 3)
	return &response, err
}

// GetCurrentPendingOrders gets current pending/trigger orders
// GET /capi/v2/order/currentPlan
// Weight(IP): 3, Weight(UID): 3
func (s *Service) GetCurrentPendingOrders(ctx context.Context, symbol string, orderId int64, startTime, endTime int64, limit, page int) ([]PlanOrder, error) {
	params := url.Values{}
	if symbol != "" {
		params.Set("symbol", symbol)
	}
	if orderId > 0 {
		params.Set("orderId", strconv.FormatInt(orderId, 10))
	}
	if startTime > 0 {
		params.Set("startTime", strconv.FormatInt(startTime, 10))
	}
	if endTime > 0 {
		params.Set("endTime", strconv.FormatInt(endTime, 10))
	}
	if limit > 0 {
		params.Set("limit", strconv.Itoa(limit))
	}
	if page > 0 {
		params.Set("page", strconv.Itoa(page))
	}

	path := "/order/currentPlan"
	if len(params) > 0 {
		path = path + "?" + params.Encode()
	}

	var orders []PlanOrder
	err := s.client.Get(ctx, path, &orders, 3, 3)
	return orders, err
}

// PlaceTpSlOrder places a take profit/stop loss order
// POST /capi/v2/order/placeTpSlOrder
// Weight(IP): 2, Weight(UID): 5
func (s *Service) PlaceTpSlOrder(ctx context.Context, req *PlaceTpSlOrderRequest) ([]PlaceTpSlOrderResultItem, error) {
	path := "/order/placeTpSlOrder"
	var response []PlaceTpSlOrderResultItem
	err := s.client.Post(ctx, path, req, &response, 2, 5)
	return response, err
}

// ModifyTpSlOrder modifies a take profit/stop loss order
// POST /capi/v2/order/modifyTpSlOrder
// Weight(IP): 2, Weight(UID): 5
func (s *Service) ModifyTpSlOrder(ctx context.Context, req *ModifyTpSlOrderRequest) (*ModifyTpSlOrderResponse, error) {
	path := "/order/modifyTpSlOrder"
	var response ModifyTpSlOrderResponse
	err := s.client.Post(ctx, path, req, &response, 2, 5)
	return &response, err
}

// ClosePositions closes all positions
// POST /capi/v2/order/closePositions
// Weight(IP): 40, Weight(UID): 50
func (s *Service) ClosePositions(ctx context.Context, req *ClosePositionsRequest) ([]ClosePositionsResultItem, error) {
	path := "/order/closePositions"
	var response []ClosePositionsResultItem
	err := s.client.Post(ctx, path, req, &response, 40, 50)
	return response, err
}

// GetSingleOrderInfo gets single order information
// GET /capi/v2/order/detail
// Weight(IP): 2, Weight(UID): 2
func (s *Service) GetSingleOrderInfo(ctx context.Context, orderId string) (*Order, error) {
	params := url.Values{}
	params.Set("orderId", orderId)
	path := "/order/detail?" + params.Encode()

	var order Order
	err := s.client.Get(ctx, path, &order, 2, 2)
	return &order, err
}

// GetOrderHistory gets order history (completed orders)
// GET /capi/v2/order/history
// Weight(IP): 10, Weight(UID): 10
func (s *Service) GetOrderHistory(ctx context.Context, symbol string, pageSize int, createDate, endCreateDate int64) ([]Order, error) {
	params := url.Values{}
	if symbol != "" {
		params.Set("symbol", symbol)
	}
	if pageSize > 0 {
		params.Set("pageSize", strconv.Itoa(pageSize))
	}
	if createDate > 0 {
		params.Set("createDate", strconv.FormatInt(createDate, 10))
	}
	if endCreateDate > 0 {
		params.Set("endCreateDate", strconv.FormatInt(endCreateDate, 10))
	}

	path := "/order/history"
	if len(params) > 0 {
		path = path + "?" + params.Encode()
	}

	var orders []Order
	err := s.client.Get(ctx, path, &orders, 10, 10)
	return orders, err
}

// GetCurrentOrderStatus gets current order status (open orders)
// GET /capi/v2/order/current
// Weight(IP): 2, Weight(UID): 2
func (s *Service) GetCurrentOrderStatus(ctx context.Context, symbol string, orderId int64, startTime, endTime int64, limit, page int) ([]Order, error) {
	params := url.Values{}
	if symbol != "" {
		params.Set("symbol", symbol)
	}
	if orderId > 0 {
		params.Set("orderId", strconv.FormatInt(orderId, 10))
	}
	if startTime > 0 {
		params.Set("startTime", strconv.FormatInt(startTime, 10))
	}
	if endTime > 0 {
		params.Set("endTime", strconv.FormatInt(endTime, 10))
	}
	if limit > 0 {
		params.Set("limit", strconv.Itoa(limit))
	}
	if page > 0 {
		params.Set("page", strconv.Itoa(page))
	}

	path := "/order/current"
	if len(params) > 0 {
		path = path + "?" + params.Encode()
	}

	var orders []Order
	err := s.client.Get(ctx, path, &orders, 2, 2)
	return orders, err
}

// GetTradeDetails gets trade fill details
// GET /capi/v2/order/fills
// Weight(IP): 5, Weight(UID): 5
func (s *Service) GetTradeDetails(ctx context.Context, symbol string, orderId int64, startTime, endTime int64, limit int) (*FillsResponse, error) {
	params := url.Values{}
	if symbol != "" {
		params.Set("symbol", symbol)
	}
	if orderId > 0 {
		params.Set("orderId", strconv.FormatInt(orderId, 10))
	}
	if startTime > 0 {
		params.Set("startTime", strconv.FormatInt(startTime, 10))
	}
	if endTime > 0 {
		params.Set("endTime", strconv.FormatInt(endTime, 10))
	}
	if limit > 0 {
		params.Set("limit", strconv.Itoa(limit))
	}

	path := "/order/fills"
	if len(params) > 0 {
		path = path + "?" + params.Encode()
	}

	var response FillsResponse
	err := s.client.Get(ctx, path, &response, 5, 5)
	if err != nil {
		// Empty response case
		return &FillsResponse{List: []Fill{}, NextFlag: false, Totals: 0}, nil
	}
	return &response, nil
}
