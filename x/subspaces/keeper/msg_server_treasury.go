package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

// GrantTreasuryAuthorization defines a rpc method for MsgGrantTreasuryAuthorization
func (k msgServer) GrantTreasuryAuthorization(goCtx context.Context, msg *types.MsgGrantTreasuryAuthorization) (*types.MsgGrantTreasuryAuthorizationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	subspace, found := k.GetSubspace(ctx, msg.SubspaceID)
	if !found {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("subspace with id %d not found", msg.SubspaceID)
	}

	if !k.HasPermission(ctx, msg.SubspaceID, types.RootSectionID, msg.Granter, types.PermissionManageTreasuryAuthorization) {
		return nil, types.ErrPermissionDenied.Wrap("you cannot manage this subspace's treasury authorizations")
	}

	treasury, err := sdk.AccAddressFromBech32(subspace.Treasury)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid treasury address: %s", treasury.String())
	}

	grantee, err := sdk.AccAddressFromBech32(msg.Grantee)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid grantee address: %s", msg.Grantee)
	}

	err = k.authzk.SaveGrant(ctx, grantee, treasury, msg.Grant.GetAuthorization(), msg.Grant.Expiration)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Granter),
		),
		sdk.NewEvent(
			types.EventTypeGrantTreasuryAuthorization,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyGranter, msg.Granter),
			sdk.NewAttribute(types.AttributeKeyGrantee, msg.Grantee),
		),
	})
	return &types.MsgGrantTreasuryAuthorizationResponse{}, nil
}

// RevokeTreasuryAuthorization defines a rpc method for MsgRevokeTreasuryAuthorization
func (k msgServer) RevokeTreasuryAuthorization(goCtx context.Context, msg *types.MsgRevokeTreasuryAuthorization) (*types.MsgRevokeTreasuryAuthorizationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	subspace, found := k.GetSubspace(ctx, msg.SubspaceID)
	if !found {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("subspace with id %d not found", msg.SubspaceID)
	}

	if !k.HasPermission(ctx, msg.SubspaceID, types.RootSectionID, msg.Granter, types.PermissionManageTreasuryAuthorization) {
		return nil, types.ErrPermissionDenied.Wrap("you cannot manage this subspace treasury authorization")
	}

	treasury, err := sdk.AccAddressFromBech32(subspace.Treasury)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid treasury address: %s", treasury.String())
	}

	grantee, err := sdk.AccAddressFromBech32(msg.Grantee)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid grantee address: %s", msg.Grantee)
	}

	err = k.authzk.DeleteGrant(ctx, grantee, treasury, msg.MsgTypeUrl)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Granter),
		),
		sdk.NewEvent(
			types.EventTypeRevokeTreasuryAuthorization,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyGranter, msg.Granter),
			sdk.NewAttribute(types.AttributeKeyGrantee, msg.Grantee),
		),
	})

	return &types.MsgRevokeTreasuryAuthorizationResponse{}, nil
}
