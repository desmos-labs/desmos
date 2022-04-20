package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v3/x/profiles/types"
)

var _ types.MsgServer = &msgServer{}

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the profiles MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{keeper}
}

func (k msgServer) SaveProfile(goCtx context.Context, msg *types.MsgSaveProfile) (*types.MsgSaveProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	profile, found, err := k.GetProfile(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	// Create a new profile if not found
	if !found {
		addr, err := sdk.AccAddressFromBech32(msg.Creator)
		if err != nil {
			return nil, err
		}

		profile, err = types.NewProfileFromAccount(msg.DTag, k.ak.GetAccount(ctx, addr), ctx.BlockTime())
		if err != nil {
			return nil, err
		}
	}

	// Update the existing profile with the values provided from the user
	updated, err := profile.Update(types.NewProfileUpdate(
		msg.DTag,
		msg.Nickname,
		msg.Bio,
		types.NewPictures(msg.ProfilePicture, msg.CoverPicture),
	))
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Validate the profile
	err = k.ValidateProfile(ctx, updated)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Save the profile
	err = k.Keeper.SaveProfile(ctx, updated)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
		sdk.NewEvent(
			types.EventTypeProfileSaved,
			sdk.NewAttribute(types.AttributeKeyProfileDTag, updated.DTag),
			sdk.NewAttribute(types.AttributeKeyProfileCreator, updated.GetAddress().String()),
			sdk.NewAttribute(types.AttributeKeyProfileCreationTime, updated.CreationDate.Format(time.RFC3339Nano)),
		),
	})

	return &types.MsgSaveProfileResponse{}, nil
}

func (k msgServer) DeleteProfile(goCtx context.Context, msg *types.MsgDeleteProfile) (*types.MsgDeleteProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.RemoveProfile(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
		sdk.NewEvent(
			types.EventTypeProfileDeleted,
			sdk.NewAttribute(types.AttributeKeyProfileCreator, msg.Creator),
		),
	})

	return &types.MsgDeleteProfileResponse{}, nil
}
