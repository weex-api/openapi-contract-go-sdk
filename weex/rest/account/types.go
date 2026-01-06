package account

import (
	"github.com/weex-api/openapi-contract-go-sdk/weex/types"
)

// AccountResponse represents the common response for GetAccountList and GetSingleAsset
// Both APIs return the same structure: account + collateral + version
type AccountResponse struct {
	Account    Account      `json:"account"`    // Account information
	Collateral []Collateral `json:"collateral"` // Collateral information
	Version    string       `json:"version"`    // Version number
}

// Account represents account configuration information
type Account struct {
	DefaultFeeSetting             FeeSetting        `json:"defaultFeeSetting"`             // Default fee configuration
	FeeSetting                    []FeeSetting      `json:"feeSetting"`                    // Fee settings per symbol
	ModeSetting                   []ModeSetting     `json:"modeSetting"`                   // Mode settings per symbol
	LeverageSetting               []LeverageSetting `json:"leverageSetting"`               // Leverage settings per symbol
	CreateOrderRateLimitPerMinute int               `json:"createOrderRateLimitPerMinute"` // Order creation rate limit per minute
	CreateOrderDelayMilliseconds  int               `json:"createOrderDelayMilliseconds"`  // Order creation delay (milliseconds)
	CreatedTime                   int64             `json:"createdTime"`                   // Creation time (Unix millisecond timestamp)
	UpdatedTime                   int64             `json:"updatedTime"`                   // Update time (Unix millisecond timestamp)
}

// FeeSetting represents fee configuration
type FeeSetting struct {
	Symbol                     string `json:"symbol,omitempty"`                // Symbol name (empty for default)
	IsSetFeeRate               bool   `json:"is_set_fee_rate"`                 // Whether fee rates are set
	TakerFeeRate               string `json:"taker_fee_rate"`                  // Taker fee rate
	MakerFeeRate               string `json:"maker_fee_rate"`                  // Maker fee rate
	IsSetFeeDiscount           bool   `json:"is_set_fee_discount"`             // Whether fee discounts are enabled
	FeeDiscount                string `json:"fee_discount"`                    // Fee rate discount
	IsSetTakerMakerFeeDiscount bool   `json:"is_set_taker_maker_fee_discount"` // Whether to apply separate fee discounts
	TakerFeeDiscount           string `json:"taker_fee_discount"`              // Taker fee rate discount
	MakerFeeDiscount           string `json:"maker_fee_discount"`              // Maker fee rate discount
}

// ModeSetting represents mode configuration
type ModeSetting struct {
	Symbol        string `json:"symbol"`        // Symbol name
	MarginMode    string `json:"marginMode"`    // Margin mode
	SeparatedMode string `json:"separatedMode"` // Position segregation mode
	PositionMode  string `json:"positionMode"`  // Position mode
}

// LeverageSetting represents leverage configuration
type LeverageSetting struct {
	Symbol                string `json:"symbol"`                  // Symbol name
	IsolatedLongLeverage  string `json:"isolated_long_leverage"`  // Isolated long position leverage
	IsolatedShortLeverage string `json:"isolated_short_leverage"` // Isolated short position leverage
	CrossLeverage         string `json:"cross_leverage"`          // Cross margin leverage
}

// SingleAccountResponse is an alias for AccountResponse (for backward compatibility)
// GetAccount returns the same structure as GetAccounts
type SingleAccountResponse = AccountResponse

