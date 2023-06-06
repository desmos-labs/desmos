package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	tokenfactorytypes "github.com/osmosis-labs/osmosis/v15/x/tokenfactory/types"
)

func NewParams(denomCreationFee sdk.Coins, denomCreationGasConsume uint64) Params {
	return Params{
		DenomCreationFee:        denomCreationFee,
		DenomCreationGasConsume: denomCreationGasConsume,
	}
}

func ToOsmosisTokenFactoryParams(p Params) tokenfactorytypes.Params {
	return tokenfactorytypes.Params{
		DenomCreationFee:        p.DenomCreationFee,
		DenomCreationGasConsume: p.DenomCreationGasConsume,
	}
}

func FromOsmosisTokenFactoryParams(p tokenfactorytypes.Params) Params {
	return Params{
		DenomCreationFee:        p.DenomCreationFee,
		DenomCreationGasConsume: p.DenomCreationGasConsume,
	}
}
