package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewMsgCreateSubspace creates a new MsgCreateSubspace instance
func NewMsgCreateSubspace(name, description, owner, treasury, creator string) *MsgCreateSubspace {
	if owner == "" {
		// If the owner is empty, set the creator as the owner
		owner = creator
	}

	return &MsgCreateSubspace{
		Name:        name,
		Description: description,
		Owner:       owner,
		Treasury:    treasury,
		Creator:     creator,
	}
}

// Route implements sdk.Msg
func (msg MsgCreateSubspace) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgCreateSubspace) Type() string { return ActionCreateSubspace }

// ValidateBasic implements sdk.Msg
func (msg MsgCreateSubspace) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	_, err = sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address")
	}

	if strings.TrimSpace(msg.Name) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "subspace name cannot be empty or blank")
	}

	if msg.Treasury != "" {
		_, err = sdk.AccAddressFromBech32(msg.Treasury)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid treasury address")
		}
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
func NewMsgEditSubspace(subspaceID uint64, name, description, owner, treasury, signer string) *MsgEditSubspace {
	return &MsgEditSubspace{
		SubspaceID:  subspaceID,
		Name:        name,
		Description: description,
		Owner:       owner,
		Treasury:    treasury,
		Signer:      signer,
	}
}

// Route implements sdk.Msg
func (msg MsgEditSubspace) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgEditSubspace) Type() string { return ActionEditSubspace }

// ValidateBasic implements sdk.Msg
func (msg MsgEditSubspace) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address")
	}

	if strings.TrimSpace(msg.Owner) != "" {
		_, err = sdk.AccAddressFromBech32(msg.Owner)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address")
		}
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

// NewMsgCreateUserGroup creates a new MsgCreateUserGroup instance
func NewMsgCreateUserGroup(subspaceID uint64, name string, permissions uint32, creator string) *MsgCreateUserGroup {
	return &MsgCreateUserGroup{
		SubspaceID:         subspaceID,
		GroupName:          name,
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
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	if strings.TrimSpace(msg.GroupName) != "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "group name cannot be empty or blank")
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

// NewMsgDeleteUserGroup creates a new MsgDeleteUserGroup instance
func NewMsgDeleteUserGroup(subspaceID uint64, group string, signer string) *MsgDeleteUserGroup {
	return &MsgDeleteUserGroup{
		SubspaceID: subspaceID,
		GroupName:  group,
		Signer:     signer,
	}
}

// Route implements sdk.Msg
func (msg MsgDeleteUserGroup) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgDeleteUserGroup) Type() string { return ActionDeleteUserGroup }

// ValidateBasic implements sdk.Msg
func (msg MsgDeleteUserGroup) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address")
	}

	if strings.TrimSpace(msg.GroupName) != "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "group name cannot be empty or blank")
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
func NewMsgAddUserToUserGroup(subspaceID uint64, group string, user string, signer string) *MsgAddUserToUserGroup {
	return &MsgAddUserToUserGroup{
		SubspaceID: subspaceID,
		GroupName:  group,
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
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address")
	}

	_, err = sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address")
	}

	if strings.TrimSpace(msg.GroupName) != "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "group name cannot be empty or blank")
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
func NewMsgRemoveUserFromUserGroup(subspaceID uint64, group string, user string, signer string) *MsgRemoveUserFromUserGroup {
	return &MsgRemoveUserFromUserGroup{
		SubspaceID: subspaceID,
		GroupName:  group,
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
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address")
	}

	_, err = sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address")
	}

	if strings.TrimSpace(msg.GroupName) != "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "group name cannot be empty or blank")
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

// NewMsgSetPermissions creates a new MsgSetPermissions instance
func NewMsgSetPermissions(subspaceID uint64, target string, permissions uint32, signer string) *MsgSetPermissions {
	return &MsgSetPermissions{
		SubspaceID:  subspaceID,
		Target:      target,
		Permissions: permissions,
		Signer:      signer,
	}
}

// Route implements sdk.Msg
func (msg MsgSetPermissions) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgSetPermissions) Type() string { return ActionSetPermissions }

// ValidateBasic implements sdk.Msg
func (msg MsgSetPermissions) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address")
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg MsgSetPermissions) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgSetPermissions) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{addr}
}
