package websocket

import (
	"encoding/json"

	"github.com/weex/openapi-contract-go-sdk/weex/types"
)

// MessageType represents the type of WebSocket message
type MessageType string

const (
	MessageTypeSubscribe   MessageType = "subscribe"
	MessageTypeUnsubscribe MessageType = "unsubscribe"
	MessageTypePing        MessageType = "ping"
	MessageTypePong        MessageType = "pong"
	MessageTypeData        MessageType = "data"
	MessageTypeError       MessageType = "error"
)

// BaseMessage represents the base WebSocket message structure
type BaseMessage struct {
	Event   string          `json:"event,omitempty"`   // Event type (subscribe, unsubscribe, error, etc.)
	Channel string          `json:"channel,omitempty"` // Channel name
	Code    string          `json:"code,omitempty"`    // Error code
	Message string          `json:"msg,omitempty"`     // Error message
	Data    json.RawMessage `json:"data,omitempty"`    // Raw data payload
}

// SubscribeRequest represents a subscription request message
type SubscribeRequest struct {
	Op   string   `json:"op"`   // "subscribe"
	Args []string `json:"args"` // Channel names
}

// UnsubscribeRequest represents an unsubscribe request message
type UnsubscribeRequest struct {
	Op   string   `json:"op"`   // "unsubscribe"
	Args []string `json:"args"` // Channel names
}

// PingMessage represents a ping message
type PingMessage struct {
	Op string `json:"op"` // "ping"
}

// PongMessage represents a pong response
type PongMessage struct {
	Op string `json:"op"` // "pong"
}

// AuthRequest represents authentication request for private channels
type AuthRequest struct {
	Op   string   `json:"op"`   // "login"
	Args []string `json:"args"` // [apiKey, passphrase, timestamp, sign]
}

// Response represents a WebSocket response
type Response struct {
	Event   string `json:"event"`             // Response event type
	Channel string `json:"channel,omitempty"` // Channel name
	Code    string `json:"code,omitempty"`    // Error code (if error)
	Message string `json:"msg,omitempty"`     // Response message
}

// ================== Public Channel Data Types ==================

// TickerData represents ticker channel data
type TickerData struct {
	Channel string       `json:"channel"`
	Data    []TickerItem `json:"data"`
}

// TickerItem represents a single ticker item
type TickerItem struct {
	Symbol             string        `json:"symbol"`
	LastPrice          types.Decimal `json:"lastPrice"`
	BestBidPrice       types.Decimal `json:"bestBidPrice"`
	BestAskPrice       types.Decimal `json:"bestAskPrice"`
	High24h            types.Decimal `json:"high24h"`
	Low24h             types.Decimal `json:"low24h"`
	Volume24h          types.Decimal `json:"volume24h"`
	QuoteVolume24h     types.Decimal `json:"quoteVolume24h"`
	PriceChange24h     types.Decimal `json:"priceChange24h"`
	PriceChangePercent types.Decimal `json:"priceChangePercent24h"`
	Timestamp          int64         `json:"timestamp"`
}

// DepthData represents depth/orderbook channel data
type DepthData struct {
	Channel string      `json:"channel"`
	Data    []DepthItem `json:"data"`
}

// DepthItem represents orderbook depth data
type DepthItem struct {
	Symbol    string           `json:"symbol"`
	Bids      []types.PriceQty `json:"bids"` // [price, quantity]
	Asks      []types.PriceQty `json:"asks"` // [price, quantity]
	Timestamp int64            `json:"timestamp"`
}

// CandlestickData represents candlestick/kline channel data
type CandlestickData struct {
	Channel string            `json:"channel"`
	Data    []CandlestickItem `json:"data"`
}

// CandlestickItem represents a single candlestick
type CandlestickItem struct {
	Symbol      string        `json:"symbol"`
	Interval    string        `json:"interval"`
	OpenTime    int64         `json:"openTime"`
	CloseTime   int64         `json:"closeTime"`
	Open        types.Decimal `json:"open"`
	High        types.Decimal `json:"high"`
	Low         types.Decimal `json:"low"`
	Close       types.Decimal `json:"close"`
	Volume      types.Decimal `json:"volume"`
	QuoteVolume types.Decimal `json:"quoteVolume"`
}

// TradesData represents trades channel data
type TradesData struct {
	Channel string      `json:"channel"`
	Data    []TradeItem `json:"data"`
}

// TradeItem represents a single trade
type TradeItem struct {
	Symbol    string        `json:"symbol"`
	TradeId   string        `json:"tradeId"`
	Price     types.Decimal `json:"price"`
	Size      types.Decimal `json:"size"`
	Side      string        `json:"side"` // "buy" or "sell"
	Timestamp int64         `json:"timestamp"`
}

