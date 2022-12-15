package types

import (
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	feegranttypes "github.com/cosmos/cosmos-sdk/x/feegrant"
	proto "github.com/gogo/protobuf/proto"
)

var (
	_ sdk.Msg = &MsgGrantUserAllowance{}
	_ sdk.Msg = &MsgRevokeUserAllowance{}
	_ sdk.Msg = &MsgGrantGroupAllowance{}
	_ sdk.Msg = &MsgRevokeGroupAllowance{}

	_ legacytx.LegacyMsg = &MsgGrantUserAllowance{}
	_ legacytx.LegacyMsg = &MsgRevokeUserAllowance{}
	_ legacytx.LegacyMsg = &MsgGrantGroupAllowance{}
	_ legacytx.LegacyMsg = &MsgRevokeGroupAllowance{}

	_ codectypes.UnpackInterfacesMessage = &MsgGrantUserAllowance{}
	_ codectypes.UnpackInterfacesMessage = &MsgGrantGroupAllowance{}
)

// NewMsgGrantUserAllowance creates a new MsgGrantUserAllowance instance
func NewMsgGrantUserAllowance(subspaceID uint64, granter string, grantee string, allowance feegranttypes.FeeAllowanceI) *MsgGrantUserAllowance {
	msg, ok := allowance.(proto.Message)
	if !ok {
		panic("cannot proto marshal allowance")
	}
	any, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		panic("failed to pack allowance to any type")
	}
	return &MsgGrantUserAllowance{
		SubspaceID: subspaceID,
		Granter:    granter,
		Grantee:    grantee,
		Allowance:  any,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgGrantUserAllowance) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", msg.SubspaceID)
	}
	_, err := sdk.AccAddressFromBech32(msg.Granter)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing granter address")
	}
	_, err = sdk.AccAddressFromBech32(msg.Grantee)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing grantee address")
	}
	if msg.Grantee == msg.Granter {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "cannot self-grant fee authorization")
	}
	allowance, err := msg.GetUnpackedAllowance()
	if err != nil {
		return err
	}

	return allowance.ValidateBasic()
}

// GetSigners implements sdk.Msg
func (msg MsgGrantUserAllowance) GetSigners() []sdk.AccAddress {
	granter, err := sdk.AccAddressFromBech32(msg.Granter)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{granter}
}

// Type implements legacytx.LegacyMsg
func (msg MsgGrantUserAllowance) Type() string {
	return ActionGrantUserAllowance
}

// Route implements legacytx.LegacyMsg
func (msg MsgGrantUserAllowance) Route() string {
	return RouterKey
}

// GetSignBytes implements legacytx.LegacyMsg
func (msg MsgGrantUserAllowance) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetUnpackedAllowance gets the unpacked allowance from the cached value of the allowance
func (msg MsgGrantUserAllowance) GetUnpackedAllowance() (feegranttypes.FeeAllowanceI, error) {
	allowance, ok := msg.Allowance.GetCachedValue().(feegranttypes.FeeAllowanceI)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid allowance type %T", allowance)
	}
	return allowance, nil
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (msg MsgGrantUserAllowance) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var allowance feegranttypes.FeeAllowanceI
	return unpacker.UnpackAny(msg.Allowance, &allowance)
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgRevokeUserAllowance creates a new MsgRevokeUserAllowance instance
func NewMsgRevokeUserAllowance(subspaceID uint64, granter string, grantee string) *MsgRevokeUserAllowance {
	return &MsgRevokeUserAllowance{
		SubspaceID: subspaceID,
		Granter:    granter,
		Grantee:    grantee,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgRevokeUserAllowance) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", msg.SubspaceID)
	}
	_, err := sdk.AccAddressFromBech32(msg.Granter)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing granter address")
	}
	_, err = sdk.AccAddressFromBech32(msg.Grantee)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing grantee address")
	}
	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgRevokeUserAllowance) GetSigners() []sdk.AccAddress {
	granter, err := sdk.AccAddressFromBech32(msg.Granter)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{granter}
}

