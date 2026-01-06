// Package market provides market data API endpoints
package market

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/weex-api/openapi-contract-go-sdk/weex/rest"
)

// Service provides access to market data API endpoints
type Service struct {
	client *rest.Client
}

// NewService creates a new market service
func NewService(client *rest.Client) *Service {
	return &Service{client: client}
}

// GetContracts gets contract information
// GET /market/contracts
// Weight(IP): 10, Weight(UID): 5
//
// Reference: /contract/Market_API/GetContractInfo.md
func (s *Service) GetContracts(ctx context.Context, req *GetContractsRequest) ([]ContractInfo, error) {
	path := "/market/contracts"

	// Add query parameters if symbol is specified
	if req != nil && req.Symbol != "" {
		params := url.Values{}
		params.Set("symbol", req.Symbol)
		path = path + "?" + params.Encode()
	}

	var contracts []ContractInfo
	err := s.client.Get(ctx, path, &contracts, 10, 5)
	return contracts, err
}

// GetTicker gets ticker information for a specific contract
// GET /market/ticker
// Weight(IP): 5, Weight(UID): 2
//
// Reference: /contract/Market_API/GetTickerInfo.md
func (s *Service) GetTicker(ctx context.Context, symbol string) (*Ticker, error) {
	params := url.Values{}
	params.Set("symbol", symbol)
	path := "/market/ticker?" + params.Encode()

	var ticker Ticker
	err := s.client.Get(ctx, path, &ticker, 5, 2)
	return &ticker, err
}

// GetAllTickers gets ticker information for all contracts
// GET /market/tickers
// Weight(IP): 20, Weight(UID): 10
//
// Reference: /contract/Market_API/GetAllTickerInfo.md
func (s *Service) GetAllTickers(ctx context.Context) ([]Ticker, error) {
	path := "/market/tickers"

	var tickers []Ticker
	err := s.client.Get(ctx, path, &tickers, 20, 10)
	return tickers, err
}

// GetDepth gets order book depth data
// GET /market/depth
// Weight(IP): 10, Weight(UID): 5
//
// Reference: /contract/Market_API/GetDepthData.md
func (s *Service) GetDepth(ctx context.Context, req *GetDepthRequest) (*Depth, error) {
	params := url.Values{}
	params.Set("symbol", req.Symbol)

	if req.Limit > 0 {
		params.Set("limit", strconv.Itoa(req.Limit))
	}

	path := "/market/depth?" + params.Encode()

	var depth Depth
	err := s.client.Get(ctx, path, &depth, 10, 5)
	return &depth, err
}

// GetKlines gets candlestick/kline data
// GET /market/klines
// Weight(IP): 10, Weight(UID): 5
//
// Reference: /contract/Market_API/GetKLineData.md
func (s *Service) GetKlines(ctx context.Context, req *GetKlinesRequest) ([]Kline, error) {
	params := url.Values{}
	params.Set("symbol", req.Symbol)
	params.Set("interval", string(req.Interval))

	if req.StartTime > 0 {
		params.Set("startTime", strconv.FormatInt(req.StartTime, 10))
	}
	if req.EndTime > 0 {
		params.Set("endTime", strconv.FormatInt(req.EndTime, 10))
	}
	if req.Limit > 0 {
		params.Set("limit", strconv.Itoa(req.Limit))
	}

	path := "/market/klines?" + params.Encode()

	var klines []Kline
	err := s.client.Get(ctx, path, &klines, 10, 5)
	return klines, err
}

// GetHistoryKlines gets historical candlestick/kline data
// GET /market/history/klines
// Weight(IP): 20, Weight(UID): 10
//
// Reference: /contract/Market_API/GetHistoryKLineData.md
func (s *Service) GetHistoryKlines(ctx context.Context, req *GetHistoryKlinesRequest) ([]Kline, error) {
	params := url.Values{}
	params.Set("symbol", req.Symbol)
	params.Set("interval", string(req.Interval))
	params.Set("startTime", strconv.FormatInt(req.StartTime, 10))
	params.Set("endTime", strconv.FormatInt(req.EndTime, 10))

	if req.Limit > 0 {
		params.Set("limit", strconv.Itoa(req.Limit))
	}

	path := "/market/history/klines?" + params.Encode()

	var klines []Kline
	err := s.client.Get(ctx, path, &klines, 20, 10)
	return klines, err
}

