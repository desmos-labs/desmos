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
		case types.MsgCreateMonoDirectionalRelationship:
			return handleMsgCreateMonoDirectionalRelationship(ctx, keeper, msg)
		case types.MsgRequestBidirectionalRelationship:
			return handleMsgRequestBiDirectionalRelationship(ctx, keeper, msg)
		case types.MsgAcceptBidirectionalRelationship:
			return handleMsgAcceptBidirectionalRelationship(ctx, keeper, msg)
		case types.MsgDenyBidirectionalRelationship:
			return handleMsgDenyBidirectionalRelationship(ctx, keeper, msg)
		case types.MsgDeleteRelationships:
			return handleMsgDeleteRelationships(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized Posts message type: %v", msg.Type())
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

///////////////////////
/////Relationships////
/////////////////////

func handleMsgCreateMonoDirectionalRelationship(ctx sdk.Context, keeper Keeper, msg types.MsgCreateMonoDirectionalRelationship) (*sdk.Result, error) {
	relationship := types.NewMonodirectionalRelationship(msg.Sender, msg.Receiver)

	// Check if the relationship exist
	if keeper.DoesRelationshipExist(ctx, relationship.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Relationship with %s has already been made", msg.Receiver))
	}

	// Save the relationship
	keeper.StoreRelationship(ctx, relationship)

	// Save user/relationship association
	keeper.SaveUserRelationshipAssociation(ctx, relationship.Sender, relationship.Id)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeMonodirectionalRelationshipCreated,
		sdk.NewAttribute(types.AttributeRelationshipSender, relationship.Sender.String()),
		sdk.NewAttribute(types.AttributeRelationshipReceiver, relationship.Receiver.String()),
	))

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(relationship.Sender),
		Events: ctx.EventManager().Events(),
	}

	return &result, nil

}

func handleMsgRequestBiDirectionalRelationship(ctx sdk.Context, keeper Keeper, msg types.MsgRequestBidirectionalRelationship) (*sdk.Result, error) {
	relationship := types.NewBiDirectionalRelationship(msg.Sender, msg.Receiver, types.Sent)

	// Check if the relationship exist
	if keeper.DoesRelationshipExist(ctx, relationship.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Relationship request to %s has already been made", msg.Receiver))
	}

	// Save the relationship
	keeper.StoreRelationship(ctx, relationship)

	// Save users/relationship association
	keeper.SaveUserRelationshipAssociation(ctx, msg.Sender, relationship.Id)
	keeper.SaveUserRelationshipAssociation(ctx, msg.Receiver, relationship.Id)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeBidirectionalRelationshipRequested,
		sdk.NewAttribute(types.AttributeRelationshipSender, relationship.Sender.String()),
		sdk.NewAttribute(types.AttributeRelationshipReceiver, relationship.Receiver.String()),
		sdk.NewAttribute(types.AttributeRelationshipStatus, relationship.Status.String()),
	))

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(relationship.Sender),
		Events: ctx.EventManager().Events(),
	}

	return &result, nil
}

func handleMsgAcceptBidirectionalRelationship(ctx sdk.Context, keeper Keeper, msg types.MsgAcceptBidirectionalRelationship) (*sdk.Result, error) {
	relationships := keeper.GetUserRelationships(ctx, msg.Receiver)

	for _, relationship := range relationships {
		if relationship.ID() == msg.Id {
			if rel, ok := relationship.(types.BidirectionalRelationship); ok {
				if rel.Status == types.Accepted {
					return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("The relationship with id: %s has already been accepted", rel.Id))
				}
				rel.Status = types.Accepted
				keeper.StoreRelationship(ctx, rel)
				break
			} else {
				return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("The relationship with id: %s is not a bidirectional relationship and cannot be accepted", rel.Id))
			}
		}
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeBidirectionalRelationshipAccepted,
		sdk.NewAttribute(types.AttributeRelationshipID, msg.Id.String()),
		sdk.NewAttribute(types.AttributeRelationshipReceiver, msg.Receiver.String()),
		sdk.NewAttribute(types.AttributeRelationshipStatus, types.Accepted.String()),
	))

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(msg.Receiver),
		Events: ctx.EventManager().Events(),
	}

	return &result, nil
}

func handleMsgDenyBidirectionalRelationship(ctx sdk.Context, keeper Keeper, msg types.MsgDenyBidirectionalRelationship) (*sdk.Result, error) {
	relationships := keeper.GetUserRelationships(ctx, msg.Receiver)

	for _, relationship := range relationships {
		if relationship.ID() == msg.Id {
			if rel, ok := relationship.(types.BidirectionalRelationship); ok {
				if rel.Status == types.Accepted {
					return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("The relationship with id: %s has already been accepted and cannot be denied now", rel.Id))
				}
				rel.Status = types.Denied
				keeper.StoreRelationship(ctx, rel)
				break
			} else {
				return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("The relationship with id: %s is not a bidirectional relationship and cannot be accepted", rel.Id))
			}
		}
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeBidirectionalRelationshipAccepted,
		sdk.NewAttribute(types.AttributeRelationshipID, msg.Id.String()),
		sdk.NewAttribute(types.AttributeRelationshipReceiver, msg.Receiver.String()),
		sdk.NewAttribute(types.AttributeRelationshipStatus, types.Denied.String()),
	))

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(msg.Receiver),
		Events: ctx.EventManager().Events(),
	}

	return &result, nil
}

func handleMsgDeleteRelationships(ctx sdk.Context, keeper Keeper, msg types.MsgDeleteRelationships) (*sdk.Result, error) {
	if err := keeper.DeleteRelationship(ctx, msg.User, msg.Counterparty); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeRelationshipsDeleted,
		sdk.NewAttribute(types.AttributeRelationshipSender, msg.User.String()),
		sdk.NewAttribute(types.AttributeRelationshipReceiver, msg.Counterparty.String()),
	))

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(msg.User),
		Events: ctx.EventManager().Events(),
	}

	return &result, nil
}
