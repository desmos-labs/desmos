package keeper

import (
	"context"
	"time"

	"cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v6/x/profiles/types"
)

// LinkChainAccount defines a rpc method for MsgLinkChainAccount
func (k MsgServer) LinkChainAccount(goCtx context.Context, msg *types.MsgLinkChainAccount) (*types.MsgLinkChainAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	srcAddrData, err := types.UnpackAddressData(k.cdc, msg.ChainAddress)
	if err != nil {
		return nil, err
	}

	link := types.NewChainLink(msg.Signer, srcAddrData, msg.Proof, msg.ChainConfig, ctx.BlockTime())
	err = k.SaveChainLink(ctx, link)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer),
		),
		sdk.NewEvent(
			types.EventTypeLinkChainAccount,
			sdk.NewAttribute(types.AttributeKeyChainLinkExternalAddress, srcAddrData.GetValue()),
			sdk.NewAttribute(types.AttributeKeyChainLinkChainName, msg.ChainConfig.Name),
			sdk.NewAttribute(types.AttributeKeyChainLinkOwner, msg.Signer),
			sdk.NewAttribute(types.AttributeKeyChainLinkCreationTime, link.CreationTime.Format(time.RFC3339Nano)),
		),
	})

	return &types.MsgLinkChainAccountResponse{}, nil
}

// UnlinkChainAccount defines a rpc method for MsgUnlinkChainAccount
func (k MsgServer) UnlinkChainAccount(goCtx context.Context, msg *types.MsgUnlinkChainAccount) (*types.MsgUnlinkChainAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the chain link
	link, found := k.GetChainLink(ctx, msg.Owner, msg.ChainName, msg.Target)
	if !found {
		return nil, errors.Wrap(sdkerrors.ErrNotFound, "chain link not found")
	}

	// Delete the link
	k.DeleteChainLink(ctx, link)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner),
		),
		sdk.NewEvent(
			types.EventTypeUnlinkChainAccount,
			sdk.NewAttribute(types.AttributeKeyChainLinkExternalAddress, msg.Target),
			sdk.NewAttribute(types.AttributeKeyChainLinkChainName, msg.ChainName),
			sdk.NewAttribute(types.AttributeKeyChainLinkOwner, msg.Owner),
		),
	})

	return &types.MsgUnlinkChainAccountResponse{}, nil
}

// SetDefaultExternalAddress defines a rpc method for MsgSetDefaultExternalAddress
func (k MsgServer) SetDefaultExternalAddress(goCtx context.Context, msg *types.MsgSetDefaultExternalAddress) (*types.MsgSetDefaultExternalAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the chain link
	_, found := k.GetChainLink(ctx, msg.Signer, msg.ChainName, msg.Target)
	if !found {
		return nil, errors.Wrap(sdkerrors.ErrNotFound, "chain link not found")
	}

	k.SaveDefaultExternalAddress(ctx, msg.Signer, msg.ChainName, msg.Target)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer),
		),
		sdk.NewEvent(
			types.EventTypeSetDefaultExternalAddress,
			sdk.NewAttribute(types.AttributeKeyChainLinkChainName, msg.ChainName),
			sdk.NewAttribute(types.AttributeKeyChainLinkExternalAddress, msg.Target),
			sdk.NewAttribute(types.AttributeKeyChainLinkOwner, msg.Signer),
		),
	})

	return &types.MsgSetDefaultExternalAddressResponse{}, nil
}
