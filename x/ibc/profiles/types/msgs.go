package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
)

func NewMsgCreateIBCAccountConnection(
	port string,
	channelID string,
	packet IBCAccountConnectionPacketData,
	timeoutTimestamp uint64,
) *MsgCreateIBCAccountConnection {
	return &MsgCreateIBCAccountConnection{
		Port:             port,
		ChannelID:        channelID,
		Packet:           packet,
		TimeoutTimestamp: timeoutTimestamp,
	}
}

// Route should return the name of the module
func (msg MsgCreateIBCAccountConnection) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateIBCAccountConnection) Type() string { return ActionIBCAccountConnection }

// ValidateBasic runs stateless checks on the message
func (msg *MsgCreateIBCAccountConnection) ValidateBasic() error {
	if err := host.PortIdentifierValidator(msg.Port); err != nil {
		return sdkerrors.Wrap(err, "invalid source port ID")
	}
	if err := host.ChannelIdentifierValidator(msg.ChannelID); err != nil {
		return sdkerrors.Wrap(err, "invalid source channel ID")
	}
	if err := msg.Packet.Validate(); err != nil {
		return sdkerrors.Wrap(err, "invalid packet data")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgCreateIBCAccountConnection) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg *MsgCreateIBCAccountConnection) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.Packet.SourceAddress)
	return []sdk.AccAddress{sender}
}

// MarshalJSON implements the json.Mashaler interface.
// This is done due to the fact that Amino does not respect omitempty clauses
func (msg MsgCreateIBCAccountConnection) MarshalJSON() ([]byte, error) {
	type temp MsgCreateIBCAccountConnection
	return json.Marshal(temp(msg))
}

// ___________________________________________________________________________________________________________________

func NewMsgCreateIBCAccountLink(
	port string,
	channelID string,
	packet IBCAccountLinkPacketData,
	timeoutTimestamp uint64,
) *MsgCreateIBCAccountLink {
	return &MsgCreateIBCAccountLink{
		Port:             port,
		ChannelID:        channelID,
		Packet:           packet,
		TimeoutTimestamp: timeoutTimestamp,
	}
}

// Route should return the name of the module
func (msg MsgCreateIBCAccountLink) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateIBCAccountLink) Type() string { return ActionIBCAccountLink }

// ValidateBasic runs stateless checks on the message
func (msg *MsgCreateIBCAccountLink) ValidateBasic() error {
	if err := host.PortIdentifierValidator(msg.Port); err != nil {
		return sdkerrors.Wrap(err, "invalid source port ID")
	}
	if err := host.ChannelIdentifierValidator(msg.ChannelID); err != nil {
		return sdkerrors.Wrap(err, "invalid source channel ID")
	}
	if err := msg.Packet.Validate(); err != nil {
		return sdkerrors.Wrap(err, "invalid packet data")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgCreateIBCAccountLink) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg *MsgCreateIBCAccountLink) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.Packet.SourceAddress)
	return []sdk.AccAddress{sender}
}