// ================== Private Channel Data Types ==================

// AccountData represents account balance update data
type AccountData struct {
	Channel string        `json:"channel"`
	Data    []AccountItem `json:"data"`
}

// AccountItem represents account balance information
type AccountItem struct {
	CoinName      string        `json:"coinName"`
	Available     types.Decimal `json:"available"`
	Frozen        types.Decimal `json:"frozen"`
	Equity        types.Decimal `json:"equity"`
	UnrealizedPnl types.Decimal `json:"unrealizedPnl"`
	RealizedPnl   types.Decimal `json:"realizedPnl"`
	MarginBalance types.Decimal `json:"marginBalance"`
	UpdateTime    int64         `json:"updateTime"`
}

// PositionData represents position update data
type PositionData struct {
	Channel string         `json:"channel"`
	Data    []PositionItem `json:"data"`
}

// PositionItem represents position information
type PositionItem struct {
	Symbol           string        `json:"symbol"`
	PositionSide     string        `json:"positionSide"` // "LONG" or "SHORT"
	Size             types.Decimal `json:"size"`
	AverageOpenPrice types.Decimal `json:"averageOpenPrice"`
	MarkPrice        types.Decimal `json:"markPrice"`
	LiquidatePrice   types.Decimal `json:"liquidatePrice"`
	UnrealizedPnl    types.Decimal `json:"unrealizedPnl"`
	RealizedPnl      types.Decimal `json:"realizedPnl"`
	Leverage         types.Decimal `json:"leverage"`
	MarginMode       int           `json:"marginMode"` // 1=shared, 3=isolated
	Margin           types.Decimal `json:"margin"`
	MarginRate       types.Decimal `json:"marginRate"`
	UpdateTime       int64         `json:"updateTime"`
}

// OrderData represents order update data
type OrderData struct {
	Channel string      `json:"channel"`
	Data    []OrderItem `json:"data"`
}

// OrderItem represents order information
type OrderItem struct {
	OrderId      string        `json:"orderId"`
	ClientOid    string        `json:"clientOid"`
	Symbol       string        `json:"symbol"`
	Type         int           `json:"type"` // Order type (1=open long, 2=open short, etc.)
	Side         string        `json:"side"` // "buy" or "sell"
	Price        types.Decimal `json:"price"`
	Size         types.Decimal `json:"size"`
	FilledSize   types.Decimal `json:"filledSize"`
	AvgFillPrice types.Decimal `json:"avgFillPrice"`
	State        int           `json:"state"`      // Order state
	OrderType    int           `json:"orderType"`  // Execution type (0=normal, 1=post-only, etc.)
	MatchPrice   int           `json:"matchPrice"` // 0=limit, 1=market
	MarginMode   int           `json:"marginMode"` // 1=shared, 3=isolated
	Leverage     types.Decimal `json:"leverage"`
	Fee          types.Decimal `json:"fee"`
	FeeCoin      string        `json:"feeCoin"`
	RealizedPnl  types.Decimal `json:"realizedPnl"`
	CreateTime   int64         `json:"createTime"`
	UpdateTime   int64         `json:"updateTime"`
}

// FillData represents fill/execution data
type FillData struct {
	Channel string     `json:"channel"`
	Data    []FillItem `json:"data"`
}

// FillItem represents a trade fill/execution
type FillItem struct {
	FillId      string        `json:"fillId"`
	OrderId     string        `json:"orderId"`
	ClientOid   string        `json:"clientOid"`
	Symbol      string        `json:"symbol"`
	Price       types.Decimal `json:"price"`
	Size        types.Decimal `json:"size"`
	Side        string        `json:"side"`      // "buy" or "sell"
	Liquidity   string        `json:"liquidity"` // "maker" or "taker"
	Fee         types.Decimal `json:"fee"`
	FeeCoin     string        `json:"feeCoin"`
	RealizedPnl types.Decimal `json:"realizedPnl"`
	Timestamp   int64         `json:"timestamp"`
}

// MessageHandler is a callback function for handling WebSocket messages
type MessageHandler func(data []byte) error

// ConnectionState represents the WebSocket connection state
type ConnectionState int

const (
	StateDisconnected ConnectionState = iota
	StateConnecting
	StateConnected
	StateReconnecting
)

// String returns the string representation of the connection state
func (s ConnectionState) String() string {
	switch s {
	case StateDisconnected:
		return "Disconnected"
	case StateConnecting:
		return "Connecting"
	case StateConnected:
		return "Connected"
	case StateReconnecting:
		return "Reconnecting"
	default:
		return "Unknown"
	}
}
