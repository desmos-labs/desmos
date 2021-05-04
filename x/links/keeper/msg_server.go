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

func (k msgServer) CreateIBCAccountConnection(
	goCtx context.Context, msg *types.MsgCreateIBCAccountConnection,
) (*types.MsgCreateIBCAccountConnectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	prefix := sdk.GetConfig().GetBech32AccountAddrPrefix()

	// Construct the packet
	packet := types.NewIBCAccountConnectionPacketData(
		prefix,
		msg.SourceAddress,
		msg.SourcePubKey,
		msg.DestinationAddress,
		msg.SourceSignature,
		msg.DestinationSignature,
	)

	height := uint64(ctx.BlockHeight())

	// Transmit the packet
	if err := k.TransmitIBCAccountConnectionPacket(
		ctx,
		packet,
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
	prefix := sdk.GetConfig().GetBech32AccountAddrPrefix()

	// Construct the packet
	packet := types.NewIBCAccountLinkPacketData(
		prefix,
		msg.SourceAddress,
		msg.SourcePubKey,
		msg.Signature,
	)

	height := uint64(ctx.BlockHeight())

	// Transmit the packet
	if err := k.TransmitIBCAccountLinkPacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelId,
		clienttypes.NewHeight(height, height+100),
		msg.TimeoutTimestamp,
	); err != nil {
		return nil, err
	}

	return &types.MsgCreateIBCAccountLinkResponse{}, nil
}