// GetTrades gets recent trades
// GET /market/trades
// Weight(IP): 10, Weight(UID): 5
//
// Reference: /contract/Market_API/GetTradeData.md
func (s *Service) GetTrades(ctx context.Context, req *GetTradesRequest) ([]Trade, error) {
	params := url.Values{}
	params.Set("symbol", req.Symbol)

	if req.Limit > 0 {
		params.Set("limit", strconv.Itoa(req.Limit))
	}

	path := "/market/trades?" + params.Encode()

	var trades []Trade
	err := s.client.Get(ctx, path, &trades, 10, 5)
	return trades, err
}

// GetServerTime gets the server time
// GET /market/time
// Weight(IP): 1, Weight(UID): 1
//
// Reference: /contract/Market_API/GetServerTime.md
func (s *Service) GetServerTime(ctx context.Context) (*ServerTime, error) {
	path := "/market/time"

	var serverTime ServerTime
	err := s.client.Get(ctx, path, &serverTime, 1, 1)
	return &serverTime, err
}

// GetIndexPrice gets the index price
// GET /market/index
// Weight(IP): 5, Weight(UID): 2
//
// Reference: /contract/Market_API/GetCurrencyIndex.md
func (s *Service) GetIndexPrice(ctx context.Context, symbol string) (*IndexPrice, error) {
	params := url.Values{}
	params.Set("symbol", symbol)
	path := "/market/index?" + params.Encode()

	var indexPrice IndexPrice
	err := s.client.Get(ctx, path, &indexPrice, 5, 2)
	return &indexPrice, err
}

// GetFundingRate gets the current funding rate
// GET /market/fundingRate
// Weight(IP): 5, Weight(UID): 2
//
// Reference: /contract/Market_API/GetCurrentFundRate.md
func (s *Service) GetFundingRate(ctx context.Context, symbol string) (*FundingRate, error) {
	params := url.Values{}
	params.Set("symbol", symbol)
	path := "/market/fundingRate?" + params.Encode()

	var fundingRate FundingRate
	err := s.client.Get(ctx, path, &fundingRate, 5, 2)
	return &fundingRate, err
}

// GetFundingHistory gets historical funding rates
// GET /market/fundingRate/history
// Weight(IP): 10, Weight(UID): 5
//
// Reference: /contract/Market_API/GetContractFundingHistory.md
func (s *Service) GetFundingHistory(ctx context.Context, req *GetFundingHistoryRequest) ([]FundingRateHistory, error) {
	params := url.Values{}
	params.Set("symbol", req.Symbol)

	if req.StartTime > 0 {
		params.Set("startTime", strconv.FormatInt(req.StartTime, 10))
	}
	if req.EndTime > 0 {
		params.Set("endTime", strconv.FormatInt(req.EndTime, 10))
	}
	if req.Limit > 0 {
		params.Set("limit", strconv.Itoa(req.Limit))
	}

	path := "/market/fundingRate/history?" + params.Encode()

	var history []FundingRateHistory
	err := s.client.Get(ctx, path, &history, 10, 5)
	return history, err
}

// GetSettlementTime gets the next settlement time
// GET /market/settlementTime
// Weight(IP): 5, Weight(UID): 2
//
// Reference: /contract/Market_API/GetNextContractSettlementTime.md
func (s *Service) GetSettlementTime(ctx context.Context, symbol string) (*SettlementTime, error) {
	params := url.Values{}
	params.Set("symbol", symbol)
	path := "/market/settlementTime?" + params.Encode()

	var settlementTime SettlementTime
	err := s.client.Get(ctx, path, &settlementTime, 5, 2)
	return &settlementTime, err
}

// GetOpenInterest gets the platform open interest
// GET /market/openInterest
// Weight(IP): 5, Weight(UID): 2
//
// Reference: /contract/Market_API/GetTotalPlatformOpenInterest.md
func (s *Service) GetOpenInterest(ctx context.Context, symbol string) (*OpenInterest, error) {
	params := url.Values{}
	params.Set("symbol", symbol)
	path := "/market/openInterest?" + params.Encode()

	var openInterest OpenInterest
	err := s.client.Get(ctx, path, &openInterest, 5, 2)
	return &openInterest, err
}

// Helper function to build query string
func buildQueryString(params url.Values) string {
	if len(params) == 0 {
		return ""
	}
	return "?" + params.Encode()
}

// Validation helpers

// ValidateSymbol checks if a symbol is valid
func ValidateSymbol(symbol string) error {
	if symbol == "" {
		return fmt.Errorf("symbol cannot be empty")
	}
	return nil
}

// ValidateInterval checks if an interval is valid
func ValidateInterval(interval string) error {
	if interval == "" {
		return fmt.Errorf("interval cannot be empty")
	}
	return nil
}
