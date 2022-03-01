package profiles

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"

	"github.com/desmos-labs/desmos/v2/x/profiles/keeper"
	"github.com/desmos-labs/desmos/v2/x/profiles/types"
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

		case *types.MsgAcceptDTagTransferRequest:
			res, err := msgServer.AcceptDTagTransferRequest(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgRefuseDTagTransferRequest:
			res, err := msgServer.RefuseDTagTransferRequest(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgCancelDTagTransferRequest:
			res, err := msgServer.CancelDTagTransferRequest(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgLinkChainAccount:
			res, err := msgServer.LinkChainAccount(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgUnlinkChainAccount:
			res, err := msgServer.UnlinkChainAccount(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgLinkApplication:
			res, err := msgServer.LinkApplication(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgUnlinkApplication:
			res, err := msgServer.UnlinkApplication(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
				"unrecognized %s message type: %v", types.ModuleName, proto.MessageName(msg))
		}
	}
}
