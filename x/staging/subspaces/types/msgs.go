package types

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/commons"
	"strings"
)

// NewMsgCreateSubspace is a constructor function for MsgCreateSubspace
func NewMsgCreateSubspace(id, name, creator string) *MsgCreateSubspace {
	return &MsgCreateSubspace{
		ID:      id,
		Name:    name,
		Creator: creator,
	}
}

// Route should return the name of the module
func (msg MsgCreateSubspace) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateSubspace) Type() string { return ActionCreateSubspace }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateSubspace) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	if !commons.IsValidSubspace(msg.ID) {
		return sdkerrors.Wrap(ErrInvalidSubspace, "subspace id must be a valid sha-256 hash")
	}

	if strings.TrimSpace(msg.Name) == "" {
		return sdkerrors.Wrap(ErrInvalidSubspaceName, "subspace name cannot be empty or blank")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateSubspace) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines the required signature
func (msg MsgCreateSubspace) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{addr}
}

// MarshalJSON implements the json.Mashaler interface.
// This is done due to the fact that Amino does not respect omitempty clauses
func (msg MsgCreateSubspace) MarshalJSON() ([]byte, error) {
	type temp MsgCreateSubspace
	return json.Marshal(temp(msg))
}

// NewMsgAddAdmin is a constructor function for MsgAddAdmin
func NewMsgAddAdmin(id, newAdmin, creator string) *MsgAddAdmin {
	return &MsgAddAdmin{
		SubspaceID: id,
		NewAdmin:   newAdmin,
		Creator:    creator,
	}
}

// Route should return the name of the module
func (msg MsgAddAdmin) Route() string { return RouterKey }

// Type should return the action
func (msg MsgAddAdmin) Type() string { return ActionAddAdmin }

// ValidateBasic runs stateless checks on the message
func (msg MsgAddAdmin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	_, err = sdk.AccAddressFromBech32(msg.NewAdmin)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid new admin address")
	}

	if !commons.IsValidSubspace(msg.SubspaceID) {
		return sdkerrors.Wrap(ErrInvalidSubspace, "subspace id must be a valid sha-256 hash")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgAddAdmin) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines the required signature
func (msg MsgAddAdmin) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{addr}
}

// MarshalJSON implements the json.Mashaler interface.
// This is done due to the fact that Amino does not respect omitempty clauses
func (msg MsgAddAdmin) MarshalJSON() ([]byte, error) {
	type temp MsgAddAdmin
	return json.Marshal(temp(msg))
}

// NewMsgRemoveAdmin is a constructor function for MsgRemoveAdmin
func NewMsgRemoveAdmin(id, admin, creator string) *MsgRemoveAdmin {
	return &MsgRemoveAdmin{
		SubspaceID: id,
		Admin:      admin,
		Creator:    creator,
	}
}

// Route should return the name of the module
func (msg MsgRemoveAdmin) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRemoveAdmin) Type() string { return ActionRemoveAdmin }

// ValidateBasic runs stateless checks on the message
func (msg MsgRemoveAdmin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address")
	}

	_, err = sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	if !commons.IsValidSubspace(msg.SubspaceID) {
		return sdkerrors.Wrap(ErrInvalidSubspace, "subspace id must be a valid sha-256 hash")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRemoveAdmin) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines the required signature
func (msg MsgRemoveAdmin) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{addr}
}

// MarshalJSON implements the json.Mashaler interface.
// This is done due to the fact that Amino does not respect omitempty clauses
func (msg MsgRemoveAdmin) MarshalJSON() ([]byte, error) {
	type temp MsgRemoveAdmin
	return json.Marshal(temp(msg))
}

// NewMsgEnableUserPosts is a constructor function for MsgEnableUserPosts
func NewMsgEnableUserPosts(user, id, admin string) *MsgEnableUserPosts {
	return &MsgEnableUserPosts{
		User:       user,
		SubspaceID: id,
		Admin:      admin,
	}
}

// Route should return the name of the module
func (msg MsgEnableUserPosts) Route() string { return RouterKey }

// Type should return the action
func (msg MsgEnableUserPosts) Type() string { return ActionAllowUserPosts }

// ValidateBasic runs stateless checks on the message
func (msg MsgEnableUserPosts) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address")
	}

	_, err = sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address")
	}

	if !commons.IsValidSubspace(msg.SubspaceID) {
		return sdkerrors.Wrap(ErrInvalidSubspace, "subspace id must be a valid sha-256 hash")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgEnableUserPosts) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines the required signature
func (msg MsgEnableUserPosts) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Admin)
	return []sdk.AccAddress{addr}
}

// MarshalJSON implements the json.Mashaler interface.
// This is done due to the fact that Amino does not respect omitempty clauses
func (msg MsgEnableUserPosts) MarshalJSON() ([]byte, error) {
	type temp MsgEnableUserPosts
	return json.Marshal(temp(msg))
}

// NewMsgDisableUserPosts is a constructor function for MsgDisableUserPosts
func NewMsgDisableUserPosts(user, id, admin string) *MsgDisableUserPosts {
	return &MsgDisableUserPosts{
		User:       user,
		SubspaceID: id,
		Admin:      admin,
	}
}

// Route should return the name of the module
func (msg MsgDisableUserPosts) Route() string { return RouterKey }

// Type should return the action
func (msg MsgDisableUserPosts) Type() string { return ActionBlockUserPosts }

// ValidateBasic runs stateless checks on the message
func (msg MsgDisableUserPosts) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address")
	}

	_, err = sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address")
	}

	if !commons.IsValidSubspace(msg.SubspaceID) {
		return sdkerrors.Wrap(ErrInvalidSubspace, "subspace id must be a valid sha-256 hash")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDisableUserPosts) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines the required signature
func (msg MsgDisableUserPosts) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Admin)
	return []sdk.AccAddress{addr}
}

// MarshalJSON implements the json.Mashaler interface.
// This is done due to the fact that Amino does not respect omitempty clauses
func (msg MsgDisableUserPosts) MarshalJSON() ([]byte, error) {
	type temp MsgDisableUserPosts
	return json.Marshal(temp(msg))
}
