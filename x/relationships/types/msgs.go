package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/commons"
)

func NewMsgCreateRelationship(sender, receiver string, subspace string) *MsgCreateRelationship {
	return &MsgCreateRelationship{
		Sender:   sender,
		Receiver: receiver,
		Subspace: subspace,
	}
}

// Route should return the name of the module
func (msg MsgCreateRelationship) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateRelationship) Type() string {
	return ActionCreateRelationship
}

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateRelationship) ValidateBasic() error {
	if len(msg.Sender) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid sender address: %s", msg.Sender))
	}

	if len(msg.Receiver) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid receiver address: %s", msg.Receiver))
	}

	if msg.Sender == msg.Receiver {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender and receiver must be different")
	}

	if !commons.IsValidSubspace(msg.Subspace) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "subspace must be a sha-256")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateRelationship) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateRelationship) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{sender}
}

// ___________________________________________________________________________________________________________________

func NewMsgDeleteRelationship(sender, receiver string, subspace string) *MsgDeleteRelationship {
	return &MsgDeleteRelationship{
		Sender:       sender,
		Counterparty: receiver,
		Subspace:     subspace,
	}
}

// Route should return the name of the module
func (msg MsgDeleteRelationship) Route() string { return RouterKey }

// Type should return the action
func (msg MsgDeleteRelationship) Type() string {
	return ActionDeleteRelationship
}

// ValidateBasic runs stateless checks on the message
func (msg MsgDeleteRelationship) ValidateBasic() error {
	if len(msg.Sender) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid sender address: %s", msg.Sender))
	}

	if len(msg.Counterparty) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid counterparty address: %s", msg.Counterparty))
	}

	if msg.Sender == msg.Counterparty {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender and receiver must be different")
	}

	if !commons.IsValidSubspace(msg.Subspace) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "subspace must be a sha-256")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDeleteRelationship) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgDeleteRelationship) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{sender}
}

// ___________________________________________________________________________________________________________________

func NewMsgBlockUser(blocker, blocked string, reason, subspace string) *MsgBlockUser {
	return &MsgBlockUser{
		Blocker:  blocker,
		Blocked:  blocked,
		Reason:   reason,
		Subspace: subspace,
	}
}

// Route should return the name of the module
func (msg MsgBlockUser) Route() string { return RouterKey }

// Type should return the action
func (msg MsgBlockUser) Type() string {
	return ActionBlockUser
}

// ValidateBasic runs stateless checks on the message
func (msg MsgBlockUser) ValidateBasic() error {
	if len(msg.Blocker) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid blocker address: %s", msg.Blocker))
	}

	if len(msg.Blocked) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid blocked address: %s", msg.Blocked))
	}

	if msg.Blocker == msg.Blocked {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "blocker and blocked must be different")
	}

	if !commons.IsValidSubspace(msg.Subspace) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "subspace must be a valid sha-256 hash")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgBlockUser) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgBlockUser) GetSigners() []sdk.AccAddress {
	blocker, _ := sdk.AccAddressFromBech32(msg.Blocker)
	return []sdk.AccAddress{blocker}
}

// ___________________________________________________________________________________________________________________

func NewMsgUnblockUser(blocker, blocked string, subspace string) *MsgUnblockUser {
	return &MsgUnblockUser{
		Blocker:  blocker,
		Blocked:  blocked,
		Subspace: subspace,
	}
}

// Route should return the name of the module
func (msg MsgUnblockUser) Route() string { return RouterKey }

// Type should return the action
func (msg MsgUnblockUser) Type() string {
	return ActionUnblockUser
}

// ValidateBasic runs stateless checks on the message
func (msg MsgUnblockUser) ValidateBasic() error {
	if len(msg.Blocker) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid blocker address: %s", msg.Blocker))
	}

	if len(msg.Blocked) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid blocked address: %s", msg.Blocked))
	}

	if msg.Blocker == msg.Blocked {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "blocker and blocked must be different")
	}

	if !commons.IsValidSubspace(msg.Subspace) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "subspace must be a valid sha-256 hash")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgUnblockUser) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgUnblockUser) GetSigners() []sdk.AccAddress {
	blocker, _ := sdk.AccAddressFromBech32(msg.Blocker)
	return []sdk.AccAddress{blocker}
}
