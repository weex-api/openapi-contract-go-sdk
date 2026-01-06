package market

import (
	"github.com/weex-api/openapi-contract-go-sdk/weex/types"
)

// ContractInfo represents contract information
type ContractInfo struct {
	Symbol              string   `json:"symbol"`              // Contract symbol (e.g., "cmt_btcusdt")
	UnderlyingIndex     string   `json:"underlying_index"`    // Underlying index (e.g., "BTC")
	QuoteCurrency       string   `json:"quote_currency"`      // Quote currency (e.g., "USDT")
	Coin                string   `json:"coin"`                // Margin token
	ContractVal         string   `json:"contract_val"`        // Contract value
	Delivery            []string `json:"delivery"`            // Delivery times
	SizeIncrement       string   `json:"size_increment"`      // Size increment
	TickSize            string   `json:"tick_size"`           // Tick size
	ForwardContractFlag bool     `json:"forwardContractFlag"` // Whether it is USDT-M futures
	PriceEndStep        float64  `json:"priceEndStep"`        // Price end step
	MinLeverage         int      `json:"minLeverage"`         // Minimum leverage
	MaxLeverage         int      `json:"maxLeverage"`         // Maximum leverage
	BuyLimitPriceRatio  string   `json:"buyLimitPriceRatio"`  // Buy limit price ratio
	SellLimitPriceRatio string   `json:"sellLimitPriceRatio"` // Sell limit price ratio
	MakerFeeRate        string   `json:"makerFeeRate"`        // Maker fee rate
	TakerFeeRate        string   `json:"takerFeeRate"`        // Taker fee rate
	MinOrderSize        string   `json:"minOrderSize"`        // Minimum order size
	MaxOrderSize        string   `json:"maxOrderSize"`        // Maximum order size
	MaxPositionSize     string   `json:"maxPositionSize"`     // Maximum position size
}

// Ticker represents ticker information
type Ticker struct {
	Symbol             string `json:"symbol"`             // Contract symbol
	Last               string `json:"last"`               // Last price
	BestAsk            string `json:"best_ask"`           // Best ask price
	BestBid            string `json:"best_bid"`           // Best bid price
	High24h            string `json:"high_24h"`           // 24h high price
	Low24h             string `json:"low_24h"`            // 24h low price
	Volume24h          string `json:"volume_24h"`         // 24h volume
	Timestamp          string `json:"timestamp"`          // Timestamp
	PriceChangePercent string `json:"priceChangePercent"` // Price change percent
	BaseVolume         string `json:"base_volume"`        // Base volume
	MarkPrice          string `json:"markPrice"`          // Mark price
	IndexPrice         string `json:"indexPrice"`         // Index price
}

// DepthEntry represents a single depth level
type DepthEntry struct {
	Price    types.Decimal `json:"price"`    // Price level
	Quantity types.Decimal `json:"quantity"` // Quantity at this price
}

// Depth represents order book depth data
type Depth struct {
	Asks      [][]string `json:"asks"`      // Sell orders (price ascending) - array of [price, quantity]
	Bids      [][]string `json:"bids"`      // Buy orders (price descending) - array of [price, quantity]
	Timestamp string     `json:"timestamp"` // Timestamp
}

// Kline represents candlestick data
// API returns array: [timestamp, open, high, low, close, base_volume, quote_volume]
type Kline []string

// Trade represents a trade record
type Trade struct {
	TicketID     string `json:"ticketId"`     // Trade ID
	Time         int64  `json:"time"`         // Trade time
	Price        string `json:"price"`        // Price
	Size         string `json:"size"`         // Size
	Value        string `json:"value"`        // Value
	Symbol       string `json:"symbol"`       // Symbol
	IsBestMatch  bool   `json:"isBestMatch"`  // Is best match
	IsBuyerMaker bool   `json:"isBuyerMaker"` // Is buyer maker
	ContractVal  string `json:"contractVal"`  // Contract value
}

// ServerTime represents server time response
type ServerTime struct {
	Epoch     string `json:"epoch"`     // Unix timestamp in seconds (decimal)
	ISO       string `json:"iso"`       // ISO 8601 format
	Timestamp int64  `json:"timestamp"` // Unix timestamp in milliseconds
}

// IndexPrice represents index price information
type IndexPrice struct {
	Symbol    string `json:"symbol"`    // Contract symbol
	Index     string `json:"index"`     // Index price
	Timestamp string `json:"timestamp"` // Timestamp
}

// FundingRate represents funding rate information
type FundingRate struct {
	Symbol       string `json:"symbol"`       // Contract symbol
	FundingRate  string `json:"fundingRate"`  // Current funding rate
	CollectCycle int64  `json:"collectCycle"` // Funding rate collection cycle (minutes)
	Timestamp    int64  `json:"timestamp"`    // Funding fee settlement time
}

// FundingRateHistory represents historical funding rate record
type FundingRateHistory struct {
	Symbol      string        `json:"symbol"`      // Contract symbol
	FundingRate types.Decimal `json:"fundingRate"` // Funding rate
	FundingTime int64         `json:"fundingTime"` // Funding time
}

// SettlementTime represents settlement time information
type SettlementTime struct {
	Symbol         string `json:"symbol"`         // Contract symbol
	SettlementTime int64  `json:"settlementTime"` // Next settlement time
}

// OpenInterest represents open interest information
type OpenInterest struct {
	Symbol       string `json:"symbol"`        // Contract symbol
	BaseVolume   string `json:"base_volume"`   // Base volume
	TargetVolume string `json:"target_volume"` // Target volume
	Timestamp    string `json:"timestamp"`     // Timestamp
}

// Request types

// GetContractsRequest is the request for GetContracts
type GetContractsRequest struct {
	Symbol string // Optional: specific contract symbol
}

// GetKlinesRequest is the request for GetKlines
type GetKlinesRequest struct {
	Symbol    string              // Required: contract symbol
	Interval  types.KlineInterval // Required: kline interval (granularity)
	Limit     int                 // Optional: number of results (default 100, max 1000)
	PriceType string              // Optional: LAST, MARK, INDEX (default: LAST)
}

// GetHistoryKlinesRequest is the request for GetHistoryKlines
type GetHistoryKlinesRequest struct {
	Symbol    string              // Required: contract symbol
	Interval  types.KlineInterval // Required: kline interval
	StartTime int64               // Required: start time (Unix timestamp in ms)
	EndTime   int64               // Required: end time (Unix timestamp in ms)
	Limit     int                 // Optional: number of results (default 500, max 1000)
}

// GetDepthRequest is the request for GetDepth
type GetDepthRequest struct {
	Symbol string // Required: contract symbol
	Limit  int    // Optional: depth levels (default 20, max 100)
}

// GetTradesRequest is the request for GetTrades
type GetTradesRequest struct {
	Symbol string // Required: contract symbol
	Limit  int    // Optional: number of trades (default 500, max 1000)
}

// GetFundingHistoryRequest is the request for GetFundingHistory
type GetFundingHistoryRequest struct {
	Symbol    string // Required: contract symbol
	StartTime int64  // Optional: start time (Unix timestamp in ms)
	EndTime   int64  // Optional: end time (Unix timestamp in ms)
	Limit     int    // Optional: number of results (default 100, max 1000)
}
