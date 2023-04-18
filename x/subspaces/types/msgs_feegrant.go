package types

import (
	"fmt"

	errors "cosmossdk.io/errors"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
	feegranttypes "github.com/cosmos/cosmos-sdk/x/feegrant"
	proto "github.com/gogo/protobuf/proto"
)

var (
	_ sdk.Msg = &MsgGrantAllowance{}
	_ sdk.Msg = &MsgRevokeAllowance{}

	_ legacytx.LegacyMsg = &MsgGrantAllowance{}
	_ legacytx.LegacyMsg = &MsgRevokeAllowance{}

	_ codectypes.UnpackInterfacesMessage = &MsgGrantAllowance{}
	_ codectypes.UnpackInterfacesMessage = &MsgRevokeAllowance{}
)

// NewMsgGrantAllowance creates a new MsgGrantAllowance instance
func NewMsgGrantAllowance(subspaceID uint64, granter string, grantee Grantee, allowance feegranttypes.FeeAllowanceI) *MsgGrantAllowance {
	allowanceProto, isProto := allowance.(proto.Message)
	if !isProto {
		panic("cannot proto marshal allowance")
	}

	allowanceAny, err := codectypes.NewAnyWithValue(allowanceProto)
	if err != nil {
		panic("failed to pack allowance to any type")
	}

	granteeAny, err := codectypes.NewAnyWithValue(grantee)
	if err != nil {
		panic("failed to pack grantee to any type")
	}

	return &MsgGrantAllowance{
		SubspaceID: subspaceID,
		Granter:    granter,
		Grantee:    granteeAny,
		Allowance:  allowanceAny,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgGrantAllowance) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", msg.SubspaceID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Granter)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "invalid granter address")
	}

	grantee := msg.Grantee.GetCachedValue().(Grantee)
	err = grantee.Validate()
	if err != nil {
		return err
	}

	if u, ok := grantee.(*UserGrantee); ok {
		if u.User == msg.Granter {
			return errors.Wrap(sdkerrors.ErrInvalidAddress, "cannot self-grant fee authorization")
		}
	}

	allowance, err := msg.GetUnpackedAllowance()
	if err != nil {
		return err
	}

	return allowance.ValidateBasic()
}

// GetSigners implements sdk.Msg
func (msg MsgGrantAllowance) GetSigners() []sdk.AccAddress {
	granter, err := sdk.AccAddressFromBech32(msg.Granter)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{granter}
}

// Type implements legacytx.LegacyMsg
func (msg MsgGrantAllowance) Type() string {
	return ActionGrantAllowance
}

// Route implements legacytx.LegacyMsg
func (msg MsgGrantAllowance) Route() string {
	return RouterKey
}

// GetSignBytes implements legacytx.LegacyMsg
func (msg MsgGrantAllowance) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetUnpackedAllowance gets the unpacked allowance from the cached value of the allowance
func (msg MsgGrantAllowance) GetUnpackedAllowance() (feegranttypes.FeeAllowanceI, error) {
	allowance, ok := msg.Allowance.GetCachedValue().(feegranttypes.FeeAllowanceI)
	if !ok {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid allowance type %T", allowance)
	}
	return allowance, nil
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (msg MsgGrantAllowance) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var allowance feegranttypes.FeeAllowanceI
	err := unpacker.UnpackAny(msg.Allowance, &allowance)
	if err != nil {
		return err
	}
	var grantee Grantee
	return unpacker.UnpackAny(msg.Grantee, &grantee)
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgRevokeAllowance creates a new MsgRevokeAllowance instance
func NewMsgRevokeAllowance(subspaceID uint64, granter string, grantee Grantee) *MsgRevokeAllowance {
	granteeAny, err := codectypes.NewAnyWithValue(grantee)
	if err != nil {
		panic("failed to pack grantee to any type")
	}

	return &MsgRevokeAllowance{
		SubspaceID: subspaceID,
		Granter:    granter,
		Grantee:    granteeAny,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgRevokeAllowance) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", msg.SubspaceID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Granter)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "invalid granter address")
	}

	return msg.Grantee.GetCachedValue().(Grantee).Validate()
}

// GetSigners implements sdk.Msg
func (msg MsgRevokeAllowance) GetSigners() []sdk.AccAddress {
	granter, err := sdk.AccAddressFromBech32(msg.Granter)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{granter}
}

// Type implements legacytx.LegacyMsg
func (msg MsgRevokeAllowance) Type() string {
	return ActionRevokeAllowance
}

// Route implements legacytx.LegacyMsg
func (msg MsgRevokeAllowance) Route() string {
	return RouterKey
}

// GetSignBytes implements legacytx.LegacyMsg
func (msg MsgRevokeAllowance) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (msg MsgRevokeAllowance) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var grantee Grantee
	return unpacker.UnpackAny(msg.Grantee, &grantee)
}
