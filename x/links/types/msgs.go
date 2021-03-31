package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewMsgIBCLink(
	sender string,
	port string,
	channelId string,
	timeoutTimestamp uint64,
	sourceAddress string,
	sourcePubkey string,
	destinationAddress string,
	sourceSignature string,
	destinationSignature string,
) *MsgIBCLink {
	return &MsgIBCLink{
		Sender:               sender,
		Port:                 port,
		ChannelId:            channelId,
		TimeoutTimestamp:     timeoutTimestamp,
		SourceAddress:        sourceAddress,
		SourcePubkey:         sourcePubkey,
		DestinationAddress:   destinationAddress,
		SourceSignature:      sourceSignature,
		DestinationSignature: destinationSignature,
	}
}

// Route should return the name of the module
func (msg MsgIBCLink) Route() string { return RouterKey }

// Type should return the action
func (msg MsgIBCLink) Type() string { return ActionIBCLink }

// ValidateBasic runs stateless checks on the message
func (msg *MsgIBCLink) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.SourceAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid source address (%s)", err)
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgIBCLink) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners defines whose signature is required
func (msg *MsgIBCLink) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.SourceAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// MarshalJSON implements the json.Mashaler interface.
// This is done due to the fact that Amino does not respect omitempty clauses
func (msg MsgIBCLink) MarshalJSON() ([]byte, error) {
	type temp MsgIBCLink
	return json.Marshal(temp(msg))
}
