package types

import (
	"encoding/json"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/commons"
)

// NewMsgCreateSubspace is a constructor function for MsgCreateSubspace
func NewMsgCreateSubspace(id, name, description, logo, creator string, subspaceType SubspaceType) *MsgCreateSubspace {
	return &MsgCreateSubspace{
		SubspaceID:   id,
		Name:         name,
		Description:  description,
		Logo:         logo,
		Creator:      creator,
		SubspaceType: subspaceType,
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

	if !IsValidSubspace(msg.SubspaceID) {
		return sdkerrors.Wrap(ErrInvalidSubspaceID, "subspace id must be a valid SHA-256 hash")
	}

	if strings.TrimSpace(msg.Name) == "" {
		return sdkerrors.Wrap(ErrInvalidSubspaceName, "subspace name cannot be empty or blank")
	}

	validLogo := commons.IsURIValid(msg.Logo)
	if !validLogo {
		return fmt.Errorf("invalid subspace logo uri provided")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateSubspace) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
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

// NewMsgEditSubspace is a constructor function for MsgEditSubspace
func NewMsgEditSubspace(subspaceID, owner, name, description, logo, editor string, subspaceType SubspaceType) *MsgEditSubspace {
	return &MsgEditSubspace{
		ID:           subspaceID,
		Owner:        owner,
		Name:         name,
		Description:  description,
		Logo:         logo,
		Editor:       editor,
		SubspaceType: subspaceType,
	}
}

// Route should return the name of the module
func (msg MsgEditSubspace) Route() string { return RouterKey }

// Type should return the action
func (msg MsgEditSubspace) Type() string { return ActionEditSubspace }

// ValidateBasic runs stateless checks on the message
func (msg MsgEditSubspace) ValidateBasic() error {
	if msg.Editor == msg.Owner {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "the owner address is equal to the editor address")
	}

	if strings.TrimSpace(msg.Name) == "" {
		return sdkerrors.Wrap(ErrInvalidSubspaceName, "subspace name cannot be empty or blank")
	}

	_, err := sdk.AccAddressFromBech32(msg.Editor)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid editor address")
	}

	if strings.TrimSpace(msg.Owner) != "" {
		_, err = sdk.AccAddressFromBech32(msg.Owner)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address")
		}
	}

	if !IsValidSubspace(msg.ID) {
		return sdkerrors.Wrap(ErrInvalidSubspaceID, "subspace id must be a valid SHA-256 hash")
	}

	if msg.Logo != DoNotModify && strings.TrimSpace(msg.Logo) != "" && !commons.IsURIValid(msg.Logo) {
		return fmt.Errorf("invalid subspace logo uri provided")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgEditSubspace) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetSigners defines the required signature
func (msg MsgEditSubspace) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Editor)
	return []sdk.AccAddress{addr}
}