// Collateral represents collateral information
type Collateral struct {
	Coin                             string `json:"coin"`                                  // Currency
	MarginMode                       string `json:"marginMode"`                            // Margin mode
	CrossSymbol                      string `json:"crossSymbol"`                           // When marginMode=CROSS, represents the symbol associated with cross margin mode
	IsolatedPositionId               int64  `json:"isolatedPositionId"`                    // When marginMode=ISOLATED, represents the position ID associated with isolated margin
	Amount                           string `json:"amount"`                                // Collateral amount
	PendingDepositAmount             string `json:"pending_deposit_amount"`                // Pending deposit amount
	PendingWithdrawAmount            string `json:"pending_withdraw_amount"`               // Pending withdrawal amount
	PendingTransferInAmount          string `json:"pending_transfer_in_amount"`            // Pending inbound transfer amount
	PendingTransferOutAmount         string `json:"pending_transfer_out_amount"`           // Pending outbound transfer amount
	IsLiquidating                    bool   `json:"is_liquidating"`                        // Whether liquidation is triggered (in progress)
	LegacyAmount                     string `json:"legacy_amount"`                         // Legacy balance (display only)
	CumDepositAmount                 string `json:"cum_deposit_amount"`                    // Accumulated deposit amount
	CumWithdrawAmount                string `json:"cum_withdraw_amount"`                   // Accumulated withdrawal amount
	CumTransferInAmount              string `json:"cum_transfer_in_amount"`                // Accumulated inbound transfer amount
	CumTransferOutAmount             string `json:"cum_transfer_out_amount"`               // Accumulated outbound transfer amount
	CumMarginMoveInAmount            string `json:"cum_margin_move_in_amount"`             // Accumulated margin deposit amount
	CumMarginMoveOutAmount           string `json:"cum_margin_move_out_amount"`            // Accumulated margin withdrawal amount
	CumPositionOpenLongAmount        string `json:"cum_position_open_long_amount"`         // Accumulated collateral amount for long position openings
	CumPositionOpenShortAmount       string `json:"cum_position_open_short_amount"`        // Accumulated collateral amount for short position openings
	CumPositionCloseLongAmount       string `json:"cum_position_close_long_amount"`        // Accumulated collateral amount for long position closings
	CumPositionCloseShortAmount      string `json:"cum_position_close_short_amount"`       // Accumulated collateral amount for short position closings
	CumPositionFillFeeAmount         string `json:"cum_position_fill_fee_amount"`          // Accumulated trading fees for filled orders
	CumPositionLiquidateFeeAmount    string `json:"cum_position_liquidate_fee_amount"`     // Accumulated liquidation fees
	CumPositionFundingAmount         string `json:"cum_position_funding_amount"`           // Accumulated funding fees
	CumOrderFillFeeIncomeAmount      string `json:"cum_order_fill_fee_income_amount"`      // Accumulated order fee income
	CumOrderLiquidateFeeIncomeAmount string `json:"cum_order_liquidate_fee_income_amount"` // Accumulated liquidation fee income
	CreatedTime                      int64  `json:"created_time"`                          // Creation time (Unix millisecond timestamp)
	UpdatedTime                      int64  `json:"updated_time"`                          // Update time (Unix millisecond timestamp)
}

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
	CoinName     string `json:"coinName"`     // Currency name
	Available    string `json:"available"`    // Available balance
	Frozen       string `json:"frozen"`       // Frozen balance
	Equity       string `json:"equity"`       // Account equity
	UnrealizePnl string `json:"unrealizePnl"` // Unrealized PnL
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
	ID                         int64  `json:"id"`                             // Position ID
	AccountID                  int64  `json:"account_id"`                     // Account ID
	CoinID                     int    `json:"coin_id"`                        // Coin ID
	ContractID                 int64  `json:"contract_id"`                    // Contract ID
	Symbol                     string `json:"symbol"`                         // Symbol
	Side                       string `json:"side"`                           // Position side (LONG/SHORT)
	MarginMode                 string `json:"margin_mode"`                    // Margin mode
	SeparatedMode              string `json:"separated_mode"`                 // Separated mode
	SeparatedOpenOrderID       int64  `json:"separated_open_order_id"`        // Separated open order ID
	Leverage                   string `json:"leverage"`                       // Leverage
	Size                       string `json:"size"`                           // Position size
	OpenValue                  string `json:"open_value"`                     // Open value
	OpenFee                    string `json:"open_fee"`                       // Open fee
	FundingFee                 string `json:"funding_fee"`                    // Funding fee
	MarginSize                 string `json:"marginSize"`                     // Margin size
	IsolatedMargin             string `json:"isolated_margin"`                // Isolated margin
	IsAutoAppendIsolatedMargin bool   `json:"is_auto_append_isolated_margin"` // Auto append isolated margin
	CumOpenSize                string `json:"cum_open_size"`                  // Cumulative open size
	CumOpenValue               string `json:"cum_open_value"`                 // Cumulative open value
	CumOpenFee                 string `json:"cum_open_fee"`                   // Cumulative open fee
	CumCloseSize               string `json:"cum_close_size"`                 // Cumulative close size
	CumCloseValue              string `json:"cum_close_value"`                // Cumulative close value
	CumCloseFee                string `json:"cum_close_fee"`                  // Cumulative close fee
	CumFundingFee              string `json:"cum_funding_fee"`                // Cumulative funding fee
	CumLiquidateFee            string `json:"cum_liquidate_fee"`              // Cumulative liquidate fee
	CreatedMatchSequenceID     int64  `json:"created_match_sequence_id"`      // Created match sequence ID
	UpdatedMatchSequenceID     int64  `json:"updated_match_sequence_id"`      // Updated match sequence ID
	CreatedTime                int64  `json:"created_time"`                   // Created time
	UpdatedTime                int64  `json:"updated_time"`                   // Updated time
	ContractVal                string `json:"contractVal"`                    // Contract value
	UnrealizePnl               string `json:"unrealizePnl"`                   // Unrealized PnL
	LiquidatePrice             string `json:"liquidatePrice"`                 // Liquidate price
}

