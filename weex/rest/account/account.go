// Package account provides account management API endpoints
package account

import (
	"context"
	"fmt"
	"net/url"

	"github.com/weex-api/openapi-contract-go-sdk/weex/rest"
)

// Service provides access to account management API endpoints
type Service struct {
	client *rest.Client
}

// NewService creates a new account service
func NewService(client *rest.Client) *Service {
	return &Service{client: client}
}

// GetAccountList gets the list of all contract accounts
// GET /account/getAccounts
// Weight(IP): 5, Weight(UID): 5
//
// Reference: /contract/Account_API/AllContractAccountsInfo.md
func (s *Service) GetAccountList(ctx context.Context) (*AccountResponse, error) {
	path := "/account/getAccounts"

	var response AccountResponse
	err := s.client.Get(ctx, path, &response, 5, 5)
	return &response, err
}

// GetAccountBalance gets account asset information
// GET /account/assets
// Weight(IP): 10, Weight(UID): 5
//
// Reference: /contract/Account_API/GetAccountBalance.md
func (s *Service) GetAccountBalance(ctx context.Context) ([]AssetBalance, error) {
	path := "/account/assets"

	var assets []AssetBalance
	err := s.client.Get(ctx, path, &assets, 10, 5)
	return assets, err
}

// GetSingleAsset gets single asset information (single account info)
// GET /account/getAccount
// Weight(IP): 1, Weight(UID): 1
//
// Reference: /contract/Account_API/GetUserSingleAssetInfo.md
// Note: Returns the same structure as GetAccountList
func (s *Service) GetSingleAsset(ctx context.Context, coin string) (*AccountResponse, error) {
	params := url.Values{}
	params.Set("coin", coin)
	path := "/account/getAccount?" + params.Encode()

	var response AccountResponse
	err := s.client.Get(ctx, path, &response, 1, 1)
	return &response, err
}

// GetAllPositions gets all positions
// GET /account/position/allPosition
// Weight(IP): 10, Weight(UID): 15
//
// Reference: /contract/Account_API/GetAllContractPositions.md
func (s *Service) GetAllPositions(ctx context.Context, req *GetAllPositionsRequest) ([]Position, error) {
	path := "/account/position/allPosition"

	var positions []Position
	err := s.client.Get(ctx, path, &positions, 10, 15)
	return positions, err
}

// GetSinglePosition gets a single position
// GET /account/position/singlePosition
// Weight(IP): 2, Weight(UID): 3
//
// Reference: /contract/Account_API/GetSingleContractPosition.md
// Note: API may return empty array [] when no position exists
func (s *Service) GetSinglePosition(ctx context.Context, symbol string) (*Position, error) {
	params := url.Values{}
	params.Set("symbol", symbol)
	path := "/account/position/singlePosition?" + params.Encode()

	// Try to unmarshal as Position first
	var position Position
	err := s.client.Get(ctx, path, &position, 2, 3)
	if err != nil {
		// If it fails, might be an empty array, return empty position
		return &Position{}, nil
	}
	return &position, nil
}

// GetBills gets account bills/transaction history
// POST /account/bills
// Weight(IP): 2, Weight(UID): 5
//
// Reference: /contract/Account_API/GetContractBills.md
func (s *Service) GetBills(ctx context.Context, req *GetBillsRequest) (*BillsResponse, error) {
	path := "/account/bills"

	var response BillsResponse
	err := s.client.Post(ctx, path, req, &response, 2, 5)
	return &response, err
}

// GetUserConfig gets user configuration for a contract
// GET /account/settings
// Weight(IP): 1, Weight(UID): 1
//
// Reference: /contract/Account_API/GetSingleContractUserConfig.md
// Returns a map of symbol to UserConfig
func (s *Service) GetUserConfig(ctx context.Context, req *GetUserConfigRequest) (map[string]*UserConfigData, error) {
	params := url.Values{}
	if req != nil && req.Symbol != "" {
		params.Set("symbol", req.Symbol)
	}

	path := "/account/settings"
	if len(params) > 0 {
		path = path + "?" + params.Encode()
	}

	var config map[string]*UserConfigData
	err := s.client.Get(ctx, path, &config, 1, 1)
	return config, err
}

// AdjustLeverage adjusts leverage for a contract
// POST /account/leverage
// Weight(IP): 10, Weight(UID): 20
//
// Reference: /contract/Account_API/AdjustLeverage.md
func (s *Service) AdjustLeverage(ctx context.Context, req *AdjustLeverageRequest) error {
	path := "/account/leverage"

	// API returns standard response (code, msg, requestTime), not data
	var response rest.APIResponse
	err := s.client.PostRaw(ctx, path, req, &response, 10, 20)
	if err != nil {
		return err
	}
	return nil
}

// AdjustMargin adjusts margin for an isolated position
// POST /account/adjustMargin
// Weight(IP): 15, Weight(UID): 30
//
// Reference: /contract/Account_API/AdjustMargin.md
func (s *Service) AdjustMargin(ctx context.Context, req *AdjustMarginRequest) error {
	path := "/account/adjustMargin"

	// API returns standard response (code, msg, requestTime), not data
	var response rest.APIResponse
	err := s.client.PostRaw(ctx, path, req, &response, 15, 30)
	if err != nil {
		return err
	}
	return nil
}

// AutoAddMargin enables/disables auto add margin for an isolated position
// POST /account/autoAddMargin
// Weight(IP): 10, Weight(UID): 5
//
// Reference: /contract/Account_API/AutoAddMargin.md
func (s *Service) AutoAddMargin(ctx context.Context, req *AutoAddMarginRequest) (*AutoAddMarginResponse, error) {
	path := "/account/autoAddMargin"

	// Validate: margin mode must be ISOLATED (3)
	if req.MarginMode != 3 {
		return nil, fmt.Errorf("margin mode must be ISOLATED (3) for auto add margin")
	}

	var response AutoAddMarginResponse
	err := s.client.Post(ctx, path, req, &response, 10, 5)
	return &response, err
}

// ModifyAccountMode modifies account mode (margin mode and position mode)
// POST /account/position/changeHoldModel
// Weight(IP): 20, Weight(UID): 50
//
// Reference: /contract/Account_API/ModifyUserAccountMode.md
func (s *Service) ModifyAccountMode(ctx context.Context, req *ModifyAccountModeRequest) error {
	path := "/account/position/changeHoldModel"

	// API returns standard response (code, msg, requestTime), not data
	var response rest.APIResponse
	err := s.client.PostRaw(ctx, path, req, &response, 20, 50)
	if err != nil {
		return err
	}
	return nil
}

// Validation helpers

// ValidateCoinId checks if a coin ID is valid
func ValidateCoinId(coinId int) error {
	if coinId <= 0 {
		return fmt.Errorf("coinId must be greater than 0")
	}
	return nil
}

// ValidatePositionSide checks if a position side is valid
func ValidatePositionSide(side string) error {
	if side != "LONG" && side != "SHORT" {
		return fmt.Errorf("positionSide must be LONG or SHORT")
	}
	return nil
}

// ValidateMarginMode checks if a margin mode is valid
func ValidateMarginMode(mode int) error {
	if mode != 1 && mode != 3 {
		return fmt.Errorf("marginMode must be 1 (SHARED) or 3 (ISOLATED)")
	}
	return nil
}
