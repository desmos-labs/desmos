package types

import (
	"encoding/json"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/commons"
)

// NewMsgCreateSubspace is a constructor function for MsgCreateSubspace
func NewMsgCreateSubspace(id, name, creator string, open bool) *MsgCreateSubspace {
	return &MsgCreateSubspace{
		SubspaceID: id,
		Name:       name,
		Creator:    creator,
		Open:       open,
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

	if !commons.IsValidSubspace(msg.SubspaceID) {
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
func NewMsgAddAdmin(id, admin, owner string) *MsgAddAdmin {
	return &MsgAddAdmin{
		SubspaceID: id,
		Admin:      admin,
		Owner:      owner,
	}
}

// Route should return the name of the module
func (msg MsgAddAdmin) Route() string { return RouterKey }

// Type should return the action
func (msg MsgAddAdmin) Type() string { return ActionAddAdmin }

// ValidateBasic runs stateless checks on the message
func (msg MsgAddAdmin) ValidateBasic() error {
	if msg.Owner == msg.Admin {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "owner address can't be equal to admin address")
	}

	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address")
	}

	_, err = sdk.AccAddressFromBech32(msg.Admin)
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
	addr, _ := sdk.AccAddressFromBech32(msg.Owner)
	return []sdk.AccAddress{addr}
}

// MarshalJSON implements the json.Mashaler interface.
// This is done due to the fact that Amino does not respect omitempty clauses
func (msg MsgAddAdmin) MarshalJSON() ([]byte, error) {
	type temp MsgAddAdmin
	return json.Marshal(temp(msg))
}

// NewMsgRemoveAdmin is a constructor function for MsgRemoveAdmin
func NewMsgRemoveAdmin(id, admin, owner string) *MsgRemoveAdmin {
	return &MsgRemoveAdmin{
		SubspaceID: id,
		Admin:      admin,
		Owner:      owner,
	}
}

// Route should return the name of the module
func (msg MsgRemoveAdmin) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRemoveAdmin) Type() string { return ActionRemoveAdmin }

// ValidateBasic runs stateless checks on the message
func (msg MsgRemoveAdmin) ValidateBasic() error {
	if msg.Owner == msg.Admin {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "owner address can't be equal to admin address")
	}

	_, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address")
	}

	_, err = sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address")
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
	addr, _ := sdk.AccAddressFromBech32(msg.Owner)
	return []sdk.AccAddress{addr}
}

// MarshalJSON implements the json.Mashaler interface.
// This is done due to the fact that Amino does not respect omitempty clauses
func (msg MsgRemoveAdmin) MarshalJSON() ([]byte, error) {
	type temp MsgRemoveAdmin
	return json.Marshal(temp(msg))
}

// NewMsgRegisterUser is a constructor function for MsgRegisterUser
func NewMsgRegisterUser(user, id, admin string) *MsgRegisterUser {
	return &MsgRegisterUser{
		User:       user,
		SubspaceID: id,
		Admin:      admin,
	}
}

// Route should return the name of the module
func (msg MsgRegisterUser) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRegisterUser) Type() string { return ActionRegisterUser }

// ValidateBasic runs stateless checks on the message
func (msg MsgRegisterUser) ValidateBasic() error {
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
func (msg MsgRegisterUser) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines the required signature
func (msg MsgRegisterUser) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Admin)
	return []sdk.AccAddress{addr}
}

// MarshalJSON implements the json.Mashaler interface.
// This is done due to the fact that Amino does not respect omitempty clauses
func (msg MsgRegisterUser) MarshalJSON() ([]byte, error) {
	type temp MsgRegisterUser
	return json.Marshal(temp(msg))
}

// NewMsgUnregisterUser is a constructor function for MsgUnregisterUser
func NewMsgUnregisterUser(user, id, admin string) *MsgUnregisterUser {
	return &MsgUnregisterUser{
		User:       user,
		SubspaceID: id,
		Admin:      admin,
	}
}

// Route should return the name of the module
func (msg MsgUnregisterUser) Route() string { return RouterKey }

// Type should return the action
func (msg MsgUnregisterUser) Type() string { return ActionUnregisterUser }

