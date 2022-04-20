package keeper

import (
	"context"
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/profiles/types"
)

func (k msgServer) LinkChainAccount(goCtx context.Context, msg *types.MsgLinkChainAccount) (*types.MsgLinkChainAccountResponse, error) {
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
			sdk.NewAttribute(types.AttributeKeyChainLinkSourceAddress, srcAddrData.GetValue()),
			sdk.NewAttribute(types.AttributeKeyChainLinkSourceChainName, msg.ChainConfig.Name),
			sdk.NewAttribute(types.AttributeKeyChainLinkDestinationAddress, msg.Signer),
			sdk.NewAttribute(types.AttributeKeyChainLinkCreationTime, link.CreationTime.Format(time.RFC3339Nano)),
		),
	})

	return &types.MsgLinkChainAccountResponse{}, nil
}

func (k msgServer) UnlinkChainAccount(goCtx context.Context, msg *types.MsgUnlinkChainAccount) (*types.MsgUnlinkChainAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the chain link
	link, found := k.GetChainLink(ctx, msg.Owner, msg.ChainName, msg.Target)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "chain link not found")
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
			sdk.NewAttribute(types.AttributeKeyChainLinkSourceAddress, msg.Target),
			sdk.NewAttribute(types.AttributeKeyChainLinkSourceChainName, msg.ChainName),
			sdk.NewAttribute(types.AttributeKeyChainLinkDestinationAddress, msg.Owner),
		),
	})

	return &types.MsgUnlinkChainAccountResponse{}, nil
}
