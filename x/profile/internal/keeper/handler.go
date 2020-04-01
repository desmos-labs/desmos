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
			return handleMsgCreateAccount(ctx, keeper, msg)
		case types.MsgEditProfile:
			return handleMsgEditAccount(ctx, keeper, msg)
		case types.MsgDeleteProfile:
			return handleMsgDeleteAccount(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized Posts message type: %v", msg.Type())
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// handleMsgCreateAccount handles the creation of an account
func handleMsgCreateAccount(ctx sdk.Context, keeper Keeper, msg types.MsgCreateProfile) (*sdk.Result, error) {

	// check if an account with the same moniker already exists
	// this check prevent the same user to create the same account multiple times
	if _, found := keeper.GetAccount(ctx, msg.Moniker); found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("An account with %s moniker already exist", msg.Moniker))
	}

	account := types.Profile{
		Name:     msg.Name,
		Surname:  msg.Surname,
		Moniker:  msg.Moniker,
		Bio:      msg.Bio,
		Pictures: msg.Pictures,
		Creator:  msg.Creator,
	}

	// Before saving this method checks if an account with the same moniker exist
	err := keeper.SaveAccount(ctx, account)
	if err != nil {
		return nil, err
	}

	createEvent := sdk.NewEvent(
		types.EventTypeAccountCreated,
		sdk.NewAttribute(types.AttributeAccountMoniker, account.Moniker),
		sdk.NewAttribute(types.AttributeAccountCreator, account.Creator.String()),
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
		account.Name = msg.Name
	}

	if msg.Surname != defaultValue {
		account.Surname = msg.Surname
	}

	if msg.Bio != defaultValue {
		account.Bio = msg.Bio
	}

	if msg.Pictures.Profile != defaultValue {
		account.Pictures.Profile = msg.Pictures.Profile
	}

	if msg.Pictures.Cover != defaultValue {
		account.Pictures.Cover = msg.Pictures.Cover
	}

	return account
}

// handleMsgEditAccount handles the edit of an account
func handleMsgEditAccount(ctx sdk.Context, keeper Keeper, msg types.MsgEditProfile) (*sdk.Result, error) {

	account, found := keeper.GetAccount(ctx, msg.PreviousMoniker)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("No existent profile with moniker: %s", msg.PreviousMoniker))
	}

	account = getEditedProfile(account, msg)

	// New moniker already taken
	err := keeper.SaveAccount(ctx, account)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	createEvent := sdk.NewEvent(
		types.EventTypeAccountEdited,
		sdk.NewAttribute(types.AttributeAccountMoniker, account.Moniker),
		sdk.NewAttribute(types.AttributeAccountCreator, account.Creator.String()),
	)

	ctx.EventManager().EmitEvent(createEvent)

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(account.Moniker),
		Events: sdk.Events{createEvent},
	}

	return &result, nil
}

// handleMsgDeleteAccount handles the deletion of an account
func handleMsgDeleteAccount(ctx sdk.Context, keeper Keeper, msg types.MsgDeleteProfile) (*sdk.Result, error) {

	// check if an account with the same moniker exists
	acc, found := keeper.GetAccount(ctx, msg.Moniker)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("An account with %s moniker doesn't exist", msg.Moniker))
	}

	// check if the creator of the message match the account creator
	if !acc.Creator.Equals(msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("You cannot delete an account that is not yours"))
	}

	keeper.DeleteAccount(ctx, msg.Moniker)

	createEvent := sdk.NewEvent(
		types.EventTypeAccountDeleted,
		sdk.NewAttribute(types.AttributeAccountMoniker, acc.Moniker),
		sdk.NewAttribute(types.AttributeAccountCreator, acc.Creator.String()),
	)

	ctx.EventManager().EmitEvent(createEvent)

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(acc.Moniker),
		Events: sdk.Events{createEvent},
	}

	return &result, nil
}
