package types

import (
	subspacestypes "github.com/desmos-labs/desmos/x/staging/subspaces/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ----------------------
// --- MsgSaveProfile
// ----------------------

func NewMsgCreateRelationship(creator, recipient, subspace string) *MsgCreateRelationship {
	return &MsgCreateRelationship{
		Sender:   creator,
		Receiver: recipient,
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
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address")
	}

	_, err = sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid receiver address")
	}

	if msg.Sender == msg.Receiver {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "sender and receiver must be different")
	}

	if !subspacestypes.IsValidSubspace(msg.Subspace) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "subspace must be a sha-256")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateRelationship) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateRelationship) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{sender}
}

// ___________________________________________________________________________________________________________________

func NewMsgDeleteRelationship(user, counterparty, subspace string) *MsgDeleteRelationship {
	return &MsgDeleteRelationship{
		User:         user,
		Counterparty: counterparty,
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
	_, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid user address")
	}

	_, err = sdk.AccAddressFromBech32(msg.Counterparty)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid counterparty address")
	}

	if msg.User == msg.Counterparty {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "user and counterparty must be different")
	}

	if !subspacestypes.IsValidSubspace(msg.Subspace) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "subspace must be a sha-256")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDeleteRelationship) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgDeleteRelationship) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.User)
	return []sdk.AccAddress{sender}
}

// ___________________________________________________________________________________________________________________

func NewMsgBlockUser(blocker, blocked, reason, subspace string) *MsgBlockUser {
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
	_, err := sdk.AccAddressFromBech32(msg.Blocker)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid blocker address")
	}

	_, err = sdk.AccAddressFromBech32(msg.Blocked)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid blocked address")
	}

	if msg.Blocker == msg.Blocked {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "blocker and blocked must be different")
	}

	if !subspacestypes.IsValidSubspace(msg.Subspace) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "subspace must be a valid sha-256 hash")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgBlockUser) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgBlockUser) GetSigners() []sdk.AccAddress {
	blocker, _ := sdk.AccAddressFromBech32(msg.Blocker)
	return []sdk.AccAddress{blocker}
}

// ___________________________________________________________________________________________________________________

func NewMsgUnblockUser(blocker, blocked, subspace string) *MsgUnblockUser {
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
	_, err := sdk.AccAddressFromBech32(msg.Blocker)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid blocker")
	}

	_, err = sdk.AccAddressFromBech32(msg.Blocked)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid blocked")
	}

	if msg.Blocker == msg.Blocked {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "blocker and blocked must be different")
	}

	if !subspacestypes.IsValidSubspace(msg.Subspace) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "subspace must be a valid sha-256 hash")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgUnblockUser) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgUnblockUser) GetSigners() []sdk.AccAddress {
	blocker, _ := sdk.AccAddressFromBech32(msg.Blocker)
	return []sdk.AccAddress{blocker}
}
