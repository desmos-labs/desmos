package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

func (k msgServer) GrantAllowance(goCtx context.Context, msg *types.MsgGrantAllowance) (*types.MsgGrantAllowanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	grantee, err := sdk.AccAddressFromBech32(msg.Grantee)
	if err != nil {
		return nil, err
	}

	granter, err := sdk.AccAddressFromBech32(msg.Granter)
	if err != nil {
		return nil, err
	}

	// Checking for duplicate entry
	if f, _ := k.Keeper.GetAllowance(ctx, granter, grantee, msg.SubspaceId); f != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "fee allowance already exists")
	}

	allowance, err := msg.GetFeeAllowanceI()
	if err != nil {
		return nil, err
	}

	err = k.Keeper.GrantAllowance(ctx, granter, grantee, msg.SubspaceId, allowance)
	if err != nil {
		return nil, err
	}

	return &types.MsgGrantAllowanceResponse{}, nil
}
