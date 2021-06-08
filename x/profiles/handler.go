package profiles

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/profiles/keeper"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

// NewHandler returns a handler for "profile" type messages.
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgSaveProfile:
			res, err := msgServer.SaveProfile(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgDeleteProfile:
			res, err := msgServer.DeleteProfile(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgRequestDTagTransfer:
			res, err := msgServer.RequestDTagTransfer(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgAcceptDTagTransfer:
			res, err := msgServer.AcceptDTagTransfer(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgRefuseDTagTransfer:
			res, err := msgServer.RefuseDTagTransfer(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgCancelDTagTransfer:
			res, err := msgServer.CancelDTagTransfer(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgCreateRelationship:
			res, err := msgServer.CreateRelationship(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgDeleteRelationship:
			res, err := msgServer.DeleteRelationship(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgBlockUser:
			res, err := msgServer.BlockUser(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgUnblockUser:
			res, err := msgServer.UnblockUser(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgLinkChainAccount:
			res, err := msgServer.LinkChainAccount(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgUnlinkChainAccount:
			res, err := msgServer.UnlinkChainAccount(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
				"unrecognized %s message type: %v", types.ModuleName, msg.Type())
		}
	}
}