// Type implements legacytx.LegacyMsg
func (msg MsgRevokeUserAllowance) Type() string {
	return ActionRevokeUserAllowance
}

// Route implements legacytx.LegacyMsg
func (msg MsgRevokeUserAllowance) Route() string {
	return RouterKey
}

// GetSignBytes implements legacytx.LegacyMsg
func (msg MsgRevokeUserAllowance) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgGrantGroupAllowance creates a new MsgGrantGroupAllowance instance
func NewMsgGrantGroupAllowance(subspaceID uint64, granter string, groupID uint32, allowance feegranttypes.FeeAllowanceI) *MsgGrantGroupAllowance {
	msg, ok := allowance.(proto.Message)
	if !ok {
		panic("cannot proto marshal allowance")
	}
	any, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		panic("failed to pack allowance to any type")
	}
	return &MsgGrantGroupAllowance{
		SubspaceID: subspaceID,
		Granter:    granter,
		GroupID:    groupID,
		Allowance:  any,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgGrantGroupAllowance) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", msg.SubspaceID)
	}
	if msg.GroupID == 0 {
		return fmt.Errorf("invalid group id: %d", msg.GroupID)
	}
	_, err := sdk.AccAddressFromBech32(msg.Granter)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing granter address")
	}

	f, err := msg.GetUnpackedAllowance()
	if err != nil {
		return err
	}
	return f.ValidateBasic()
}

// GetSigners implements sdk.Msg
func (msg MsgGrantGroupAllowance) GetSigners() []sdk.AccAddress {
	granter, err := sdk.AccAddressFromBech32(msg.Granter)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{granter}
}

// Type implements legacytx.LegacyMsg
func (msg MsgGrantGroupAllowance) Type() string {
	return ActionGrantGroupAllowance
}

// Route implements legacytx.LegacyMsg
func (msg MsgGrantGroupAllowance) Route() string {
	return RouterKey
}

// GetSignBytes implements legacytx.LegacyMsg
func (msg MsgGrantGroupAllowance) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetUnpackedAllowance gets the unpacked allowance from the cached value of the allowance
func (msg MsgGrantGroupAllowance) GetUnpackedAllowance() (feegranttypes.FeeAllowanceI, error) {
	allowance, ok := msg.Allowance.GetCachedValue().(feegranttypes.FeeAllowanceI)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid allowance type %T", allowance)
	}

	return allowance, nil
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (msg MsgGrantGroupAllowance) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var allowance feegranttypes.FeeAllowanceI
	return unpacker.UnpackAny(msg.Allowance, &allowance)
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgRevokeGroupAllowance creates a new MsgRevokeGroupAllowance instance
func NewMsgRevokeGroupAllowance(subspaceID uint64, granter string, groupID uint32) *MsgRevokeGroupAllowance {
	return &MsgRevokeGroupAllowance{
		SubspaceID: subspaceID,
		Granter:    granter,
		GroupID:    groupID,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgRevokeGroupAllowance) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", msg.SubspaceID)
	}
	if msg.GroupID == 0 {
		return fmt.Errorf("invalid group id: %d", msg.GroupID)
	}
	_, err := sdk.AccAddressFromBech32(msg.Granter)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing granter address")
	}
	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgRevokeGroupAllowance) GetSigners() []sdk.AccAddress {
	granter, err := sdk.AccAddressFromBech32(msg.Granter)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{granter}
}

// Type implements sdk.Msg
func (msg MsgRevokeGroupAllowance) Type() string {
	return ActionRevokeGroupAllowance
}

// Route implements sdk.Msg
func (msg MsgRevokeGroupAllowance) Route() string {
	return RouterKey
}

// GetSignBytes implements sdk.Msg
func (msg MsgRevokeGroupAllowance) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}
