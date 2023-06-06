package keeper

import (
	"context"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	tokenfactorytypes "github.com/osmosis-labs/osmosis/v15/x/tokenfactory/types"

	subspacestypes "github.com/desmos-labs/desmos/v5/x/subspaces/types"
	"github.com/desmos-labs/desmos/v5/x/tokenfactory/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the stored MsgServer interface for the provided keeper
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

// CreateDenom defines a rpc method for MsgCreateDenom
func (k msgServer) CreateDenom(goCtx context.Context, msg *types.MsgCreateDenom) (*types.MsgCreateDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	subspace, exists := k.sk.GetSubspace(ctx, msg.SubspaceID)
	if !exists {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check the permission to manage the subspace tokens
	if !k.sk.HasPermission(ctx, msg.SubspaceID, subspacestypes.RootSectionID, msg.Sender, types.PermissionManageSubspaceTokens) {
		return nil, errors.Wrap(subspacestypes.ErrPermissionDenied, "you cannot manage the subspace tokens inside this subspace")
	}

	denom, err := k.tfk.CreateDenom(ctx, subspace.Treasury, msg.Subdenom)
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateDenomResponse{
		NewTokenDenom: denom,
	}, nil
}

// Mint defines a rpc method for MsgMint
func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	subspace, exists := k.sk.GetSubspace(ctx, msg.SubspaceID)
	if !exists {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check the permission to manage the subspace tokens
	if !k.sk.HasPermission(ctx, msg.SubspaceID, subspacestypes.RootSectionID, msg.Sender, types.PermissionManageSubspaceTokens) {
		return nil, errors.Wrap(subspacestypes.ErrPermissionDenied, "you cannot manage the subspace tokens inside this subspace")
	}

	// Check if the denom exists
	_, denomExists := k.bk.GetDenomMetaData(ctx, msg.Amount.Denom)
	if !denomExists {
		return nil, tokenfactorytypes.ErrDenomDoesNotExist.Wrapf("denom: %s", msg.Amount.Denom)
	}

	authorityMetadata, err := k.tfk.GetAuthorityMetadata(ctx, msg.Amount.GetDenom())
	if err != nil {
		return nil, err
	}

	// Check if the subspace treasury is the admin of the denom
	if subspace.Treasury != authorityMetadata.GetAdmin() {
		return nil, tokenfactorytypes.ErrUnauthorized
	}

	err = k.tfk.MintTo(ctx, msg.Amount, msg.MintToAddress)
	if err != nil {
		return nil, err
	}

	return &types.MsgMintResponse{}, nil
}

// Burn defines a rpc method for MsgBurn
func (k msgServer) Burn(goCtx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	subspace, exists := k.sk.GetSubspace(ctx, msg.SubspaceID)
	if !exists {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check the permission to manage the subspace tokens
	if !k.sk.HasPermission(ctx, msg.SubspaceID, subspacestypes.RootSectionID, msg.Sender, types.PermissionManageSubspaceTokens) {
		return nil, errors.Wrap(subspacestypes.ErrPermissionDenied, "you cannot manage the subspace tokens inside this subspace")
	}

	authorityMetadata, err := k.tfk.GetAuthorityMetadata(ctx, msg.Amount.GetDenom())
	if err != nil {
		return nil, err
	}

	// Check if the subspace treasury is the admin of the denom
	if subspace.Treasury != authorityMetadata.GetAdmin() {
		return nil, tokenfactorytypes.ErrUnauthorized
	}

	err = k.tfk.BurnFrom(ctx, msg.Amount, subspace.Treasury)
	if err != nil {
		return nil, err
	}

	return &types.MsgBurnResponse{}, nil
}

// SetDenomMetadata defines a rpc method for MsgSetDenomMetadata
func (k msgServer) SetDenomMetadata(goCtx context.Context, msg *types.MsgSetDenomMetadata) (*types.MsgSetDenomMetadataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	subspace, exists := k.sk.GetSubspace(ctx, msg.SubspaceID)
	if !exists {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check the permission to manage the subspace tokens
	if !k.sk.HasPermission(ctx, msg.SubspaceID, subspacestypes.RootSectionID, msg.Sender, types.PermissionManageSubspaceTokens) {
		return nil, errors.Wrap(subspacestypes.ErrPermissionDenied, "you cannot manage the subspace tokens inside this subspace")
	}

	authorityMetadata, err := k.tfk.GetAuthorityMetadata(ctx, msg.Metadata.Base)
	if err != nil {
		return nil, err
	}

	// Check if the subspace treasury is the admin of the denom
	if subspace.Treasury != authorityMetadata.GetAdmin() {
		return nil, tokenfactorytypes.ErrUnauthorized
	}

	k.bk.SetDenomMetaData(ctx, msg.Metadata)

	return &types.MsgSetDenomMetadataResponse{}, nil
}

// UpdateParams defines a rpc method for MsgUpdateParams
func (k msgServer) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	authority := k.authority
	if authority != msg.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", authority, msg.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	k.tfk.SetParams(ctx, types.ToOsmosisTokenFactoryParams(msg.Params))

	return &types.MsgUpdateParamsResponse{}, nil
}
