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
	if !found {
		profile = types.NewProfile(msg.Creator)
	}
	// replace all editable fields (clients should autofill existing values)
	profile = profile.
		WithMoniker(msg.Moniker).
		WithName(msg.Name).
		WithSurname(msg.Surname).
		WithBio(msg.Bio).
		WithPictures(msg.ProfilePic, msg.ProfileCov)

	err := ValidateProfile(ctx, keeper, profile)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	err = keeper.SaveProfile(ctx, profile)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	createEvent := sdk.NewEvent(
		types.EventTypeProfileSaved,
		sdk.NewAttribute(types.AttributeProfileMoniker, profile.Moniker),
		sdk.NewAttribute(types.AttributeProfileCreator, profile.Creator.String()),
	)

	ctx.EventManager().EmitEvent(createEvent)

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(profile.Moniker),
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

	keeper.DeleteProfile(ctx, profile.Creator, profile.Moniker)

	createEvent := sdk.NewEvent(
		types.EventTypeProfileDeleted,
		sdk.NewAttribute(types.AttributeProfileMoniker, profile.Moniker),
		sdk.NewAttribute(types.AttributeProfileCreator, profile.Creator.String()),
	)

	ctx.EventManager().EmitEvent(createEvent)

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(profile.Moniker),
		Events: ctx.EventManager().Events(),
	}

	return &result, nil
}