// ValidateBasic runs stateless checks on the message
func (msg MsgUnregisterUser) ValidateBasic() error {
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
func (msg MsgUnregisterUser) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines the required signature
func (msg MsgUnregisterUser) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Admin)
	return []sdk.AccAddress{addr}
}

// MarshalJSON implements the json.Mashaler interface.
// This is done due to the fact that Amino does not respect omitempty clauses
func (msg MsgUnregisterUser) MarshalJSON() ([]byte, error) {
	type temp MsgUnregisterUser
	return json.Marshal(temp(msg))
}

// NewMsgBlockUser is a constructor function for MsgBlockUser
func NewMsgBlockUser(user, id, admin string) *MsgBlockUser {
	return &MsgBlockUser{
		User:       user,
		SubspaceID: id,
		Admin:      admin,
	}
}

// Route should return the name of the module
func (msg MsgBlockUser) Route() string { return RouterKey }

// Type should return the action
func (msg MsgBlockUser) Type() string { return ActionBlockUser }

// ValidateBasic runs stateless checks on the message
func (msg MsgBlockUser) ValidateBasic() error {
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
func (msg MsgBlockUser) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines the required signature
func (msg MsgBlockUser) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Admin)
	return []sdk.AccAddress{addr}
}

// MarshalJSON implements the json.Mashaler interface.
// This is done due to the fact that Amino does not respect omitempty clauses
func (msg MsgBlockUser) MarshalJSON() ([]byte, error) {
	type temp MsgBlockUser
	return json.Marshal(temp(msg))
}

// NewMsgUnblockUser is a constructor function for MsgUnblockUser
func NewMsgUnblockUser(user, id, admin string) *MsgUnblockUser {
	return &MsgUnblockUser{
		User:       user,
		SubspaceID: id,
		Admin:      admin,
	}
}

// Route should return the name of the module
func (msg MsgUnblockUser) Route() string { return RouterKey }

// Type should return the action
func (msg MsgUnblockUser) Type() string { return ActionUnblockUser }

// ValidateBasic runs stateless checks on the message
func (msg MsgUnblockUser) ValidateBasic() error {
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
func (msg MsgUnblockUser) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines the required signature
func (msg MsgUnblockUser) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Admin)
	return []sdk.AccAddress{addr}
}

// MarshalJSON implements the json.Mashaler interface.
// This is done due to the fact that Amino does not respect omitempty clauses
func (msg MsgUnblockUser) MarshalJSON() ([]byte, error) {
	type temp MsgUnblockUser
	return json.Marshal(temp(msg))
}

// NewMsgEditSubspace is a constructor function for MsgEditSubspace
func NewMsgEditSubspace(subspaceID, newOwner, newName, owner string) *MsgEditSubspace {
	return &MsgEditSubspace{
		ID:       subspaceID,
		NewOwner: newOwner,
		NewName:  newName,
		Owner:    owner,
	}
}

// Route should return the name of the module
func (msg MsgEditSubspace) Route() string { return RouterKey }

// Type should return the action
func (msg MsgEditSubspace) Type() string { return ActionEditSubspace }

// ValidateBasic runs stateless checks on the message
func (msg MsgEditSubspace) ValidateBasic() error {
	if msg.Owner == msg.NewOwner {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "the new owner address is equal to the owner address")
	}

	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address")
	}

	if strings.TrimSpace(msg.NewOwner) != "" {
		_, err = sdk.AccAddressFromBech32(msg.NewOwner)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid new owner address")
		}
	}

	if !commons.IsValidSubspace(msg.ID) {
		return sdkerrors.Wrap(ErrInvalidSubspace, "subspace id must be a valid sha-256 hash")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgEditSubspace) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines the required signature
func (msg MsgEditSubspace) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Owner)
	return []sdk.AccAddress{addr}
}

// MarshalJSON implements the json.Mashaler interface.
// This is done due to the fact that Amino does not respect omitempty clauses
func (msg MsgEditSubspace) MarshalJSON() ([]byte, error) {
	type temp MsgEditSubspace
	return json.Marshal(temp(msg))
}
