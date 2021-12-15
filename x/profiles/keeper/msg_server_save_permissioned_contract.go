package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/v2/x/profiles/types"
)

func (k msgServer) SavePermissionedContractReference(goCtx context.Context,
	msg *types.MsgSavePermissionedContractReference) (*types.MsgSavePermissionedContractReferenceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if k.DoesPermissionedContractExist(ctx, msg.Admin, msg.Address) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"Permissioned contract reference already stored by admin %s", msg.Admin)
	}

	k.SavePermissionedContract(ctx, types.NewPermissionedContract(msg.Admin, msg.Address))

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypePermissionedContractSaved,
		sdk.NewAttribute(types.AttributeContractAddress, msg.Address),
		sdk.NewAttribute(types.AttributeContractAdmin, msg.Admin),
	))

	return &types.MsgSavePermissionedContractReferenceResponse{}, nil
}
