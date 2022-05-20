package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewMsgCreateSubspace creates a new MsgCreateSubspace instance
func NewMsgCreateSubspace(name, description, treasury, owner, creator string) *MsgCreateSubspace {
	if owner == "" {
		// If the owner is empty, set the creator as the owner
		owner = creator
	}

	return &MsgCreateSubspace{
		Name:        name,
		Description: description,
		Treasury:    treasury,
		Owner:       owner,
		Creator:     creator,
	}
}

// Route implements sdk.Msg
func (msg MsgCreateSubspace) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgCreateSubspace) Type() string { return ActionCreateSubspace }

// ValidateBasic implements sdk.Msg
func (msg MsgCreateSubspace) ValidateBasic() error {
	if strings.TrimSpace(msg.Name) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "subspace name cannot be empty or blank")
	}

	if msg.Treasury != "" {
		_, err := sdk.AccAddressFromBech32(msg.Treasury)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid treasury address")
		}
	}

	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address")
	}

	_, err = sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg MsgCreateSubspace) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgCreateSubspace) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{addr}
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgEditSubspace creates a new MsgEditSubspace instance
func NewMsgEditSubspace(subspaceID uint64, name, description, treasury, owner, signer string) *MsgEditSubspace {
	return &MsgEditSubspace{
		SubspaceID:  subspaceID,
		Name:        name,
		Description: description,
		Treasury:    treasury,
		Owner:       owner,
		Signer:      signer,
	}
}

// Route implements sdk.Msg
func (msg MsgEditSubspace) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgEditSubspace) Type() string { return ActionEditSubspace }

// ValidateBasic implements sdk.Msg
func (msg MsgEditSubspace) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address")
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg MsgEditSubspace) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgEditSubspace) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{addr}
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgDeleteSubspace returns a new MsgDeleteSubspace instance
func NewMsgDeleteSubspace(subspaceID uint64, signer string) *MsgDeleteSubspace {
	return &MsgDeleteSubspace{
		SubspaceID: subspaceID,
		Signer:     signer,
	}
}

// Route implements sdk.Msg
func (msg MsgDeleteSubspace) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgDeleteSubspace) Type() string { return ActionDeleteSubspace }

// ValidateBasic implements sdk.Msg
func (msg MsgDeleteSubspace) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address")
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg MsgDeleteSubspace) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgDeleteSubspace) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{addr}
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgCreateUserGroup creates a new MsgCreateUserGroup instance
func NewMsgCreateUserGroup(subspaceID uint64, name, description string, permissions Permissions, creator string) *MsgCreateUserGroup {
	return &MsgCreateUserGroup{
		SubspaceID:         subspaceID,
		Name:               name,
		Description:        description,
		DefaultPermissions: permissions,
		Creator:            creator,
	}
}

// Route implements sdk.Msg
func (msg MsgCreateUserGroup) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgCreateUserGroup) Type() string { return ActionCreateUserGroup }

// ValidateBasic implements sdk.Msg
func (msg MsgCreateUserGroup) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if strings.TrimSpace(msg.Name) == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid group name: %s", msg.Name)
	}

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg MsgCreateUserGroup) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgCreateUserGroup) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{addr}
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgEditUserGroup returns a new NewMsgEditUserGroup instance
func NewMsgEditUserGroup(subspaceID uint64, groupID uint32, name, description string, signer string) *MsgEditUserGroup {
	return &MsgEditUserGroup{
		SubspaceID:  subspaceID,
		GroupID:     groupID,
		Name:        name,
		Description: description,
		Signer:      signer,
	}
}

// Route implements sdk.Msg
func (msg MsgEditUserGroup) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgEditUserGroup) Type() string { return ActionEditUserGroup }

// ValidateBasic implements sdk.Msg
func (msg MsgEditUserGroup) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address")
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg MsgEditUserGroup) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgEditUserGroup) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{addr}
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgSetUserGroupPermissions returns a new MsgSetUserGroupPermissions instance
func NewMsgSetUserGroupPermissions(subspaceID uint64, groupID uint32, permissions Permissions, signer string) *MsgSetUserGroupPermissions {
	return &MsgSetUserGroupPermissions{
		SubspaceID:  subspaceID,
		GroupID:     groupID,
		Permissions: permissions,
		Signer:      signer,
	}
}

