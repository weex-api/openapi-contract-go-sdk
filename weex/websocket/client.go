package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/weex-api/openapi-contract-go-sdk/weex"
)

const (
	// Default WebSocket configuration
	DefaultPingInterval    = 30 * time.Second
	DefaultPongWait        = 60 * time.Second
	DefaultReconnectDelay  = 1 * time.Second
	DefaultMaxReconnect    = 10
	DefaultWriteWait       = 10 * time.Second
	DefaultReadBufferSize  = 1024 * 1024
	DefaultWriteBufferSize = 1024 * 1024
)

// Client represents a WebSocket client for WEEX Contract API
type Client struct {
	config *weex.Config
	auth   *weex.Authenticator
	logger weex.Logger

	// Connection
	mu        sync.RWMutex
	conn      *websocket.Conn
	state     ConnectionState
	url       string
	isPrivate bool

	// Subscription management
	subscriptions *SubscriptionManager

	// Control channels
	done      chan struct{}
	reconnect chan struct{}
	writeChan chan []byte

	// Reconnection settings
	reconnectDelay time.Duration
	maxReconnect   int
	reconnectCount int

	// Heartbeat settings
	pingInterval time.Duration
	pongWait     time.Duration
	writeWait    time.Duration

	// Callbacks
	onConnect    func()
	onDisconnect func(error)
	onError      func(error)
}

// NewClient creates a new WebSocket client for public channels
func NewClient(config *weex.Config) *Client {
	return newClient(config, nil, false)
}

// NewPrivateClient creates a new WebSocket client for private channels (requires authentication)
func NewPrivateClient(config *weex.Config, auth *weex.Authenticator) *Client {
	return newClient(config, auth, true)
}

// newClient creates a new WebSocket client
func newClient(config *weex.Config, auth *weex.Authenticator, isPrivate bool) *Client {
	// Determine WebSocket URL
	var url string
	if isPrivate {
		url = config.WSPrivateURL
	} else {
		url = config.WSPublicURL
	}

	return &Client{
		config:         config,
		auth:           auth,
		logger:         config.Logger,
		state:          StateDisconnected,
		url:            url,
		isPrivate:      isPrivate,
		subscriptions:  NewSubscriptionManager(),
		done:           make(chan struct{}),
		reconnect:      make(chan struct{}, 1),
		writeChan:      make(chan []byte, 256),
		reconnectDelay: DefaultReconnectDelay,
		maxReconnect:   DefaultMaxReconnect,
		pingInterval:   DefaultPingInterval,
		pongWait:       DefaultPongWait,
		writeWait:      DefaultWriteWait,
	}
}

// Connect establishes a WebSocket connection
func (c *Client) Connect(ctx context.Context) error {
	c.mu.Lock()
	if c.state == StateConnected || c.state == StateConnecting {
		c.mu.Unlock()
		return fmt.Errorf("already connected or connecting")
	}
	c.setState(StateConnecting)
	c.mu.Unlock()

	c.logger.Info("Connecting to WebSocket: %s", c.url)

	// Create WebSocket connection
	dialer := websocket.Dialer{
		ReadBufferSize:  DefaultReadBufferSize,
		WriteBufferSize: DefaultWriteBufferSize,
	}

	conn, _, err := dialer.DialContext(ctx, c.url, nil)
	if err != nil {
		c.setState(StateDisconnected)
		return fmt.Errorf("failed to connect to WebSocket: %w", err)
	}

	c.mu.Lock()
	c.conn = conn
	c.setState(StateConnected)
	c.reconnectCount = 0
	c.mu.Unlock()

	c.logger.Info("WebSocket connected successfully")

	// Authenticate for private channels
	if c.isPrivate && c.auth != nil {
		if err := c.authenticate(); err != nil {
			c.Close()
			return fmt.Errorf("authentication failed: %w", err)
		}
		c.logger.Info("WebSocket authenticated successfully")
	}

	// Start goroutines for read/write/ping
	go c.readPump()
	go c.writePump()
	go c.pingPump()

	// Trigger onConnect callback
	if c.onConnect != nil {
		go c.onConnect()
	}

	return nil
}