// Bill represents an account bill/transaction
type Bill struct {
	BillId         int64  `json:"billId"`         // Bill ID
	Coin           string `json:"coin"`           // Currency name
	Symbol         string `json:"symbol"`         // Contract symbol
	Amount         string `json:"amount"`         // Amount
	BusinessType   string `json:"businessType"`   // Business type
	Balance        string `json:"balance"`        // Balance after transaction
	FillFee        string `json:"fillFee"`        // Transaction fee
	TransferReason string `json:"transferReason"` // Transfer reason
	CTime          int64  `json:"cTime"`          // Creation time (Unix millisecond timestamp)
}

// BillsResponse represents the paginated bills response
type BillsResponse struct {
	HasNextPage bool   `json:"hasNextPage"` // Has next page
	Items       []Bill `json:"items"`       // Bills list
}

// UserConfigData represents user configuration data for a single contract
type UserConfigData struct {
	IsolatedLongLeverage  string `json:"isolated_long_leverage"`  // Isolated long leverage
	IsolatedShortLeverage string `json:"isolated_short_leverage"` // Isolated short leverage
	CrossLeverage         string `json:"cross_leverage"`          // Cross leverage
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
	Coin         string `json:"coin,omitempty"`         // Optional: currency name
	Symbol       string `json:"symbol,omitempty"`       // Optional: contract symbol
	BusinessType string `json:"businessType,omitempty"` // Optional: business type
	StartTime    int64  `json:"startTime,omitempty"`    // Optional: start time (Unix timestamp in ms)
	EndTime      int64  `json:"endTime,omitempty"`      // Optional: end time (Unix timestamp in ms)
	Limit        int    `json:"limit,omitempty"`        // Optional: page size (default 20, max 100)
}

// GetUserConfigRequest is the request for GetUserConfig
type GetUserConfigRequest struct {
	Symbol string // Optional: contract symbol (if not specified, returns all)
}

// AdjustLeverageRequest is the request for AdjustLeverage
type AdjustLeverageRequest struct {
	Symbol        string `json:"symbol"`        // Required: contract symbol
	MarginMode    int    `json:"marginMode"`    // Required: margin mode (1=Cross, 3=Isolated)
	LongLeverage  string `json:"longLeverage"`  // Required: long position leverage
	ShortLeverage string `json:"shortLeverage"` // Required: short position leverage (must equal longLeverage in Cross mode)
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
	IsolatedPositionId int64  `json:"isolatedPositionId"` // Required: isolated position ID
	CollateralAmount   string `json:"collateralAmount"`   // Required: collateral amount (positive=increase, negative=decrease)
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
	Symbol        string `json:"symbol"`                  // Required: trading pair
	MarginMode    int    `json:"marginMode"`              // Required: margin mode (1=Cross, 3=Isolated)
	SeparatedMode int    `json:"separatedMode,omitempty"` // Optional: position segregation mode (1=Combined)
}

// ModifyAccountModeResponse is the response for ModifyAccountMode
type ModifyAccountModeResponse struct {
	MarginMode   int `json:"marginMode"`   // Margin mode
	PositionMode int `json:"positionMode"` // Position mode
}
