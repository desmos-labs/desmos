package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v3/x/reports/types"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
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

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check if the reason exists
	if !k.HasReason(ctx, msg.SubspaceID, msg.ReasonID) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "reason with id %d not found inside subspace %d", msg.ReasonID, msg.SubspaceID)
	}

	reporter, err := sdk.AccAddressFromBech32(msg.Reporter)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid reporter address: %s", msg.Reporter)
	}

	// Check the permission to report
	if !k.HasPermission(ctx, msg.SubspaceID, reporter, subspacestypes.PermissionReportContent) {
		return nil, sdkerrors.Wrap(subspacestypes.ErrPermissionDenied, "you cannot report content inside this subspace")
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
		msg.ReasonID,
		msg.Message,
		msg.Data.GetCachedValue().(types.ReportData),
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

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Reporter),
		),
		sdk.NewEvent(
			types.EventTypeCreateReport,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyReportID, fmt.Sprintf("%d", report.ID)),
			sdk.NewAttribute(types.AttributeKeyReporter, msg.Reporter),
			sdk.NewAttribute(types.AttributeKeyCreationTime, report.CreationDate.Format(time.RFC3339)),
		),
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
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check if the report exists
	report, found := k.GetReport(ctx, msg.SubspaceID, msg.ReportID)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "report with id %d not found inside subspace %d", msg.ReportID, msg.SubspaceID)
	}

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address: %s", msg.Signer)
	}

	// Check the permission to delete reports
	isModerator := k.HasPermission(ctx, msg.SubspaceID, signer, subspacestypes.PermissionManageReports)
	canDelete := report.Reporter == msg.Signer && k.HasPermission(ctx, msg.SubspaceID, signer, subspacestypes.PermissionDeleteOwnReports)
	if !isModerator && !canDelete {
		return nil, sdkerrors.Wrap(subspacestypes.ErrPermissionDenied, "you cannot delete reports inside this subspace")
	}

	// Delete the post
	k.Keeper.DeleteReport(ctx, msg.SubspaceID, msg.ReportID)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer),
		),
		sdk.NewEvent(
			types.EventTypeDeleteReport,
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
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check if the standard reason exists
	standardReason, found := k.GetStandardReason(ctx, msg.StandardReasonID)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "standard reason with id %d could not be found", msg.StandardReasonID)
	}

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address: %s", msg.Signer)
	}

	// Check the permission to manage reasons
	if !k.HasPermission(ctx, msg.SubspaceID, signer, subspacestypes.PermissionManageReasons) {
		return nil, sdkerrors.Wrap(subspacestypes.ErrPermissionDenied, "you cannot manage reasons inside this subspace")
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
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Store the reason
	k.SaveReason(ctx, reason)

	// Update the id for the next reason
	k.SetNextReasonID(ctx, msg.SubspaceID, reason.ID+1)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer),
		),
		sdk.NewEvent(
			types.EventTypeSupportStandardReason,
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
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address: %s", msg.Signer)
	}

	// Check the permission to manage reasons
	if !k.HasPermission(ctx, msg.SubspaceID, signer, subspacestypes.PermissionManageReasons) {
		return nil, sdkerrors.Wrap(subspacestypes.ErrPermissionDenied, "you cannot manage reasons inside this subspace")
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
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Store the reason
	k.SaveReason(ctx, reason)

	// Update the id for the next reason
	k.SetNextReasonID(ctx, msg.SubspaceID, reason.ID+1)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer),
		),
		sdk.NewEvent(
			types.EventTypeAddReason,
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
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check if the reason exists
	if !k.HasReason(ctx, msg.SubspaceID, msg.ReasonID) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "reason with id %d does not existing inside subspace %d", msg.ReasonID, msg.SubspaceID)
	}

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address: %s", msg.Signer)
	}

	// Check the permission to manage reasons
	if !k.HasPermission(ctx, msg.SubspaceID, signer, subspacestypes.PermissionManageReasons) {
		return nil, sdkerrors.Wrap(subspacestypes.ErrPermissionDenied, "you cannot manage reasons inside this subspace")
	}

	// Delete the reason
	k.DeleteReason(ctx, msg.SubspaceID, msg.ReasonID)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer),
		),
		sdk.NewEvent(
			types.EventTypeRemoveReason,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyReasonID, fmt.Sprintf("%d", msg.ReasonID)),
		),
	})

	return &types.MsgRemoveReasonResponse{}, nil
}
