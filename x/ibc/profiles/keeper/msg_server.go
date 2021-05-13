package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	"github.com/desmos-labs/desmos/x/ibc/profiles/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) CreateIBCAccountConnection(
	goCtx context.Context, msg *types.MsgCreateIBCAccountConnection,
) (*types.MsgCreateIBCAccountConnectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	height := uint64(ctx.BlockHeight())

	// Transmit the packet
	if err := k.TransmitIBCAccountConnectionPacket(
		ctx,
		msg.Packet,
		msg.Port,
		msg.ChannelId,
		clienttypes.NewHeight(height, height+100),
		msg.TimeoutTimestamp,
	); err != nil {
		return nil, err
	}

	return &types.MsgCreateIBCAccountConnectionResponse{}, nil
}

// ___________________________________________________________________________________________________________________

func (k msgServer) CreateIBCAccountLink(
	goCtx context.Context, msg *types.MsgCreateIBCAccountLink,
) (*types.MsgCreateIBCAccountLinkResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	height := uint64(ctx.BlockHeight())

	// Transmit the packet
	if err := k.TransmitIBCAccountLinkPacket(
		ctx,
		msg.Packet,
		msg.Port,
		msg.ChannelId,
		clienttypes.NewHeight(height, height+100),
		msg.TimeoutTimestamp,
	); err != nil {
		return nil, err
	}

	return &types.MsgCreateIBCAccountLinkResponse{}, nil
}
