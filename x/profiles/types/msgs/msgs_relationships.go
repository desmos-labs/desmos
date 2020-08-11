package msgs

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/profiles/types"
	"github.com/desmos-labs/desmos/x/profiles/types/models"
)

// Creates a monodirectional relationship between the sender and
// the receiver. Monodirectional relationship do not need the receiver to accept in order to be valid.
// An example of monodirectional relationship is the follow on Twitter or the subscribe on YouTube.
type MsgCreateMonoDirectionalRelationship struct {
	Sender   sdk.AccAddress `json:"sender" yaml:"sender"`
	Receiver sdk.AccAddress `json:"receiver" yaml:"receiver"`
}

func NewMsgCreateMonoDirectionalRelationship(sender, receiver sdk.AccAddress) MsgCreateMonoDirectionalRelationship {
	return MsgCreateMonoDirectionalRelationship{
		Sender:   sender,
		Receiver: receiver,
	}
}

// Route should return the name of the module
func (msg MsgCreateMonoDirectionalRelationship) Route() string { return models.RouterKey }

// Type should return the action
func (msg MsgCreateMonoDirectionalRelationship) Type() string {
	return models.ActionCreateMonoDirectionalRelationship
}

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateMonoDirectionalRelationship) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid sender address: %s", msg.Sender))
	}

	if msg.Receiver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid receiver address: %s", msg.Receiver))
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateMonoDirectionalRelationship) GetSignBytes() []byte {
	return sdk.MustSortJSON(MsgsCodec.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateMonoDirectionalRelationship) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// MsgRequestBidirectionalRelationship allows the specified Sender to ask for a bidirectional
// relationship to the specified Receiver, attaching the given (optional) message to the request.
// In order to accept, the Receiver must be user either MsgAcceptBidirectionalRelationship or
// MsgDenyBidirectionalRelationship.
type MsgRequestBidirectionalRelationship struct {
	Sender   sdk.AccAddress `json:"sender" yaml:"sender"`
	Receiver sdk.AccAddress `json:"receiver" yaml:"receiver"`
	Message  string         `json:"message,omitempty" yaml:"message,omitempty"`
}

func NewMsgRequestBidirectionalRelationship(sender, receiver sdk.AccAddress, message string) MsgRequestBidirectionalRelationship {
	return MsgRequestBidirectionalRelationship{
		Sender:   sender,
		Receiver: receiver,
		Message:  message,
	}
}

// Route should return the name of the module
func (msg MsgRequestBidirectionalRelationship) Route() string { return models.RouterKey }

// Type should return the action
func (msg MsgRequestBidirectionalRelationship) Type() string {
	return models.ActionRequestBiDirectionalRelationship
}

// ValidateBasic runs stateless checks on the message
func (msg MsgRequestBidirectionalRelationship) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid sender address: %s", msg.Sender))
	}

	if msg.Receiver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid receiver address: %s", msg.Receiver))
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRequestBidirectionalRelationship) GetSignBytes() []byte {
	return sdk.MustSortJSON(MsgsCodec.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgRequestBidirectionalRelationship) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// MsgAcceptBidirectionalRelationship allows the receiver of a bidirectional relationship request
// to accept that request leading to the effective creation of the relationship itself.
type MsgAcceptBidirectionalRelationship struct {
	Id       types.RelationshipID `json:"id" yaml:"id"`
	Receiver sdk.AccAddress       `json:"receiver" yaml:"receiver"`
}

func NewMsgAcceptBidirectionalRelationship(id types.RelationshipID, receiver sdk.AccAddress) MsgAcceptBidirectionalRelationship {
	return MsgAcceptBidirectionalRelationship{
		Id:       id,
		Receiver: receiver,
	}
}

// Route should return the name of the module
func (msg MsgAcceptBidirectionalRelationship) Route() string { return models.RouterKey }

// Type should return the action
func (msg MsgAcceptBidirectionalRelationship) Type() string {
	return models.ActionAcceptBiDirectionalRelationship
}

// ValidateBasic runs stateless checks on the message
func (msg MsgAcceptBidirectionalRelationship) ValidateBasic() error {
	if msg.Id.Valid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid relationship's id: %s", msg.Id))
	}

	if msg.Receiver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid receiver address: %s", msg.Receiver))
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgAcceptBidirectionalRelationship) GetSignBytes() []byte {
	return sdk.MustSortJSON(MsgsCodec.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgAcceptBidirectionalRelationship) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Receiver}
}

// MsgDenyBidirectionalRelationship allows the receiver of a bidirectional relationship request
// to deny that request.
type MsgDenyBidirectionalRelationship struct {
	Id       types.RelationshipID `json:"id" yaml:"id"`
	Receiver sdk.AccAddress       `json:"receiver" yaml:"receiver"`
}

func NewMsgDenyBidirectionalRelationship(id types.RelationshipID, receiver sdk.AccAddress) MsgDenyBidirectionalRelationship {
	return MsgDenyBidirectionalRelationship{
		Id:       id,
		Receiver: receiver,
	}
}

// Route should return the name of the module
func (msg MsgDenyBidirectionalRelationship) Route() string { return models.RouterKey }

// Type should return the action
func (msg MsgDenyBidirectionalRelationship) Type() string {
	return models.ActionDenyBiDirectionalRelationship
}

// ValidateBasic runs stateless checks on the message
func (msg MsgDenyBidirectionalRelationship) ValidateBasic() error {
	if msg.Id.Valid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid relationship's id: %s", msg.Id))
	}

	if msg.Receiver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid receiver address: %s", msg.Receiver))
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDenyBidirectionalRelationship) GetSignBytes() []byte {
	return sdk.MustSortJSON(MsgsCodec.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgDenyBidirectionalRelationship) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Receiver}
}

// MsgDeleteRelationships allows the specified User to cut off the relationship he had prevously
// created with the specified Counterparty.
// If the relationship was a monodirectional relationship, the user must be the original Sender of
// the relationship, otherwise, if it was a bidirectional one, it can be either one of the two users
// taking part to it.
type MsgDeleteRelationships struct {
	User         sdk.AccAddress `json:"user" yaml:"user"`
	Counterparty sdk.AccAddress `json:"counterparty" yaml:"counterparty"`
}

func NewMsgDeleteRelationships(user, counterpart sdk.AccAddress) MsgDeleteRelationships {
	return MsgDeleteRelationships{
		User:         user,
		Counterparty: counterpart,
	}
}

// Route should return the name of the module
func (msg MsgDeleteRelationships) Route() string { return models.RouterKey }

// Type should return the action
func (msg MsgDeleteRelationships) Type() string {
	return models.ActionDeleteRelationships
}

// ValidateBasic runs stateless checks on the message
func (msg MsgDeleteRelationships) ValidateBasic() error {
	if msg.User.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid user address: %s", msg.User))
	}

	if msg.Counterparty.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid counterparty address: %s", msg.Counterparty))
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDeleteRelationships) GetSignBytes() []byte {
	return sdk.MustSortJSON(MsgsCodec.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgDeleteRelationships) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.User}
}
