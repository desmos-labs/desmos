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
		switch msg := msg.(type) {
		case types.MsgCreateProfile:
			return handleMsgCreateProfile(ctx, keeper, msg)
		case types.MsgEditProfile:
			return handleMsgEditProfile(ctx, keeper, msg)
		case types.MsgDeleteProfile:
			return handleMsgDeleteProfile(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized Posts message type: %v", msg.Type())
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// handleMsgCreateProfile handles the creation of a profile
func handleMsgCreateProfile(ctx sdk.Context, keeper Keeper, msg types.MsgCreateProfile) (*sdk.Result, error) {
	// check if an account with the same moniker already exists
	// this check prevent the same user to create the same account multiple times
	if address := keeper.GetMonikerRelatedAddress(ctx, msg.Moniker); address.Equals(msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("An account with %s moniker already exists", msg.Moniker))
	}

	account := types.NewProfile(msg.Moniker, msg.Creator).
		WithName(msg.Name).
		WithSurname(msg.Surname).
		WithBio(msg.Bio).
		WithPictures(msg.Pictures)

	// Before saving this method checks if an account with the same moniker exist
	err := keeper.SaveProfile(ctx, account)
	if err != nil {
		return nil, err
	}

	createEvent := sdk.NewEvent(
		types.EventTypeProfileCreated,
		sdk.NewAttribute(types.AttributeProfileMoniker, account.Moniker),
		sdk.NewAttribute(types.AttributeProfileCreator, account.Creator.String()),
	)

	ctx.EventManager().EmitEvent(createEvent)

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(account.Moniker),
		Events: sdk.Events{createEvent},
	}

	return &result, nil
}

const (
	defaultValue = "default"
)

// returns the profile with the proper edited fields
// default string is used to let user replace previous inserted values with blank or empty ones
func getEditedProfile(account types.Profile, msg types.MsgEditProfile) types.Profile {
	account.Moniker = msg.NewMoniker

	if msg.Name != defaultValue {
		account = account.WithName(msg.Name)
	}

	if msg.Surname != defaultValue {
		account = account.WithSurname(msg.Surname)
	}

	if msg.Bio != defaultValue {
		account = account.WithBio(msg.Bio)
	}

	if msg.ProfilePic != defaultValue && msg.ProfileCov != defaultValue {
		pictures := types.NewPictures(msg.ProfilePic, msg.ProfileCov)
		account = account.WithPictures(&pictures)
	} else {
		if msg.ProfilePic != defaultValue && msg.ProfileCov == defaultValue {
			account.Pictures.Profile = msg.ProfilePic
		}
		if msg.ProfileCov != defaultValue && msg.ProfilePic == defaultValue {
			account.Pictures.Cover = msg.ProfileCov
		}
		account = account.WithPictures(account.Pictures)
	}

	return account
}

// handleMsgEditProfile handles the edit of a profile
func handleMsgEditProfile(ctx sdk.Context, keeper Keeper, msg types.MsgEditProfile) (*sdk.Result, error) {
	account, found := keeper.GetProfile(ctx, msg.Creator)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("No existent profile to edit for address: %s", msg.Creator))
	}
	previousMoniker := account.Moniker

	account = getEditedProfile(account, msg)

	err := keeper.SaveProfile(ctx, account)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("An account with moniker: %s has already been created", msg.NewMoniker))
	}

	keeper.DeleteMonikerAddressAssociation(ctx, previousMoniker)

	createEvent := sdk.NewEvent(
		types.EventTypeProfileEdited,
		sdk.NewAttribute(types.AttributeProfileMoniker, account.Moniker),
		sdk.NewAttribute(types.AttributeProfileCreator, account.Creator.String()),
	)

	ctx.EventManager().EmitEvent(createEvent)

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(account.Moniker),
		Events: sdk.Events{createEvent},
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

	keeper.DeleteProfile(ctx, profile.Creator, profile.Moniker)

	createEvent := sdk.NewEvent(
		types.EventTypeProfileDeleted,
		sdk.NewAttribute(types.AttributeProfileMoniker, profile.Moniker),
		sdk.NewAttribute(types.AttributeProfileCreator, profile.Creator.String()),
	)

	ctx.EventManager().EmitEvent(createEvent)

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(profile.Moniker),
		Events: sdk.Events{createEvent},
	}

	return &result, nil
}
