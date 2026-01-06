package websocket

import (
	"sync"
)

// Subscription represents a channel subscription
type Subscription struct {
	Channel string
	Handler MessageHandler
}

// SubscriptionManager manages WebSocket channel subscriptions
type SubscriptionManager struct {
	mu            sync.RWMutex
	subscriptions map[string]*Subscription
}

// NewSubscriptionManager creates a new subscription manager
func NewSubscriptionManager() *SubscriptionManager {
	return &SubscriptionManager{
		subscriptions: make(map[string]*Subscription),
	}
}

// Add adds a new subscription
func (sm *SubscriptionManager) Add(channel string, handler MessageHandler) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sm.subscriptions[channel] = &Subscription{
		Channel: channel,
		Handler: handler,
	}
}

// Remove removes a subscription
func (sm *SubscriptionManager) Remove(channel string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	delete(sm.subscriptions, channel)
}

// Get retrieves a subscription by channel name
func (sm *SubscriptionManager) Get(channel string) (*Subscription, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	sub, exists := sm.subscriptions[channel]
	return sub, exists
}

// GetAll returns all subscriptions
func (sm *SubscriptionManager) GetAll() []*Subscription {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	subs := make([]*Subscription, 0, len(sm.subscriptions))
	for _, sub := range sm.subscriptions {
		subs = append(subs, sub)
	}
	return subs
}

// GetChannels returns all subscribed channel names
func (sm *SubscriptionManager) GetChannels() []string {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	channels := make([]string, 0, len(sm.subscriptions))
	for channel := range sm.subscriptions {
		channels = append(channels, channel)
	}
	return channels
}

// Clear removes all subscriptions
func (sm *SubscriptionManager) Clear() {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sm.subscriptions = make(map[string]*Subscription)
}

// Count returns the number of active subscriptions
func (sm *SubscriptionManager) Count() int {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	return len(sm.subscriptions)
}

// Exists checks if a channel subscription exists
func (sm *SubscriptionManager) Exists(channel string) bool {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	_, exists := sm.subscriptions[channel]
	return exists
}
