package types

import (
	"fmt"
	"strings"

	errors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// SubspaceMsg represents a generic message that is related to a subspace
type SubspaceMsg interface {
	sdk.Msg

	// GetSubspaceID returns the subspace id associated to this message
	GetSubspaceID() uint64
}

var (
	_ sdk.Msg = &MsgCreateSubspace{}
	_ sdk.Msg = &MsgEditSubspace{}
	_ sdk.Msg = &MsgDeleteSubspace{}
	_ sdk.Msg = &MsgCreateSection{}
	_ sdk.Msg = &MsgEditSection{}
	_ sdk.Msg = &MsgMoveSection{}
	_ sdk.Msg = &MsgDeleteSection{}
	_ sdk.Msg = &MsgCreateUserGroup{}
	_ sdk.Msg = &MsgEditUserGroup{}
	_ sdk.Msg = &MsgMoveUserGroup{}
	_ sdk.Msg = &MsgSetUserGroupPermissions{}
	_ sdk.Msg = &MsgDeleteUserGroup{}
	_ sdk.Msg = &MsgAddUserToUserGroup{}
	_ sdk.Msg = &MsgRemoveUserFromUserGroup{}
	_ sdk.Msg = &MsgSetUserPermissions{}
	_ sdk.Msg = &MsgGrantTreasuryAuthorization{}
)

// NewMsgCreateSubspace creates a new MsgCreateSubspace instance
func NewMsgCreateSubspace(
	name string,
	description string,
	owner string,
	creator string,
) *MsgCreateSubspace {
	if owner == "" {
		// If the owner is empty, set the creator as the owner
		owner = creator
	}

	return &MsgCreateSubspace{
		Name:        name,
		Description: description,
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
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "subspace name cannot be empty or blank")
	}

	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address")
	}

	_, err = sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address")
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
func NewMsgEditSubspace(
	subspaceID uint64,
	name string,
	description,
	owner string,
	signer string,
) *MsgEditSubspace {
	return &MsgEditSubspace{
		SubspaceID:  subspaceID,
		Name:        name,
		Description: description,
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
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address")
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
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address")
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

// NewMsgCreateSection returns a new MsgCreateSection instance
func NewMsgCreateSection(
	subspaceID uint64,
	name string,
	description string,
	parentID uint32,
	creator string,
) *MsgCreateSection {
	return &MsgCreateSection{
		SubspaceID:  subspaceID,
		Name:        name,
		Description: description,
		ParentID:    parentID,
		Creator:     creator,
	}
}

// Route implements sdk.Msg
func (msg MsgCreateSection) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgCreateSection) Type() string { return ActionCreateSection }

// ValidateBasic implements sdk.Msg
func (msg MsgCreateSection) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if strings.TrimSpace(msg.Name) == "" {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid section name: %s", msg.Name)
	}

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg MsgCreateSection) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgCreateSection) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{addr}
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgEditSection returns a new MsgEditSection instance
func NewMsgEditSection(
	subspaceID uint64,
	sectionID uint32,
	name string,
	description string,
	editor string,
) *MsgEditSection {
	return &MsgEditSection{
		SubspaceID:  subspaceID,
		SectionID:   sectionID,
		Name:        name,
		Description: description,
		Editor:      editor,
	}
}

// Route implements sdk.Msg
func (msg MsgEditSection) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgEditSection) Type() string { return ActionEditSection }

// ValidateBasic implements sdk.Msg
func (msg MsgEditSection) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if strings.TrimSpace(msg.Name) == "" {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid section name: %s", msg.Name)
	}

	_, err := sdk.AccAddressFromBech32(msg.Editor)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid editor address")
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg MsgEditSection) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgEditSection) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Editor)
	return []sdk.AccAddress{addr}
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgMoveSection returns a new MsgMoveSection instance
func NewMsgMoveSection(
	subspaceID uint64,
	sectionID uint32,
	newParentID uint32,
	signer string,
) *MsgMoveSection {
	return &MsgMoveSection{
		SubspaceID:  subspaceID,
		SectionID:   sectionID,
		NewParentID: newParentID,
		Signer:      signer,
	}
}

// Route implements sdk.Msg
func (msg MsgMoveSection) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgMoveSection) Type() string { return ActionMoveSection }

