package market

import (
	"github.com/weex-api/openapi-contract-go-sdk/weex/types"
)

// ContractInfo represents contract information
type ContractInfo struct {
	Symbol                string        `json:"symbol"`                // Contract symbol (e.g., "cmt_btcusdt")
	BaseCoin              string        `json:"baseCoin"`              // Base coin (e.g., "BTC")
	QuoteCoin             string        `json:"quoteCoin"`             // Quote coin (e.g., "USDT")
	SettleCoin            string        `json:"settleCoin"`            // Settlement coin
	ContractSize          types.Decimal `json:"contractSize"`          // Contract size
	DeliveryDate          int64         `json:"deliveryDate"`          // Delivery date (Unix timestamp)
	MakerFee              types.Decimal `json:"makerFee"`              // Maker fee rate
	TakerFee              types.Decimal `json:"takerFee"`              // Taker fee rate
	PriceTick             types.Decimal `json:"priceTick"`             // Minimum price increment
	LotSize               types.Decimal `json:"lotSize"`               // Minimum order size
	ContractType          string        `json:"contractType"`          // Contract type (e.g., "PERPETUAL")
	Status                int           `json:"status"`                // Contract status
	MaxLeverage           types.Decimal `json:"maxLeverage"`           // Maximum leverage
	MinLeverage           types.Decimal `json:"minLeverage"`           // Minimum leverage
	MaintenanceMarginRate types.Decimal `json:"maintenanceMarginRate"` // Maintenance margin rate
	InitialMarginRate     types.Decimal `json:"initialMarginRate"`     // Initial margin rate
}

// Ticker represents ticker information
type Ticker struct {
	Symbol             string        `json:"symbol"`             // Contract symbol
	PriceChange        types.Decimal `json:"priceChange"`        // Price change
	PriceChangePercent types.Decimal `json:"priceChangePercent"` // Price change percentage
	LastPrice          types.Decimal `json:"lastPrice"`          // Last traded price
	MarkPrice          types.Decimal `json:"markPrice"`          // Mark price
	IndexPrice         types.Decimal `json:"indexPrice"`         // Index price
	OpenPrice          types.Decimal `json:"openPrice"`          // Open price (24h)
	HighPrice          types.Decimal `json:"highPrice"`          // High price (24h)
	LowPrice           types.Decimal `json:"lowPrice"`           // Low price (24h)
	Volume             types.Decimal `json:"volume"`             // Trading volume (24h)
	QuoteVolume        types.Decimal `json:"quoteVolume"`        // Quote volume (24h)
	OpenTime           int64         `json:"openTime"`           // Open time
	CloseTime          int64         `json:"closeTime"`          // Close time
	FirstTradeId       int64         `json:"firstTradeId"`       // First trade ID
	TradeCount         int64         `json:"tradeCount"`         // Number of trades
	BidPrice           types.Decimal `json:"bidPrice"`           // Best bid price
	BidQty             types.Decimal `json:"bidQty"`             // Best bid quantity
	AskPrice           types.Decimal `json:"askPrice"`           // Best ask price
	AskQty             types.Decimal `json:"askQty"`             // Best ask quantity
}

// DepthEntry represents a single depth level
type DepthEntry struct {
	Price    types.Decimal `json:"price"`    // Price level
	Quantity types.Decimal `json:"quantity"` // Quantity at this price
}

// Depth represents order book depth data
type Depth struct {
	Symbol string       `json:"symbol"` // Contract symbol
	Bids   []DepthEntry `json:"bids"`   // Buy orders (price descending)
	Asks   []DepthEntry `json:"asks"`   // Sell orders (price ascending)
	Time   int64        `json:"time"`   // Timestamp
}

// Kline represents candlestick data
type Kline struct {
	OpenTime            int64         `json:"openTime"`            // Open time
	Open                types.Decimal `json:"open"`                // Open price
	High                types.Decimal `json:"high"`                // High price
	Low                 types.Decimal `json:"low"`                 // Low price
	Close               types.Decimal `json:"close"`               // Close price
	Volume              types.Decimal `json:"volume"`              // Trading volume
	CloseTime           int64         `json:"closeTime"`           // Close time
	QuoteVolume         types.Decimal `json:"quoteVolume"`         // Quote asset volume
	TradeCount          int64         `json:"tradeCount"`          // Number of trades
	TakerBuyBaseVolume  types.Decimal `json:"takerBuyBaseVolume"`  // Taker buy base volume
	TakerBuyQuoteVolume types.Decimal `json:"takerBuyQuoteVolume"` // Taker buy quote volume
}

// Trade represents a trade record
type Trade struct {
	ID           int64         `json:"id"`           // Trade ID
	Price        types.Decimal `json:"price"`        // Trade price
	Qty          types.Decimal `json:"qty"`          // Trade quantity
	QuoteQty     types.Decimal `json:"quoteQty"`     // Quote quantity
	Time         int64         `json:"time"`         // Trade time
	IsBuyerMaker bool          `json:"isBuyerMaker"` // Whether buyer is maker
}

// ServerTime represents server time response
type ServerTime struct {
	ServerTime int64 `json:"serverTime"` // Server timestamp in milliseconds
}

// IndexPrice represents index price information
type IndexPrice struct {
	Symbol     string        `json:"symbol"`     // Contract symbol
	IndexPrice types.Decimal `json:"indexPrice"` // Index price
	Time       int64         `json:"time"`       // Timestamp
}

// FundingRate represents funding rate information
type FundingRate struct {
	Symbol          string        `json:"symbol"`          // Contract symbol
	FundingRate     types.Decimal `json:"fundingRate"`     // Current funding rate
	FundingTime     int64         `json:"fundingTime"`     // Funding time
	NextFundingRate types.Decimal `json:"nextFundingRate"` // Next funding rate (predicted)
	NextFundingTime int64         `json:"nextFundingTime"` // Next funding time
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
	Symbol            string        `json:"symbol"`            // Contract symbol
	OpenInterest      types.Decimal `json:"openInterest"`      // Total open interest
	OpenInterestValue types.Decimal `json:"openInterestValue"` // Open interest value
	Time              int64         `json:"time"`              // Timestamp
}

// Request types

// GetContractsRequest is the request for GetContracts
type GetContractsRequest struct {
	Symbol string // Optional: specific contract symbol
}

// GetKlinesRequest is the request for GetKlines
type GetKlinesRequest struct {
	Symbol    string              // Required: contract symbol
	Interval  types.KlineInterval // Required: kline interval
	StartTime int64               // Optional: start time (Unix timestamp in ms)
	EndTime   int64               // Optional: end time (Unix timestamp in ms)
	Limit     int                 // Optional: number of results (default 500, max 1000)
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
