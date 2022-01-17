package subspaces

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"

	keeper2 "github.com/desmos-labs/desmos/v2/x/subspaces/keeper"
	types2 "github.com/desmos-labs/desmos/v2/x/subspaces/types"
)

// NewHandler returns a handler for subspaces type messages
func NewHandler(k keeper2.Keeper) sdk.Handler {
	msgServer := keeper2.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types2.MsgCreateSubspace:
			res, err := msgServer.CreateSubspace(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types2.MsgEditSubspace:
			res, err := msgServer.EditSubspace(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types2.MsgAddAdmin:
			res, err := msgServer.AddAdmin(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types2.MsgRemoveAdmin:
			res, err := msgServer.RemoveAdmin(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types2.MsgRegisterUser:
			res, err := msgServer.RegisterUser(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types2.MsgUnregisterUser:
			res, err := msgServer.UnregisterUser(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types2.MsgBanUser:
			res, err := msgServer.BanUser(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types2.MsgUnbanUser:
			res, err := msgServer.UnbanUser(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
				"unrecognized %s message type: %v", types2.ModuleName, proto.MessageName(msg))
		}
	}
}
