package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// OnRecvPacket processes packet reception
func (k Keeper) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	data types.LinkChainAccountPacketData,
) (packetAck types.LinkChainAccountPacketAck, err error) {
	// validate packet data upon receiving
	if err := data.Validate(); err != nil {
		return packetAck, err
	}

	srcAddrData, err := types.UnpackAddressData(k.cdc, data.SourceAddress)
	if err != nil {
		return packetAck, err
	}

	if err := data.DestinationProof.Verify(k.cdc); err != nil {
		return packetAck, err
	}

	chainLink := types.NewChainLink(
		srcAddrData,
		data.SourceProof,
		data.SourceChainConfig,
		ctx.BlockTime(),
	)

	if err := k.StoreChainLink(ctx, data.DestinationAddress, chainLink); err != nil {
		return packetAck, err
	}

	packetAck.SourceAddress = srcAddrData.GetAddress()
	return packetAck, nil
}
