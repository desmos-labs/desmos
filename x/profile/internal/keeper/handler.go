package keeper

import (
	"fmt"

	"github.com/desmos-labs/desmos/x/profile/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler returns a handler for "profile" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case types.MsgSaveProfile:
			return handleMsgSaveProfile(ctx, keeper, msg)
		case types.MsgDeleteProfile:
			return handleMsgDeleteProfile(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized Posts message type: %v", msg.Type())
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// handleMsgSaveProfile handles the creation/edit of a profile
func handleMsgSaveProfile(ctx sdk.Context, keeper Keeper, msg types.MsgSaveProfile) (*sdk.Result, error) {
	profile, found := keeper.GetProfile(ctx, msg.Creator)

	// If it's found and the DTag is not the same, return an error
	if found && profile.DTag != msg.Dtag {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "wrong dtag provided. Make sure to use the current one")
	}

	// Create a new profile if not found
	if !found {
		profile = types.NewProfile(msg.Dtag, msg.Creator)
	}

	// Replace all editable fields (clients should autofill existing values)
	// We do not replace the tag since we do not want it to be editable
	profile = profile.
		WithMoniker(msg.Moniker).
		WithBio(msg.Bio).
		WithPictures(msg.ProfilePic, msg.CoverPic)
	if err := profile.Validate(); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Save the profile
	if err := keeper.SaveProfile(ctx, profile); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeProfileSaved,
		sdk.NewAttribute(types.AttributeProfileDtag, profile.DTag),
		sdk.NewAttribute(types.AttributeProfileCreator, profile.Creator.String()),
	))

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(profile.DTag),
		Events: ctx.EventManager().Events(),
	}

	return &result, nil
}

// handleMsgDeleteProfile handles the deletion of a profile
func handleMsgDeleteProfile(ctx sdk.Context, keeper Keeper, msg types.MsgDeleteProfile) (*sdk.Result, error) {
	profile, found := keeper.GetProfile(ctx, msg.Creator)

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("No profile associated with this address: %s", msg.Creator))
	}

	keeper.DeleteProfile(ctx, profile.Creator, profile.DTag)

	createEvent := sdk.NewEvent(
		types.EventTypeProfileDeleted,
		sdk.NewAttribute(types.AttributeProfileDtag, profile.DTag),
		sdk.NewAttribute(types.AttributeProfileCreator, profile.Creator.String()),
	)

	ctx.EventManager().EmitEvent(createEvent)

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(profile.DTag),
		Events: ctx.EventManager().Events(),
	}

	return &result, nil
}
