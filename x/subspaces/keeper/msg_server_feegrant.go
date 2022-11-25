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
	if _, found, _ := k.GetUserAllowance(ctx, msg.SubspaceID, msg.Granter, msg.Grantee); found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "fee allowance already exists")
	}
	allowance, err := msg.GetFeeAllowanceI()
	if err != nil {
		return nil, err
	}
	err = k.Keeper.GrantUserAllowance(ctx, msg.SubspaceID, msg.Granter, msg.Grantee, allowance)
	if err != nil {
		return nil, err
	}
	return &types.MsgGrantUserAllowanceResponse{}, nil
}

func (k msgServer) GrantGroupAllowance(goCtx context.Context, msg *types.MsgGrantGroupAllowance) (*types.MsgGrantGroupAllowanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Checking for duplicate entry
	if _, found, _ := k.GetGroupAllowance(ctx, msg.SubspaceID, msg.Granter, msg.GroupID); found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "fee allowance already exists")
	}
	allowance, err := msg.GetFeeAllowanceI()
	if err != nil {
		return nil, err
	}
	err = k.Keeper.GrantGroupAllowance(ctx, msg.SubspaceID, msg.Granter, msg.GroupID, allowance)
	if err != nil {
		return nil, err
	}
	return &types.MsgGrantGroupAllowanceResponse{}, nil
}
