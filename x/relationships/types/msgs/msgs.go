package msgs

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/relationships/types/models"
)

// Creates a relationship between the sender and
// the receiver.
// An example of relationship is the follow on Twitter or the subscribe on YouTube.
type MsgCreateRelationship struct {
	Sender   sdk.AccAddress `json:"sender" yaml:"sender"`
	Receiver sdk.AccAddress `json:"receiver" yaml:"receiver"`
}

func NewMsgCreateRelationship(sender, receiver sdk.AccAddress) MsgCreateRelationship {
	return MsgCreateRelationship{
		Sender:   sender,
		Receiver: receiver,
	}
}

// Route should return the name of the module
func (msg MsgCreateRelationship) Route() string { return models.RouterKey }

// Type should return the action
func (msg MsgCreateRelationship) Type() string {
	return models.ActionCreateRelationship
}

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateRelationship) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid sender address: %s", msg.Sender))
	}

	if msg.Receiver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid receiver address: %s", msg.Receiver))
	}

	if msg.Sender.Equals(msg.Receiver) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender and receiver must be different")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateRelationship) GetSignBytes() []byte {
	return sdk.MustSortJSON(MsgsCodec.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateRelationship) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// MsgDeleteRelationship allows the specified Sender to cut off the relationship he had previously
// created with the specified Counterparty.
type MsgDeleteRelationship struct {
	Sender       sdk.AccAddress `json:"sender" yaml:"sender"`
	Counterparty sdk.AccAddress `json:"counterparty" yaml:"counterparty"`
}

func NewMsgDeleteRelationship(sender, receiver sdk.AccAddress) MsgDeleteRelationship {
	return MsgDeleteRelationship{
		Sender:       sender,
		Counterparty: receiver,
	}
}

// Route should return the name of the module
func (msg MsgDeleteRelationship) Route() string { return models.RouterKey }

// Type should return the action
func (msg MsgDeleteRelationship) Type() string {
	return models.ActionDeleteRelationship
}

// ValidateBasic runs stateless checks on the message
func (msg MsgDeleteRelationship) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid sender address: %s", msg.Sender))
	}

	if msg.Counterparty.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid counterparty address: %s", msg.Counterparty))
	}

	if msg.Sender.Equals(msg.Counterparty) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender and receiver must be different")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDeleteRelationship) GetSignBytes() []byte {
	return sdk.MustSortJSON(MsgsCodec.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgDeleteRelationship) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}
