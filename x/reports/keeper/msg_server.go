package keeper

import (
	"context"
	"fmt"
	"time"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/desmos-labs/desmos/v6/x/reports/types"
	subspacestypes "github.com/desmos-labs/desmos/v6/x/subspaces/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the stored MsgServer interface for the provided keeper
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = &msgServer{}

// CreateReport defines the rpc method for Msg/CreateReport
func (k msgServer) CreateReport(goCtx context.Context, msg *types.MsgCreateReport) (*types.MsgCreateReportResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the reporter has a profile
	if !k.HasProfile(ctx, msg.Reporter) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "you cannot create a report without having a profile")
	}

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check if the reasons exist
	for _, reasonID := range msg.ReasonsIDs {
		if !k.HasReason(ctx, msg.SubspaceID, reasonID) {
			return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "reason with id %d not found inside subspace %d", reasonID, msg.SubspaceID)
		}
	}

	// Check the permission to report
	if !k.HasPermission(ctx, msg.SubspaceID, msg.Reporter, types.PermissionReportContent) {
		return nil, errors.Wrap(subspacestypes.ErrPermissionDenied, "you cannot report content inside this subspace")
	}

	target, ok := msg.Target.GetCachedValue().(types.ReportTarget)
	if !ok {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid target type: %s", msg.Target)
	}

	// Make sure the report is not duplicated
	if k.HasReported(ctx, msg.SubspaceID, msg.Reporter, target) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "you have already reported this target")
	}

	// Get the next report id
	reportID, err := k.GetNextReportID(ctx, msg.SubspaceID)
	if err != nil {
		return nil, err
	}

	// Create and validate the report
	report := types.NewReport(
		msg.SubspaceID,
		reportID,
		msg.ReasonsIDs,
		msg.Message,
		target,
		msg.Reporter,
		ctx.BlockTime(),
	)
	err = k.ValidateReport(ctx, report)
	if err != nil {
		return nil, err
	}

	// Store the report
	k.SaveReport(ctx, report)

	// Update the id for the next report
	k.SetNextReportID(ctx, msg.SubspaceID, report.ID+1)

	// Get the reporting event (different based on the target)
	var reportEvent sdk.Event
	switch target := target.(type) {
	case *types.PostTarget:
		reportEvent = sdk.NewEvent(
			types.EventTypeReportedPost,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyPostID, fmt.Sprintf("%d", target.PostID)),
			sdk.NewAttribute(types.AttributeKeyReporter, msg.Reporter),
		)
	case *types.UserTarget:
		reportEvent = sdk.NewEvent(
			types.EventTypeReportedUser,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyUser, target.User),
			sdk.NewAttribute(types.AttributeKeyReporter, msg.Reporter),
		)
	default:
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid report target type: %T", msg.Target)
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreatedReport,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyReportID, fmt.Sprintf("%d", report.ID)),
			sdk.NewAttribute(types.AttributeKeyReporter, msg.Reporter),
			sdk.NewAttribute(types.AttributeKeyCreationTime, report.CreationDate.Format(time.RFC3339)),
		),
		reportEvent,
	})

	return &types.MsgCreateReportResponse{
		ReportID:     report.ID,
		CreationDate: report.CreationDate,
	}, nil
}

// DeleteReport defines the rpc method for Msg/DeleteReport
func (k msgServer) DeleteReport(goCtx context.Context, msg *types.MsgDeleteReport) (*types.MsgDeleteReportResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check if the report exists
	report, found := k.GetReport(ctx, msg.SubspaceID, msg.ReportID)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "report with id %d not found inside subspace %d", msg.ReportID, msg.SubspaceID)
	}

	// Check the permission to delete reports
	isModerator := k.HasPermission(ctx, msg.SubspaceID, msg.Signer, types.PermissionManageReports)
	canDelete := report.Reporter == msg.Signer && k.HasPermission(ctx, msg.SubspaceID, msg.Signer, types.PermissionDeleteOwnReports)
	if !isModerator && !canDelete {
		return nil, errors.Wrap(subspacestypes.ErrPermissionDenied, "you cannot delete reports inside this subspace")
	}

	// Delete the report
	k.Keeper.DeleteReport(ctx, msg.SubspaceID, msg.ReportID)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeDeletedReport,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyReportID, fmt.Sprintf("%d", msg.ReportID)),
		),
	})

	return &types.MsgDeleteReportResponse{}, nil
}

