package keeper

import (
	"errors"
	"fmt"

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

	if err := data.SourceProof.Verify(k.cdc); err != nil {
		return packetAck, err
	}

	if err := data.DestinationProof.Verify(k.cdc); err != nil {
		return packetAck, err
	}

	// Check if address has the profile
	profile, found, err := k.GetProfile(ctx, data.DestinationAddress)
	if err != nil {
		return packetAck, err
	}
	if !found {
		return packetAck, fmt.Errorf("address does not have any profile")
	}

	chainLink := types.NewChainLink(
		data.SourceAddress,
		data.SourceProof,
		data.SourceChainConfig,
		ctx.BlockTime(),
	)

	if err := k.StoreChainLink(ctx, chainLink); err != nil {
		return packetAck, err
	}

	// Store chain link to the profile
	profile.ChainsLinks = append(profile.ChainsLinks, chainLink)
	if err := k.StoreProfile(ctx, profile); err != nil {
		k.DeleteChainLink(ctx, chainLink.ChainConfig.Name, chainLink.Address)
		return packetAck, err
	}

	packetAck.SourceAddress = data.SourceAddress
	return packetAck, nil
}

// OnAcknowledgementPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementPacket(ctx sdk.Context,
	packet channeltypes.Packet,
	data types.LinkChainAccountPacketData,
	ack channeltypes.Acknowledgement,
) error {
	switch ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:
		return nil
	case *channeltypes.Acknowledgement_Result:
		dispatchedAck := ack.Response.(*channeltypes.Acknowledgement_Result)
		var packetAck types.LinkChainAccountPacketAck
		err := packetAck.Unmarshal(dispatchedAck.Result)
		if err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}
		// the acknowledgement succeeded on the receiving chain so nothing
		// needs to be executed and no error needs to be returned
		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}
