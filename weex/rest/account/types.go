package account

import (
	"github.com/weex/openapi-contract-go-sdk/weex/types"
)

// AccountInfo represents account information
type AccountInfo struct {
	CoinId       int           `json:"coinId"`       // Currency ID
	CoinName     string        `json:"coinName"`     // Currency name
	Available    types.Decimal `json:"available"`    // Available balance
	Frozen       types.Decimal `json:"frozen"`       // Frozen balance
	Equity       types.Decimal `json:"equity"`       // Account equity
	Unrealized   types.Decimal `json:"unrealized"`   // Unrealized PnL
	MarginMode   int           `json:"marginMode"`   // Margin mode
	PositionMode int           `json:"positionMode"` // Position mode
}

// AssetBalance represents account asset balance
type AssetBalance struct {
	CoinId            int           `json:"coinId"`            // Currency ID
	CoinName          string        `json:"coinName"`          // Currency name
	Available         types.Decimal `json:"available"`         // Available balance
	Frozen            types.Decimal `json:"frozen"`            // Frozen balance
	Equity            types.Decimal `json:"equity"`            // Account equity
	UnrealizedPnl     types.Decimal `json:"unrealizedPnl"`     // Unrealized PnL
	MarginBalance     types.Decimal `json:"marginBalance"`     // Margin balance
	MarginRate        types.Decimal `json:"marginRate"`        // Margin rate
	MaintenanceMargin types.Decimal `json:"maintenanceMargin"` // Maintenance margin
	InitialMargin     types.Decimal `json:"initialMargin"`     // Initial margin
	RiskLevel         int           `json:"riskLevel"`         // Risk level (1-5)
}

// SingleAssetInfo represents single asset information
type SingleAssetInfo struct {
	CoinId            int           `json:"coinId"`            // Currency ID
	CoinName          string        `json:"coinName"`          // Currency name
	Available         types.Decimal `json:"available"`         // Available balance
	Frozen            types.Decimal `json:"frozen"`            // Frozen balance
	Equity            types.Decimal `json:"equity"`            // Account equity
	UnrealizedPnl     types.Decimal `json:"unrealizedPnl"`     // Unrealized PnL
	MarginBalance     types.Decimal `json:"marginBalance"`     // Margin balance
	MarginRate        types.Decimal `json:"marginRate"`        // Margin rate
	MaintenanceMargin types.Decimal `json:"maintenanceMargin"` // Maintenance margin
	InitialMargin     types.Decimal `json:"initialMargin"`     // Initial margin
}

// Position represents a contract position
type Position struct {
	PositionId        int64         `json:"positionId"`        // Position ID
	Symbol            string        `json:"symbol"`            // Contract symbol
	MarginMode        string        `json:"marginMode"`        // Margin mode (SHARED/ISOLATED)
	PositionSide      string        `json:"positionSide"`      // Position side (LONG/SHORT)
	Size              types.Decimal `json:"size"`              // Position size
	Available         types.Decimal `json:"available"`         // Available size
	Frozen            types.Decimal `json:"frozen"`            // Frozen size
	AverageOpenPrice  types.Decimal `json:"averageOpenPrice"`  // Average open price
	UnrealizedPnl     types.Decimal `json:"unrealizedPnl"`     // Unrealized PnL
	RealizedPnl       types.Decimal `json:"realizedPnl"`       // Realized PnL
	Leverage          types.Decimal `json:"leverage"`          // Leverage
	MarginBalance     types.Decimal `json:"marginBalance"`     // Margin balance
	MaintenanceMargin types.Decimal `json:"maintenanceMargin"` // Maintenance margin
	LiquidatePrice    types.Decimal `json:"liquidatePrice"`    // Liquidation price
	MarkPrice         types.Decimal `json:"markPrice"`         // Mark price
	CreateTime        int64         `json:"createTime"`        // Create time
	UpdateTime        int64         `json:"updateTime"`        // Update time
	AutoAddMargin     bool          `json:"autoAddMargin"`     // Auto add margin enabled
}

// Bill represents an account bill/transaction
type Bill struct {
	Id           int64         `json:"id"`           // Bill ID
	CoinId       int           `json:"coinId"`       // Currency ID
	CoinName     string        `json:"coinName"`     // Currency name
	Type         int           `json:"type"`         // Bill type
	Amount       types.Decimal `json:"amount"`       // Amount
	Balance      types.Decimal `json:"balance"`      // Balance after transaction
	Fee          types.Decimal `json:"fee"`          // Fee
	Symbol       string        `json:"symbol"`       // Contract symbol
	PositionSide string        `json:"positionSide"` // Position side
	CreateTime   int64         `json:"createTime"`   // Create time
	Remark       string        `json:"remark"`       // Remark
}

// BillsResponse represents the paginated bills response
type BillsResponse struct {
	Bills    []Bill `json:"bills"`    // Bills list
	NextFlag bool   `json:"nextFlag"` // Has next page
	Totals   int    `json:"totals"`   // Total count
}