// Route implements sdk.Msg
func (msg MsgSetUserGroupPermissions) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgSetUserGroupPermissions) Type() string { return ActionSetUserGroupPermissions }

// ValidateBasic implements sdk.Msg
func (msg MsgSetUserGroupPermissions) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address")
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg MsgSetUserGroupPermissions) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgSetUserGroupPermissions) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{addr}
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgDeleteUserGroup creates a new MsgDeleteUserGroup instance
func NewMsgDeleteUserGroup(subspaceID uint64, groupID uint32, signer string) *MsgDeleteUserGroup {
	return &MsgDeleteUserGroup{
		SubspaceID: subspaceID,
		GroupID:    groupID,
		Signer:     signer,
	}
}

// Route implements sdk.Msg
func (msg MsgDeleteUserGroup) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgDeleteUserGroup) Type() string { return ActionDeleteUserGroup }

// ValidateBasic implements sdk.Msg
func (msg MsgDeleteUserGroup) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.GroupID == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid group id: %d", msg.GroupID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address")
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg MsgDeleteUserGroup) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgDeleteUserGroup) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{addr}
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgAddUserToUserGroup creates a new MsgAddUserToUserGroup instance
func NewMsgAddUserToUserGroup(subspaceID uint64, groupID uint32, user string, signer string) *MsgAddUserToUserGroup {
	return &MsgAddUserToUserGroup{
		SubspaceID: subspaceID,
		GroupID:    groupID,
		User:       user,
		Signer:     signer,
	}
}

// Route implements sdk.Msg
func (msg MsgAddUserToUserGroup) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgAddUserToUserGroup) Type() string { return ActionAddUserToUserGroup }

// ValidateBasic implements sdk.Msg
func (msg MsgAddUserToUserGroup) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.GroupID == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid group id: %d", msg.GroupID)
	}

	_, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address")
	}

	_, err = sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address")
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg MsgAddUserToUserGroup) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgAddUserToUserGroup) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{addr}
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgRemoveUserFromUserGroup creates a new MsgRemoveUserFromUserGroup instance
func NewMsgRemoveUserFromUserGroup(subspaceID uint64, groupID uint32, user string, signer string) *MsgRemoveUserFromUserGroup {
	return &MsgRemoveUserFromUserGroup{
		SubspaceID: subspaceID,
		GroupID:    groupID,
		User:       user,
		Signer:     signer,
	}
}

// Route implements sdk.Msg
func (msg MsgRemoveUserFromUserGroup) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgRemoveUserFromUserGroup) Type() string { return ActionRemoveUserFromUserGroup }

// ValidateBasic implements sdk.Msg
func (msg MsgRemoveUserFromUserGroup) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.GroupID == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid group id: %d", msg.GroupID)
	}

	_, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address")
	}

	_, err = sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address")
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg MsgRemoveUserFromUserGroup) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgRemoveUserFromUserGroup) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{addr}
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgSetUserPermissions creates a new MsgSetUserPermissions instance
func NewMsgSetUserPermissions(subspaceID uint64, user string, permissions Permissions, signer string) *MsgSetUserPermissions {
	return &MsgSetUserPermissions{
		SubspaceID:  subspaceID,
		User:        user,
		Permissions: permissions,
		Signer:      signer,
	}
}

// Route implements sdk.Msg
func (msg MsgSetUserPermissions) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgSetUserPermissions) Type() string { return ActionSetUserPermissions }

// ValidateBasic implements sdk.Msg
func (msg MsgSetUserPermissions) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	_, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address")
	}

	_, err = sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address")
	}

	if !ArePermissionsValid(msg.Permissions) {
		return fmt.Errorf("invalid permissions value: %s", msg.Permissions)
	}

	if msg.User == msg.Signer {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "cannot set the permissions for yourself")
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg MsgSetUserPermissions) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgSetUserPermissions) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{addr}
}
