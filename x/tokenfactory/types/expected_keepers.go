package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	tokenfactorytypes "github.com/osmosis-labs/osmosis/v15/x/tokenfactory/types"
)

type TokenFactoryKeeper interface {
	CreateDenom(ctx sdk.Context, creatorAddr string, subdenom string) (newTokenDenom string, err error)
	GetAuthorityMetadata(ctx sdk.Context, denom string) (tokenfactorytypes.DenomAuthorityMetadata, error)
	BurnFrom(ctx sdk.Context, amount sdk.Coin, burnFrom string) error
	MintTo(ctx sdk.Context, amount sdk.Coin, mintTo string) error
	GetParams(ctx sdk.Context) (params tokenfactorytypes.Params)
	SetParams(ctx sdk.Context, params tokenfactorytypes.Params)
}
