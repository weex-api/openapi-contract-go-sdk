// Package public provides helper functions for subscribing to public WebSocket channels.
package public

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/weex/openapi-contract-go-sdk/weex"
	"github.com/weex/openapi-contract-go-sdk/weex/websocket"
)

// TickerCallback is called when ticker data is received
type TickerCallback func(ticker *websocket.TickerData) error

// DepthCallback is called when depth/orderbook data is received
type DepthCallback func(depth *websocket.DepthData) error

// CandlestickCallback is called when candlestick/kline data is received
type CandlestickCallback func(kline *websocket.CandlestickData) error

// TradesCallback is called when trade data is received
type TradesCallback func(trades *websocket.TradesData) error

// Client provides convenient methods for subscribing to public channels
type Client struct {
	ws *websocket.Client
}

// NewClient creates a new public WebSocket client
func NewClient(config *weex.Config) *Client {
	return &Client{
		ws: websocket.NewClient(config),
	}
}

// Connect establishes the WebSocket connection
func (c *Client) Connect(ctx context.Context) error {
	return c.ws.Connect(ctx)
}

// Close closes the WebSocket connection
func (c *Client) Close() error {
	return c.ws.Close()
}

// SubscribeTicker subscribes to ticker updates for a symbol
//
// Channel format: ticker.{symbol}
// Example: ticker.cmt_btcusdt
func (c *Client) SubscribeTicker(symbol string, callback TickerCallback) error {
	channel := fmt.Sprintf("ticker.%s", symbol)

	handler := func(data []byte) error {
		var ticker websocket.TickerData
		if err := json.Unmarshal(data, &ticker); err != nil {
			return fmt.Errorf("failed to unmarshal ticker data: %w", err)
		}
		return callback(&ticker)
	}

	return c.ws.Subscribe(channel, handler)
}

// SubscribeDepth subscribes to order book depth updates for a symbol
//
// Channel format: depth.{symbol}
// Example: depth.cmt_btcusdt
func (c *Client) SubscribeDepth(symbol string, callback DepthCallback) error {
	channel := fmt.Sprintf("depth.%s", symbol)

	handler := func(data []byte) error {
		var depth websocket.DepthData
		if err := json.Unmarshal(data, &depth); err != nil {
			return fmt.Errorf("failed to unmarshal depth data: %w", err)
		}
		return callback(&depth)
	}

	return c.ws.Subscribe(channel, handler)
}

// SubscribeCandlestick subscribes to candlestick/kline updates
//
// Channel format: candlestick.{symbol}.{interval}
// Example: candlestick.cmt_btcusdt.1m
//
// Supported intervals: 1m, 5m, 15m, 30m, 1h, 4h, 1d, 1w
func (c *Client) SubscribeCandlestick(symbol, interval string, callback CandlestickCallback) error {
	channel := fmt.Sprintf("candlestick.%s.%s", symbol, interval)

	handler := func(data []byte) error {
		var kline websocket.CandlestickData
		if err := json.Unmarshal(data, &kline); err != nil {
			return fmt.Errorf("failed to unmarshal candlestick data: %w", err)
		}
		return callback(&kline)
	}

	return c.ws.Subscribe(channel, handler)
}

// SubscribeTrades subscribes to recent trades for a symbol
//
// Channel format: trades.{symbol}
// Example: trades.cmt_btcusdt
func (c *Client) SubscribeTrades(symbol string, callback TradesCallback) error {
	channel := fmt.Sprintf("trades.%s", symbol)

	handler := func(data []byte) error {
		var trades websocket.TradesData
		if err := json.Unmarshal(data, &trades); err != nil {
			return fmt.Errorf("failed to unmarshal trades data: %w", err)
		}
		return callback(&trades)
	}

	return c.ws.Subscribe(channel, handler)
}

// Unsubscribe unsubscribes from a channel
func (c *Client) Unsubscribe(channel string) error {
	return c.ws.Unsubscribe(channel)
}

// UnsubscribeTicker unsubscribes from ticker updates
func (c *Client) UnsubscribeTicker(symbol string) error {
	channel := fmt.Sprintf("ticker.%s", symbol)
	return c.ws.Unsubscribe(channel)
}

// UnsubscribeDepth unsubscribes from depth updates
func (c *Client) UnsubscribeDepth(symbol string) error {
	channel := fmt.Sprintf("depth.%s", symbol)
	return c.ws.Unsubscribe(channel)
}

// UnsubscribeCandlestick unsubscribes from candlestick updates
func (c *Client) UnsubscribeCandlestick(symbol, interval string) error {
	channel := fmt.Sprintf("candlestick.%s.%s", symbol, interval)
	return c.ws.Unsubscribe(channel)
}

// UnsubscribeTrades unsubscribes from trades updates
func (c *Client) UnsubscribeTrades(symbol string) error {
	channel := fmt.Sprintf("trades.%s", symbol)
	return c.ws.Unsubscribe(channel)
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