// ValidateBasic implements sdk.Msg
func (msg MsgMoveSection) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.SectionID == RootSectionID {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid section id: %d", msg.SectionID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address")
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg MsgMoveSection) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgMoveSection) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{addr}
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgDeleteSection returns a new MsgDeleteSection instance
func NewMsgDeleteSection(subspaceID uint64, sectionID uint32, signer string) *MsgDeleteSection {
	return &MsgDeleteSection{
		SubspaceID: subspaceID,
		SectionID:  sectionID,
		Signer:     signer,
	}
}

// Route implements sdk.Msg
func (msg MsgDeleteSection) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgDeleteSection) Type() string { return ActionDeleteSection }

// ValidateBasic implements sdk.Msg
func (msg MsgDeleteSection) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.SectionID == RootSectionID {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid section id: %d", msg.SectionID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address")
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg MsgDeleteSection) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgDeleteSection) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{addr}
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgCreateUserGroup creates a new MsgCreateUserGroup instance
func NewMsgCreateUserGroup(
	subspaceID uint64,
	sectionID uint32,
	name string,
	description string,
	permissions Permissions,
	initialMembers []string,
	creator string,
) *MsgCreateUserGroup {
	return &MsgCreateUserGroup{
		SubspaceID:         subspaceID,
		SectionID:          sectionID,
		Name:               name,
		Description:        description,
		DefaultPermissions: permissions,
		InitialMembers:     initialMembers,
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
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if strings.TrimSpace(msg.Name) == "" {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid group name: %s", msg.Name)
	}

	if !ArePermissionsValid(msg.DefaultPermissions) {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid permissions: %s", msg.DefaultPermissions)
	}

	for _, member := range msg.InitialMembers {
		_, err := sdk.AccAddressFromBech32(member)
		if err != nil {
			return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid member address: %s", member)
		}
	}

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address")
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
func NewMsgEditUserGroup(
	subspaceID uint64,
	groupID uint32,
	name string,
	description string,
	signer string,
) *MsgEditUserGroup {
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
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address")
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

// NewMsgMoveUserGroup returns a new NewMsgMoveUserGroup instance
func NewMsgMoveUserGroup(
	subspaceID uint64,
	groupID uint32,
	newSectionID uint32,
	signer string,
) *MsgMoveUserGroup {
	return &MsgMoveUserGroup{
		SubspaceID:   subspaceID,
		GroupID:      groupID,
		NewSectionID: newSectionID,
		Signer:       signer,
	}
}

// Route implements sdk.Msg
func (msg MsgMoveUserGroup) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgMoveUserGroup) Type() string { return ActionMoveUserGroup }

// ValidateBasic implements sdk.Msg
func (msg MsgMoveUserGroup) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.GroupID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid group id: %d", msg.GroupID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address")
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg MsgMoveUserGroup) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgMoveUserGroup) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{addr}
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgSetUserGroupPermissions returns a new MsgSetUserGroupPermissions instance
func NewMsgSetUserGroupPermissions(
	subspaceID uint64,
	groupID uint32,
	permissions Permissions,
	signer string,
) *MsgSetUserGroupPermissions {
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
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if !ArePermissionsValid(msg.Permissions) {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid permissions: %s", msg.Permissions)
	}

	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address")
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
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.GroupID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid group id: %d", msg.GroupID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address")
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
func NewMsgAddUserToUserGroup(
	subspaceID uint64,
	groupID uint32,
	user string,
	signer string,
) *MsgAddUserToUserGroup {
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
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.GroupID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid group id: %d", msg.GroupID)
	}

	_, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address")
	}

	_, err = sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address")
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
func NewMsgRemoveUserFromUserGroup(
	subspaceID uint64,
	groupID uint32,
	user string,
	signer string,
) *MsgRemoveUserFromUserGroup {
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
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.GroupID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid group id: %d", msg.GroupID)
	}

	_, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address")
	}

	_, err = sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address")
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
func NewMsgSetUserPermissions(
	subspaceID uint64,
	sectionID uint32,
	user string,
	permissions Permissions,
	signer string,
) *MsgSetUserPermissions {
	return &MsgSetUserPermissions{
		SubspaceID:  subspaceID,
		SectionID:   sectionID,
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
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	_, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address")
	}

	_, err = sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address")
	}

	if !ArePermissionsValid(msg.Permissions) {
		return fmt.Errorf("invalid permissions value: %s", msg.Permissions)
	}

	if msg.User == msg.Signer {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "cannot set the permissions for yourself")
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
