package keeper

import (
	"bytes"
	"encoding/hex"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// OnRecvLinkChainAccountPacket processes the reception of a LinkChainAccountPacket
func (k Keeper) OnRecvLinkChainAccountPacket(
	ctx sdk.Context,
	data types.LinkChainAccountPacketData,
) (packetAck types.LinkChainAccountPacketAck, err error) {
	// Validate the packet data upon receiving
	if err := data.Validate(); err != nil {
		return packetAck, err
	}

	srcAddrData, err := types.UnpackAddressData(k.cdc, data.SourceAddress)
	if err != nil {
		return packetAck, err
	}

	// Get the destination address and make sure it has a profile
	addr, err := sdk.AccAddressFromBech32(data.DestinationAddress)
	if err != nil {
		return packetAck, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, data.DestinationAddress)
	}

	account := k.ak.GetAccount(ctx, addr)
	profile, ok := account.(*types.Profile)
	if !ok {
		return packetAck, sdkerrors.Wrap(types.ErrProfileNotFound, addr.String())
	}

	// Get the destination proof public key
	var pubKey cryptotypes.PubKey
	err = k.cdc.UnpackAny(data.DestinationProof.PubKey, &pubKey)
	if err != nil {
		return packetAck, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "invalid public key type")
	}

	// Make sure the profile public key and the one provided are equals
	if !bytes.Equal(pubKey.Bytes(), profile.GetPubKey().Bytes()) {
		return packetAck, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid pub key value: expected %s but got %s instead",
			hex.EncodeToString(profile.GetPubKey().Bytes()), hex.EncodeToString(pubKey.Bytes()))
	}

	// Verify the destination proof
	destAddrData := types.NewBech32Address(data.DestinationAddress, sdk.GetConfig().GetBech32AccountAddrPrefix())
	err = data.DestinationProof.Verify(k.cdc, destAddrData)
	if err != nil {
		return packetAck, err
	}

	// Store the link
	chainLink := types.NewChainLink(data.DestinationAddress, srcAddrData, data.SourceProof, data.SourceChainConfig, ctx.BlockTime())
	err = k.SaveChainLink(ctx, chainLink)
	if err != nil {
		return packetAck, err
	}

	packetAck.SourceAddress = srcAddrData.GetValue()
	return packetAck, nil
}
