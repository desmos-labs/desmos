package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewMsgCreateRelationship returns a new MsgCreateRelationship instance
func NewMsgCreateRelationship(signer, counterparty string, subspaceID uint64) *MsgCreateRelationship {
	return &MsgCreateRelationship{
		Signer:       signer,
		Counterparty: counterparty,
		SubspaceID:   subspaceID,
	}
}

// Route should return the name of the module
func (msg *MsgCreateRelationship) Route() string { return RouterKey }

// Type should return the action
func (msg *MsgCreateRelationship) Type() string {
	return ActionCreateRelationship
}

// ValidateBasic runs stateless checks on the message
func (msg *MsgCreateRelationship) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "invalid signer address")
	}

	_, err = sdk.AccAddressFromBech32(msg.Counterparty)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "invalid counterparty address")
	}

	if msg.Signer == msg.Counterparty {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "signer and counterparty must be different")
	}

	if msg.SubspaceID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgCreateRelationship) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg *MsgCreateRelationship) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{sender}
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgDeleteRelationship returns a new MsgDeleteRelationship instance
func NewMsgDeleteRelationship(signer, counterparty string, subspaceID uint64) *MsgDeleteRelationship {
	return &MsgDeleteRelationship{
		Signer:       signer,
		Counterparty: counterparty,
		SubspaceID:   subspaceID,
	}
}

// Route should return the name of the module
func (msg *MsgDeleteRelationship) Route() string { return RouterKey }

// Type should return the action
func (msg *MsgDeleteRelationship) Type() string {
	return ActionDeleteRelationship
}

// ValidateBasic runs stateless checks on the message
func (msg *MsgDeleteRelationship) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "invalid user address")
	}

	_, err = sdk.AccAddressFromBech32(msg.Counterparty)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "invalid counterparty address")
	}

	if msg.Signer == msg.Counterparty {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "signer and counterparty must be different")
	}

	if msg.SubspaceID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgDeleteRelationship) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg *MsgDeleteRelationship) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{sender}
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgBlockUser returns a new MsgBlockUser instance
func NewMsgBlockUser(blocker, blocked, reason string, subspaceID uint64) *MsgBlockUser {
	return &MsgBlockUser{
		Blocker:    blocker,
		Blocked:    blocked,
		Reason:     reason,
		SubspaceID: subspaceID,
	}
}

// Route should return the name of the module
func (msg *MsgBlockUser) Route() string { return RouterKey }

// Type should return the action
func (msg *MsgBlockUser) Type() string {
	return ActionBlockUser
}

// ValidateBasic runs stateless checks on the message
func (msg *MsgBlockUser) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Blocker)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "invalid blocker address")
	}

	_, err = sdk.AccAddressFromBech32(msg.Blocked)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "invalid blocked address")
	}

	if msg.Blocker == msg.Blocked {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "blocker and blocked must be different")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgBlockUser) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg *MsgBlockUser) GetSigners() []sdk.AccAddress {
	blocker, _ := sdk.AccAddressFromBech32(msg.Blocker)
	return []sdk.AccAddress{blocker}
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgUnblockUser returns a new MsgUnblockUser instance
func NewMsgUnblockUser(blocker, blocked string, subspaceID uint64) *MsgUnblockUser {
	return &MsgUnblockUser{
		Blocker:    blocker,
		Blocked:    blocked,
		SubspaceID: subspaceID,
	}
}

// Route should return the name of the module
func (msg *MsgUnblockUser) Route() string { return RouterKey }

// Type should return the action
func (msg *MsgUnblockUser) Type() string {
	return ActionUnblockUser
}

// ValidateBasic runs stateless checks on the message
func (msg *MsgUnblockUser) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Blocker)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "invalid blocker")
	}

	_, err = sdk.AccAddressFromBech32(msg.Blocked)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "invalid blocked")
	}

	if msg.Blocker == msg.Blocked {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "blocker and blocked must be different")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgUnblockUser) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg *MsgUnblockUser) GetSigners() []sdk.AccAddress {
	blocker, _ := sdk.AccAddressFromBech32(msg.Blocker)
	return []sdk.AccAddress{blocker}
}
