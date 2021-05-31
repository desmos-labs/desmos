package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (k msgServer) LinkChainAccount(goCtx context.Context, msg *types.MsgLinkChainAccount) (*types.LinkChainAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	if err := msg.SourceProof.Verify(k.cdc); err != nil {
		return nil, err
	}

	if err := msg.DestinationProof.Verify(k.cdc); err != nil {
		return nil, err
	}

	// Check if address has the profile
	profile, found, err := k.GetProfile(ctx, msg.DestinationAddress)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, fmt.Errorf("address does not have any profile")
	}

	chainLink := types.NewChainLink(
		msg.SourceAddress,
		msg.SourceProof,
		msg.SourceChainConfig,
		ctx.BlockTime(),
	)

	if err := k.StoreChainLink(ctx, chainLink); err != nil {
		return nil, err
	}

	// Store chain link to the profile
	profile.ChainsLinks = append(profile.ChainsLinks, chainLink)
	if err := k.StoreProfile(ctx, profile); err != nil {
		k.DeleteChainLink(ctx, chainLink.ChainConfig.Name, chainLink.Address)
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeLinkChainAccount,
		sdk.NewAttribute(types.AttributeChainLinkAccountTarget, msg.SourceAddress),
		sdk.NewAttribute(types.AttributeChainLinkSourceChainName, msg.SourceChainConfig.Name),
		sdk.NewAttribute(types.AttributeChainLinkAccountOwner, msg.DestinationAddress),
		sdk.NewAttribute(types.AttributeChainLinkCreated, chainLink.CreationTime.Format(time.RFC3339Nano)),
	))

	return &types.LinkChainAccountResponse{}, nil
}

func (k msgServer) UnlinkChainAccount(goCtx context.Context, msg *types.MsgUnlinkChainAccount) (*types.UnlinkChainAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Check if owner has the profile and get the profile
	profile, found, err := k.GetProfile(ctx, msg.Owner)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, ("non existent profile on destination address"))
	}

	isTargetExist := false
	newChainsLinks := []types.ChainLink{}

	// Try to find the target link
	for _, link := range profile.ChainsLinks {
		chainName := link.ChainConfig.Name
		address := link.Address
		if chainName == msg.ChainName && address == msg.Target {
			isTargetExist = true
			continue
		}
		newChainsLinks = append(newChainsLinks, link)
	}

	if !isTargetExist {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, ("non existent target link in the profile"))
	}

	// Update profile status
	profile.ChainsLinks = newChainsLinks
	err = k.StoreProfile(ctx, profile)
	if err != nil {
		return nil, err
	}

	k.DeleteChainLink(ctx, msg.ChainName, msg.Target)

	return &types.UnlinkChainAccountResponse{}, nil
}
