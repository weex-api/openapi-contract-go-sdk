// Package account provides account management API endpoints
package account

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/weex/openapi-contract-go-sdk/weex/rest"
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
// GET /account/accounts
// Weight(IP): 10, Weight(UID): 5
//
// Reference: /contract/Account_API/AllContractAccountsInfo.md
func (s *Service) GetAccountList(ctx context.Context) ([]AccountInfo, error) {
	path := "/account/accounts"

	var accounts []AccountInfo
	err := s.client.Get(ctx, path, &accounts, 10, 5)
	return accounts, err
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

// GetSingleAsset gets single asset information
// GET /account/asset
// Weight(IP): 5, Weight(UID): 2
//
// Reference: /contract/Account_API/GetUserSingleAssetInfo.md
func (s *Service) GetSingleAsset(ctx context.Context, req *GetSingleAssetRequest) (*SingleAssetInfo, error) {
	params := url.Values{}
	params.Set("coinId", strconv.Itoa(req.CoinId))
	path := "/account/asset?" + params.Encode()

	var asset SingleAssetInfo
	err := s.client.Get(ctx, path, &asset, 5, 2)
	return &asset, err
}

// GetAllPositions gets all positions
// GET /account/positions
// Weight(IP): 20, Weight(UID): 10
//
// Reference: /contract/Account_API/GetAllContractPositions.md
func (s *Service) GetAllPositions(ctx context.Context, req *GetAllPositionsRequest) ([]Position, error) {
	params := url.Values{}

	if req != nil {
		if req.Symbol != "" {
			params.Set("symbol", req.Symbol)
		}
		if req.MarginMode > 0 {
			params.Set("marginMode", strconv.Itoa(req.MarginMode))
		}
		if req.PositionSide != "" {
			params.Set("positionSide", req.PositionSide)
		}
	}

	path := "/account/positions"
	if len(params) > 0 {
		path = path + "?" + params.Encode()
	}

	var positions []Position
	err := s.client.Get(ctx, path, &positions, 20, 10)
	return positions, err
}

// GetSinglePosition gets a single position
// GET /account/position
// Weight(IP): 5, Weight(UID): 2
//
// Reference: /contract/Account_API/GetSingleContractPosition.md
func (s *Service) GetSinglePosition(ctx context.Context, req *GetSinglePositionRequest) (*Position, error) {
	params := url.Values{}
	params.Set("symbol", req.Symbol)
	params.Set("marginMode", strconv.Itoa(req.MarginMode))
	params.Set("positionSide", req.PositionSide)
	path := "/account/position?" + params.Encode()

	var position Position
	err := s.client.Get(ctx, path, &position, 5, 2)
	return &position, err
}

// GetBills gets account bills/transaction history
// GET /account/bills
// Weight(IP): 20, Weight(UID): 10
//
// Reference: /contract/Account_API/GetContractBills.md
func (s *Service) GetBills(ctx context.Context, req *GetBillsRequest) (*BillsResponse, error) {
	params := url.Values{}

	if req != nil {
		if req.CoinId > 0 {
			params.Set("coinId", strconv.Itoa(req.CoinId))
		}
		if req.Symbol != "" {
			params.Set("symbol", req.Symbol)
		}
		if req.Type > 0 {
			params.Set("type", strconv.Itoa(req.Type))
		}
		if req.StartTime > 0 {
			params.Set("startTime", strconv.FormatInt(req.StartTime, 10))
		}
		if req.EndTime > 0 {
			params.Set("endTime", strconv.FormatInt(req.EndTime, 10))
		}
		if req.Limit > 0 {
			params.Set("limit", strconv.Itoa(req.Limit))
		}
		if req.Offset > 0 {
			params.Set("offset", strconv.Itoa(req.Offset))
		}
	}

	path := "/account/bills"
	if len(params) > 0 {
		path = path + "?" + params.Encode()
	}

	var response BillsResponse
	err := s.client.Get(ctx, path, &response, 20, 10)
	return &response, err
}

// GetUserConfig gets user configuration for a contract
// GET /account/config
// Weight(IP): 5, Weight(UID): 2
//
// Reference: /contract/Account_API/GetSingleContractUserConfig.md
func (s *Service) GetUserConfig(ctx context.Context, req *GetUserConfigRequest) (*UserConfig, error) {
	params := url.Values{}
	params.Set("symbol", req.Symbol)
	path := "/account/config?" + params.Encode()

	var config UserConfig
	err := s.client.Get(ctx, path, &config, 5, 2)
	return &config, err
}

// AdjustLeverage adjusts leverage for a contract
// POST /account/leverage
// Weight(IP): 10, Weight(UID): 5
//
// Reference: /contract/Account_API/AdjustLeverage.md
func (s *Service) AdjustLeverage(ctx context.Context, req *AdjustLeverageRequest) (*AdjustLeverageResponse, error) {
	path := "/account/leverage"

	var response AdjustLeverageResponse
	err := s.client.Post(ctx, path, req, &response, 10, 5)
	return &response, err
}

// AdjustMargin adjusts margin for an isolated position
// POST /account/margin
// Weight(IP): 10, Weight(UID): 5
//
// Reference: /contract/Account_API/AdjustMargin.md
func (s *Service) AdjustMargin(ctx context.Context, req *AdjustMarginRequest) (*AdjustMarginResponse, error) {
	path := "/account/margin"

	// Validate: margin mode must be ISOLATED (3)
	if req.MarginMode != 3 {
		return nil, fmt.Errorf("margin mode must be ISOLATED (3) for margin adjustment")
	}

	var response AdjustMarginResponse
	err := s.client.Post(ctx, path, req, &response, 10, 5)
	return &response, err
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
// POST /account/mode
// Weight(IP): 10, Weight(UID): 5
//
// Reference: /contract/Account_API/ModifyUserAccountMode.md
func (s *Service) ModifyAccountMode(ctx context.Context, req *ModifyAccountModeRequest) (*ModifyAccountModeResponse, error) {
	path := "/account/mode"

	var response ModifyAccountModeResponse
	err := s.client.Post(ctx, path, req, &response, 10, 5)
	return &response, err
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
