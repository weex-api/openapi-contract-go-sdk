// Package private provides helper functions for subscribing to private WebSocket channels.
package private

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/weex-api/openapi-contract-go-sdk/weex"
	"github.com/weex-api/openapi-contract-go-sdk/weex/websocket"
)

// AccountCallback is called when account balance data is received
type AccountCallback func(account *websocket.AccountData) error

// PositionCallback is called when position data is received
type PositionCallback func(position *websocket.PositionData) error

// OrderCallback is called when order data is received
type OrderCallback func(order *websocket.OrderData) error

// FillCallback is called when fill/execution data is received
type FillCallback func(fill *websocket.FillData) error

// Client provides convenient methods for subscribing to private channels
type Client struct {
	ws *websocket.Client
}

// NewClient creates a new private WebSocket client (requires authentication)
func NewClient(config *weex.Config, auth *weex.Authenticator) *Client {
	return &Client{
		ws: websocket.NewPrivateClient(config, auth),
	}
}

// Connect establishes the WebSocket connection and authenticates
func (c *Client) Connect(ctx context.Context) error {
	return c.ws.Connect(ctx)
}

// Close closes the WebSocket connection
func (c *Client) Close() error {
	return c.ws.Close()
}

// SubscribeAccount subscribes to account balance updates
//
// Channel: account
// Receives updates when account balances change
func (c *Client) SubscribeAccount(callback AccountCallback) error {
	channel := "account"

	handler := func(data []byte) error {
		var account websocket.AccountData
		if err := json.Unmarshal(data, &account); err != nil {
			return fmt.Errorf("failed to unmarshal account data: %w", err)
		}
		return callback(&account)
	}

	return c.ws.Subscribe(channel, handler)
}

// SubscribePositions subscribes to position updates
//
// Channel: positions
// Receives updates when positions change
func (c *Client) SubscribePositions(callback PositionCallback) error {
	channel := "positions"

	handler := func(data []byte) error {
		var position websocket.PositionData
		if err := json.Unmarshal(data, &position); err != nil {
			return fmt.Errorf("failed to unmarshal position data: %w", err)
		}
		return callback(&position)
	}

	return c.ws.Subscribe(channel, handler)
}

// SubscribeOrders subscribes to order updates
//
// Channel: orders
// Receives updates when orders are created, updated, filled, or canceled
func (c *Client) SubscribeOrders(callback OrderCallback) error {
	channel := "orders"

	handler := func(data []byte) error {
		var order websocket.OrderData
		if err := json.Unmarshal(data, &order); err != nil {
			return fmt.Errorf("failed to unmarshal order data: %w", err)
		}
		return callback(&order)
	}

	return c.ws.Subscribe(channel, handler)
}

// SubscribeFills subscribes to fill/execution updates
//
// Channel: fill
// Receives notifications when orders are executed (filled)
func (c *Client) SubscribeFills(callback FillCallback) error {
	channel := "fill"

	handler := func(data []byte) error {
		var fill websocket.FillData
		if err := json.Unmarshal(data, &fill); err != nil {
			return fmt.Errorf("failed to unmarshal fill data: %w", err)
		}
		return callback(&fill)
	}

	return c.ws.Subscribe(channel, handler)
}

// UnsubscribeAccount unsubscribes from account updates
func (c *Client) UnsubscribeAccount() error {
	return c.ws.Unsubscribe("account")
}

// UnsubscribePositions unsubscribes from position updates
func (c *Client) UnsubscribePositions() error {
	return c.ws.Unsubscribe("positions")
}

// UnsubscribeOrders unsubscribes from order updates
func (c *Client) UnsubscribeOrders() error {
	return c.ws.Unsubscribe("orders")
}

// UnsubscribeFills unsubscribes from fill updates
func (c *Client) UnsubscribeFills() error {
	return c.ws.Unsubscribe("fill")
}

// IsConnected returns true if the WebSocket is connected
func (c *Client) IsConnected() bool {
	return c.ws.IsConnected()
}

// GetState returns the current connection state
func (c *Client) GetState() websocket.ConnectionState {
	return c.ws.GetState()
}

// SetOnConnect sets the callback for connection events
func (c *Client) SetOnConnect(callback func()) {
	c.ws.SetOnConnect(callback)
}

// SetOnDisconnect sets the callback for disconnection events
func (c *Client) SetOnDisconnect(callback func(error)) {
	c.ws.SetOnDisconnect(callback)
}

// SetOnError sets the callback for error events
func (c *Client) SetOnError(callback func(error)) {
	c.ws.SetOnError(callback)
}
