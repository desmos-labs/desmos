package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

func (k msgServer) GrantUserAllowance(goCtx context.Context, msg *types.MsgGrantUserAllowance) (*types.MsgGrantUserAllowanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Checking for duplicate entry
	_, found, err := k.GetUserAllowance(ctx, msg.SubspaceID, msg.Granter, msg.Grantee)
	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "fee allowance already exists")
	}
	if err != nil {
		return nil, err
	}

	allowance, err := msg.GetUnpackedAllowance()
	if err != nil {
		return nil, err
	}
	err = k.Keeper.GrantUserAllowance(ctx, msg.SubspaceID, msg.Granter, msg.Grantee, allowance)
	if err != nil {
		return nil, err
	}
	return &types.MsgGrantUserAllowanceResponse{}, nil
}

func (k msgServer) RevokeUserAllowance(goCtx context.Context, msg *types.MsgRevokeUserAllowance) (*types.MsgRevokeUserAllowanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := k.RevokeUserGrant(ctx, msg.SubspaceID, msg.Granter, msg.Grantee)
	if err != nil {
		return nil, err
	}
	return &types.MsgRevokeUserAllowanceResponse{}, nil
}

func (k msgServer) GrantGroupAllowance(goCtx context.Context, msg *types.MsgGrantGroupAllowance) (*types.MsgGrantGroupAllowanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Checking for duplicate entry
	if _, found, _ := k.GetGroupAllowance(ctx, msg.SubspaceID, msg.Granter, msg.GroupID); found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "fee allowance already exists")
	}
	allowance, err := msg.GetUnpackedAllowance()
	if err != nil {
		return nil, err
	}
	err = k.Keeper.GrantGroupAllowance(ctx, msg.SubspaceID, msg.Granter, msg.GroupID, allowance)
	if err != nil {
		return nil, err
	}
	return &types.MsgGrantGroupAllowanceResponse{}, nil
}

func (k msgServer) RevokeGroupAllowance(goCtx context.Context, msg *types.MsgRevokeGroupAllowance) (*types.MsgRevokeGroupAllowanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Checking for duplicate entry
	err := k.RevokeGroupGrant(ctx, msg.SubspaceID, msg.Granter, msg.GroupID)
	if err != nil {
		return nil, err
	}
	return &types.MsgRevokeGroupAllowanceResponse{}, nil
}