// SupportStandardReason defines the rpc method for Msg/SupportStandardReason
func (k msgServer) SupportStandardReason(goCtx context.Context, msg *types.MsgSupportStandardReason) (*types.MsgSupportStandardReasonResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check if the standard reason exists
	standardReason, found := k.GetStandardReason(ctx, msg.StandardReasonID)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "standard reason with id %d could not be found", msg.StandardReasonID)
	}

	// Check the permission to manage reasons
	if !k.HasPermission(ctx, msg.SubspaceID, msg.Signer, types.PermissionManageReasons) {
		return nil, errors.Wrap(subspacestypes.ErrPermissionDenied, "you cannot manage reasons inside this subspace")
	}

	// Get the next reason id for the subspace
	reasonID, err := k.GetNextReasonID(ctx, msg.SubspaceID)
	if err != nil {
		return nil, err
	}

	// Create the reason and validate it
	reason := types.NewReason(msg.SubspaceID, reasonID, standardReason.Title, standardReason.Description)
	err = reason.Validate()
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Store the reason
	k.SaveReason(ctx, reason)

	// Update the id for the next reason
	k.SetNextReasonID(ctx, msg.SubspaceID, reason.ID+1)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSupportedStandardReason,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyStandardReasonID, fmt.Sprintf("%d", msg.StandardReasonID)),
			sdk.NewAttribute(types.AttributeKeyReasonID, fmt.Sprintf("%d", reason.ID)),
		),
	})

	return &types.MsgSupportStandardReasonResponse{
		ReasonsID: reason.ID,
	}, nil
}

// AddReason defines the rpc method for Msg/AddReason
func (k msgServer) AddReason(goCtx context.Context, msg *types.MsgAddReason) (*types.MsgAddReasonResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check the permission to manage reasons
	if !k.HasPermission(ctx, msg.SubspaceID, msg.Signer, types.PermissionManageReasons) {
		return nil, errors.Wrap(subspacestypes.ErrPermissionDenied, "you cannot manage reasons inside this subspace")
	}

	// Get the next reason id for the subspace
	reasonID, err := k.GetNextReasonID(ctx, msg.SubspaceID)
	if err != nil {
		return nil, err
	}

	// Create the reason and validate it
	reason := types.NewReason(msg.SubspaceID, reasonID, msg.Title, msg.Description)
	err = reason.Validate()
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Store the reason
	k.SaveReason(ctx, reason)

	// Update the id for the next reason
	k.SetNextReasonID(ctx, msg.SubspaceID, reason.ID+1)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeAddedReportingReason,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyReasonID, fmt.Sprintf("%d", reason.ID)),
		),
	})

	return &types.MsgAddReasonResponse{
		ReasonID: reason.ID,
	}, nil
}

// RemoveReason defines the rpc method for Msg/RemoveReason
func (k msgServer) RemoveReason(goCtx context.Context, msg *types.MsgRemoveReason) (*types.MsgRemoveReasonResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check if the reason exists
	if !k.HasReason(ctx, msg.SubspaceID, msg.ReasonID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "reason with id %d does not existing inside subspace %d", msg.ReasonID, msg.SubspaceID)
	}

	// Check the permission to manage reasons
	if !k.HasPermission(ctx, msg.SubspaceID, msg.Signer, types.PermissionManageReasons) {
		return nil, errors.Wrap(subspacestypes.ErrPermissionDenied, "you cannot manage reasons inside this subspace")
	}

	// Delete the reason
	k.DeleteReason(ctx, msg.SubspaceID, msg.ReasonID)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeRemovedReportingReason,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyReasonID, fmt.Sprintf("%d", msg.ReasonID)),
		),
	})

	return &types.MsgRemoveReasonResponse{}, nil
}

// UpdateParams updates the module parameters
func (m msgServer) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	authority := m.authority
	if authority != msg.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", authority, msg.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	m.SetParams(ctx, msg.Params)

	return &types.MsgUpdateParamsResponse{}, nil
}
