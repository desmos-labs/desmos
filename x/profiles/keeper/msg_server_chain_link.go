package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v2/x/profiles/types"
)

func (k msgServer) LinkChainAccount(goCtx context.Context, msg *types.MsgLinkChainAccount) (*types.MsgLinkChainAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	srcAddrData, err := types.UnpackAddressData(k.Cdc, msg.ChainAddress)
	if err != nil {
		return nil, err
	}

	link := types.NewChainLink(msg.Signer, srcAddrData, msg.Proof, msg.ChainConfig, ctx.BlockTime())
	err = k.SaveChainLink(ctx, link)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeLinkChainAccount,
		sdk.NewAttribute(types.AttributeChainLinkSourceAddress, srcAddrData.GetValue()),
		sdk.NewAttribute(types.AttributeChainLinkSourceChainName, msg.ChainConfig.Name),
		sdk.NewAttribute(types.AttributeChainLinkDestinationAddress, msg.Signer),
		sdk.NewAttribute(types.AttributeChainLinkCreationTime, link.CreationTime.Format(time.RFC3339Nano)),
	))

	return &types.MsgLinkChainAccountResponse{}, nil
}

func (k msgServer) UnlinkChainAccount(goCtx context.Context, msg *types.MsgUnlinkChainAccount) (*types.MsgUnlinkChainAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.DeleteChainLink(ctx, msg.Owner, msg.ChainName, msg.Target); err != nil {
		return &types.MsgUnlinkChainAccountResponse{}, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeUnlinkChainAccount,
		sdk.NewAttribute(types.AttributeChainLinkSourceAddress, msg.Target),
		sdk.NewAttribute(types.AttributeChainLinkSourceChainName, msg.ChainName),
		sdk.NewAttribute(types.AttributeChainLinkDestinationAddress, msg.Owner),
	))

	return &types.MsgUnlinkChainAccountResponse{}, nil
}
