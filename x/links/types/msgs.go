package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewMsgIBCAccountConnection(
	port string,
	channelId string,
	timeoutTimestamp uint64,
	sourceAddress string,
	sourcePubKey string,
	destinationAddress string,
	sourceSignature string,
	destinationSignature string,
) *MsgIBCAccountConnection {
	return &MsgIBCAccountConnection{
		Port:                 port,
		ChannelId:            channelId,
		TimeoutTimestamp:     timeoutTimestamp,
		SourceAddress:        sourceAddress,
		SourcePubKey:         sourcePubKey,
		DestinationAddress:   destinationAddress,
		SourceSignature:      sourceSignature,
		DestinationSignature: destinationSignature,
	}
}

// Route should return the name of the module
func (msg MsgIBCAccountConnection) Route() string { return RouterKey }

// Type should return the action
func (msg MsgIBCAccountConnection) Type() string { return ActionIBCAccountConnection }

// ValidateBasic runs stateless checks on the message
func (msg *MsgIBCAccountConnection) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.SourceAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid source address (%s)", err)
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgIBCAccountConnection) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners defines whose signature is required
func (msg *MsgIBCAccountConnection) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.SourceAddress)
	return []sdk.AccAddress{addr}
}

// MarshalJSON implements the json.Mashaler interface.
// This is done due to the fact that Amino does not respect omitempty clauses
func (msg MsgIBCAccountConnection) MarshalJSON() ([]byte, error) {
	type temp MsgIBCAccountConnection
	return json.Marshal(temp(msg))
}
