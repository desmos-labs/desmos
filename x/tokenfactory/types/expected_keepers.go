package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	subspacestypes "github.com/desmos-labs/desmos/v5/x/subspaces/types"
	tokenfactorytypes "github.com/osmosis-labs/osmosis/v15/x/tokenfactory/types"
)

type BankKeeper interface {
	GetDenomMetaData(ctx sdk.Context, denom string) (banktypes.Metadata, bool)
	SetDenomMetaData(ctx sdk.Context, denomMetaData banktypes.Metadata)
}

type TokenFactoryKeeper interface {
	CreateDenom(ctx sdk.Context, creatorAddr string, subdenom string) (newTokenDenom string, err error)
	GetAuthorityMetadata(ctx sdk.Context, denom string) (tokenfactorytypes.DenomAuthorityMetadata, error)
	BurnFrom(ctx sdk.Context, amount sdk.Coin, burnFrom string) error
	MintTo(ctx sdk.Context, amount sdk.Coin, mintTo string) error
	GetParams(ctx sdk.Context) (params tokenfactorytypes.Params)
	SetParams(ctx sdk.Context, params tokenfactorytypes.Params)
	GetDenomsFromCreator(ctx sdk.Context, creator string) []string
	InitGenesis(ctx sdk.Context, genState tokenfactorytypes.GenesisState)
	ExportGenesis(ctx sdk.Context) *tokenfactorytypes.GenesisState
}

type SubspacesKeeper interface {
	GetSubspace(ctx sdk.Context, subspaceID uint64) (subspace subspacestypes.Subspace, found bool)
	HasPermission(ctx sdk.Context, subspaceID uint64, sectionID uint32, user string, permission subspacestypes.Permission) bool
}