// UserConfig represents user configuration
type UserConfig struct {
	Symbol            string        `json:"symbol"`            // Contract symbol
	MarginMode        int           `json:"marginMode"`        // Margin mode
	PositionMode      int           `json:"positionMode"`      // Position mode
	SplitPositionMode int           `json:"splitPositionMode"` // Split position mode
	Leverage          types.Decimal `json:"leverage"`          // Leverage
	LongLeverage      types.Decimal `json:"longLeverage"`      // Long leverage (separated mode)
	ShortLeverage     types.Decimal `json:"shortLeverage"`     // Short leverage (separated mode)
}

// Request types

// GetSingleAssetRequest is the request for GetSingleAsset
type GetSingleAssetRequest struct {
	CoinId int // Required: currency ID
}

// GetAllPositionsRequest is the request for GetAllPositions
type GetAllPositionsRequest struct {
	Symbol       string // Optional: contract symbol
	MarginMode   int    // Optional: margin mode
	PositionSide string // Optional: position side (LONG/SHORT)
}

// GetSinglePositionRequest is the request for GetSinglePosition
type GetSinglePositionRequest struct {
	Symbol       string // Required: contract symbol
	MarginMode   int    // Required: margin mode
	PositionSide string // Required: position side (LONG/SHORT)
}

// GetBillsRequest is the request for GetBills
type GetBillsRequest struct {
	CoinId    int    // Optional: currency ID
	Symbol    string // Optional: contract symbol
	Type      int    // Optional: bill type
	StartTime int64  // Optional: start time (Unix timestamp in ms)
	EndTime   int64  // Optional: end time (Unix timestamp in ms)
	Limit     int    // Optional: page size (default 100, max 500)
	Offset    int    // Optional: offset for pagination
}

// GetUserConfigRequest is the request for GetUserConfig
type GetUserConfigRequest struct {
	Symbol string // Required: contract symbol
}

// AdjustLeverageRequest is the request for AdjustLeverage
type AdjustLeverageRequest struct {
	Symbol            string        `json:"symbol"`                      // Required: contract symbol
	MarginMode        int           `json:"marginMode"`                  // Required: margin mode
	Leverage          types.Decimal `json:"leverage,omitempty"`          // Leverage (combined mode)
	LongLeverage      types.Decimal `json:"longLeverage,omitempty"`      // Long leverage (separated mode)
	ShortLeverage     types.Decimal `json:"shortLeverage,omitempty"`     // Short leverage (separated mode)
	SplitPositionMode int           `json:"splitPositionMode,omitempty"` // Split position mode
}

// AdjustLeverageResponse is the response for AdjustLeverage
type AdjustLeverageResponse struct {
	Symbol            string        `json:"symbol"`            // Contract symbol
	MarginMode        int           `json:"marginMode"`        // Margin mode
	Leverage          types.Decimal `json:"leverage"`          // Leverage
	LongLeverage      types.Decimal `json:"longLeverage"`      // Long leverage
	ShortLeverage     types.Decimal `json:"shortLeverage"`     // Short leverage
	SplitPositionMode int           `json:"splitPositionMode"` // Split position mode
}

// AdjustMarginRequest is the request for AdjustMargin
type AdjustMarginRequest struct {
	Symbol       string        `json:"symbol"`       // Required: contract symbol
	MarginMode   int           `json:"marginMode"`   // Required: margin mode (must be ISOLATED)
	PositionSide string        `json:"positionSide"` // Required: position side
	Amount       types.Decimal `json:"amount"`       // Required: adjustment amount (positive to add, negative to reduce)
}

// AdjustMarginResponse is the response for AdjustMargin
type AdjustMarginResponse struct {
	Symbol        string        `json:"symbol"`        // Contract symbol
	MarginMode    string        `json:"marginMode"`    // Margin mode
	PositionSide  string        `json:"positionSide"`  // Position side
	Amount        types.Decimal `json:"amount"`        // Adjusted amount
	MarginBalance types.Decimal `json:"marginBalance"` // New margin balance
}

// AutoAddMarginRequest is the request for AutoAddMargin
type AutoAddMarginRequest struct {
	Symbol        string `json:"symbol"`        // Required: contract symbol
	MarginMode    int    `json:"marginMode"`    // Required: margin mode (must be ISOLATED)
	PositionSide  string `json:"positionSide"`  // Required: position side
	AutoAddMargin bool   `json:"autoAddMargin"` // Required: enable/disable auto add margin
}

// AutoAddMarginResponse is the response for AutoAddMargin
type AutoAddMarginResponse struct {
	Symbol        string `json:"symbol"`        // Contract symbol
	MarginMode    string `json:"marginMode"`    // Margin mode
	PositionSide  string `json:"positionSide"`  // Position side
	AutoAddMargin bool   `json:"autoAddMargin"` // Auto add margin status
}

// ModifyAccountModeRequest is the request for ModifyAccountMode
type ModifyAccountModeRequest struct {
	MarginMode   int `json:"marginMode"`   // Required: margin mode
	PositionMode int `json:"positionMode"` // Required: position mode
}

// ModifyAccountModeResponse is the response for ModifyAccountMode
type ModifyAccountModeResponse struct {
	MarginMode   int `json:"marginMode"`   // Margin mode
	PositionMode int `json:"positionMode"` // Position mode
}