// MarshalJSON implements the json.Mashaler interface.
// This is done due to the fact that Amino does not respect omitempty clauses
func (msg MsgEditSubspace) MarshalJSON() ([]byte, error) {
	type temp MsgEditSubspace
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

	if !IsValidSubspace(msg.SubspaceID) {
		return sdkerrors.Wrap(ErrInvalidSubspaceID, "subspace id must be a valid SHA-256 hash")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgAddAdmin) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
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

	if !IsValidSubspace(msg.SubspaceID) {
		return sdkerrors.Wrap(ErrInvalidSubspaceID, "subspace id must be a valid SHA-256 hash")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRemoveAdmin) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
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
func NewMsgRegisterUser(id, user, admin string) *MsgRegisterUser {
	return &MsgRegisterUser{
		SubspaceID: id,
		User:       user,
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

	if !IsValidSubspace(msg.SubspaceID) {
		return sdkerrors.Wrap(ErrInvalidSubspaceID, "subspace id must be a valid SHA-256 hash")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRegisterUser) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
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
func NewMsgUnregisterUser(id, user, admin string) *MsgUnregisterUser {
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

	if !IsValidSubspace(msg.SubspaceID) {
		return sdkerrors.Wrap(ErrInvalidSubspaceID, "subspace id must be a valid SHA-256 hash")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgUnregisterUser) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
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

// NewMsgBanUser is a constructor function for MsgBanUser
func NewMsgBanUser(id, user, admin string) *MsgBanUser {
	return &MsgBanUser{
		User:       user,
		SubspaceID: id,
		Admin:      admin,
	}
}

// Route should return the name of the module
func (msg MsgBanUser) Route() string { return RouterKey }

// Type should return the action
func (msg MsgBanUser) Type() string { return ActionBlockUser }

// ValidateBasic runs stateless checks on the message
func (msg MsgBanUser) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address")
	}

	_, err = sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address")
	}

	if !IsValidSubspace(msg.SubspaceID) {
		return sdkerrors.Wrap(ErrInvalidSubspaceID, "subspace id must be a valid SHA-256 hash")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgBanUser) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetSigners defines the required signature
func (msg MsgBanUser) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Admin)
	return []sdk.AccAddress{addr}
}

// MarshalJSON implements the json.Mashaler interface.
// This is done due to the fact that Amino does not respect omitempty clauses
func (msg MsgBanUser) MarshalJSON() ([]byte, error) {
	type temp MsgBanUser
	return json.Marshal(temp(msg))
}

// NewMsgUnbanUser is a constructor function for MsgUnbanUser
func NewMsgUnbanUser(id, user, admin string) *MsgUnbanUser {
	return &MsgUnbanUser{
		User:       user,
		SubspaceID: id,
		Admin:      admin,
	}
}

// Route should return the name of the module
func (msg MsgUnbanUser) Route() string { return RouterKey }

// Type should return the action
func (msg MsgUnbanUser) Type() string { return ActionUnblockUser }

// ValidateBasic runs stateless checks on the message
func (msg MsgUnbanUser) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address")
	}

	_, err = sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address")
	}

	if !IsValidSubspace(msg.SubspaceID) {
		return sdkerrors.Wrap(ErrInvalidSubspaceID, "subspace id must be a valid SHA-256 hash")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgUnbanUser) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetSigners defines the required signature
func (msg MsgUnbanUser) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Admin)
	return []sdk.AccAddress{addr}
}

// MarshalJSON implements the json.Mashaler interface.
// This is done due to the fact that Amino does not respect omitempty clauses
func (msg MsgUnbanUser) MarshalJSON() ([]byte, error) {
	type temp MsgUnbanUser
	return json.Marshal(temp(msg))
}

// NewMsgSaveTokenomics is a constructor for MsgSaveTokenomics
func NewMsgSaveTokenomics(subspaceID, contractAddress, admin string, message []byte) *MsgSaveTokenomics {
	return &MsgSaveTokenomics{
		SubspaceID:      subspaceID,
		ContractAddress: contractAddress,
		Message:         message,
		Admin:           admin,
	}
}

// Route should return the name of the module
func (msg MsgSaveTokenomics) Route() string {
	return RouterKey
}

// Type should return the action
func (msg MsgSaveTokenomics) Type() string {
	return ActionSaveTokenomics
}

// ValidateBasic runs stateless checks on the message
func (msg MsgSaveTokenomics) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address")
	}

	_, err = sdk.AccAddressFromBech32(msg.ContractAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid contract address")
	}

	if !IsValidSubspace(msg.SubspaceID) {
		return sdkerrors.Wrap(ErrInvalidSubspaceID, "subspace id must be a valid SHA-256 hash")
	}

	if msg.Message == nil {
		return sdkerrors.Wrap(ErrInvalidTokenomics, "empty message bytes")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSaveTokenomics) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetSigners defines the required signature
func (msg MsgSaveTokenomics) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Admin)
	return []sdk.AccAddress{addr}
}
