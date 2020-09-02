package msgs

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/commons"
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

// MsgBlockUser allows the given Blocker to block the specified Blocked user
// for the (optional) reason.
type MsgBlockUser struct {
	Blocker  sdk.AccAddress `json:"blocker" yaml:"blocker"`
	Blocked  sdk.AccAddress `json:"blocked" yaml:"blocked"`
	Reason   string         `json:"reason,omitempty" yaml:"reason,omitempty"`
	Subspace string         `json:"subspace" yaml:"subspace"`
}

func NewMsgBlockUser(blocker, blocked sdk.AccAddress, reason, subspace string) MsgBlockUser {
	return MsgBlockUser{
		Blocker:  blocker,
		Blocked:  blocked,
		Reason:   reason,
		Subspace: subspace,
	}
}

// Route should return the name of the module
func (msg MsgBlockUser) Route() string { return models.RouterKey }

// Type should return the action
func (msg MsgBlockUser) Type() string {
	return models.ActionBlockUser
}

// ValidateBasic runs stateless checks on the message
func (msg MsgBlockUser) ValidateBasic() error {
	if msg.Blocker.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid blocker address: %s", msg.Blocker))
	}

	if msg.Blocked.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid blocked address: %s", msg.Blocked))
	}

	if msg.Blocker.Equals(msg.Blocked) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "blocker and blocked must be different")
	}

	if !commons.IsValidSubspace(msg.Subspace) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "subspace must be a valid sha-256 hash")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgBlockUser) GetSignBytes() []byte {
	return sdk.MustSortJSON(MsgsCodec.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgBlockUser) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Blocker}
}

// MsgUnblockUser allows the given original Blocker to unblock the specified Blocked user.
type MsgUnblockUser struct {
	Blocker  sdk.AccAddress `json:"blocker" yaml:"blocker"`
	Blocked  sdk.AccAddress `json:"blocked" yaml:"blocked"`
	Subspace string         `json:"subspace" yaml:"subspace"`
}

func NewMsgUnblockUser(blocker, blocked sdk.AccAddress, subspace string) MsgUnblockUser {
	return MsgUnblockUser{
		Blocker:  blocker,
		Blocked:  blocked,
		Subspace: subspace,
	}
}

// Route should return the name of the module
func (msg MsgUnblockUser) Route() string { return models.RouterKey }

// Type should return the action
func (msg MsgUnblockUser) Type() string {
	return models.ActionUnblockUser
}

// ValidateBasic runs stateless checks on the message
func (msg MsgUnblockUser) ValidateBasic() error {
	if msg.Blocker.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid blocker address: %s", msg.Blocker))
	}

	if msg.Blocked.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid blocked address: %s", msg.Blocked))
	}

	if msg.Blocker.Equals(msg.Blocked) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "blocker and blocked must be different")
	}

	if !commons.IsValidSubspace(msg.Subspace) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "subspace must be a valid sha-256 hash")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgUnblockUser) GetSignBytes() []byte {
	return sdk.MustSortJSON(MsgsCodec.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgUnblockUser) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Blocker}
}
