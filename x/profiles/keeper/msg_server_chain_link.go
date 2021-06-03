package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (k msgServer) LinkChainAccount(goCtx context.Context, msg *types.MsgLinkChainAccount) (*types.LinkChainAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	srcAddrData, err := types.UnpackAddressData(k.cdc, msg.SourceAddress)
	if err != nil {
		return nil, err
	}

	if err := srcAddrData.Validate(); err != nil {
		return nil, err
	}

	if err := msg.SourceProof.Verify(k.cdc); err != nil {
		return nil, err
	}

	if err := msg.DestinationProof.Verify(k.cdc); err != nil {
		return nil, err
	}

	link := types.NewChainLink(
		srcAddrData,
		msg.SourceProof,
		msg.SourceChainConfig,
		ctx.BlockTime(),
	)

	if err := k.StoreChainLink(ctx, msg.DestinationAddress, link); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeLinkChainAccount,
		sdk.NewAttribute(types.AttributeChainLinkAccountTarget, srcAddrData.GetAddress()),
		sdk.NewAttribute(types.AttributeChainLinkSourceChainName, msg.SourceChainConfig.Name),
		sdk.NewAttribute(types.AttributeChainLinkAccountOwner, msg.DestinationAddress),
		sdk.NewAttribute(types.AttributeChainLinkCreated, link.CreationTime.Format(time.RFC3339Nano)),
	))

	return &types.LinkChainAccountResponse{}, nil
}

func (k msgServer) UnlinkChainAccount(goCtx context.Context, msg *types.MsgUnlinkChainAccount) (*types.UnlinkChainAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.DeleteChainLink(ctx, msg.Owner, msg.ChainName, msg.Target); err != nil {
		return &types.UnlinkChainAccountResponse{}, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeUnlinkChainAccount,
		sdk.NewAttribute(types.AttributeChainLinkAccountTarget, msg.Target),
		sdk.NewAttribute(types.AttributeChainLinkSourceChainName, msg.ChainName),
		sdk.NewAttribute(types.AttributeChainLinkAccountOwner, msg.Owner),
	))

	return &types.UnlinkChainAccountResponse{}, nil
}
