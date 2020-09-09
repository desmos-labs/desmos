package keeper

import (
	"fmt"
	"regexp"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/profiles/types"
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
		case types.MsgRequestDTagTransfer:
			return handleMsgRequestDTagTransfer(ctx, keeper, msg)
		case types.MsgAcceptDTagTransfer:
			return handleMsgAcceptDTagTransfer(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized Profiles message type: %v", msg.Type())
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// ValidateProfile checks if the given profile is valid according to the current profile's module params
func ValidateProfile(ctx sdk.Context, keeper Keeper, profile types.Profile) error {
	params := keeper.GetParams(ctx)

	minMonikerLen := params.MonikerParams.MinMonikerLen.Int64()
	maxMonikerLen := params.MonikerParams.MaxMonikerLen.Int64()

	if profile.Moniker != nil {
		nameLen := int64(len(*profile.Moniker))
		if nameLen < minMonikerLen {
			return fmt.Errorf("profile moniker cannot be less than %d characters", minMonikerLen)
		}
		if nameLen > maxMonikerLen {
			return fmt.Errorf("profile moniker cannot exceed %d characters", maxMonikerLen)
		}
	}

	dTagRegEx := regexp.MustCompile(params.DtagParams.RegEx)
	minDtagLen := params.DtagParams.MinDtagLen.Int64()
	maxDtagLen := params.DtagParams.MaxDtagLen.Int64()
	dtagLen := int64(len(profile.DTag))

	if !dTagRegEx.MatchString(profile.DTag) {
		return fmt.Errorf("invalid profile dtag, it should match the following regEx %s", dTagRegEx)
	}

	if dtagLen < minDtagLen {
		return fmt.Errorf("profile dtag cannot be less than %d characters", minDtagLen)
	}

	if dtagLen > maxDtagLen {
		return fmt.Errorf("profile dtag cannot exceed %d characters", maxDtagLen)
	}

	maxBioLen := params.MaxBioLen.Int64()
	if profile.Bio != nil && int64(len(*profile.Bio)) > maxBioLen {
		return fmt.Errorf("profile biography cannot exceed %d characters", maxBioLen)
	}

	if err := profile.Validate(); err != nil {
		return err
	}

	return nil
}

// handleMsgSaveProfile handles the creation/edit of a profile
func handleMsgSaveProfile(ctx sdk.Context, keeper Keeper, msg types.MsgSaveProfile) (*sdk.Result, error) {
	profile, found := keeper.GetProfile(ctx, msg.Creator)

	// Create a new profile if not found
	if !found {
		profile = types.NewProfile(msg.Dtag, msg.Creator, ctx.BlockTime())
	}

	// Replace all editable fields (clients should autofill existing values)
	// We do not replace the tag since we do not want it to be editable
	profile = profile.
		WithDTag(msg.Dtag).
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
		sdk.NewAttribute(types.AttributeProfileCreationTime, profile.CreationDate.Format(time.RFC3339)),
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
			fmt.Sprintf("no profile associated with this address: %s", msg.Creator))
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

// handleMsgRequestDTagTransfer handles the request of a dTag transfer
func handleMsgRequestDTagTransfer(ctx sdk.Context, keeper Keeper, msg types.MsgRequestDTagTransfer) (*sdk.Result, error) {
	transferRequest := types.NewDTagTransferRequest(msg.CurrentOwner, msg.ReceivingUser)

	if err := keeper.SaveDTagTransferRequest(ctx, transferRequest); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeDTagTransferRequest,
		sdk.NewAttribute(types.AttributeCurrentOwner, transferRequest.CurrentOwner.String()),
		sdk.NewAttribute(types.AttributeCurrentOwner, transferRequest.ReceivingUser.String()),
	))

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(transferRequest),
		Events: ctx.EventManager().Events(),
	}

	return &result, nil
}

// handleMsgAcceptDTagTransfer handle the acceptance of a dTag transfer request
func handleMsgAcceptDTagTransfer(ctx sdk.Context, keeper Keeper, msg types.MsgAcceptDTagTransfer) (*sdk.Result, error) {
	requests := keeper.GetDTagTransferRequests(ctx, msg.CurrentOwner)

	// Check if the receiving user request is present, if not return error
	found := false
	for _, req := range requests {
		if req.ReceivingUser.Equals(msg.ReceivingUser) {
			found = true
			break
		}
	}

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("No request made from %s", msg.ReceivingUser))
	}

	// Get the current owner profile
	currentOwnerProfile, exist := keeper.GetProfile(ctx, msg.CurrentOwner)
	if !exist {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Profile of %s doesn't exist", msg.CurrentOwner))
	}

	// Save the current owner profile with his new dTag
	currentOwnerProfile = currentOwnerProfile.WithDTag(msg.NewDTag)
	err := keeper.SaveProfile(ctx, currentOwnerProfile)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// check for an existent profile of the receiving user
	receiverProfile, exist := keeper.GetProfile(ctx, msg.ReceivingUser)
	if !exist {
		receiverProfile = types.NewProfile(currentOwnerProfile.DTag, msg.ReceivingUser, ctx.BlockTime())
	} else {
		receiverProfile = receiverProfile.WithDTag(currentOwnerProfile.DTag)
	}

	err = keeper.SaveProfile(ctx, receiverProfile)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeDTagTransferReqAccepted,
		sdk.NewAttribute(types.AttributeCurrentOwner, msg.CurrentOwner.String()),
		sdk.NewAttribute(types.AttributeCurrentOwner, msg.ReceivingUser.String()),
	))

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(msg.CurrentOwner),
		Events: ctx.EventManager().Events(),
	}

	return &result, nil

}