// Close closes the WebSocket connection
func (c *Client) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.state == StateDisconnected {
		return nil
	}

	c.logger.Info("Closing WebSocket connection")

	close(c.done)

	if c.conn != nil {
		// Send close message
		c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		c.conn.Close()
		c.conn = nil
	}

	c.setState(StateDisconnected)
	return nil
}

// Subscribe subscribes to a channel with a message handler
func (c *Client) Subscribe(channel string, handler MessageHandler) error {
	c.mu.RLock()
	if c.state != StateConnected {
		c.mu.RUnlock()
		return fmt.Errorf("not connected")
	}
	c.mu.RUnlock()

	// Add subscription
	c.subscriptions.Add(channel, handler)

	// Send subscribe request
	req := SubscribeRequest{
		Op:   "subscribe",
		Args: []string{channel},
	}

	data, err := json.Marshal(req)
	if err != nil {
		c.subscriptions.Remove(channel)
		return fmt.Errorf("failed to marshal subscribe request: %w", err)
	}

	if err := c.write(data); err != nil {
		c.subscriptions.Remove(channel)
		return fmt.Errorf("failed to send subscribe request: %w", err)
	}

	c.logger.Info("Subscribed to channel: %s", channel)
	return nil
}

// Unsubscribe unsubscribes from a channel
func (c *Client) Unsubscribe(channel string) error {
	c.mu.RLock()
	if c.state != StateConnected {
		c.mu.RUnlock()
		return fmt.Errorf("not connected")
	}
	c.mu.RUnlock()

	// Remove subscription
	c.subscriptions.Remove(channel)

	// Send unsubscribe request
	req := UnsubscribeRequest{
		Op:   "unsubscribe",
		Args: []string{channel},
	}

	data, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal unsubscribe request: %w", err)
	}

	if err := c.write(data); err != nil {
		return fmt.Errorf("failed to send unsubscribe request: %w", err)
	}

	c.logger.Info("Unsubscribed from channel: %s", channel)
	return nil
}

// authenticate sends authentication message for private channels
func (c *Client) authenticate() error {
	timestamp := time.Now().Unix()
	path := "/users/self/verify"
	sign := c.auth.SignWebSocket(timestamp, "GET", path, "")

	req := AuthRequest{
		Op:   "login",
		Args: []string{c.auth.GetAPIKey(), c.auth.GetPassphrase(), fmt.Sprintf("%d", timestamp), sign},
	}

	data, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal auth request: %w", err)
	}

	return c.write(data)
}

// write sends data to the WebSocket connection
func (c *Client) write(data []byte) error {
	select {
	case c.writeChan <- data:
		return nil
	case <-c.done:
		return fmt.Errorf("connection closed")
	case <-time.After(c.writeWait):
		return fmt.Errorf("write timeout")
	}
}

// readPump reads messages from the WebSocket connection
func (c *Client) readPump() {
	defer func() {
		c.handleDisconnect(nil)
	}()

	c.conn.SetReadDeadline(time.Now().Add(c.pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(c.pongWait))
		return nil
	})

	for {
		select {
		case <-c.done:
			return
		default:
		}

		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				c.logger.Error("WebSocket read error: %v", err)
			}
			return
		}

		c.handleMessage(message)
	}
}

// writePump writes messages to the WebSocket connection
func (c *Client) writePump() {
	defer func() {
		c.handleDisconnect(nil)
	}()

	for {
		select {
		case <-c.done:
			return
		case message := <-c.writeChan:
			c.conn.SetWriteDeadline(time.Now().Add(c.writeWait))
			if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
				c.logger.Error("WebSocket write error: %v", err)
				return
			}
		}
	}
}

// pingPump sends periodic ping messages
func (c *Client) pingPump() {
	ticker := time.NewTicker(c.pingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-c.done:
			return
		case <-ticker.C:
			ping := PingMessage{Op: "ping"}
			data, _ := json.Marshal(ping)
			if err := c.write(data); err != nil {
				c.logger.Error("Failed to send ping: %v", err)
				return
			}
		}
	}
}

