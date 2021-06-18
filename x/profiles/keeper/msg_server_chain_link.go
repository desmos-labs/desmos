package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (k msgServer) LinkChainAccount(goCtx context.Context, msg *types.MsgLinkChainAccount) (*types.LinkChainAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	srcAddrData, err := types.UnpackAddressData(k.cdc, msg.ChainAddress)
	if err != nil {
		return nil, err
	}

	link := types.NewChainLink(msg.Signer, srcAddrData, msg.Proof, msg.ChainConfig, ctx.BlockTime())
	err = k.StoreChainLink(ctx, link)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeLinkChainAccount,
		sdk.NewAttribute(types.AttributeChainLinkSourceAddress, srcAddrData.GetAddress()),
		sdk.NewAttribute(types.AttributeChainLinkSourceChainName, msg.ChainConfig.Name),
		sdk.NewAttribute(types.AttributeChainLinkDestinationAddress, msg.Signer),
		sdk.NewAttribute(types.AttributeChainLinkCreationTime, link.CreationTime.Format(time.RFC3339Nano)),
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
		sdk.NewAttribute(types.AttributeChainLinkSourceAddress, msg.Target),
		sdk.NewAttribute(types.AttributeChainLinkSourceChainName, msg.ChainName),
		sdk.NewAttribute(types.AttributeChainLinkDestinationAddress, msg.Owner),
	))

	return &types.UnlinkChainAccountResponse{}, nil
}
