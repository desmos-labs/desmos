package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewParams creates a new params instance
func NewParams(denomCreationFee sdk.Coins) Params {
	return Params{
		DenomCreationFee: denomCreationFee,
	}
}

// DefaultParams creates a default params instance
func DefaultParams() Params {
	return Params{
		DenomCreationFee: sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 10_000_000)), // 10 DSM
	}
}

// Validate implements fmt.Validator
func (p Params) Validate() error {
	if err := p.DenomCreationFee.Validate(); err != nil {
		return err
	}

	return nil
}
