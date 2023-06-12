package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	tokenfactorytypes "github.com/osmosis-labs/osmosis/v15/x/tokenfactory/types"
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

// ToOsmosisTokenFactoryParams converts desmos tokenfactory Params into osmosis tokenfactory Params
func ToOsmosisTokenFactoryParams(p Params) tokenfactorytypes.Params {
	return tokenfactorytypes.Params{
		DenomCreationFee: p.DenomCreationFee,
	}
}

// FromOsmosisTokenFactoryParams converts osmosis tokenfactory Params into desmos tokenfactory Params
func FromOsmosisTokenFactoryParams(p tokenfactorytypes.Params) Params {
	return Params{
		DenomCreationFee: p.DenomCreationFee,
	}
}