// handleMessage processes incoming WebSocket messages
func (c *Client) handleMessage(message []byte) {
	// Parse base message to determine type
	var base BaseMessage
	if err := json.Unmarshal(message, &base); err != nil {
		c.logger.Error("Failed to parse WebSocket message: %v", err)
		return
	}

	// Handle pong response
	if base.Event == "pong" {
		return
	}

	// Handle subscription response
	if base.Event == "subscribe" || base.Event == "unsubscribe" {
		if base.Code != "" && base.Code != "0" {
			c.logger.Error("Subscription error: code=%s, msg=%s", base.Code, base.Message)
			if c.onError != nil {
				go c.onError(fmt.Errorf("subscription error: %s", base.Message))
			}
		}
		return
	}

	// Handle error
	if base.Event == "error" {
		c.logger.Error("WebSocket error: code=%s, msg=%s", base.Code, base.Message)
		if c.onError != nil {
			go c.onError(fmt.Errorf("websocket error: %s", base.Message))
		}
		return
	}

	// Route to subscription handler
	if base.Channel != "" {
		if sub, exists := c.subscriptions.Get(base.Channel); exists {
			if err := sub.Handler(message); err != nil {
				c.logger.Error("Handler error for channel %s: %v", base.Channel, err)
			}
		}
	}
}

// handleDisconnect handles connection disconnection and triggers reconnection
func (c *Client) handleDisconnect(err error) {
	c.mu.Lock()
	if c.state == StateDisconnected {
		c.mu.Unlock()
		return
	}

	oldState := c.state
	c.setState(StateDisconnected)

	if c.conn != nil {
		c.conn.Close()
		c.conn = nil
	}
	c.mu.Unlock()

	c.logger.Warn("WebSocket disconnected")

	// Trigger onDisconnect callback
	if c.onDisconnect != nil && oldState == StateConnected {
		go c.onDisconnect(err)
	}

	// Trigger reconnection
	select {
	case c.reconnect <- struct{}{}:
	default:
	}

	c.attemptReconnect()
}

// attemptReconnect attempts to reconnect with exponential backoff
func (c *Client) attemptReconnect() {
	c.mu.Lock()
	if c.reconnectCount >= c.maxReconnect {
		c.mu.Unlock()
		c.logger.Error("Max reconnection attempts reached")
		return
	}
	c.reconnectCount++
	count := c.reconnectCount
	c.mu.Unlock()

	delay := c.reconnectDelay * time.Duration(count)
	if delay > 30*time.Second {
		delay = 30 * time.Second
	}

	c.logger.Info("Reconnecting in %v (attempt %d/%d)", delay, count, c.maxReconnect)
	time.Sleep(delay)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := c.Connect(ctx); err != nil {
		c.logger.Error("Reconnection failed: %v", err)
		c.attemptReconnect()
		return
	}

	// Resubscribe to all channels
	c.resubscribe()
}

// resubscribe resubscribes to all channels after reconnection
func (c *Client) resubscribe() {
	channels := c.subscriptions.GetChannels()
	if len(channels) == 0 {
		return
	}

	c.logger.Info("Resubscribing to %d channels", len(channels))

	req := SubscribeRequest{
		Op:   "subscribe",
		Args: channels,
	}

	data, err := json.Marshal(req)
	if err != nil {
		c.logger.Error("Failed to marshal resubscribe request: %v", err)
		return
	}

	if err := c.write(data); err != nil {
		c.logger.Error("Failed to send resubscribe request: %v", err)
	}
}

// setState sets the connection state
func (c *Client) setState(state ConnectionState) {
	c.state = state
	c.logger.Debug("WebSocket state changed to: %s", state.String())
}

// GetState returns the current connection state
func (c *Client) GetState() ConnectionState {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.state
}

// IsConnected returns true if connected
func (c *Client) IsConnected() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.state == StateConnected
}

// SetOnConnect sets the callback for connection events
func (c *Client) SetOnConnect(callback func()) {
	c.onConnect = callback
}

// SetOnDisconnect sets the callback for disconnection events
func (c *Client) SetOnDisconnect(callback func(error)) {
	c.onDisconnect = callback
}

// SetOnError sets the callback for error events
func (c *Client) SetOnError(callback func(error)) {
	c.onError = callback
}

// GetSubscriptions returns all active subscriptions
func (c *Client) GetSubscriptions() []string {
	return c.subscriptions.GetChannels()
}
