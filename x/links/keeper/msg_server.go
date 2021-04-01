package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	"github.com/desmos-labs/desmos/x/links/types"
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

func (k msgServer) IBCAccountConnection(goCtx context.Context, msg *types.MsgIBCAccountConnection) (*types.MsgIBCAccountConnectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Construct the packet
	var packet types.IBCAccountConnectionPacketData

	packet.SourceChainPrefix = msg.SourceChainPrefix
	packet.SourceAddress = msg.SourceAddress
	packet.SourcePubkey = msg.SourcePubkey
	packet.DestinationAddress = msg.DestinationAddress
	packet.SourceSignature = msg.SourceSignature
	packet.DestinationSignature = msg.DestinationSignature

	// Transmit the packet
	err := k.TransmitIBCAccountConnectionPacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelId,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgIBCAccountConnectionResponse{}, nil
}
