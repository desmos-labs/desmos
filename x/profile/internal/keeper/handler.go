package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
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

// ValidateProfile checks if the given profile is valid according to the current profile's module params
func ValidateProfile(ctx sdk.Context, keeper Keeper, profile types.Profile) error {
	params := keeper.GetParams(ctx)

	minNameSurnameLen := params.NameSurnameLengths.MinNameSurnameLen.Int64()
	maxNameSurnameLen := params.NameSurnameLengths.MaxNameSurnameLen.Int64()

	if profile.Name != nil {
		nameLen := int64(len(*profile.Name))
		if nameLen < minNameSurnameLen {
			return fmt.Errorf("Profile name cannot be less than %d characters", minNameSurnameLen)
		}
		if nameLen > maxNameSurnameLen {
			return fmt.Errorf("Profile name cannot exceed %d characters", maxNameSurnameLen)
		}
	}

	if profile.Surname != nil {
		surNameLen := int64(len(*profile.Surname))
		if surNameLen < minNameSurnameLen {
			return fmt.Errorf("Profile surname cannot be less than %d characters", minNameSurnameLen)
		}
		if surNameLen > maxNameSurnameLen {
			return fmt.Errorf("Profile surname cannot exceed %d characters", maxNameSurnameLen)
		}
	}

	minMonikerLen := params.MonikerLengths.MinMonikerLen.Int64()
	maxMonikerLen := params.MonikerLengths.MaxMonikerLen.Int64()
	monikerLen := int64(len(profile.Moniker))

	if monikerLen < minMonikerLen {
		return fmt.Errorf("Profile moniker cannot be less than %d characters", minMonikerLen)
	}

	if monikerLen > maxMonikerLen {
		return fmt.Errorf("Profile moniker cannot exceed %d characters", maxMonikerLen)
	}

	maxBioLen := params.MaxBioLen.Int64()
	if profile.Bio != nil && int64(len(*profile.Bio)) > maxBioLen {
		return fmt.Errorf("Profile biography cannot exceed %d characters", maxBioLen)
	}

	if err := profile.Validate(); err != nil {
		return err
	}

	return nil
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
		profile = types.NewProfile(msg.Dtag, msg.Creator, ctx.BlockTime())
	}

	// Replace all editable fields (clients should autofill existing values)
	// We do not replace the tag since we do not want it to be editable
	profile = profile.
		WithMoniker(msg.Moniker).
		WithBio(msg.Bio).
		WithPictures(msg.ProfilePic, msg.CoverPic)
	err := ValidateProfile(ctx, keeper, profile)
	if err != nil {
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
